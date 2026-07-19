# Localization Implementation Plan

Implementation-ready, staged, idempotent plan to build the localization design
([spec](spec.md), [architecture](architecture.md)) to A+ / military-grade
quality. Every stage compiles, is race-clean, and leaves the test suite green on
its own; stages are ordered so each builds only on committed, verified work.
There are **no grey areas**: ambiguous points are resolved in §2 and inside each
stage. This plan does not change any runtime behavior until a stage explicitly
says so, and every behavioral change is gated by a golden test.

Module: `github.com/maniartech/gotime/v2`. Root package: `gotime`. New runtime
package: `locales`. Internal: `internal/nites`, `internal/cache`,
`internal/utils`. Maintainer tooling: `internal/localegen` (extractor +
generator, never imported by the library).

---

## 1. Global Invariants (hold after *every* stage)

These are non-negotiable acceptance gates checked at the end of each stage:

- **G1 — Byte-exact backward compatibility.** For all inputs:
  - `FormatIn(t, layout, locales.MustGet("en")) == Format(t, layout)`
  - `TimeAgoIn(t, locales.MustGet("en"), base...) == TimeAgo(t, base...)`
  Existing public signatures are unchanged. The single intended behavior change
  (sub-hour relative time, §2.6) is applied to **both** `TimeAgo` and the `en`
  locale simultaneously, so the equality above still holds; it is gated by an
  explicitly updated golden in Stage 3.
- **G2 — Compiles & vets clean.** `go build ./...` and `go vet ./...` succeed.
- **G3 — Race-clean & green.** `go test -race ./...` passes.
- **G4 — No allocation regression.** English-path benchmarks (`Format`,
  `TimeAgo`) show no new allocations vs. the Stage-0 baseline; localized-path
  benchmarks meet the budgets in §2.7.
- **G5 — Deterministic & idempotent codegen** (from Stage 6): running
  `go generate ./...` twice produces an empty `git diff`.
- **G6 — TinyGo-safe core.** The default build and the `en` locale never import
  `encoding/json`; `GOOS=js`/TinyGo-relevant packages stay reflection-free.
- **G7 — No panics on bad input.** Every public entry point returns
  deterministically (string or error); fuzz targets never panic.

Idempotency definition for this plan: re-running a stage's build/test/generate
steps on its committed output yields identical artifacts and identical pass/fail
results (no hidden state, no time/locale/network dependence; all test times use
explicit `time.UTC` + fixed base instants).

---

## 2. Resolved Decisions (the "no grey area" section)

Every micro-decision an implementer would otherwise have to guess:

**2.1 `@`-preset detection & escaping.**
- A layout is a preset iff its **first byte is `@`**. The preset name is the
  remainder, lowercased. Example: `"@longdate"` → name `longdate`.
- A preset name is the **entire** layout argument; presets are not embedded in a
  larger layout in Phase 1.
