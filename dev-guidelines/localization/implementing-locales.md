# Implementing Locales

This guide defines the expectations for adding a new locale implementation.

## 1. What a Locale Must Provide

A locale implementation should provide:

- Locale code (e.g., `"es"`, `"pl"`)
- Month names:
  - short (`mmm`)
  - full (`mmmm`)
- Weekday names:
  - short (`www`)
  - full (`wwww`)
- Ordinal behavior for day/month ordinals (`dt`, `mt`)
- Relative-time phrase set:
  - just now / a minute / few minutes
  - yesterday / tomorrow
  - last/next singular unit (week/month/year/etc.)
  - numeric unit forms (e.g., “%d days ago”, “In %d days”)

## 2. Data-First Locale Pattern

Recommended structure:

- Tables:
  - `monthsShort[12]`, `monthsFull[12]`
  - `weekdaysShort[7]`, `weekdaysFull[7]`
- Small rule helpers:
  - `plural(unit, n) string` (needed for Slavic languages)
  - `ordinal(n, kind) string`

Keep locale code minimal and deterministic.

## 3. Testing Expectations

When adding a locale:

- Add tests for:
  - at least one `TimeAgo` past and future sample
  - at least one `Format` sample using `mmm` / `mmmm`
  - at least one ordinal output for `dt` or `mt`

Avoid tests that depend on the system locale/timezone.

## 4. Notes for Complex Languages (e.g., Polish)

Some languages require:

- multiple plural forms (e.g., 1, few, many)
- case inflection for month names depending on context

For Phase 1, prefer a consistent display form even if it’s not perfect in every grammatical context. Improve grammar rules later if needed.
