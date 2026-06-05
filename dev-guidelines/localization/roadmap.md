# Localization Roadmap

This roadmap is designed for gradual, maintainable delivery by a solo maintainer.

## Phase 0: Documentation (now)

- Define scope, goals, non-goals.
- Describe architecture and phased plan.

## Phase 1: Opt-in localized output

- Add locale registry + built-in `en`.
- Add locale-aware variants of:
  - `TimeAgo`
  - `Format` (for `dt`, `mt`, `mmm`, `mmmm`, `www`, `wwww`)
- Add tests that prove:
  - existing APIs remain English
  - locale-aware APIs change output

## Phase 2: Expand locale set

- Add `es` and `pl` as reference locales.
- Add contributor docs for adding more locales.

## Phase 3 (optional): Localized parsing

- Support parsing month/weekday names per locale.
- Add fallback logic (`es-MX` → `es`).

## Phase 4 (optional): Better grammar

- Context-aware month forms where applicable.
- More natural relative-time phrasing.
