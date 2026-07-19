# Localization Architecture

How localization is built so it is incremental, maintainable by a solo
developer, safe for contributors, and meets the allocation/performance bar — and
so the locale **data** serves both GoTime (Go) and UExL (a JSON-consuming DSL)
from a single source. Grounded in the current code: `time_ago.go`, `format.go`,
`internal/nites/convert.go`, `internal/nites/format.go`.

## 1. Design Principles

1. **Opt-in.** Default behavior stays English; `*In` variants are the only
   entry points that consult a locale.
2. **One data source, multiple consumers.** Locale data is authored once as slim
   JSON. GoTime consumes it via generated Go (fast path) and via a runtime
   loader (`RegisterJSON`, dynamic path). UExL does **not** read files — its
   host loads the data and injects it as a key/value **context map**; UExL
   evaluates expressions against that context. All consumers must agree.
3. **Isolated locale logic.** Rules live in `locales/` and generated
   subpackages; core formatting never hardcodes language beyond the existing
   English fast path.
4. **Data-first.** A locale is tables + a plural-rule id + an ordinal-rule id +
   phrase tables + presets. The only *logic* is a small set of shared plural
   functions in core.
5. **Allocation-conscious.** The hot path is single-pass over a pooled buffer;
   name/ordinal/preset resolution is table/array access. No `fmt.Sprint`.
6. **No heavy runtime deps.** CLDR *model*, not a CLDR runtime library.

## 2. The Locale Contract

Defined in package `locales`. Narrow; every question is a cheap lookup or a small
pure function.

```go
package locales

type Width uint8;        const ( Short Width = iota; Long )            // mmm/www vs mmmm/wwww
type Context uint8;      const ( Format Context = iota; StandAlone )   // CLDR format vs stand-alone
type OrdinalKind uint8;  const ( OrdinalDay OrdinalKind = iota; OrdinalMonth )
type PluralCategory uint8; const ( Zero PluralCategory = iota; One; Two; Few; Many; Other )
type Unit uint8;         const ( UnitSecond Unit = iota; UnitMinute; UnitHour; UnitDay; UnitWeek; UnitMonth; UnitYear )
type Style uint8;        const ( StyleAuto Style = iota; StyleNumeric )

// Locale is immutable after registration and safe for concurrent use.
type Locale interface {
    Code() string

    // Names carry grammatical context (Format = in-date/genitive, StandAlone =
    // nominative). Phase-1 formatting passes Format for mmmm; StandAlone data is
    // stored for a future stand-alone token.
    Month(m time.Month, w Width, c Context) string   // table lookup, no alloc
    Weekday(d time.Weekday, w Width, c Context) string
    DayPeriod(pm bool, w Width) string               // a/aa → "am"/"p. m." etc.
    Ordinal(n int, kind OrdinalKind) string          // en: "7th"
    Plural(n int) PluralCategory                     // shared rule func (see §5)
    Relative(unit Unit, qty int, style Style) string // signed qty; uses Plural

    Pattern(name string) (layout string, ok bool)    // @preset → NITES layout
}
```

The `Context` argument is what makes Slavic/Greek/etc. correct (CLDR's
`format` vs `stand-alone`); it is in the contract from day one so the data and
signatures never need a breaking change to support stand-alone rendering.

A shared `tableLocale` struct implements this from data; most locales are pure
data + a plural-rule id. Behavior-rich locales (future grammar) can implement the
interface directly or via an optional extension interface (§9).

## 3. Public Wiring

`TimeAgoIn` / `FormatIn` live in `gotime`; they adapt `locales.Locale` to the
internal renderer and own the relative-time classification (§6). `TimeAgo` /
`Format` are reimplemented as the English locale through the same path, so there
is exactly one code path and English is just a registered locale.

## 4. Locale Data Pipeline (single source, dual path)

This is the heart of the design. **Slim JSON is the canonical, versioned source
of truth**; GoTime (two paths) and the UExL host are its consumers.

```
        CLDR (full, external)
              │  Stage 0: EXTRACT  (maintainer-only, rare; on CLDR version bump)
              ▼
   locales/data/<code>.json   ← SINGLE SOURCE OF TRUTH (slim, committed, versioned)
              │
      ┌───────┴───────────────┬───────────────────────────┐
      │ 1a: GENERATE          │ 1b: LOAD (runtime, Go)     │ 1c: CONTEXT (UExL host)
      ▼                       ▼                            ▼
 locales/<code>/<code>.go   gotime.RegisterJSON(data)   host parses JSON → map
 (Go literals,              (drop-in / custom Go        and injects as UExL
  compile-checked)           locales)                   context (key/value)
      │                       │                            │
      ▼                       ▼                            ▼
 zero-alloc compiled       parse-once → same           UExL evaluates expressions
 tables                    tableLocale                 (incl. plural rules) on it
```

