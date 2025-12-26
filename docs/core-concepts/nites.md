[Home](../README.md) > [Core Concepts](README.md) > NITES

# Natural and Intuitive Time Expression Syntax (NITES)

NITES is the core innovation that makes GoTime intuitive and powerful. It replaces Go's cryptic reference time (`2006-01-02 15:04:05`) with human-readable, logical format specifiers.

> **📄 NITES Specification Paper**: A comprehensive technical specification and research paper on NITES is currently under development and will be released soon. This will include formal grammar definitions, implementation guidelines, and comparative analysis with other time formatting systems.

## Philosophy Behind NITES

### The Problem with Reference Time

Go's standard library uses a reference time approach:
```go
// Go's magic reference time
time.Now().Format("2006-01-02 15:04:05")
```

**Issues:**
- **Arbitrary**: Why January 2, 2006? Why 15:04:05?
- **Hard to Remember**: No logical pattern
- **Error-Prone**: Easy to mix up numbers
- **Case-Sensitive**: `MM` vs `mm` has different meanings

### NITES Solution

```go
// Intuitive, logical format specifiers
gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")
```

**Benefits:**
- **Logical**: `yyyy` = 4-digit year, `mm` = 2-digit month
- **Memorable**: Format looks like what you want
- **Case-Insensitive**: No need to remember case rules
- **Self-Documenting**: Format intent is clear
- **Intuitive**: Natural patterns that developers expect
- **Hackable**: Easy to experiment and build custom formats

## NITES Specification

> **⚠️ Note**: This is a practical implementation reference. The complete formal specification with EBNF grammar, semantic rules, and implementation guidelines is being prepared for academic publication.

### Case Insensitivity

NITES is completely case-insensitive. These are all equivalent:
```go
gotime.Format(dt, "YYYY-MM-DD")
gotime.Format(dt, "yyyy-mm-dd")
gotime.Format(dt, "Yyyy-Mm-Dd")
gotime.Format(dt, "YyYy-mM-DD")
```

### Repetition Logic

The number of characters determines the output format:
- More characters = more detailed/padded output
- Fewer characters = condensed output

## Complete Format Reference

### Date Formats

| Format | Output | Description | Example |
|--------|--------|-------------|---------|
| `y` | `6` | 1-digit year (last digit) | 2026 → 6 |
| `yy` | `06` | 2-digit year with leading zero | 2026 → 26 |
| `yyyy` | `2006` | 4-digit year | 2026 → 2026 |
| `m` | `1` | Month without leading zero | January → 1 |
| `mm` | `01` | Month with leading zero | January → 01 |
| `mt` | `1st` | Month in ordinal format | January → 1st |
| `mmm` | `Jan` | Month short name | January → Jan |
| `mmmm` | `January` | Month full name | January → January |
| `d` | `2` | Day without leading zero | 2nd → 2 |
| `dd` | `02` | Day with leading zero | 2nd → 02 |
| `db` | ` 2` | Day blank-padded | 2nd → ` 2` |
| `dt` | `2nd` | Day in ordinal format | 2nd → 2nd |
| `ddd` | `002` | Day of year (zero-padded) | Jan 2 → 002 |
| `www` | `Mon` | Weekday short name | Monday → Mon |
| `wwww` | `Monday` | Weekday full name | Monday → Monday |

### Time Formats

| Format | Output | Description | Example |
|--------|--------|-------------|---------|
| `h` | `3` | Hour 12-format, no leading zero | 3 AM → 3 |
| `hh` | `03` | Hour 12-format, with leading zero | 3 AM → 03 |
| `hhhh` | `15` | Hour 24-format with leading zero | 3 PM → 15 |
| `i` | `4` | Minute without leading zero | 04:04 → 4 |
| `ii` | `04` | Minute with leading zero | 04:04 → 04 |
| `s` | `5` | Second without leading zero | 05 seconds → 5 |
| `ss` | `05` | Second with leading zero | 05 seconds → 05 |
| `a` | `pm` | AM/PM lowercase | 3 PM → pm |
| `aa` | `PM` | AM/PM uppercase | 3 PM → PM |
| `.0` | `.000` | Microseconds with leading zeros | - |
| `.9` | `.999` | Microseconds without trailing zeros | - |

### Timezone Formats

| Format | Output | Description | Example |
|--------|--------|-------------|---------|
| `z` | `Z` | UTC indicator | UTC → Z |
| `zz` | `MST` | Timezone abbreviation | Mountain → MST |
| `o` | `±07` | Timezone offset (hours only) | UTC+7 → +07 |
| `oo` | `±0700` | Timezone offset (no colon) | UTC+7 → +0700 |
| `ooo` | `±07:00` | Timezone offset (with colon) | UTC+7 → +07:00 |

