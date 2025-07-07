# Calendar Math Functions

## Overview

The GoTime library provides efficient calendar math utilities for common date calculations. These functions follow Go's standard library patterns and provide excellent performance with zero memory allocations.

## Functions

### DayOfYear

```go
func DayOfYear(t time.Time) int
```

Returns the day number (1-366) within the year for the given date.

**Performance:** ~7.4 ns/op, 0 allocations

**Examples:**
```go
// New Year's Day
gotime.DayOfYear(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))  // Returns 1

// Independence Day
gotime.DayOfYear(time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC))  // Returns 185

// New Year's Eve (non-leap year)
gotime.DayOfYear(time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)) // Returns 365

// New Year's Eve (leap year)
gotime.DayOfYear(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)) // Returns 366
```

### WeekOfMonth

```go
func WeekOfMonth(t time.Time) int
```

Returns the week number (1-5) within the month for the given date. The first week starts on the 1st of the month, regardless of the day of the week.

**Performance:** ~31.8 ns/op, 0 allocations

**Examples:**
```go
// First week of July 2025
gotime.WeekOfMonth(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))  // Returns 1
gotime.WeekOfMonth(time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC))  // Returns 1

// Second week of July 2025
gotime.WeekOfMonth(time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC))  // Returns 2
gotime.WeekOfMonth(time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC)) // Returns 2

// Last week of July 2025
gotime.WeekOfMonth(time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC)) // Returns 5
```

### IsFirstDayOfMonth

```go
func IsFirstDayOfMonth(t time.Time) bool
```

Returns true if the date is the first day of its month.

**Performance:** ~6.0 ns/op, 0 allocations

**Examples:**
```go
// First day of month
gotime.IsFirstDayOfMonth(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))  // Returns true

// Not first day
gotime.IsFirstDayOfMonth(time.Date(2025, 7, 2, 0, 0, 0, 0, time.UTC))  // Returns false
gotime.IsFirstDayOfMonth(time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC)) // Returns false
```

### IsLastDayOfMonth

```go
func IsLastDayOfMonth(t time.Time) bool
```

Returns true if the date is the last day of its month. This function correctly handles leap years and varying month lengths.

**Performance:** ~36.2 ns/op, 0 allocations

**Examples:**
```go
// Last day of February (non-leap year)
gotime.IsLastDayOfMonth(time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC))  // Returns true
gotime.IsLastDayOfMonth(time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC))  // Returns false

// Last day of February (leap year)
gotime.IsLastDayOfMonth(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC))  // Returns true
gotime.IsLastDayOfMonth(time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC))  // Returns false

// Last day of July
gotime.IsLastDayOfMonth(time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC))  // Returns true
gotime.IsLastDayOfMonth(time.Date(2025, 7, 30, 0, 0, 0, 0, time.UTC))  // Returns false

// Last day of April (30 days)
gotime.IsLastDayOfMonth(time.Date(2025, 4, 30, 0, 0, 0, 0, time.UTC))  // Returns true
gotime.IsLastDayOfMonth(time.Date(2025, 4, 29, 0, 0, 0, 0, time.UTC))  // Returns false
```

## Performance Characteristics

All calendar math functions are highly optimized:

| Function | Performance | Memory | Use Case |
|----------|-------------|---------|----------|
| `DayOfYear` | ~7.4 ns/op | 0 allocs | Day numbering, progress tracking |
| `WeekOfMonth` | ~31.8 ns/op | 0 allocs | Calendar layouts, week-based reporting |
| `IsFirstDayOfMonth` | ~6.0 ns/op | 0 allocs | Month boundary detection |
| `IsLastDayOfMonth` | ~36.2 ns/op | 0 allocs | Month-end processing, billing cycles |

## Common Use Cases

### Progress Tracking
```go
now := time.Now()
dayOfYear := gotime.DayOfYear(now)
daysInYear := 365
if isLeapYear(now.Year()) {
    daysInYear = 366
}
progress := float64(dayOfYear) / float64(daysInYear) * 100
fmt.Printf("%.1f%% of the year completed", progress)
```

### Week-Based Reports
```go
date := time.Now()
week := gotime.WeekOfMonth(date)
fmt.Printf("Week %d of %s", week, date.Month())
```

### Month Boundary Processing
```go
date := time.Now()
if gotime.IsFirstDayOfMonth(date) {
    // Start of month processing
    processMonthlyReports()
}
if gotime.IsLastDayOfMonth(date) {
    // End of month processing
    generateBillingCycles()
}
```

## Design Principles

1. **Zero Allocations:** All functions avoid memory allocations for optimal performance
2. **Standard Library Alignment:** Uses Go's time package patterns and conventions
3. **Edge Case Handling:** Properly handles leap years, varying month lengths, and timezone considerations
4. **Business Logic Focused:** Designed for common business and calendar applications
