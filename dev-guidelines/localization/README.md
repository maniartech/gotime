# Localization (i18n/l10n) Guidelines

This directory specifies how GoTime adds **locale-aware human-language output**
without compromising its core promises: backward compatibility, low/zero
allocation on the hot path, and high throughput. The design target is parity
with professional datetime libraries (Intl.RelativeTimeFormat, Luxon, java.time,
.NET, CLDR), adapted to idiomatic, allocation-conscious Go.

A defining constraint: the locale **data** is a shared asset. It is consumed not
only by GoTime (Go) but also by **UExL**, our expression-engine DSL. UExL has no
file access — its host loads the data and supplies it as a key/value **context
map**, against which UExL evaluates expressions. The architecture therefore
treats locale data as a **single canonical source (slim JSON) with multiple
consumers** — generated Go for GoTime's fast path, a runtime `RegisterJSON`
loader for dynamic Go use, and a host-injected context map for UExL.

## Scope

GoTime is *mostly* locale-neutral. A small, well-defined set of features produce
human-language output and are therefore locale-sensitive:

- **Relative time** text (`TimeAgo`) — "just now", "5 minutes ago", "yesterday".
- **Ordinal suffixes** for the `dt` / `mt` NITES tokens — `1st`, `2nd`.
- **Month / weekday names** for `mmm`, `mmmm`, `www`, `wwww`, including CLDR
  **format vs stand-alone** grammatical context (genitive vs nominative).
- **Day periods** (`a`, `aa`) — localized am/pm markers.
- **Named format presets** (`@date`, `@time`, …) whose *layout* (field order,
  separators) is locale-defined — this is how GoTime does locale-driven ordering
  without an ICU skeleton engine; `@datetime` uses the locale's date-time glue.

Everything else (numeric fields, timezone offsets, arithmetic) stays
locale-neutral and is never routed through the locale layer.

## Goals

- **Backward compatible.** Existing APIs keep their current English output,
  byte-for-byte. Localization is opt-in.
- **Best-in-class behavior.** CLDR plural categories, a unit-based relative-time
  engine, locale-driven format presets, deterministic fallback.
- **One data set, many consumers.** A single versioned JSON schema feeds GoTime
  (via generated Go or `RegisterJSON`) and UExL (via a host-injected context
  map). No divergent data.
- **Cheap to add a locale.** A new locale is one slim JSON file (extracted from
  CLDR). GoTime's generated Go and UExL's host-supplied context both derive from
  it.
- **Allocation-conscious & fast.** The formatting hot path is single-pass and
  buffer-based; name/ordinal lookups are table reads. Generated Go locales have
  zero startup parse cost. English fast paths are not regressed.
- **Robust & concurrent-safe.** Lock-free reads on the hot path; identical
  fallback across every locale-aware API; compile-checked data on the generated
  path; CI-validated data on the JSON path.
- **No heavy runtime deps.** No CLDR runtime dependency; we adopt the CLDR
  *model* (plural categories, fallback) via slim hand-extracted data.

## Non-Goals (initially)

- Locale-aware **parsing** of month/weekday names (Phase 3; the data model is
  designed so parsing is additive).
- Automatic locale detection from OS/environment.
- Localized **numerics** (digits, grouping). Numerics stay locale-neutral; see
  the roadmap for a tightly-scoped future track.
- Localized **timezone display names** (`zz` keeps Go's abbreviation).
- **Bidi/RTL** isolate insertion — output is the logical string; callers handle
  bidi wrapping (see [spec §5](spec.md)).
- Week-calculation localization (first-day-of-week/weekend) — data is carried but
  week APIs stay neutral in Phase 1.
- Full CLDR feature surface (calendar systems, number systems, transliteration,
  era/narrow names).

## Key Design Decisions (normative)

These are settled; the rest of the docs elaborate them.

1. **Single canonical data source.** Slim JSON (extracted from CLDR) is the
   versioned source of truth, shipped as a first-class artifact. See
   [architecture.md §4](architecture.md).
2. **One source, many consumers.** GoTime ships **generated Go** locales
   (zero-alloc, compile-checked, TinyGo-clean) *and* a **runtime JSON loader**
   (`RegisterJSON`) for dynamic/custom Go locales; **UExL** consumes the same
   JSON as a host-injected **context map** (it does not read files). All derive
   from the same JSON; a cross-path conformance test proves the Go readers agree.
3. **One API style.** Locale is passed as a value: `TimeAgoIn`, `FormatIn`.
   A package-level default (`locales.SetDefault`) provides ergonomics.
4. **CLDR plural categories from day one.** The contract speaks
   `zero/one/two/few/many/other`. Rules are referenced by a canonical id; GoTime
   maps the id to a compile-checked Go function. The JSON also carries the rule
   as a UExL-syntax expression, which the UExL host evaluates with the CLDR
   operands supplied as context.
5. **Locale-driven format presets via `@` sigil.** `@date`, `@time`, `@datetime`,
   etc. resolve to a per-locale NITES layout. The `@` disambiguates a preset
   name from a literal layout.
6. **Relative time is unit-based.** Rendering is decoupled from a locale-neutral
   `(unit, signed quantity, style)` triple.
7. **Deterministic fallback.** BCP-47 subtag truncation right-to-left, then `en`
   (`zh-Hant-HK → zh-Hant → zh → en`), identically across all locale-aware APIs.
8. **Allocation budget is a contract.** The hot path's budget is enforced by
   `-benchmem` benchmarks.

## Documents

- [Specification](spec.md) — *what* locale support means: API contract, the JSON
  schema contract, requirements, acceptance criteria.
- [Architecture](architecture.md) — *how* it is built: interfaces, registry, the
  single-source data pipeline, dual consumption paths, formatter hot path,
  relative-time engine, `internal/nites` layering.
- [Locale Implementer Guide](implementing-locales.md) — how to add a locale
  (author/extract the JSON), plural/ordinal rules, presets, required tests.
- [Roadmap](roadmap.md) — phased, solo-maintainer-friendly delivery plan.
