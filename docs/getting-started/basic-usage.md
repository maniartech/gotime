[Home](../README.md) > [Getting Started](README.md) > Basic Usage

# Basic Usage

This guide covers the fundamental concepts and usage patterns of GoTime, building on the [Quick Start Guide](quick-start.md) with more detailed explanations and examples.

## Core Concepts

### 1. NITES - The Foundation

GoTime's power comes from **Natural and Intuitive Time Expression Syntax (NITES)**. Unlike Go's cryptic reference time (`2006-01-02`), NITES uses logical, memorable patterns:

```go
// Standard Go - cryptic and hard to remember
time.Now().Format("2006-01-02 15:04:05")

// GoTime - intuitive and readable
gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")
```

**Key NITES Principles:**
- **Logical**: `yyyy` = 4-digit year, `mm` = 2-digit month, `dd` = 2-digit day
- **Case-insensitive**: `YYYY`, `yyyy`, `Yyyy` all work the same
- **Repetition-based**: More characters = more detail/padding

### 2. Three Core Operations

GoTime revolves around three fundamental operations:

1. **Formatting**: Convert `time.Time` to string
2. **Parsing**: Convert string to `time.Time`
3. **Converting**: Transform string from one format to another

```go
// 1. Formatting
formatted := gotime.Format(time.Now(), "yyyy-mm-dd")

// 2. Parsing
parsed, err := gotime.Parse("2025-07-07", "yyyy-mm-dd")

// 3. Converting
converted, err := gotime.Convert("2025-07-07", "yyyy-mm-dd", "mm/dd/yyyy")
```

## Working with Dates

### Common Date Formats

```go
dt := time.Date(2025, 7, 7, 14, 30, 45, 0, time.UTC)

// ISO 8601 (recommended for APIs)
iso := gotime.Format(dt, "yyyy-mm-dd")               // "2025-07-07"
isoTime := gotime.Format(dt, "yyyy-mm-dd hh:ii:ss")  // "2025-07-07 14:30:45"

// US Format
us := gotime.Format(dt, "mm/dd/yyyy")                // "07/07/2025"
usTime := gotime.Format(dt, "mm/dd/yyyy h:ii aa")    // "07/07/2025 2:30 PM"

// European Format
euro := gotime.Format(dt, "dd/mm/yyyy")              // "07/07/2025"
euroTime := gotime.Format(dt, "dd.mm.yyyy hh:ii")    // "07.07.2025 14:30"

// Human-readable
readable := gotime.Format(dt, "mmmm dt, yyyy")       // "July 7th, 2025"
full := gotime.Format(dt, "wwww, mmmm dt, yyyy")     // "Monday, July 7th, 2025"
```

### Parsing User Input

```go
func parseUserDate(input string) (time.Time, error) {
    // Define supported formats in order of preference
    formats := []string{
        "yyyy-mm-dd",           // ISO format
        "mm/dd/yyyy",           // US format
        "dd/mm/yyyy",           // European format
        "m/d/yyyy",             // US short
        "d/m/yyyy",             // European short
        "mmmm dt, yyyy",        // "July 7th, 2025"
        "mmmm d, yyyy",         // "July 7, 2025"
    }

    for _, format := range formats {
        if date, err := gotime.Parse(input, format); err == nil {
            return date, nil
        }
    }

    return time.Time{}, fmt.Errorf("unrecognized date format: %s", input)
}

// Examples
date1, _ := parseUserDate("2025-07-07")      // ISO
date2, _ := parseUserDate("07/07/2025")      // US
date3, _ := parseUserDate("7/7/2025")        // US short
date4, _ := parseUserDate("July 7th, 2025")  // Readable
```

## Working with Time

### Time-Only Operations

```go
// Parse time strings
morning, _ := gotime.Parse("09:30", "hh:ii")           // 9:30 AM
afternoon, _ := gotime.Parse("2:30 PM", "h:ii aa")     // 2:30 PM
evening, _ := gotime.Parse("18:45", "hhh:ii")          // 6:45 PM

// Format time values
now := time.Now()
time12 := gotime.Format(now, "h:ii aa")                // "2:30 PM"
time24 := gotime.Format(now, "hhh:ii")                 // "14:30"
withSeconds := gotime.Format(now, "hhh:ii:ss")         // "14:30:45"
```