## Common Format Patterns

### Standard Formats

```go
// ISO 8601
gotime.Format(dt, "yyyy-mm-dd")           // 2025-07-07
gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")  // 2025-07-07 14:30:45

// US Format
gotime.Format(dt, "mm/dd/yyyy")           // 07/07/2025
gotime.Format(dt, "m/d/yyyy")             // 7/7/2025

// European Format
gotime.Format(dt, "dd/mm/yyyy")           // 07/07/2025
gotime.Format(dt, "dd.mm.yyyy")           // 07.07.2025

// Readable Formats
gotime.Format(dt, "mmmm dt, yyyy")        // July 7th, 2025
gotime.Format(dt, "wwww, mmmm dt, yyyy")  // Monday, July 7th, 2025
```

### Time-Only Formats

```go
// 12-hour format
gotime.Format(dt, "h:ii aa")              // 2:30 PM
gotime.Format(dt, "hh:ii:ss aa")          // 02:30:45 PM

// 24-hour format
gotime.Format(dt, "hhhh:ii")               // 14:30
gotime.Format(dt, "hhhh:ii:ss")            // 14:30:45
```

### Database Formats

```go
// MySQL DATETIME
gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")  // 2025-07-07 14:30:45

// PostgreSQL TIMESTAMP
gotime.Format(dt, "yyyy-mm-dd hh:ii:ss.0") // 2025-07-07 14:30:45.000

// SQLite
gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")  // 2025-07-07 14:30:45
```

## Practical Examples

### User Interface Displays

```go
package main

import (
    "fmt"
    "time"
    "github.com/maniartech/gotime/v2"
)

func displayFormats(dt time.Time) {
    // Short formats
    fmt.Println("Short Date:", gotime.Format(dt, "mm/dd/yy"))
    fmt.Println("ISO Date:", gotime.Format(dt, "yyyy-mm-dd"))

    // Readable formats
    fmt.Println("Friendly:", gotime.Format(dt, "mmmm dt, yyyy"))
    fmt.Println("Full:", gotime.Format(dt, "wwww, mmmm dt, yyyy"))

    // Time formats
    fmt.Println("12-hour:", gotime.Format(dt, "h:ii aa"))
    fmt.Println("24-hour:", gotime.Format(dt, "hhhh:ii"))

    // Combined
    fmt.Println("Log format:", gotime.Format(dt, "yyyy-mm-dd hh:ii:ss aa"))
}
```

### API Responses

```go
type Event struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    StartDate string `json:"start_date"`
    StartTime string `json:"start_time"`
    CreatedAt string `json:"created_at"`
}

func formatEvent(event Event, startTime time.Time, createdAt time.Time) Event {
    return Event{
        ID:        event.ID,
        Name:      event.Name,
        StartDate: gotime.Format(startTime, "yyyy-mm-dd"),
        StartTime: gotime.Format(startTime, "hh:ii aa"),
        CreatedAt: gotime.Format(createdAt, "yyyy-mm-dd hh:ii:ss"),
    }
}
```

### Log Formatting

```go
func logWithTimestamp(message string) {
    timestamp := gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss.000")
    fmt.Printf("[%s] %s\n", timestamp, message)
}
```

## Parsing with NITES

The same format specifiers work for parsing:

```go
// Parse various formats
date1, _ := gotime.Parse("2025-07-07", "yyyy-mm-dd")
date2, _ := gotime.Parse("07/07/2025", "mm/dd/yyyy")
date3, _ := gotime.Parse("July 7th, 2025", "mmmm dt, yyyy")
date4, _ := gotime.Parse("2025-07-07 14:30:45", "yyyy-mm-dd hh:ii:ss")
```

### Flexible Input Parsing

```go
func parseFlexibleDate(input string) (time.Time, error) {
    formats := []string{
        "yyyy-mm-dd",
        "mm/dd/yyyy",
        "dd/mm/yyyy",
        "mmmm dt, yyyy",
        "yyyy-mm-dd hh:ii:ss",
    }

    for _, format := range formats {
        if parsed, err := gotime.Parse(input, format); err == nil {
            return parsed, nil
        }
    }
    return time.Time{}, fmt.Errorf("unable to parse date: %s", input)
}
```

## Built-in Format Constants

GoTime provides access to all standard Go time format constants:

