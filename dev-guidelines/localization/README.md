# Localization (i18n/l10n) Guidelines

This directory contains the specification and architecture for adding locale-aware (human-language) output to GoTime in a maintainable way.

## Scope

GoTime is *mostly* locale-neutral for calculations and numeric formats. However, a few features produce human-language output and are therefore locale-sensitive:

- Relative time text (e.g., "Just now", "Yesterday", "In a few minutes")
- Ordinal suffixes for `dt` / `mt` formatting (e.g., `1st`, `2nd`, `3rd`)
- Month/weekday names when using name-based NITES tokens (`mmm`, `mmmm`, `www`, `wwww`)

## Goals

- Keep existing behavior stable (English output remains the default).
- Add locale support in a way that scales to many locales without rewriting the library.
- Make adding a new locale a small, isolated change (ideally: 1 file of data + a small amount of rule code).
- Avoid heavy dependencies.

## Non-Goals (initially)

- Locale-aware parsing of month/weekday names (formatting-only localization is the first milestone).
- Automatic locale detection from OS/env.

## Documents

- [Specification](spec.md)
- [Architecture](architecture.md)
- [Locale Implementer Guide](implementing-locales.md)
- [Roadmap](roadmap.md)
