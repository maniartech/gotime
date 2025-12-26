[Home](../README.md) > [Getting Started](README.md) > Quick Start

# Quick Start Guide

Get started with GoTime in just a few minutes! This guide will walk you through the basic setup and most common operations.

## Installation

```bash
go get github.com/maniartech/gotime/v2@v2.0.2
```

## Basic Operations

### Import the Package

```go
import "github.com/maniartech/gotime/v2"
```

### 1. Format Dates with Human-Friendly Specifiers

```go
package main

import (
    "fmt"
    "time"
    "github.com/maniartech/gotime/v2"
)

func main() {
    dt := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

    // Format using intuitive specifiers
    formatted := gotime.Format(dt, "dt mmmm, yyyy")
    fmt.Println(formatted) // Output: 1st January, 2012

    // Different formats
    fmt.Println(gotime.Format(dt, "yyyy-mm-dd"))    // 2012-01-01
    fmt.Println(gotime.Format(dt, "dd/mm/yyyy"))    // 01/01/2012
    fmt.Println(gotime.Format(dt, "wwww, mmmm dt")) // Sunday, January 1st
}
```

### 2. Parse Dates Easily

```go
// Parse dates using the same intuitive format
date, err := gotime.Parse("2012-01-01", "yyyy-mm-dd")
if err != nil {
    log.Fatal(err)
}
fmt.Println(date) // 2012-01-01 00:00:00 +0000 UTC

// Parse with timezone
tz := time.FixedZone("IST", 5*60*60+30*60) // UTC+5:30
date, err = gotime.ParseInLocation("01/01/2020", "dd/mm/yyyy", tz)
if err != nil {
    log.Fatal(err)
}
fmt.Println(date) // 2020-01-01 00:00:00 +0530 IST
```

### 3. Convert Between Date Formats

```go
// Convert from one format to another
converted, err := gotime.Convert("2012-01-01", "yyyy-mm-dd", "wwww, dt mmmm, yyyy")
if err != nil {
    log.Fatal(err)
}
fmt.Println(converted) // Sunday, 1st January, 2012
```

### 4. Get Relative Time

```go
// Time ago functionality
pastTime := time.Now().Add(-5 * time.Minute)
timeAgo := gotime.TimeAgo(pastTime)
fmt.Println(timeAgo) // "5 minutes ago"

// Future time
futureTime := time.Now().Add(2 * time.Hour)
timeAgo = gotime.TimeAgo(futureTime)
fmt.Println(timeAgo) // "In 2 hours"
```

### 5. Quick Date Calculations

```go
// Get common dates
fmt.Println(gotime.Yesterday())  // Yesterday's date
fmt.Println(gotime.Tomorrow())   // Tomorrow's date
fmt.Println(gotime.NextWeek())   // Date one week from now

// Date arithmetic
d1 := gotime.Days(10)     // 10 days from now
d2 := gotime.Days(-2)     // 2 days ago
d3 := gotime.Months(3)    // 3 months from now
```

## Most Common Use Cases

### 1. Display User-Friendly Dates

```go
func formatUserDate(t time.Time) string {
    return gotime.Format(t, "wwww, mmmm dt, yyyy")
}

// Usage
userDate := formatUserDate(time.Now())
// Output: "Monday, July 7th, 2025"
```

### 2. Parse User Input

```go
func parseUserInput(dateStr string) (time.Time, error) {
    // Try common formats
    formats := []string{
        "yyyy-mm-dd",
        "dd/mm/yyyy",
        "mm/dd/yyyy",
        "mmmm dt, yyyy",
    }

    for _, format := range formats {
        if date, err := gotime.Parse(dateStr, format); err == nil {
            return date, nil
        }
    }
    return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}
```

### 3. API Response Formatting

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
```

### 4. Working with Business Days

```go
// Define working days (Monday to Friday)
workingDays := [7]bool{false, true, true, true, true, true, false}

// Calculate 10 business days from today
businessDate, err := gotime.WorkDay(time.Now(), 10, workingDays)
if err != nil {
    log.Fatal(err)
}

// Count business days between two dates
count, err := gotime.NetWorkDays(start, end, workingDays)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Business days: %d\n", count)
```

## Key Advantages

1. **Intuitive**: Format specifiers like `yyyy-mm-dd` instead of `2006-01-02`
2. **Case Insensitive**: No need to remember upper/lowercase rules
3. **Comprehensive**: All common date operations in one package
4. **Compatible**: Works with Go's standard `time` package
5. **Zero Dependencies**: No external dependencies required

## Next Steps

- Learn about [NITES (Natural and Intuitive Time Expression Syntax)](../core-concepts/nites.md)
- Explore the complete [API Reference](../api-reference/)
- Check out [Common Use Cases](../examples/common-use-cases.md)
- Understand [Why GoTime?](../core-concepts/why-gotime.md)

---

Ready to dive deeper? Continue with the [Basic Usage Guide](basic-usage.md) or jump to any section that interests you!
