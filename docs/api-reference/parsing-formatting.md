# Date Parsing and Formatting

GoTime provides powerful, intuitive functions for parsing and formatting dates using IDFS (Intuitive Date Format Specifiers).

## Core Functions

### Parse

Parses a date string using IDFS format specifiers.

```go
func Parse(value, layout string) (time.Time, error)
```

**Parameters:**
- `value`: The date string to parse
- `layout`: IDFS format specifier describing the input format

**Returns:**
- `time.Time`: Parsed date in UTC
- `error`: Error if parsing fails

**Examples:**

```go
// Basic date parsing
date, err := gotime.Parse("2025-07-07", "yyyy-mm-dd")
// Result: 2025-07-07 00:00:00 +0000 UTC

// Parse with time
datetime, err := gotime.Parse("2025-07-07 14:30:45", "yyyy-mm-dd hh:ii:ss")
// Result: 2025-07-07 14:30:45 +0000 UTC

// Different formats
usDate, err := gotime.Parse("07/07/2025", "mm/dd/yyyy")
euroDate, err := gotime.Parse("07/07/2025", "dd/mm/yyyy")
readableDate, err := gotime.Parse("July 7th, 2025", "mmmm dt, yyyy")
```

### ParseInLocation

Parses a date string in a specific timezone.

```go
func ParseInLocation(value, layout string, loc *time.Location) (time.Time, error)
```

**Parameters:**
- `value`: The date string to parse
- `layout`: IDFS format specifier
- `loc`: Target timezone location

**Returns:**
- `time.Time`: Parsed date in specified timezone
- `error`: Error if parsing fails

**Examples:**

```go
// Parse in specific timezone
est, _ := time.LoadLocation("America/New_York")
date, err := gotime.ParseInLocation("2025-07-07 14:30", "yyyy-mm-dd hh:ii", est)
// Result: 2025-07-07 14:30:00 -0400 EDT

// Parse in custom timezone
tz := time.FixedZone("IST", 5*60*60+30*60) // UTC+5:30
date, err = gotime.ParseInLocation("07/07/2025", "dd/mm/yyyy", tz)
// Result: 2025-07-07 00:00:00 +0530 IST
```

### Format

Formats a time.Time using IDFS format specifiers.

```go
func Format(dt time.Time, layout string) string
```

**Parameters:**
- `dt`: The time.Time to format
- `layout`: IDFS format specifier (empty string defaults to RFC3339)

**Returns:**
- `string`: Formatted date string

**Examples:**

```go
dt := time.Date(2025, 7, 7, 14, 30, 45, 0, time.UTC)

// Basic formatting
formatted := gotime.Format(dt, "yyyy-mm-dd")              // "2025-07-07"
formatted = gotime.Format(dt, "mm/dd/yyyy")               // "07/07/2025"
formatted = gotime.Format(dt, "dd.mm.yyyy")               // "07.07.2025"

// With time
formatted = gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")      // "2025-07-07 14:30:45"
formatted = gotime.Format(dt, "mm/dd/yyyy h:ii aa")       // "07/07/2025 2:30 PM"

// Readable formats
formatted = gotime.Format(dt, "mmmm dt, yyyy")            // "July 7th, 2025"
formatted = gotime.Format(dt, "wwww, mmmm dt, yyyy")      // "Monday, July 7th, 2025"

// Default format (RFC3339)
formatted = gotime.Format(dt, "")                         // "2025-07-07T14:30:45Z"
```

### FormatTimestamp

Formats a Unix timestamp using IDFS format specifiers.

```go
func FormatTimestamp(timestamp int64, layout string) string
```

**Parameters:**
- `timestamp`: Unix timestamp in seconds
- `layout`: IDFS format specifier (empty string defaults to RFC3339)

**Returns:**
- `string`: Formatted date string

**Examples:**

```go
timestamp := int64(1720357845) // July 7, 2025 14:30:45 UTC

formatted := gotime.FormatTimestamp(timestamp, "yyyy-mm-dd")           // "2025-07-07"
formatted = gotime.FormatTimestamp(timestamp, "yyyy-mm-dd hh:ii:ss")   // "2025-07-07 14:30:45"
formatted = gotime.FormatTimestamp(timestamp, "mmmm dt, yyyy")         // "July 7th, 2025"
```

### FormatUnix

Formats Unix time with seconds and nanoseconds.

```go
func FormatUnix(sec int64, nsec int64, layout string) string
```

**Parameters:**
- `sec`: Unix seconds
- `nsec`: Nanoseconds
- `layout`: IDFS format specifier

**Returns:**
- `string`: Formatted date string

**Examples:**

```go
sec := int64(1720357845)
nsec := int64(123456789)

// Format with nanosecond precision
formatted := gotime.FormatUnix(sec, nsec, "yyyy-mm-dd hh:ii:ss.0")
// Result: "2025-07-07 14:30:45.123"
```

## Common Use Cases

### 1. API Date Handling

