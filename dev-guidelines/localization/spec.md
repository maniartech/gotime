# Localization Specification

This document defines **what** locale support means for GoTime: the public API
contract, the locale-data JSON contract, functional requirements, and acceptance
criteria. The *how* lives in [architecture.md](architecture.md).

## 1. Terminology

- **Locale**: A language plus optional region/script variant. Codes follow
  BCP-47 short forms: `en`, `en-US`, `es`, `es-MX`, `pl`. Matching is
  case-insensitive; `_` and `-` separators are both accepted.
- **Locale-aware output**: Output whose text/format varies by language — words,
  pluralization, ordinals, names, and preset field ordering.
- **Locale-neutral output**: Numeric and timezone output, independent of human
  language. Never routed through a locale.
- **CLDR plural category**: One of `zero`, `one`, `two`, `few`, `many`, `other`.
- **Canonical data source**: The slim JSON (extracted from CLDR) that is the
  single versioned source of truth (§4), consumed by GoTime (generated Go or the
  `RegisterJSON` runtime loader) and by UExL (its host parses the JSON and injects
  it as a key/value **context map**; UExL has no file access).

## 2. Locale-Sensitive Areas

1. **Relative time** (`TimeAgo`) — units `second … year`, signed quantity,
   two styles (Auto / Numeric). See architecture §6.
2. **Ordinals** (`dt`, `mt`) — currently hardcoded English `st/nd/rd/th`; becomes
   a locale rule.
3. **Month / weekday names** (`mmm/mmmm/www/wwww`) — from per-locale tables.
   Names have **two grammatical contexts** that the data model must carry from
   day one (CLDR `format` vs `stand-alone`):
   - **format** context — the name as used *inside a full date* (often genitive:
     Spanish "7 de **enero**", Russian "7 **января**"). This is what `mmm/mmmm`
     render in Phase 1.
   - **stand-alone** context — the name used *alone* (nominative: "**Январь**").
     Phase 1 stores it; a stand-alone token (e.g. `LLLL`-equivalent) may be added
     later without re-extracting data.
4. **Day periods** (`a`, `aa`) — am/pm markers are locale-sensitive ("a. m.",
   German "vorm./nachm."). Currently `a/aa` emit English via Go; they become a
   locale lookup.
5. **Named format presets** (`@date`, `@time`, …) — each resolves to a
   **per-locale NITES layout**, so field ordering and separators are
   locale-driven. `@datetime` additionally uses the locale's **date-time glue**
   pattern (CLDR `dateTimeFormats`, e.g. "{date} at {time}"), not naive
   concatenation.

## 3. Requirements

### 3.1 Backward Compatibility

- Existing public functions keep current **English** output byte-for-byte.
- Localization is **opt-in** via new functions; no existing signature changes.
- English fast paths must not regress in allocations or latency.

### 3.2 Public API

```go
// Relative time, localized. baseTime optional (defaults to time.Now()).
func TimeAgoIn(t time.Time, loc locales.Locale, baseTime ...time.Time) string

// NITES formatting, localized (dt, mt, mmm, mmmm, www, wwww, and @presets).
func FormatIn(t time.Time, layout string, loc locales.Locale) string
```

`locales` package:

```go
func Get(code string) (Locale, bool) // resolve w/ fallback; ok=false ⇒ fell back to en
func MustGet(code string) Locale      // resolve w/ fallback; never nil
func Default() Locale
func SetDefault(code string) error
func Available() []string             // sorted codes of registered locales

// Runtime JSON path (drop-in / custom Go locales; UExL host parses JSON itself):
func RegisterJSON(data []byte) error  // parse slim JSON, validate, register
func Register(l Locale)               // register a pre-built (e.g. generated) locale
```

### 3.3 Named Format Presets (`@` sigil)

- A layout consisting of a registered preset name prefixed with `@`
  (e.g. `"@longdate"`) is resolved to that **locale's** NITES layout before
  formatting. The `@` distinguishes a preset from a literal layout.
- The canonical preset vocabulary is **normative** (callers rely on the same
  names across locales) and mirrors CLDR's four standard *lengths* so the
  semantics are principled, not ad-hoc. Phase-1 set:
  - Dates: `@datefull` (CLDR *full*), `@datelong` (*long*), `@date`/`@datemedium`
    (*medium*, the default), `@dateshort` (*short*).
    `@longdate`/`@shortdate` are accepted aliases for `@datelong`/`@dateshort`.
  - Times: `@time` (*medium*), `@timeshort`, `@time24` (24-hour short).
  - Combined: `@datetime`, `@datetime24` — composed via the locale's date-time
    **glue** pattern (§2.5), not string concatenation.
  - Components: `@monthyear`, `@month`, `@weekday`.
- A locale that does not define a preset falls back per §3.4, then to a neutral
  ISO form. A preset name is the *entire* layout argument, not embedded in a
  larger layout.