UExL never touches the filesystem: its host (a Go program, typically using the
same parse step as 1b) supplies the locale values as a plain map. The JSON parses
naturally into `map[string]any`, which is exactly UExL's context shape.

### 4.1 The JSON schema (versioned contract)

```jsonc
{
  "schemaVersion": 1,
  "code": "pl",
  // Names carry CLDR format vs stand-alone context, each short/long (12 / 7 each,
  // weekdays Sunday=0). Languages without a distinction repeat the same arrays.
  "months": {
    "format":      { "short": ["sty", ...], "long": ["stycznia", ...] },   // in-date (genitive)
    "stand-alone": { "short": ["sty", ...], "long": ["styczeń", ...] }      // nominative
  },
  "weekdays": {
    "format":      { "short": ["niedz", ...], "long": ["niedziela", ...] },
    "stand-alone": { "short": ["niedz", ...], "long": ["niedziela", ...] }
  },
  "dayPeriods": { "am": { "short": "AM", "long": "AM" },
                  "pm": { "short": "PM", "long": "PM" } },
  "ordinal":  "none",            // ordinal rule id (see §5)
  "firstDay": 1,                 // optional, forward-looking: 0=Sun … 6=Sat (Phase 1 unused)
  "weekend":  [6, 0],            // optional, forward-looking
  "plural": {
    "id": "west-slavic",         // canonical id → compiled Go func (GoTime)
    "rules": {                   // authored in UExL expression syntax; UExL host
                                 // evaluates these with CLDR operands (n,i,v,w,f,t)
                                 // supplied as context. GoTime ignores them.
      "one":  "i == 1 && v == 0",
      "few":  "v == 0 && (i % 10) >= 2 && (i % 10) <= 4 && !((i % 100) >= 12 && (i % 100) <= 14)",
      "many": "true"
    }
  },
  "relative": {                  // [unit][tense][category] → template with {n}
    "day": { "past": { "one": "{n} dzień temu", "few": "{n} dni temu", "many": "{n} dni temu" },
             "future": { "one": "za {n} dzień", "few": "za {n} dni", "many": "za {n} dni" } }
    // ... other units; plus "specials" for auto-style (now/yesterday/last week)
  },
  "presets": {                   // @preset → NITES layout (CLDR length semantics)
    "datemedium": "d.mm.yyyy", "datelong": "d mmmm yyyy", "time24": "hhhh:ii",
    "datetimeGlue": "{date}, {time}",   // locale date-time combiner for @datetime
    "...": "..."
  }
}
```

- **`schemaVersion`** is checked by both readers; unknown major ⇒ clear error.
- The schema is **additive-only** within a major version; breaking changes bump
  the major and are coordinated with UExL.
- `plural.rules` are authored in **UExL expression syntax** so the UExL host can
  evaluate them by supplying the CLDR operands (`n, i, v, w, f, t`) as context;
  GoTime ignores them and uses `plural.id` (§5). Authoring in UExL syntax (not
  raw CLDR notation like `2..4`) is required so UExL can parse them directly.
  Note the operator vocabulary: UExL uses `&&` / `||` / `!` and `==` — the words
  `and` / `or` / `not` and single `=` are **not** UExL operators (they lex as
  identifiers) and must not appear in rules. The extractor (Stage 0) is
  responsible for translating CLDR rule notation into this valid-UExL form.

### 4.2 Stage 0 — extract (maintainer, rare)

A maintainer tool reads full CLDR (downloaded, **not** committed) and emits the
slim `data/<code>.json` (only the fields above). Output is committed and
reviewable. Bumping CLDR = rerun, review the JSON diff.

### 4.3 Stage 1a — generate Go (GoTime fast path)

`go generate` reads each committed `data/<code>.json` and emits
`locales/<code>/<code>.go`: a `tableLocale` built from Go literals, with the
plural rule wired by **symbol** (compile-checked), registering itself in
`init()`. Generated files are marked `// Code generated … DO NOT EDIT.` and are
never hand-edited. Consumers tree-shake by import:

