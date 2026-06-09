# Localization Roadmap

Gradual, maintainable delivery by a solo maintainer. Each phase is independently
shippable and backward compatible. The locale **data** is a single canonical
slim-JSON source with multiple consumers (generated Go and `RegisterJSON` for
GoTime; a host-injected context map for UExL) — see
[architecture.md §4](architecture.md).

## Phase 0: Documentation (done)

- Scope, goals, non-goals, key decisions ([README](README.md)).
- Public + JSON-schema contracts, dual-path acceptance criteria ([spec](spec.md)).
- Interfaces, registry, single-source data pipeline, dual consumption paths,
  formatter hot path, relative-time refactor, `@`-presets, `internal/nites`
  layering ([architecture](architecture.md)).
- Implementer guide: author slim JSON, plural/ordinal by id, presets, TinyGo
  build-gate, required tests incl. dual-path conformance ([implementing-locales](implementing-locales.md)).

## Phase 1: Opt-in localized output (the core)

The vertical slice that makes the design real.

1. **`locales` package skeleton**
   - `Locale` interface + enums; `tableLocale` data-driven implementation.
   - Lock-free COW registry (`Register`/`RegisterJSON`/`Get`/`MustGet`/`Default`/
     `SetDefault`, `Available`), BCP-47 subtag-truncation fallback to `en`.
   - Shared plural functions + `pluralFuncs` id map; ordinal rules by id.
   - Built-in `en` as **Go literals** (TinyGo-safe, no parse), reproducing
     today's English output exactly.

2. **JSON contract + readers**
   - Freeze `schemaVersion: 1` and the slim-JSON schema.
   - `RegisterJSON` (runtime loader) → `tableLocale`, with schema + id validation.
   - Cross-path conformance scaffolding (generated vs JSON).

3. **Relative-time engine refactor**
   - `relativeClassify(d, base, t) (Unit, qty)` — neutral, signed, `second→year`,
     calendar-aware yesterday/tomorrow.
   - `loc.Relative(unit, qty, style)` (Auto + Numeric).
   - Reimplement `TimeAgo` as `TimeAgoIn(t, en, …)`; add public `TimeAgoIn`.

4. **Localized formatter + presets**
   - Split `mmm/mmmm/www/wwww` into fragments on the localized path; English path
     untouched.
   - `internal/nites.FormatLocalized` — single-pass, buffer-based, no
     `fmt.Sprint`.
   - `@preset` resolution (`@name` → `Pattern(name)`); replace hardcoded ordinal
     branch with `Ordinal(n, kind)`; add public `FormatIn`.

5. **Tests + benchmarks**
   - English output byte-identical (regression tests).
   - Allocation gates (architecture §7, `-benchmem`).
   - Plural/fallback/style/preset coverage; **dual-path conformance**; schema
     validation.

**Exit:** all spec §6 acceptance criteria pass.

## Phase 2: Data pipeline + reference locales

- **Stage-0 extractor:** maintainer tool, full CLDR → slim `data/<code>.json`
  (pinned CLDR version; vendored, output committed and reviewable).
- **Stage-1a generator (`go generate`):** slim JSON → `locales/<code>/<code>.go`
  (Go literals, plural wired by symbol, `// Code generated`); emits CLDR plural
  conformance tests.
- Add `es` (simple plural) and `pl` (CLDR `one/few/many`) as reference locales,
  each one slim JSON + generated subpackage + tests.
- Contributor docs/examples for adding locales and for consuming the JSON as a
  UExL context map (host-side).

## Phase 3 (optional): Localized parsing

- Build reverse name lookups from existing tables (enabled by the
  parsing-readiness rules, implementer guide §6).
- Locale-aware parsing of `mmm/mmmm/www/wwww`; confirm fallback parity with
  formatting.

## Phase 4 (optional): Grammar richness & numerics

- Context-aware month/weekday forms (e.g. Polish genitive) via an **optional
  extension interface**, without changing the core `Locale` contract.
- Tightly-scoped localized **numerics** (decimal/grouping separators, optional
  non-Latin digits) — only if demand justifies the allocation-budget work; keep
  the Latin-digit fast path unchanged for the common case.