- `Locale.Pattern(name) (layout string, ok bool)` exposes resolution.

### 3.4 Locale Resolution & Fallback

- `en` is always registered and is the ultimate fallback; it cannot be removed.
- Resolution is deterministic and **identical for every locale-aware API**. It
  truncates BCP-47 subtags **right-to-left**, then falls to `en` — so script
  variants resolve correctly, not just region:
  - `es-MX → es → en`
  - `zh-Hant-HK → zh-Hant → zh → en`
  - `pt-BR → pt → en`
- `Get(code) (Locale, bool)` — `ok=true` when any truncation matched a
  registered locale, `false` when it fell through to `en`.
- `MustGet(code) Locale` — never errors (worst case `en`).
- Codes are normalized (case-insensitive, `_`→`-`) before resolution.
- No locale-aware API may panic on an unknown or malformed code.

### 3.5 Concurrency & Robustness

- Registry allows **lock-free reads** on the hot path (copy-on-write behind an
  atomic pointer).
- `Register`/`RegisterJSON` are safe concurrently and from `init()`; last
  registration for a code wins deterministically.
- Locales are immutable after registration (read-only tables), safe to share
  across goroutines.

### 3.6 Plural Rules

- A locale references its plural rule by a **canonical id** (e.g.
  `"one-other"`, `"west-slavic"`).
- **Generated-Go path:** the id maps to a compile-checked Go function in core.
- **`RegisterJSON` path:** the runtime loader maps the id to the same Go function.
- **UExL path:** the JSON additionally carries the rule's **condition
  expressions** authored in **UExL syntax**; the UExL host evaluates them with the
  CLDR operands (`n, i, v, w, f, t`) supplied as context — no Go function needed.
- Phrase tables must define an entry for **every** category the rule can return.

### 3.7 Relative-Time Semantics (deterministic & documented)

To match best-in-class libraries, the engine's quantization rules are explicit,
not incidental:

- **Unit selection** picks the largest unit whose magnitude is ≥ 1.
- **Rounding** of the quantity is **round-half-to-even at the chosen unit**
  (documented; the legacy `TimeAgo` used `math.Round`). The chosen rule must be
  stated in the doc comment so output is predictable.
- **Auto-style thresholds** (when `±1 day` reads as "yesterday/tomorrow", etc.)
  are **fixed and documented** in Phase 1, computed as a calendar condition (not
  a fixed-seconds bucket). A future `TimeAgoOptions{Thresholds, Style}` may make
  them tunable without changing the Phase-1 default behavior.
- **Sign/tense** derive from the signed quantity; `0` maps to the locale's
  "now" phrase in Auto style.

## 4. The Locale-Data JSON Contract

The slim JSON is a **versioned public contract** (consumed by GoTime, UExL, and
potentially third parties). See architecture §4 for the full schema; normative
requirements:

- A top-level `schemaVersion` (integer). Consumers reject unknown major
  versions with a clear error.
- **Required fields:**
  - `code`
  - month names: `format` and `stand-alone`, each `short`/`long` (§2.3)
  - weekday names: `format` and `stand-alone`, each `short`/`long`
  - day-period names (`am`/`pm`, short + long) (§2.4)
  - `ordinal` rule id
  - `plural`: rule `id` + UExL-syntax condition expressions
  - `relative` phrase tables (by unit/tense/category) + auto-style `specials`
  - `presets` (incl. the date-time **glue** pattern for `@datetime`)
- **Optional fields (forward-looking, populated from CLDR now even if unused in
  Phase 1):** `firstDay` (0=Sunday…6=Saturday) and `weekend` days. These let
  week-oriented APIs become locale-aware later without re-extraction; Phase 1
  does **not** change week-calculation behavior (see §5).
- The schema is **additive-only** within a major version; breaking changes bump
  the major version and are coordinated with UExL.
- A **machine-readable JSON Schema** (`locales/data/schema.json`) is published
  and validated in CI, so third parties (and UExL) can validate independently.

## 5. Out of Scope (Phase 1) — explicit decisions

Stated deliberately so the boundary is clear (a best-in-class spec declares its
non-goals as decisions, not omissions):

- **Parsing** localized names (Phase 3; tables are defined so a reverse lookup is
  derivable — see implementer guide).
- **Localized numerics** (non-Latin digits, decimal/grouping separators). Numbers
  render with ASCII digits; revisit in a later, tightly-scoped track.
