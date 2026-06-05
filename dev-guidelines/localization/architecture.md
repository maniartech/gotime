# Localization Architecture

This document proposes an architecture that is **incremental**, **maintainable for a solo developer**, and **safe for contributors**.

## 1. Design Principles

1. **Opt-in localization**: default behavior stays English.
2. **Isolated locale logic**: locale-specific rules are not spread across the codebase.
3. **Data-first**: most locales should be representable as tables + small rule functions.
4. **Small public surface area**: avoid introducing many new public types.
5. **No heavy deps**: don’t require CLDR libraries to get started.

## 2. Proposed High-Level Shape

### 2.1 Introduce a Locale Provider

Add a small locale abstraction that supplies:

- phrases for relative time
- unit pluralization
- ordinal rules
- month/weekday names

This provider can be implemented by:

- built-in English locale
- built-in future locales (Spanish/Polish)
- third-party locales provided by library users

### 2.2 Keep Existing APIs

Do not change existing function signatures like `TimeAgo(...)` and `Format(...)`. Introduce parallel APIs that accept locale input.

Example (names TBD):

- `TimeAgoIn(t time.Time, locale Locale, baseTime ...time.Time) string`
- `FormatIn(t time.Time, layout string, locale Locale) string`

Alternative pattern (options struct):

- `TimeAgoWithOptions(t, TimeAgoOptions{Locale: ...})`

## 3. Where Localization Hooks In

### 3.1 Relative Time (`TimeAgo`)

Refactor relative-time formatting to use locale phrases and pluralization rules. The time-difference math remains unchanged.

### 3.2 Ordinals (`dt`, `mt`)

Replace hardcoded English suffix logic with:

- `locale.Ordinal(n, kind)`

English can keep the current behavior; other locales provide their own.

### 3.3 Month/Weekday Names (`mmm/mmmm/www/wwww`)

Today, these tokens map to Go’s `time.Format` layouts (`Jan`, `January`, `Mon`, `Monday`). This is always English.

Architecture change:

- Treat `mmm/mmmm/www/wwww` as *tokens* at formatting time.
- For numeric portions, keep using Go `time.Format`.
- When emitting these name tokens, call the locale provider.

This avoids re-implementing all of Go’s formatting while still making names locale-aware.

## 4. Locale Registry

Provide a small registry keyed by locale code:

- `Register(locale)` for adding locales
- `Get(code)` for retrieving

The library ships with `en` registered by default.

Registry rules:

- registration should be safe to call from `init()`
- the default English locale must always be available

## 5. Phased Delivery Plan

### Phase 1 (formatting only)

- Localized `TimeAgo` output
- Localized `dt/mt` ordinals
- Localized `mmm/mmmm/www/wwww` output

No localized parsing.

### Phase 2 (optional)

- Localized parsing for month/weekday names
- Better locale fallbacks (e.g., `es-MX` → `es`)

## 6. Contribution Safety

To keep PRs manageable:

- Each locale should live in a single file (e.g., `locales/pl.go`).
- A locale PR should not require touching core logic.
- Add focused tests per locale output (snapshot-ish tests for known inputs).

## 7. Compatibility Notes

- Go’s standard library does not provide locale-aware month/weekday names in `time.Format`.
- Therefore, name-based localization must come from GoTime’s own locale tables.