### Combining Date and Time

```go
// Create datetime from separate components
dateStr := "2025-07-07"
timeStr := "14:30"
datetimeStr := dateStr + " " + timeStr
datetime, _ := gotime.Parse(datetimeStr, "yyyy-mm-dd hh:ii")

// Format datetime in different ways
iso8601 := gotime.Format(datetime, "yyyy-mm-dd'T'hh:ii:ss'Z'")
readable := gotime.Format(datetime, "mmmm dt, yyyy at h:ii aa")
log := gotime.Format(datetime, "yyyy-mm-dd hh:ii:ss")
```

## Timezone Handling

### Parsing with Timezones

```go
// Parse in specific timezone
est, _ := time.LoadLocation("America/New_York")
date, err := gotime.ParseInLocation("2025-07-07 14:30", "yyyy-mm-dd hh:ii", est)

// Parse in custom timezone
tz := time.FixedZone("IST", 5*60*60+30*60) // UTC+5:30
date, err = gotime.ParseInLocation("07/07/2025", "dd/mm/yyyy", tz)

// Parse UTC (default)
utcDate, err := gotime.Parse("2025-07-07 14:30", "yyyy-mm-dd hh:ii")
```

### Formatting with Timezone Information

```go
now := time.Now()

// Basic timezone info
withTZ := gotime.Format(now, "yyyy-mm-dd hh:ii:ss zz")    // With timezone name
withOffset := gotime.Format(now, "yyyy-mm-dd hh:ii:ss ooo") // With offset

// UTC indicator
utc := gotime.Format(now.UTC(), "yyyy-mm-dd hh:ii:ss z")  // Adds 'Z' for UTC
```

## Relative Time Operations

### Getting Common Dates

```go
// Basic relative dates
today := time.Now()
yesterday := gotime.Yesterday()
tomorrow := gotime.Tomorrow()

// Week operations
lastWeek := gotime.LastWeek()
nextWeek := gotime.NextWeek()
weekStart := gotime.WeekStart()       // Start of current week (Monday)
weekEnd := gotime.WeekEnd()           // End of current week (Sunday)

// Month operations
monthStart := gotime.MonthStart()     // First day of current month
monthEnd := gotime.MonthEnd()         // Last day of current month
lastMonth := gotime.LastMonth()
nextMonth := gotime.NextMonth()

// Year operations
yearStart := gotime.YearStart()       // January 1st of current year
yearEnd := gotime.YearEnd()           // December 31st of current year
lastYear := gotime.LastYear()
nextYear := gotime.NextYear()
```

### Date Arithmetic

```go
// Add/subtract days
tenDaysLater := gotime.Days(10)          // 10 days from now
twoDaysAgo := gotime.Days(-2)            // 2 days ago
tenDaysFromDate := gotime.Days(10, someDate) // 10 days from specific date

// Add/subtract weeks
twoWeeksLater := gotime.Weeks(2)         // 2 weeks from now
lastWeek := gotime.Weeks(-1)             // 1 week ago

// Add/subtract months
threeMonthsLater := gotime.Months(3)     // 3 months from now
sixMonthsAgo := gotime.Months(-6)        // 6 months ago

// Add/subtract years
nextYear := gotime.Years(1)              // 1 year from now
fiveYearsAgo := gotime.Years(-5)         // 5 years ago
```

### Start and End of Periods

```go
// Day boundaries
startOfDay := gotime.SoD()                    // Start of today (00:00:00)
endOfDay := gotime.EoD()                      // End of today (23:59:59)
startOfSpecificDay := gotime.SoD(someDate)    // Start of specific day

// Week boundaries with custom start day
mondayStart := gotime.WeekStartOn(time.Monday)
sundayEnd := gotime.WeekEndOn(time.Sunday)

// Month and year boundaries for specific dates
monthStartForDate := gotime.MonthStart(someDate)
yearEndForDate := gotime.YearEnd(someDate)
```