```go
package api

import (
    "encoding/json"
    "github.com/maniartech/gotime"
    "time"
)

type Event struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    StartDate time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`

    // Formatted fields for JSON
    StartDateStr string `json:"start_date"`
    CreatedAtStr string `json:"created_at"`
}

func (e *Event) MarshalJSON() ([]byte, error) {
    type Alias Event
    return json.Marshal(&struct {
        *Alias
        StartDateStr string `json:"start_date"`
        CreatedAtStr string `json:"created_at"`
    }{
        Alias:        (*Alias)(e),
        StartDateStr: gotime.Format(e.StartDate, "yyyy-mm-dd"),
        CreatedAtStr: gotime.Format(e.CreatedAt, "yyyy-mm-dd hh:ii:ss"),
    })
}

func (e *Event) UnmarshalJSON(data []byte) error {
    type Alias Event
    aux := &struct {
        *Alias
        StartDateStr string `json:"start_date"`
        CreatedAtStr string `json:"created_at"`
    }{
        Alias: (*Alias)(e),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    var err error
    e.StartDate, err = gotime.Parse(aux.StartDateStr, "yyyy-mm-dd")
    if err != nil {
        return err
    }

    e.CreatedAt, err = gotime.Parse(aux.CreatedAtStr, "yyyy-mm-dd hh:ii:ss")
    return err
}
```

### 2. User Input Parsing

```go
func parseUserDate(input string) (time.Time, error) {
    // Common user input formats
    formats := []string{
        "yyyy-mm-dd",           // 2025-07-07
        "mm/dd/yyyy",           // 07/07/2025
        "dd/mm/yyyy",           // 07/07/2025
        "m/d/yyyy",             // 7/7/2025
        "d/m/yyyy",             // 7/7/2025
        "mmmm dt, yyyy",        // July 7th, 2025
        "mmmm d, yyyy",         // July 7, 2025
        "dd mmmm yyyy",         // 07 July 2025
        "yyyy-mm-dd hh:ii",     // 2025-07-07 14:30
        "yyyy-mm-dd hh:ii:ss",  // 2025-07-07 14:30:45
    }

    for _, format := range formats {
        if parsed, err := gotime.Parse(input, format); err == nil {
            return parsed, nil
        }
    }

    return time.Time{}, fmt.Errorf("unable to parse date: %s", input)
}
```

### 3. Log Formatting

```go
package logger

import (
    "fmt"
    "github.com/maniartech/gotime"
    "time"
)

type Logger struct {
    timestampFormat string
}

func NewLogger() *Logger {
    return &Logger{
        timestampFormat: "yyyy-mm-dd hh:ii:ss.000",
    }
}

func (l *Logger) Info(message string) {
    timestamp := gotime.Format(time.Now(), l.timestampFormat)
    fmt.Printf("[%s] INFO: %s\n", timestamp, message)
}

func (l *Logger) Error(message string) {
    timestamp := gotime.Format(time.Now(), l.timestampFormat)
    fmt.Printf("[%s] ERROR: %s\n", timestamp, message)
}

func (l *Logger) SetTimestampFormat(format string) {
    l.timestampFormat = format
}
```

### 4. Database Integration

```go
package db

import (
    "database/sql"
    "github.com/maniartech/gotime"
    "time"
)

// Custom scanner for database dates
type Date struct {
    time.Time
}

func (d *Date) Scan(value interface{}) error {
    switch v := value.(type) {
    case string:
        // Try common database formats
        formats := []string{
            "yyyy-mm-dd hh:ii:ss",      // MySQL DATETIME
            "yyyy-mm-dd hh:ii:ss.0",    // MySQL DATETIME(3)
            "yyyy-mm-dd",               // DATE
            time.RFC3339,               // ISO 8601
        }

        for _, format := range formats {
            if parsed, err := gotime.Parse(v, format); err == nil {
                d.Time = parsed
                return nil
            }
        }
        return fmt.Errorf("cannot parse date: %s", v)

    case time.Time:
        d.Time = v
        return nil

    case nil:
        d.Time = time.Time{}
        return nil

    default:
        return fmt.Errorf("cannot scan %T into Date", value)
    }
}

func (d Date) Value() (driver.Value, error) {
    if d.Time.IsZero() {
        return nil, nil
    }
    return gotime.Format(d.Time, "yyyy-mm-dd hh:ii:ss"), nil
}
```

### 5. Configuration Files

```go
package config

import (
    "github.com/maniartech/gotime"
    "time"
)

type Config struct {
    StartDate    time.Time `yaml:"start_date"`
    EndDate      time.Time `yaml:"end_date"`
    ScheduleTime time.Time `yaml:"schedule_time"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
    type rawConfig struct {
        StartDate    string `yaml:"start_date"`
        EndDate      string `yaml:"end_date"`
        ScheduleTime string `yaml:"schedule_time"`
    }

    var raw rawConfig
    if err := unmarshal(&raw); err != nil {
        return err
    }

    var err error
    c.StartDate, err = gotime.Parse(raw.StartDate, "yyyy-mm-dd")
    if err != nil {
        return fmt.Errorf("invalid start_date: %v", err)
    }

    c.EndDate, err = gotime.Parse(raw.EndDate, "yyyy-mm-dd")
    if err != nil {
        return fmt.Errorf("invalid end_date: %v", err)
    }

    c.ScheduleTime, err = gotime.Parse(raw.ScheduleTime, "hh:ii")
    if err != nil {
        return fmt.Errorf("invalid schedule_time: %v", err)
    }

    return nil
}

func (c Config) MarshalYAML() (interface{}, error) {
    return map[string]string{
        "start_date":    gotime.Format(c.StartDate, "yyyy-mm-dd"),
        "end_date":      gotime.Format(c.EndDate, "yyyy-mm-dd"),
        "schedule_time": gotime.Format(c.ScheduleTime, "hh:ii"),
    }, nil
}
```

## Error Handling

### Common Parsing Errors

```go
// Invalid format
_, err := gotime.Parse("2025-13-01", "yyyy-mm-dd")
// Error: parsing time "2025-13-01": month out of range

// Format mismatch
_, err = gotime.Parse("07/07/2025", "yyyy-mm-dd")
// Error: parsing time "07/07/2025" as "2006-01-02": cannot parse "07/07/2025" as "2006"

// Empty input
_, err = gotime.Parse("", "yyyy-mm-dd")
// Error: parsing time "": extra text: ""
```

### Robust Error Handling

```go
func parseWithFallback(input string, primaryFormat string, fallbackFormats []string) (time.Time, error) {
    // Try primary format first
    if parsed, err := gotime.Parse(input, primaryFormat); err == nil {
        return parsed, nil
    }

    // Try fallback formats
    for _, format := range fallbackFormats {
        if parsed, err := gotime.Parse(input, format); err == nil {
            return parsed, nil
        }
    }

    return time.Time{}, fmt.Errorf("unable to parse date '%s' with any known format", input)
}

// Usage
date, err := parseWithFallback(
    "07/07/2025",
    "yyyy-mm-dd",
    []string{"mm/dd/yyyy", "dd/mm/yyyy", "m/d/yyyy"},
)
```

## Performance Tips

### 1. Cache Frequently Used Formats

```go
const (
    APIDateFormat = "yyyy-mm-dd"
    LogFormat     = "yyyy-mm-dd hh:ii:ss"
    UserFormat    = "mmmm dt, yyyy"
)

// Use constants instead of string literals
formatted := gotime.Format(time.Now(), APIDateFormat)
```

### 2. Reuse Format Strings

```go
// ❌ Creates new format string each time
for _, event := range events {
    event.FormattedDate = gotime.Format(event.Date, "yyyy-mm-dd")
}

// ✅ Reuses cached format
format := "yyyy-mm-dd"
for _, event := range events {
    event.FormattedDate = gotime.Format(event.Date, format)
}
```

### 3. Batch Operations

```go
func formatDates(dates []time.Time, format string) []string {
    results := make([]string, len(dates))
    for i, date := range dates {
        results[i] = gotime.Format(date, format)
    }
    return results
}
```

## Best Practices

### 1. Use Descriptive Format Constants

```go
const (
    DatabaseTimestamp = "yyyy-mm-dd hh:ii:ss"
    APIDate          = "yyyy-mm-dd"
    UserFriendly     = "mmmm dt, yyyy"
    LogTimestamp     = "yyyy-mm-dd hh:ii:ss.000"
)
```

### 2. Validate Input Early

```go
func createEvent(dateStr string) (*Event, error) {
    date, err := gotime.Parse(dateStr, "yyyy-mm-dd")
    if err != nil {
        return nil, fmt.Errorf("invalid date format: %v", err)
    }

    return &Event{Date: date}, nil
}
```

### 3. Handle Timezones Appropriately

```go
// For APIs, always use UTC internally
func parseAPIDate(input string) (time.Time, error) {
    return gotime.Parse(input, "yyyy-mm-dd")
}

// For user input, consider user's timezone
func parseUserDate(input string, userTZ *time.Location) (time.Time, error) {
    return gotime.ParseInLocation(input, "yyyy-mm-dd", userTZ)
}
```

### 4. Document Your Date Formats

```go
// UserEvent represents a user-created event
type UserEvent struct {
    // Date in YYYY-MM-DD format for API consistency
    Date string `json:"date" example:"2025-07-07"`

    // Time in HH:MM format (24-hour)
    Time string `json:"time" example:"14:30"`
}
```

---

## Summary

GoTime's parsing and formatting functions provide:

- **Intuitive IDFS format specifiers** instead of cryptic reference times
- **Comprehensive timezone support** with ParseInLocation
- **Flexible input handling** for various date formats
- **Performance optimization** through format caching
- **Robust error handling** for invalid inputs

These functions form the foundation for all date/time operations in GoTime, making them both powerful and easy to use.

---

Next: [Date Conversion](conversion.md) | [Back to API Reference](../api-reference/)