```go
import _ "github.com/maniartech/gotime/v2/locales/pl" // only "pl" compiled in
```

### 4.4 Stage 1b — load JSON at runtime (GoTime dynamic path)

`locales.RegisterJSON(data)` parses the slim JSON into the same `tableLocale` at
runtime and registers it. This is the escape hatch for user-supplied or
dynamically-loaded **Go** locales. It accepts the parse cost in exchange for
runtime flexibility.

### 4.5 Stage 1c — UExL context (host-injected map)

UExL is an expression engine: it has no file access and consumes data only as a
**context map** (key/value) provided by its host. So a Go generator produces
nothing UExL can use, and UExL does not call `RegisterJSON`. Instead the host
parses the slim JSON (`json.Unmarshal` → `map[string]any`, which is already
UExL's context shape) and injects the relevant locale values. UExL then evaluates
expressions against that context — including the `plural.rules` (authored in UExL
syntax, §4.1) with the CLDR operands supplied as context variables.

### 4.6 Why one source, many consumers, and why this is cheap

JSON is **mandatory** for the UExL host regardless (it is the interchange format
a non-Go engine can consume as a map). Given the JSON exists, GoTime generates Go
*from it* as an optimization and offers `RegisterJSON` for dynamic Go use. You
maintain **one data set with multiple consumers**, not multiple pipelines. The
marginal cost over a JSON-only design is the generator (one `text/template`
program) plus a cross-path conformance test (§8).

## 5. Plural Rules (logic in Go, referenced by id)

Plural selection is logic and lives in core, written once; data references it by
id. Many languages share a rule, so a handful of functions cover dozens of
locales.

```go
// locales/plural.go
func pluralOneOther(n int) PluralCategory { if n == 1 { return One }; return Other }
func pluralFrench(n int) PluralCategory   { if n == 0 || n == 1 { return One }; return Other }
func pluralWestSlavic(n int) PluralCategory {
    if n == 1 { return One }
    m10, m100 := n%10, n%100
    if m10 >= 2 && m10 <= 4 && !(m100 >= 12 && m100 <= 14) { return Few }
    return Many
}
var pluralFuncs = map[string]func(int) PluralCategory{
    "one-other": pluralOneOther, "french": pluralFrench, "west-slavic": pluralWestSlavic,
}
```

- **Generated-Go path:** the generator emits `plural: pluralWestSlavic` (a
  symbol) — a wrong/missing id is a **compile error**.
- **JSON path:** `RegisterJSON` looks up `plural.id` in `pluralFuncs`; an unknown
  id is a validation error, caught by the CI test that loads every locale.
- **UExL path:** the host evaluates the `plural.rules` expressions (UExL syntax)
  by supplying the CLDR operands (`n, i, v, …`) as context — no Go func needed.
- Ordinal rules follow the same id pattern (`"english"`, `"none"`, …).

The generator also emits CLDR plural **test data** as Go tests, proving each
rule against Unicode's expectations.

## 6. Relative-Time Engine (full refactor)

The current `TimeAgo` couples bucket selection to English wording (everything
< 1h → "Few minutes ago"; numeric ladder only covers hours+). It is split:

### 6.1 Locale-neutral classification (package gotime, internal)

```go
func relativeClassify(d time.Duration, base, t time.Time) (unit locales.Unit, qty int)
```

- `qty` signed: negative past, positive future.
- Units span `second → year`. "Yesterday"/"tomorrow" detected as a **calendar**
  condition (a ±1-day unit at a day boundary, reusing existing
  `yesterdayOrTomorrow` logic), expressed as `(UnitDay, ±1)` — the locale decides
  the wording. No English strings here.
- **Quantization is explicit and documented** (spec §3.7): largest unit with
  magnitude ≥ 1; round-half-to-even at that unit; fixed, documented auto-style
  thresholds. The function's doc comment states these so output is predictable
  and stable across releases.

### 6.2 Locale rendering — `loc.Relative(unit, qty, style)`

- **Auto:** consult the locale's `specials` (now, ±1 day, ±1 week/month/year);
  else render numeric.
- **Numeric:** `cat = Plural(abs(qty))`; pick `relative[unit][tense][cat]`;
  splice the number via `strconv.AppendInt` around a single `{n}` marker — no
  `fmt.Sprintf`.

`TimeAgo(t, …)` == `TimeAgoIn(t, en, …)` with `StyleAuto`.

## 7. Formatter Hot Path, Presets & `internal/nites` Layering

`internal/nites/format.go`'s `formatStrs` builds `[]any` + `fmt.Sprint`
(reflection, multi-alloc) and calls `dt.Format` per fragment. The localized
formatter replaces this with a single-pass, buffer-based writer.

```go
// internal/nites — a minimal interface so nites needn't import locales
type Localizer interface {
    Month(m time.Month, long, standAlone bool) string
    Weekday(d time.Weekday, long, standAlone bool) string
    DayPeriod(pm, long bool) string
    Ordinal(n int, monthKind bool) string
}
func FormatLocalized(dt time.Time, layout string, loc Localizer) string
```

- `locales.Locale` is adapted to `nites.Localizer` in `gotime` → **no import
  cycle** (`nites` imports nothing from `locales`).
- **Token splitting:** for the localized path, the locale-sensitive tokens
  `mmm/mmmm/www/wwww` **and `a/aa`** (day periods) are emitted as **fragments**
  (like `dt`/`mt` already are) instead of being baked into the Go layout. The
  split layout is **locale-independent** (tokens, not words), so the existing
  `internal/cache` still works and is shared across locales; only final
  substitution differs per locale. The English `Format` path is untouched (names
  stay baked, one `time.Format` call) — zero regression.
- **`@presets`:** resolved in `gotime`/`nites` *before* layout conversion — a
  layout equal to `@name` is replaced by `loc.Pattern(name)`’s NITES layout, then
  formatted normally. The `@` sigil disambiguates from literal layouts. Resolved
  preset layouts are cacheable (locale-independent once resolved).
- **Allocation budget (contract):** English `Format`/`TimeAgo` unchanged;
  `FormatIn` ≤ 1 alloc beyond the result string for typical layouts (no
  `fmt.Sprint`); `TimeAgoIn` numeric == 1 alloc. Locked by `-benchmem`.

## 8. Locale Registry & Conformance

Registry (package `locales`): copy-on-write map behind `atomic.Pointer` for
**lock-free reads** on the hot path; `Register`/`RegisterJSON` clone-and-store
under a small write mutex. `en` registered in `init()`, never removable.
Resolution truncates BCP-47 subtags right-to-left then falls to `en`
(`zh-Hant-HK → zh-Hant → zh → en`); see spec §3.4. `Available()` returns the
sorted registered codes.

**Robustness:** `RegisterJSON` and the host-side parse are **fuzz-tested**
(malformed/adversarial JSON must error, never panic). The published
`locales/data/schema.json` (JSON Schema) is validated in CI against every
committed locale file.

**Cross-path conformance test (required):** for each locale, load it via the
**generated-Go** subpackage and via **`RegisterJSON`** of the same
`data/<code>.json`, and assert identical output across the token/relative-time
matrix. This guarantees the two Go readers never drift. Because UExL consumes the
**same committed `data/<code>.json`** (as a host-injected context map), agreement
with UExL follows from a single source of truth rather than a duplicated test —
optionally, a golden-output fixture per locale lets the UExL repo assert the same
expected strings. A second test loads every committed `data/*.json` and validates
schema + plural-id resolution (replacing the compile-time safety the JSON path
lacks).

## 9. Compatibility & Extension Notes

- Go's stdlib has no locale-aware names in `time.Format`; localized names come
  from GoTime tables (why §7's token-splitting is required).
- Grammar richness (e.g. Polish genitive month forms) is a future track via an
  **optional extension interface**, added without changing the core `Locale`
  contract — existing locales keep working.

## 10. `internal/nites` Layering Summary

| Concern | Lives in | Notes |
|---|---|---|
| Layout token parsing + cache | `internal/nites` (existing) | locale-independent; shared |
| English-only fast format | `internal/nites` (existing) | untouched, zero regression |
| Localized single-pass format | `internal/nites` (new `FormatLocalized`) | takes `Localizer`; no `locales` import |
| `@preset` resolution | `gotime` (+ pass-through) | `@name` → `Pattern(name)` before convert |
| `Locale`, registry, plural funcs, `RegisterJSON` | `locales/` | core data + runtime loader |
| Generated locale data | `locales/<code>/` | `// Code generated`; import to include |
| Slim JSON source of truth | `locales/data/` | versioned; UExL host injects it as a context map |
| Public `FormatIn`/`TimeAgoIn`, relative classify | `gotime` | adapts `Locale`→`Localizer` |
