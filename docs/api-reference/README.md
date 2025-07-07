[Home](../README.md) > API Reference

# API Reference

Complete documentation for all GoTime functions and their usage patterns.

## Quick Navigation

### Core Operations
- **[Parsing & Formatting](parsing-formatting.md)** - Parse, Format, FormatTimestamp, FormatUnix functions
- **[Date Conversion](conversion.md)** - Convert function and format transformation patterns

### Time Calculations
- **[Relative Time Functions](relative-time.md)** - TimeAgo, Days, Weeks, Months, Years functions
- **[Time Calculations](time-calculations.md)** - Latest, Earliest, Diff, WorkDay, business day functions
- **[Date Range Operations](date-ranges.md)** - IsBetween, range validation, overlap detection

### Utilities & Helpers
- **[Utility Functions](utilities.md)** - TruncateTime, DateValue, timezone helpers, time boundaries
- **[Calendar Math](calendar-math.md)** - DayOfYear, WeekOfMonth, month boundary checks

## Function Categories

### 🔍 Parsing Functions
Parse date/time strings in various formats:
- `Parse(dateString, format)` - Parse with IDFS format
- `ParseInLocation(dateString, format, location)` - Parse in specific timezone

### 🎨 Formatting Functions
Format time.Time values to strings:
- `Format(time, format)` - Format with IDFS format
- `FormatTimestamp(timestamp, format)` - Format Unix timestamp
- `FormatUnix(timestamp, format)` - Format Unix timestamp (alias)

### 🔄 Conversion Functions
Transform between different formats:
- `Convert(dateString, fromFormat, toFormat)` - Convert format while preserving value

### ⏰ Relative Time Functions
Human-readable time differences:
- `TimeAgo(time)` - "2 hours ago", "3 days ago", etc.
- `Days(n, from)` - Add/subtract days
- `Weeks(n, from)` - Add/subtract weeks
- `Months(n, from)` - Add/subtract months
- `Years(n, from)` - Add/subtract years

### ⏱️ Time Arithmetic Functions
Precise time-level arithmetic operations:
- `Hours(n, from)` - Add/subtract hours
- `Minutes(n, from)` - Add/subtract minutes
- `Seconds(n, from)` - Add/subtract seconds

### 📊 Calculation Functions
Advanced time calculations:
- `Latest(times...)` - Find latest time
- `Earliest(times...)` - Find earliest time
- `Diff(time1, time2, unit)` - Calculate difference
- `WorkDay(n, from)` - Add/subtract work days
- `PrevWorkDay(time)` - Find previous work day
- `NetWorkDays(start, end)` - Count business days

### 🎂 Age Calculation Functions
Age and time difference calculations:
- `Age(birthDate)` - Calculate current age
- `YearsBetween(start, end)` - Years between dates
- `MonthsBetween(start, end)` - Months between dates
- `DaysBetween(start, end)` - Days between dates
- `WeeksBetween(start, end)` - Weeks between dates
- `DurationInWords(start, end)` - Human-readable duration
- `IsValidAge(age)` - Validate age reasonableness

### 📊 Quarter Functions
Quarterly time operations for business applications:
- `QuarterStart(time)` - Start of quarter
- `QuarterEnd(time)` - End of quarter
- `LastQuarter(time)` - Previous quarter start
- `NextQuarter(time)` - Next quarter start
- `Quarters(start, end)` - List quarters in range
- `QuarterOfYear(time)` - Get quarter number (1-4)

### 📅 Range Functions
Date range operations:
- `IsBetween(time, start, end)` - Check if time is in range
- `IsBetweenDates(time, start, end)` - Check if date is in range (ignoring time)

### 🛠️ Utility Functions
Helper functions for common tasks:
- `TruncateTime(time)` - Remove time component
- `DateValue(input)` - Convert various types to time.Time
- `SoD(time)` / `EoD(time)` - Start/end of day
- `WeekStart(time)` / `WeekEnd(time)` - Week boundaries
- `MonthStart(time)` / `MonthEnd(time)` - Month boundaries
- `IsLeapYear(year)` - Check leap year
- `DaysInMonth(year, month)` - Days in month
- `IsWeekend(time)` / `IsWeekday(time)` - Day type checking