- **Localized timezone names** (`zz` stays Go's `MST`-style abbreviation). CLDR
  timezone display names are a large surface and excluded for now.
- **Bidi/RTL shaping.** Output is the logical string; GoTime does **not** insert
  Unicode bidi isolates. Callers rendering mixed LTR/RTL content are responsible
  for bidi wrapping (as with most formatting libraries). Documented so it is a
  known boundary, not a surprise.
- **Week-calculation localization** (first-day-of-week / weekend). The data is
  carried (§4) but week APIs remain neutral in Phase 1.
- **Era and narrow-width names**, transliteration, alternate calendar/number
  systems. (NITES has no era/narrow tokens today; adding them is a separate NITES
  change.)

## 6. Acceptance Criteria

Phase 1 is "done" when **all** hold:

1. **Compat:** existing APIs produce identical English output; English-path
   benchmarks show no regression.
2. **Localization:** `TimeAgoIn`, `FormatIn`, and `@presets` change output
   correctly for a non-English locale across **all** locale-sensitive tokens
   (`mmm/mmmm/www/wwww`, `dt/mt`, `a/aa`) and the relative-time unit range, in
   both tenses and both styles.
3. **Plural correctness:** a locale with non-trivial rules (Polish) selects the
   right category for representative counts (1, 2, 5, 22, 0), proven against
   CLDR plural test data.
4. **Fallback:** subtag truncation works — `es-MX → es`, `zh-Hant-HK → zh-Hant`
   when registered, `→ en` when not — identical through `TimeAgoIn` and
   `FormatIn`.
4a. **Grammatical context:** a language with distinct format vs stand-alone names
   (e.g. Russian) renders the format (genitive) name inside a full date.
5. **Dual-path conformance:** for a given locale, the **generated-Go** and
   **runtime-JSON** paths produce **identical** output (cross-path test).
6. **Schema contract:** the JSON validates against `schemaVersion`; an unknown
   major version is rejected with a clear error.
7. **Isolation:** adding a locale required no edits to core formatting logic.
8. **Allocation budget:** the localized `FormatIn` path meets its documented
   budget under `-benchmem`; generated-Go locales add no startup parse cost.

## 7. API Contract Examples

### 7.1 Default behavior (backward compatible)

```go
fmt.Println(gotime.TimeAgo(time.Now().Add(-2 * time.Hour)))                 // "2 hours ago"
fmt.Println(gotime.Format(time.Date(2025,7,7,0,0,0,0,time.UTC), "wwww, mmmm dt")) // "Monday, July 7th"
```

### 7.2 Locale-aware relative time

```go
t := time.Now().Add(-5 * time.Minute)
fmt.Println(gotime.TimeAgoIn(t, locales.MustGet("es"))) // "hace 5 minutos"
fmt.Println(gotime.TimeAgoIn(t, locales.MustGet("pl"))) // "5 minut temu" (category many)

loc := locales.MustGet("en")
fmt.Println(loc.Relative(locales.UnitDay, -1, locales.StyleAuto))    // "yesterday"
fmt.Println(loc.Relative(locales.UnitDay, -1, locales.StyleNumeric)) // "1 day ago"
```

> The unit-based refactor makes `"5 minutes ago"` the real English output too;
> the pre-localization code collapsed all sub-hour values to "Few minutes ago".

### 7.3 Locale-aware formatting + presets

```go
dt := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)

fmt.Println(gotime.FormatIn(dt, "wwww, mmmm dt", locales.MustGet("en"))) // "Monday, July 7th"
fmt.Println(gotime.FormatIn(dt, "@longdate", locales.MustGet("en")))     // "July 7, 2025"
fmt.Println(gotime.FormatIn(dt, "@longdate", locales.MustGet("es")))     // "7 de julio de 2025"
fmt.Println(gotime.FormatIn(dt, "@shortdate", locales.MustGet("en")))    // "7/7/2025"
fmt.Println(gotime.FormatIn(dt, "@shortdate", locales.MustGet("de")))    // "07.07.2025"
```

Field order and separators differ per locale because each locale defines its own
layout for `@longdate` / `@shortdate`. Numeric tokens still render identically.

### 7.4 Runtime JSON — dynamic / custom Go locale

```go
//go:embed mylocale.json
var raw []byte
if err := locales.RegisterJSON(raw); err != nil { /* malformed/unknown schema */ }
loc := locales.MustGet("xx-custom")
_ = gotime.FormatIn(dt, "@date", loc)
```

### 7.5 UExL — host injects the same data as a context map

UExL does not read files or call `RegisterJSON`. The host parses the same slim
JSON and supplies it as UExL's evaluation context (key/value); UExL evaluates
expressions — including the `plural.rules` (UExL syntax) with operands as
context:

```go
// host side (Go), pseudo-illustrative
var data map[string]any
_ = json.Unmarshal(raw, &data)        // same data/<code>.json
ctx := map[string]any{"locale": data, "n": 5, "i": 5, "v": 0}
// UExL evaluates e.g. data.plural.rules.many against ctx → selects category
```

### 7.6 Unknown-locale handling (deterministic)

```go
loc, ok := locales.Get("zz")          // ok == false; loc == en
_ = gotime.FormatIn(dt, "mmmm", loc)  // never panics; English on fallback
```