- To emit a literal leading `@`, escape it NITES-style: `\@longdate` formats the
  literal text `@longdate` (the existing `\` escape in `convertLayout` handles
  this; preset detection runs only on an unescaped leading `@`).
- **Unknown preset:** resolution applies the locale fallback chain; `en` is
  **required by test to define every canonical preset** (§2.5), so any canonical
  name always resolves. A **non-canonical** `@xyz` (not in the canonical set and
  undefined by `en`) is emitted **verbatim** as the literal string `@xyz`
  (documented, never panics).

**2.2 Canonical preset vocabulary (frozen).**
`datefull, datelong, datemedium (== date), dateshort, time (== timemedium),
timeshort, time24, datetime, datetime24, monthyear, month, weekday`.
Aliases accepted: `longdate→datelong`, `shortdate→dateshort`.
`@datetime`/`@datetime24` are composed from the date+time presets via the
locale's `datetimeGlue` (`{date}`/`{time}` placeholders).

**2.3 Day-period (`a`/`aa`) mapping (byte-exact for en).**
NITES `a`/`aa` differ by **case**, not width (NITES has no wide day-period
token). Resolve by refining the architecture's `DayPeriod` to:
`DayPeriod(pm bool, upper bool) string`, where token `a` → `upper=false`, `aa` →
`upper=true`. The JSON stores both forms explicitly:
`dayPeriods.am = {lower, upper}` (en: `{"am","AM"}`, `{"pm","PM"}`). en therefore
reproduces Go's `pm`/`PM` exactly. (This refines architecture §2/§7; the stored
pair removes any casing ambiguity for locales like `"a. m."`.)

**2.4 Name grammatical context.**
`Month(m, w, c)` / `Weekday(d, w, c)` take `Context ∈ {Format, StandAlone}`.
Phase-1 formatting of `mmm/mmmm/www/wwww` passes **`Format`**. `StandAlone` is
stored and reachable via the API but no NITES token selects it yet (reserved).
Locales without a distinction store identical arrays for both.

**2.5 Fallback (frozen algorithm).**
`normalize(code)`: trim, lowercase, `_`→`-`, collapse repeats. Resolution:
repeatedly strip the last `-`-separated subtag until a registered code matches;
if none, return `en`. `Get` returns `ok=false` exactly when it returned `en`
without a non-`en` match. `en` is registered in `locales` package `init()` and
cannot be removed.

**2.6 Relative-time algorithm (frozen, deterministic).**
Given signed `Δ = base.Sub(t)` (note `TimeAgo` treats *past* as positive elapsed;
internal `qty` is signed with **negative = past**, positive = future):
1. **Calendar specials first** (Auto style only): compute calendar-day delta in
   `t`'s location. If 0 → `now`/today special; if −1 → "yesterday"; if +1 →
   "tomorrow" (reusing existing `yesterdayOrTomorrow` semantics, generalized).
2. Otherwise classify by duration: for units `year→second`, compute
   `q = roundHalfEven(|Δ| / unit)`; choose the **largest** unit with `q ≥ 1`.
   `qty = sign(Δ_future?+:−) * q`. Sub-1-second → `(UnitSecond, 0)`.
3. Render via `loc.Relative(unit, qty, style)`.
- **Rounding is round-half-to-even**, documented in the doc comment. This is the
  one intended change from the legacy `math.Round` ladder and the only G1 golden
  that is deliberately updated (sub-hour values now read e.g. "5 minutes ago"
  instead of "Few minutes ago", in both `TimeAgo` and `en`).

**2.7 Allocation budgets (frozen, `-benchmem`).**
- `Format`, `TimeAgo` (English): **0 new allocs** vs. Stage-0 baseline.
- `FormatIn` typical layout: **≤ 1 alloc** beyond the returned string.
- `TimeAgoIn` numeric: **= 1 alloc** (the returned string).
- Hot path uses a pooled `[]byte`/`strings.Builder`; **no `fmt.Sprint`/`fmt.Sprintf`**.

**2.8 `RegisterJSON` validation (frozen error set).** Returns a non-nil error
(never panics) on: unknown/zero `schemaVersion` major; missing required field;
month array length ≠ 12 or weekday ≠ 7 (any context/width); unknown `plural.id`;
unknown `ordinal` id; missing relative template for a category the plural rule
can emit; missing canonical preset. Errors are wrapped and identify the locale
`code` and offending field. Registration is last-wins and idempotent per code.

---

## 3. Stages

Each stage: **Goal → Files → Public surface → Tests/Gates → Idempotency → DoD**.
Each ends with a single commit. Recommended commit prefix in parentheses.

### Stage 0 — Baseline guardrails (no behavior change) `(test)`

- **Goal:** lock current behavior so later refactors are provably non-regressing.
- **Files:** `time_ago_golden_test.go`, `format_golden_test.go`,
  `relative_bench_test.go`, `format_bench_test.go` (root package).
- **Public surface:** none.
- **Tests/Gates:** golden tables capturing current `TimeAgo` and `Format` outputs
  across a fixed matrix (fixed UTC base instant; all unit ranges; all NITES name
  tokens incl. `a/aa`, `dt/mt`). Benchmarks with `b.ReportAllocs()` recording the
  English-path baseline alloc counts (documented in the test as the G4 baseline).
- **Idempotency:** pure reads; deterministic (fixed times).
- **DoD:** G2, G3 pass; goldens committed; baseline alloc numbers recorded.

### Stage 1 — `locales` core types + registry + `en` `(feat)`

- **Goal:** the package skeleton and a working `en` locale, not yet wired into
  `gotime`.
- **Files:** `locales/locale.go` (interface + enums: `Width, Context,
  OrdinalKind, PluralCategory, Unit, Style`), `locales/registry.go`
  (`atomic.Pointer[map[string]Locale]` COW; `Register/Get/MustGet/Default/
  SetDefault/Available`; `normalize`, subtag-truncation fallback),
  `locales/errors.go`, `locales/en.go` (Go-literal `en`, no JSON),
  `locales/tablelocale.go` (the data-backed implementation `en` uses).
- **Public surface:** the registry funcs + `Locale` interface (architecture §2,
  refined per §2.3 here).
- **Tests/Gates:** fallback matrix (`es-MX→es`, `zh-Hant-HK→zh-Hant→zh→en`, `→en`
  for unknown); `Available()` sorted & stable; **concurrency test** (`-race`:
  N goroutines `Get` while M `Register`); `en` table self-consistency (12/7
  lengths, every preset present).
- **Idempotency:** `Register` last-wins, re-registering `en` is a no-op-equivalent
  (same result); registry reads are pure.
- **DoD:** G1 (vacuous — not wired yet), G2, G3, G6 (no `encoding/json`).

### Stage 2 — Plural subsystem `(feat)`

- **Goal:** shared plural logic, id-addressed.
- **Files:** `locales/plural.go` (`pluralOneOther, pluralFrench,
  pluralWestSlavic, …`; `pluralFuncs map[string]func(int)PluralCategory`;
  `lookupPlural(id) (fn, ok)`), `locales/ordinal.go` (ordinal rule ids:
  `english`, `none`; `lookupOrdinal`).
- **Public surface:** none new (ids are data); funcs are package-internal.
- **Tests/Gates:** unit tests per rule across category-boundary counts; a
  **CLDR-plural-conformance harness** (`plural_conformance_test.go`) that will be
  fed CLDR test data in Stage 6 — wired now with a couple of hand cases so the
  harness itself is tested.
- **Idempotency:** pure functions; deterministic.
- **DoD:** G2, G3.

### Stage 3 — Relative-time engine + `TimeAgoIn` `(refactor/feat)`

- **Goal:** unit-based engine; reimplement `TimeAgo` on it; add `TimeAgoIn`.
- **Files:** `relative.go` (new: `relativeClassify(d, base, t) (Unit, qty)` per
  §2.6; `roundHalfEven`), `time_ago.go` (rewire `TimeAgo` →
  `TimeAgoIn(t, en, base...)` with `StyleAuto`), `time_ago_in.go`
  (`TimeAgoIn`), `locales/tablelocale.go` (`Relative(unit, qty, style)` impl,
  `{n}` splice via `strconv.AppendInt`), `locales/en.go` (relative phrase
  tables + specials).
- **Public surface:** `func TimeAgoIn(t time.Time, loc locales.Locale, baseTime ...time.Time) string`.
- **Tests/Gates:** update the Stage-0 relative golden for the **documented**
  sub-hour change and assert G1 (`TimeAgoIn(·,en) == TimeAgo(·)`); past/future ×
  Auto/Numeric; rounding edge cases (90s, 1.5h, half-even ties); alloc bench
  (numeric == 1 alloc).
- **Idempotency:** deterministic (fixed base instants).
- **DoD:** G1, G2, G3, G4; golden diff reviewed and intentional.

### Stage 4 — Localized formatter + `@`-presets + `FormatIn` `(feat)`

- **Goal:** localized formatting without regressing the English path.
- **Files:** `internal/nites/format_localized.go` (`Localizer` interface per
  §2.3; `FormatLocalized(dt, layout, loc)` single-pass buffer writer; token
  splitting for `mmm/mmmm/www/wwww/a/aa` + existing `dt/mt`), minimal edits to
  `internal/nites/convert.go` (emit those tokens as fragments **only** for the
  localized path; English path untouched), `format_in.go`
  (`FormatIn`; adapter `locales.Locale → nites.Localizer`; `@`-preset resolution
  per §2.1; `datetimeGlue` composition), `locales/tablelocale.go`
  (`Pattern`, `Ordinal`, `DayPeriod`).
- **Public surface:** `func FormatIn(t time.Time, layout string, loc locales.Locale) string`.
- **Tests/Gates:** **G1 keystone** — `FormatIn(·,·,en) == Format(·,·)` over the
  Stage-0 format matrix (covers names, `dt/mt`, `a/aa`, presets); preset
  resolution incl. unknown→verbatim and `\@` escape; alloc bench (≤1 alloc); the
  English `Format` path bench unchanged (G4).
- **Idempotency:** layout cache keyed by layout string (locale-independent);
  deterministic.
- **DoD:** G1, G2, G3, G4, G6.

### Stage 5 — JSON schema + `RegisterJSON` + machine schema `(feat)`

- **Goal:** the runtime JSON path and the published contract.
- **Files:** `locales/json.go` (`jsonLocale` structs; `RegisterJSON(data)
  (error)`; full §2.8 validation; build a `tableLocale`), `locales/data/schema.json`
  (machine-readable JSON Schema), `locales/json_test.go`,
  `locales/fuzz_test.go` (`FuzzRegisterJSON` — never panics; corpus of
  truncated/garbage/oversized inputs), `locales/schema_test.go` (validate any
  committed `data/*.json` against `schema.json`).
- **Public surface:** `func RegisterJSON(data []byte) error`,
  `func Register(l Locale)`.
- **Tests/Gates:** round-trip a synthetic test locale; every §2.8 error path
  asserted; fuzz target builds & runs in CI (`-run=^$ -fuzz=Fuzz... -fuzztime`).
  This file is the **only** place `encoding/json` is imported (build-gated off
  TinyGo); G6 verified by a test/CI grep that core + `en` don't import it.
- **Idempotency:** `RegisterJSON` last-wins per code; parsing is pure.
- **DoD:** G2, G3, G7 (fuzz), G6.

### Stage 6 — Data pipeline + reference locales (`es`, `pl`) `(feat/chore)`

- **Goal:** authoritative data + generated Go locales, proven equivalent to JSON.
- **Files:**
  - `internal/localegen/extract.go` — Stage-0 extractor: pinned, **vendored**
    CLDR (committed under `third_party/cldr/`, version recorded) → slim
    `locales/data/<code>.json`. Deterministic output: sorted keys, stable
    formatting.
  - `internal/localegen/generate.go` — slim JSON → `locales/<code>/<code>.go`
    (Go literals; plural wired by **symbol**; `// Code generated … DO NOT EDIT.`;
    `init(){ locales.Register(...) }`; `//go:build !tinygo` if JSON-derived but
    note generated Go itself is literal/TinyGo-safe — gate only if size-driven),
    formatted via `go/format`.
  - `//go:generate` directive (`locales/gen.go`), invoked via `go run` resolving
    to the module version.
  - `locales/data/es.json`, `locales/data/pl.json`; generated `locales/es/`,
    `locales/pl/`.
  - CLDR plural test data → fed into the Stage-2 conformance harness.
- **Public surface:** `import _ ".../locales/es"` (and `pl`).
- **Tests/Gates:** **cross-path conformance** — for `es` and `pl`, load via the
  generated subpackage and via `RegisterJSON(data/<code>.json)`; assert identical
  output across the full matrix; `pl` plural conformance vs CLDR data; Russian-
  style format-vs-stand-alone covered (add `ru` if data permits, else assert the
  mechanism with a fixture).
- **Idempotency (critical here):** **G5** — `go generate ./...` then `git diff`
  must be empty. Generation is a pure function of committed slim JSON; the
  extractor is a pure function of vendored CLDR. No network at generate time.
- **DoD:** G1–G6 all green; G5 enforced in CI.

### Stage 7 — Docs, examples, CI wiring `(docs/chore)`

- **Goal:** ship-quality polish.
- **Files:** package doc comments with runnable `Example*` tests (godoc); update
  `docs/` localization pages; CI steps: `go generate` idempotency check (G5),
  fuzz smoke, `-benchmem` budget assertions (G4), schema validation.
- **DoD:** examples run as tests; CI gates G1–G7.

---

## 4. Stage Dependency & Ordering

```
0 ─▶ 1 ─▶ 2 ─▶ 3 ─▶ 4 ─▶ 5 ─▶ 6 ─▶ 7
        └─────────┘   (3 and 4 both depend on 1+2; 4 also needs 3's en tables)
```
Strictly linear for safety; 3 and 4 are separable but 4's `en` reuses 3's
relative tables, so keep the order. No stage depends on a later stage.

## 5. Verification Commands (run at each stage's DoD)

```
go build ./...                              # G2
go vet ./...                                # G2
go test -race ./...                         # G1, G3
go test -run XXX -bench . -benchmem ./...   # G4 (compare to Stage-0 baseline)
go generate ./...  &&  git diff --exit-code # G5 (Stage 6+)
go test -run=^$ -fuzz=FuzzRegisterJSON -fuzztime=30s ./locales  # G7 (Stage 5+)
```

## 6. Rollback & Safety

- Each stage is one commit; revert is a single `git revert`.
- No stage mutates existing public signatures; `*In` APIs are purely additive, so
  a partial rollout (Stages 0–4 without 6's data) still ships a correct,
  English-only library.
- The English path is never routed through new code without the G1 keystone test
  proving byte-identical output, so a regression cannot reach users silently.
```