### 📊 Calendar Math Functions
Calendar calculation utilities:
- `DayOfYear(time)` - Day number within year (1-366)
- `WeekOfMonth(time)` - Week number within month (1-5)
- `IsFirstDayOfMonth(time)` - Check if first day of month
- `IsLastDayOfMonth(time)` - Check if last day of month

## Common Patterns

### Basic Usage
```go
// Parse and format
date, _ := gotime.Parse("2025-07-07", "yyyy-mm-dd")
formatted := gotime.Format(date, "wwww, mmmm dd, yyyy")

// Convert formats
converted, _ := gotime.Convert("07/07/2025", "mm/dd/yyyy", "yyyy-mm-dd")

// Relative time
relative := gotime.TimeAgo(date)

// Time arithmetic
futureTime := gotime.Hours(3, time.Now())    // 3 hours from now
pastTime := gotime.Minutes(-30, time.Now())  // 30 minutes ago
```

### Advanced Operations
```go
// Business day calculations
workDay := gotime.WorkDay(5, time.Now())         // 5 business days from now
businessDays := gotime.NetWorkDays(start, end)   // Count business days

// Age calculations
age := gotime.Age(birthDate)                     // Current age
years := gotime.YearsBetween(start, end)         // Years between dates
duration := gotime.DurationInWords(start, end)   // "2 years, 3 months"

// Quarter operations
quarterStart := gotime.QuarterStart(time.Now()) // Start of current quarter
quarters := gotime.Quarters(start, end)         // All quarters in range

// Range operations
inRange := gotime.IsBetween(check, start, end)   // Check if in range
latest := gotime.Latest(time1, time2, time3)    // Find latest time

// Date boundaries
dayStart := gotime.SoD(date)                     // Start of day
monthEnd := gotime.MonthEnd(date)                // End of month
```

### Error Handling
```go
date, err := gotime.Parse("invalid", "yyyy-mm-dd")
if err != nil {
    // Handle parsing error
    fmt.Printf("Parse error: %v\n", err)
}

converted, err := gotime.Convert("2025-13-01", "yyyy-mm-dd", "mm/dd/yyyy")
if err != nil {
    // Handle invalid date
    fmt.Printf("Invalid date: %v\n", err)
}
```

## IDFS (Intuitive Date Format Specifiers)

GoTime uses human-readable format specifiers instead of Go's reference time approach:

| IDFS | Description | Example |
|------|-------------|---------|
| `yyyy` | 4-digit year | 2025 |
| `yy` | 2-digit year | 25 |
| `mm` | Month (zero-padded) | 07 |
| `mmm` | Month abbreviation | Jul |
| `mmmm` | Full month name | July |
| `dd` | Day (zero-padded) | 07 |
| `d` | Day (no padding) | 7 |
| `hh` | Hour (24-hour, zero-padded) | 14 |
| `h` | Hour (24-hour, no padding) | 14 |
| `ii` | Minutes (zero-padded) | 30 |
| `i` | Minutes (no padding) | 30 |
| `ss` | Seconds (zero-padded) | 45 |
| `s` | Seconds (no padding) | 45 |
| `www` | Weekday abbreviation | Mon |
| `wwww` | Full weekday name | Monday |

See [Core Concepts - IDFS](../core-concepts/idfs.md) for complete reference.

## Performance Tips

1. **Cache parsed formats** for repeated operations
2. **Use TruncateTime()** for date-only comparisons
3. **Batch timezone conversions** when possible
4. **Pre-validate input formats** in data processing pipelines
5. **Use IsBetweenDates()** for date ranges ignoring time

## Error Handling Best Practices

1. **Always check errors** from Parse and Convert functions
2. **Validate date ranges** before processing
3. **Use default values** for optional parameters
4. **Log parsing errors** with context for debugging
5. **Implement fallback formats** for flexible parsing

---

## Getting Help

- **[Getting Started Guide](../getting-started/)** - Quick introduction
- **[Core Concepts](../core-concepts/)** - Understanding GoTime principles
- **[Examples & Use Cases](../examples/)** - Real-world implementations
- **[GitHub Issues](https://github.com/maniartech/gotime/issues)** - Bug reports and feature requests

---

*Navigate to specific function documentation using the links above, or browse the complete API reference for detailed examples and use cases.*
