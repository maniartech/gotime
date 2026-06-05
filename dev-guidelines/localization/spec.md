# Localization Specification

This document defines **what** “locale support” means for GoTime, without prescribing implementation details beyond the public API contract and functional requirements.

## 1. Terminology

- **Locale**: A language + optional region/culture variant (examples: `en`, `en-US`, `es`, `pl`).
- **Locale-aware output**: Output that can differ by language rules (words, pluralization, ordinals).
- **Locale-neutral output**: Numeric/timezone output that is independent of human language.

## 2. Current Locale-Sensitive Areas

These areas must be considered locale-sensitive because they produce English words today:

1. **Relative time** (`TimeAgo`)
   - Past phrases: "Just now", "A minute ago", "Few minutes ago", "Yesterday", "Last week", "5 days ago"
   - Future phrases: "In a few seconds", "In a minute", "In a few minutes", "Tomorrow", "In a week", "In 5 days"

2. **Ordinals** (`dt`, `mt` formatting)
   - English-only suffix rules (`st`, `nd`, `rd`, `th`) are embedded in formatting logic.

3. **Month/weekday names** (`mmm`, `mmmm`, `www`, `wwww`)
   - Go’s `time.Format` emits English names; this is not locale-aware.

## 3. Requirements

### 3.1 Backward Compatibility

- Existing public functions must keep default behavior (English output).
- Locale support must be opt-in.

### 3.2 API Requirements

- Provide a way to produce localized output for:
  - Relative time (`TimeAgo`)
  - Formatting tokens that produce words (`dt`, `mt`, `mmm`, `mmmm`, `www`, `wwww`)

- Locale selection must be explicit (e.g., pass a locale object or locale code).

### 3.3 Locale Rules Requirements

A locale must be able to define:

- **Fixed phrases** (e.g., “Yesterday”, “Tomorrow”, “Just now”).
- **Unit forms** (pluralization rules may vary by locale).
- **Ordinal formatting** for day/month numbers.
- **Month names** (short/full) and **weekday names** (short/full).

### 3.4 Error Handling

- Default English locale must always be available.
- If a locale code is unknown/unregistered, behavior should be deterministic:
  - either fall back to English, or return an error (decision left to implementation, but must be consistent across APIs).

## 4. Out of Scope (Phase 1)

- Parsing localized month/weekday names.
- Transliteration, calendar system changes, or CLDR-level features.

## 5. Acceptance Criteria

A Phase 1 implementation is “done” when:

- English behavior remains unchanged for existing APIs.
- A caller can request a different locale for:
  - `TimeAgo`
  - formatting of `dt`, `mt`, `mmm`, `mmmm`, `www`, `wwww`
- Adding a new locale can be done without modifying core formatting logic (beyond registration).

## 6. Examples (API Contract)

These examples show the intended *usage* of locale-aware output. Names are illustrative; the implementation may choose either a “pass locale object” style or an “options struct” style, but the behavior must match.

### 6.1 Default Behavior (Backward Compatible)

Existing APIs remain English by default:

```go
package main

import (
  "fmt"
  "time"

  "github.com/maniartech/gotime/v2"
)

func main() {
  fmt.Println(gotime.TimeAgo(time.Now().Add(-2 * time.Hour)))
  // Example output: "2 hours ago"

  fmt.Println(gotime.Format(time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), "wwww, mmmm dt"))
  // Example output: "Monday, July 7th"
}
```

### 6.2 Locale-Aware Relative Time

Callers can request localized relative-time text:

```go
package main

import (
  "fmt"
  "time"

  "github.com/maniartech/gotime/v2"
  "github.com/maniartech/gotime/v2/locales"
)

func main() {
  t := time.Now().Add(-5 * time.Minute)

  // Style A: pass a locale value
  fmt.Println(gotime.TimeAgoIn(t, locales.MustGet("es")))
  // Example output (Spanish): "hace 5 minutos"

  // Style B: pass locale via options
  fmt.Println(gotime.TimeAgoWithOptions(t, gotime.TimeAgoOptions{Locale: "pl"}))
  // Example output (Polish): "5 minut temu"
}
```

### 6.3 Locale-Aware Formatting Tokens

Localized formatting applies to tokens that emit language:

- `dt`, `mt` (ordinals)
- `mmm`, `mmmm` (month names)
- `www`, `wwww` (weekday names)

```go
package main

import (
  "fmt"
  "time"

  "github.com/maniartech/gotime/v2"
  "github.com/maniartech/gotime/v2/locales"
)

func main() {
  dt := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)

  fmt.Println(gotime.FormatIn(dt, "wwww, mmmm dt", locales.MustGet("en")))
  // Example output: "Monday, July 7th"

  fmt.Println(gotime.FormatIn(dt, "wwww, mmmm dt", locales.MustGet("es")))
  // Example output: "lunes, julio 7" (exact grammar may vary)

  fmt.Println(gotime.FormatIn(dt, "wwww, mmmm dt", locales.MustGet("pl")))
  // Example output: "poniedziałek, lipiec 7" (exact grammar may vary)
}
```

### 6.4 Unknown Locale Handling

Unknown locale handling must be deterministic:

```go
// Either: fallback to English
fmt.Println(gotime.TimeAgoWithOptions(t, gotime.TimeAgoOptions{Locale: "xx"}))

// Or: return an error on locale resolution (if the API design supports it)
loc, err := locales.Get("xx")
if err != nil {
  // handle unsupported locale
}
_ = loc
```