## Time Ago Functionality

### Basic Usage

```go
// Simple time ago
fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
timeAgo := gotime.TimeAgo(fiveMinutesAgo)     // "5 minutes ago"

// Future time
twoHoursLater := time.Now().Add(2 * time.Hour)
timeFromNow := gotime.TimeAgo(twoHoursLater)  // "In 2 hours"

// Using custom base time
baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)
pastTime := baseTime.Add(-30 * time.Minute)
relative := gotime.TimeAgo(pastTime, baseTime) // "30 minutes ago"
```

### Real-world Examples

```go
// For user interfaces
func formatLastActivity(lastSeen time.Time) string {
    return gotime.TimeAgo(lastSeen)
}

// For logs with relative time
func logEvent(message string, eventTime time.Time) {
    relative := gotime.TimeAgo(eventTime)
    fmt.Printf("[%s] %s\n", relative, message)
}

// For API responses
type User struct {
    Name       string `json:"name"`
    LastActive string `json:"last_active"`
}

func formatUser(name string, lastActive time.Time) User {
    return User{
        Name:       name,
        LastActive: gotime.TimeAgo(lastActive),
    }
}
```

## Working with Date Ranges

### Date Comparison

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
check := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

// Check if date is between two dates
inRange := gotime.IsBetween(check, start, end)        // true

// Date-only comparison (ignores time)
inDateRange := gotime.IsBetweenDates(check, start, end) // true
```

### Finding Earliest/Latest Dates

```go
date1 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
date3 := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

earliest := gotime.Earliest(date1, date2, date3)     // date1
latest := gotime.Latest(date1, date2, date3)         // date3

// Works with any number of dates
earliest = gotime.Earliest(date1, date2)
latest = gotime.Latest(date1, date2, date3, time.Now())
```

## Utility Functions

### Date Information

```go
// Leap year checking
isLeap2024 := gotime.IsLeapYear(2024)        // true
isLeap2025 := gotime.IsLeapYear(2025)        // false

// Days in month/year
daysInFeb := gotime.DaysInMonth(2024, 2)     // 29 (leap year)
daysIn2024 := gotime.DaysInYear(2024)        // 366

// Quarter information
q1Days := gotime.DaysInQuarter(2024, 1)      // 91 (Jan+Feb+Mar in leap year)
```

### Date Construction

```go
// Create date with timezone
utc := time.UTC
est, _ := time.LoadLocation("America/New_York")

date1 := gotime.NewDate(2025, 7, 7, utc)     // July 7, 2025 in UTC
date2 := gotime.NewDate(2025, 7, 7, est)     // July 7, 2025 in EST

// Create time (date will be zero value)
timeOnly := gotime.NewTime(14, 30, 0, utc)   // 14:30:00 UTC
```

### Date/Time Manipulation

```go
original := time.Date(2025, 7, 7, 14, 30, 45, 0, time.UTC)

// Replace date part, keep time
newDate := gotime.ReplaceDate(original, 2026, 8, 15)
// Result: 2026-08-15 14:30:45 UTC

// Replace time part, keep date
newTime := gotime.ReplaceTime(original, 9, 15, 30)
// Result: 2025-07-07 09:15:30 UTC

// Truncate operations
dateOnly := gotime.TruncateTime(original)    // 2025-07-07 00:00:00 UTC
```

## Business Day Calculations

### Basic Business Days

```go
// Define working days (Monday to Friday)
workingDays := [7]bool{false, true, true, true, true, true, false}
//                      Sun   Mon  Tue  Wed  Thu  Fri  Sat

// Calculate business days from today
businessDate, err := gotime.WorkDay(time.Now(), 10, workingDays)

// Calculate previous business days
prevDate, err := gotime.PrevWorkDay(time.Now(), 5, workingDays)

// Count business days between dates
count, err := gotime.NetWorkDays(startDate, endDate, workingDays)
```

### With Holidays

```go
// Define holidays
holidays := []time.Time{
    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),   // New Year's Day
    time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),   // Independence Day
    time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
}

// Calculate 10 business days excluding holidays
businessDate, err := gotime.WorkDay(time.Now(), 10, workingDays, holidays...)