```go
// Standard Go formats work too
formatted := gotime.Format(dt, time.RFC3339)     // 2025-07-07T14:30:45Z
formatted = gotime.Format(dt, time.Kitchen)      // 2:30PM
formatted = gotime.Format(dt, time.RFC822)       // 07 Jul 25 14:30 UTC
```

## Advanced Usage

### Custom Separators

```go
// Different separators
gotime.Format(dt, "yyyy.mm.dd")           // 2025.07.07
gotime.Format(dt, "yyyy_mm_dd")           // 2025_07_07
gotime.Format(dt, "yyyy mm dd")           // 2025 07 07
gotime.Format(dt, "dd-mm-yyyy")           // 07-07-2025
```

### Mixed Formats

```go
// Combine different elements
gotime.Format(dt, "wwww the dt of mmmm, yyyy")  // Monday the 7th of July, 2025
gotime.Format(dt, "yyyy-mm-dd at hh:ii aa")     // 2025-07-07 at 02:30 PM
gotime.Format(dt, "mmmm yyyy (yyyy-mm-dd)")     // July 2025 (2025-07-07)
```

### International Formats

```go
// Various international standards
gotime.Format(dt, "dd.mm.yyyy")           // German: 07.07.2025
gotime.Format(dt, "dd/mm/yyyy")           // UK: 07/07/2025
gotime.Format(dt, "yyyy年mm月dd日")         // Japanese: 2025年07月07日
```

## Performance Considerations

### Caching

NITES formats are cached internally for performance:
```go
// First call converts and caches
fmt1 := gotime.Format(dt, "yyyy-mm-dd")  // ~1200ns

// Subsequent calls use cache
fmt2 := gotime.Format(dt, "yyyy-mm-dd")  // ~800ns
```

### Memory Usage

- Format conversion is done once and cached
- No memory leaks from format string processing
- Efficient string building for complex formats

## Best Practices

### 1. Use Descriptive Formats

```go
// ❌ Unclear intent
gotime.Format(dt, "dd/mm/yy")

// ✅ Clear intent
gotime.Format(dt, "dd/mm/yyyy")  // Full year preferred
```

### 2. Be Consistent

```go
// ❌ Mixed formats in same application
apiDate := gotime.Format(dt, "yyyy-mm-dd")
logDate := gotime.Format(dt, "mm/dd/yyyy")

// ✅ Consistent format choice
apiDate := gotime.Format(dt, "yyyy-mm-dd")
logDate := gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")
```

### 3. Consider Locale

```go
// For US applications
usFormat := gotime.Format(dt, "mm/dd/yyyy")

// For European applications
euroFormat := gotime.Format(dt, "dd/mm/yyyy")

// For international/APIs
isoFormat := gotime.Format(dt, "yyyy-mm-dd")
```

### 4. Document Your Formats

```go
const (
    // API response format for dates
    APIDateFormat = "yyyy-mm-dd"

    // Log timestamp format
    LogTimestampFormat = "yyyy-mm-dd hh:ii:ss.000"

    // User display format
    UserFriendlyFormat = "mmmm dt, yyyy"
)
```

## Migration from Standard Time

### Simple Replacement

```go
// Before
formatted := time.Now().Format("2006-01-02")

// After
formatted := gotime.Format(time.Now(), "yyyy-mm-dd")
```

### Complex Formats

```go
// Before
formatted := time.Now().Format("Monday, January 2, 2006 at 3:04 PM")

// After
formatted := gotime.Format(time.Now(), "wwww, mmmm dt, yyyy at h:ii aa")
```

---

## Summary

NITES makes date formatting intuitive by:

1. **Logical Patterns**: `yyyy` for year, `mm` for month, `dd` for day
2. **Case Insensitivity**: No need to remember case rules
3. **Self-Documentation**: Format strings are readable
4. **Flexibility**: Rich set of specifiers for any need
5. **Performance**: Cached conversions for efficiency

**NITES transforms cryptic format strings into readable, maintainable code.**

### Formal Specification

The complete NITES specification includes:

- **📋 Grammar Definition**: Formal EBNF grammar for all format specifiers
- **📊 Performance Analysis**: Benchmarks against standard library and other systems
- **🔬 Cognitive Load Study**: Research on developer comprehension and error rates
- **🌐 Internationalization**: Unicode support and locale-aware formatting
- **🔧 Implementation Guide**: Reference implementations for multiple languages

**Coming Soon**: Watch for the official NITES specification paper for implementers and researchers.

---

Next: Learn about [Design Philosophy](design-philosophy.md) or jump to [API Reference](../api-reference/) to see NITES in action.