// Count net working days excluding holidays
count, err := gotime.NetWorkDays(start, end, workingDays, holidays...)
```

## Common Patterns

### 1. API Date Handling

```go
type APIResponse struct {
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

func formatAPIResponse(created, updated time.Time) APIResponse {
    return APIResponse{
        CreatedAt: gotime.Format(created, "yyyy-mm-dd hh:ii:ss"),
        UpdatedAt: gotime.Format(updated, "yyyy-mm-dd hh:ii:ss"),
    }
}

func parseAPIRequest(data APIResponse) (time.Time, time.Time, error) {
    created, err := gotime.Parse(data.CreatedAt, "yyyy-mm-dd hh:ii:ss")
    if err != nil {
        return time.Time{}, time.Time{}, err
    }

    updated, err := gotime.Parse(data.UpdatedAt, "yyyy-mm-dd hh:ii:ss")
    if err != nil {
        return time.Time{}, time.Time{}, err
    }

    return created, updated, nil
}
```

### 2. User-Friendly Display

```go
func displayEvent(event Event) string {
    // Format for user display
    date := gotime.Format(event.Date, "wwww, mmmm dt")
    time := gotime.Format(event.Date, "h:ii aa")

    return fmt.Sprintf("%s at %s", date, time)
    // Example: "Monday, July 7th at 2:30 PM"
}

func displayRelativeTime(timestamp time.Time) string {
    // Show relative time for recent events
    if time.Since(timestamp) < 24*time.Hour {
        return gotime.TimeAgo(timestamp)
    }

    // Show absolute date for older events
    return gotime.Format(timestamp, "mmmm dt, yyyy")
}
```

### 3. Configuration and Settings

```go
type AppSettings struct {
    DateFormat    string
    TimeFormat    string
    ShowRelative  bool
}

func (s *AppSettings) FormatDate(t time.Time) string {
    if s.ShowRelative && time.Since(t) < 7*24*time.Hour {
        return gotime.TimeAgo(t)
    }
    return gotime.Format(t, s.DateFormat)
}

func (s *AppSettings) FormatTime(t time.Time) string {
    return gotime.Format(t, s.TimeFormat)
}

// Usage
settings := &AppSettings{
    DateFormat:   "yyyy-mm-dd",
    TimeFormat:   "hh:ii aa",
    ShowRelative: true,
}

display := settings.FormatDate(someDate)
```

## Error Handling

### Graceful Degradation

```go
func safeFormat(t time.Time, format string) string {
    if t.IsZero() {
        return "N/A"
    }
    return gotime.Format(t, format)
}

func safeParse(input, format string) time.Time {
    if parsed, err := gotime.Parse(input, format); err == nil {
        return parsed
    }
    return time.Time{} // Return zero value on error
}
```

### Validation

```go
func validateDateRange(start, end string, format string) error {
    startDate, err := gotime.Parse(start, format)
    if err != nil {
        return fmt.Errorf("invalid start date: %v", err)
    }

    endDate, err := gotime.Parse(end, format)
    if err != nil {
        return fmt.Errorf("invalid end date: %v", err)
    }

    if endDate.Before(startDate) {
        return fmt.Errorf("end date must be after start date")
    }

    return nil
}
```

---

## Summary

This basic usage guide covered:

1. **NITES fundamentals** - The intuitive format system
2. **Core operations** - Format, Parse, Convert
3. **Date and time handling** - Common patterns and formats
4. **Timezone management** - Parsing and formatting with timezones
5. **Relative operations** - Date arithmetic and time ago
6. **Date ranges** - Comparisons and utilities
7. **Business logic** - Working days and practical calculations
8. **Real-world patterns** - APIs, user interfaces, configuration

**Next Steps:**
- Explore the complete [NITES specification](../core-concepts/nites.md)
- Check out [Common Use Cases](../examples/common-use-cases.md) for more examples
- Read about [Best Practices](../advanced/best-practices.md) for production use

---

Continue to: [NITES Documentation](../core-concepts/nites.md) | [API Reference](../api-reference/)
