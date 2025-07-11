[Home](../README.md) > [API Reference](README.md) > Relative Time

# Relative Time Functions

GoTime provides comprehensive functions for working with relative dates and times, making it easy to calculate dates relative to now or any given point in time.

## Core Relative Time Function

### TimeAgo

Returns human-readable relative time descriptions.

```go
func TimeAgo(t time.Time, baseTime ...time.Time) string
```

**Parameters:**
- `t`: The time to compare
- `baseTime`: Optional base time for comparison (defaults to `time.Now()`)

**Returns:**
- `string`: Human-readable relative time description

**Examples:**

```go
now := time.Now()

// Past times
fiveMinutesAgo := now.Add(-5 * time.Minute)
fmt.Println(gotime.TimeAgo(fiveMinutesAgo))     // "5 minutes ago"

oneHourAgo := now.Add(-1 * time.Hour)
fmt.Println(gotime.TimeAgo(oneHourAgo))         // "1 hour ago"

yesterday := now.Add(-24 * time.Hour)
fmt.Println(gotime.TimeAgo(yesterday))          // "Yesterday"

oneWeekAgo := now.Add(-7 * 24 * time.Hour)
fmt.Println(gotime.TimeAgo(oneWeekAgo))         // "Last week"

// Future times
fiveMinutesLater := now.Add(5 * time.Minute)
fmt.Println(gotime.TimeAgo(fiveMinutesLater))   // "In 5 minutes"

twoHoursLater := now.Add(2 * time.Hour)
fmt.Println(gotime.TimeAgo(twoHoursLater))      // "In 2 hours"

tomorrow := now.Add(24 * time.Hour)
fmt.Println(gotime.TimeAgo(tomorrow))           // "Tomorrow"

// With custom base time
baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)
pastTime := baseTime.Add(-30 * time.Minute)
fmt.Println(gotime.TimeAgo(pastTime, baseTime)) // "30 minutes ago"
```

**TimeAgo Output Examples:**

| Time Difference | Output |
|----------------|--------|
| < 10 seconds | "Just now" / "In a few seconds" |
| < 60 seconds | "A minute ago" / "In a minute" |
| < 3600 seconds | "Few minutes ago" / "In a few minutes" |
| Yesterday/Tomorrow | "Yesterday" / "Tomorrow" |
| 1 hour | "1 hour ago" / "In 1 hour" |
| 2-23 hours | "X hours ago" / "In X hours" |
| 1 day | "Yesterday" / "Tomorrow" |
| 2-6 days | "X days ago" / "In X days" |
| 1 week | "Last week" / "In a week" |
| 2-4 weeks | "X weeks ago" / "In X weeks" |
| 1 month | "Last month" / "In a month" |
| 2-11 months | "X months ago" / "In X months" |
| 1 year | "Last year" / "In a year" |
| 2+ years | "X years ago" / "In X years" |

## Date Arithmetic Functions

### Days

Add or subtract days from a date.

```go
func Days(days int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 10 days from now
future := gotime.Days(10)

// 5 days ago
past := gotime.Days(-5)

// 10 days from specific date
specificDate := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
futureFromSpecific := gotime.Days(10, specificDate)

// Zero is a no-op (returns the same date)
sameDay := gotime.Days(0, specificDate) // Returns specificDate unchanged
```

### Weeks

Add or subtract weeks from a date.

```go
func Weeks(weeks int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 2 weeks from now
twoWeeksLater := gotime.Weeks(2)

// 3 weeks ago
threeWeeksAgo := gotime.Weeks(-3)

// From specific date
nextMonth := gotime.Weeks(4, specificDate)

// Zero is a no-op (returns the same date)
sameWeek := gotime.Weeks(0, specificDate) // Returns specificDate unchanged
```

### Months

Add or subtract months from a date.

```go
func Months(months int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 3 months from now
quarterly := gotime.Months(3)

// 6 months ago
halfYearAgo := gotime.Months(-6)

// Handle month-end edge cases properly
endOfJan := time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC)
endOfFeb := gotime.Months(1, endOfJan) // Properly handles Feb 28/29

// Zero is a no-op (returns the same date)
sameMonth := gotime.Months(0, specificDate) // Returns specificDate unchanged
```

### Years

Add or subtract years from a date.

```go
func Years(years int, dt ...time.Time) time.Time
```

**Examples:**

```go
// Next year
nextYear := gotime.Years(1)

// 5 years ago
fiveYearsAgo := gotime.Years(-5)

// Leap year handling
leapDay := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
nextYearFromLeap := gotime.Years(1, leapDay) // Handles Feb 29 -> Feb 28

// Zero is a no-op (returns the same date)
sameYear := gotime.Years(0, specificDate) // Returns specificDate unchanged
```

## Time Arithmetic Functions

### Hours

Add or subtract hours from a time.

```go
func Hours(hours int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 3 hours from now
threeHoursLater := gotime.Hours(3)

// 2 hours ago
twoHoursAgo := gotime.Hours(-2)

// From specific time
specificTime := time.Date(2025, 7, 7, 14, 30, 0, 0, time.UTC)
futureTime := gotime.Hours(5, specificTime) // 2025-07-07 19:30:00

// Multiple hours
fullWorkDay := gotime.Hours(8, specificTime) // Add full work day
```

### Minutes

Add or subtract minutes from a time.

```go
func Minutes(minutes int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 30 minutes from now
halfHourLater := gotime.Minutes(30)

// 15 minutes ago
quarterHourAgo := gotime.Minutes(-15)

// Precise scheduling
meetingTime := time.Date(2025, 7, 7, 14, 0, 0, 0, time.UTC)
reminderTime := gotime.Minutes(-5, meetingTime) // 5 minutes before meeting
followupTime := gotime.Minutes(60, meetingTime) // 1 hour after meeting
```

### Seconds

Add or subtract seconds from a time.

```go
func Seconds(seconds int, dt ...time.Time) time.Time
```

**Examples:**

```go
// 30 seconds from now
halfMinuteLater := gotime.Seconds(30)

// 10 seconds ago
tenSecondsAgo := gotime.Seconds(-10)

// Precise timing
processStart := time.Now()
timeout := gotime.Seconds(30, processStart) // 30-second timeout
checkpoint := gotime.Seconds(5, processStart) // 5-second checkpoint
```

## Quick Date Shortcuts

### Day Functions

```go
// Get specific days
yesterday := gotime.Yesterday()    // Yesterday's date
tomorrow := gotime.Tomorrow()      // Tomorrow's date

// Start and end of day
startOfToday := gotime.SoD()       // Today at 00:00:00
endOfToday := gotime.EoD()         // Today at 23:59:59

// For specific date
specificDay := time.Date(2025, 7, 7, 14, 30, 0, 0, time.UTC)
startOfDay := gotime.SoD(specificDay)  // 2025-07-07 00:00:00
endOfDay := gotime.EoD(specificDay)    // 2025-07-07 23:59:59
```

### Week Functions

```go
// Basic week functions
lastWeek := gotime.LastWeek()      // 7 days ago
nextWeek := gotime.NextWeek()      // 7 days from now

// Week boundaries (default: Sunday-Saturday)
weekStart := gotime.WeekStart()    // Start of current week
weekEnd := gotime.WeekEnd()        // End of current week

// Custom week start day
mondayStart := gotime.WeekStartOn(time.Monday)
sundayEnd := gotime.WeekEndOn(time.Sunday)

// For specific date
specificDate := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
weekStartForDate := gotime.WeekStart(specificDate)
mondayStartForDate := gotime.WeekStartOn(time.Monday, specificDate)
```

### Month Functions

```go
// Basic month functions
lastMonth := gotime.LastMonth()    // Same day last month
nextMonth := gotime.NextMonth()    // Same day next month

// Month boundaries
monthStart := gotime.MonthStart()  // First day of current month
monthEnd := gotime.MonthEnd()      // Last day of current month

// For specific date
specificDate := time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC)
monthStartForDate := gotime.MonthStart(specificDate) // 2025-07-01
monthEndForDate := gotime.MonthEnd(specificDate)     // 2025-07-31
```

### Year Functions

```go
// Basic year functions
lastYear := gotime.LastYear()      // Same date last year
nextYear := gotime.NextYear()      // Same date next year

// Year boundaries
yearStart := gotime.YearStart()    // January 1st of current year
yearEnd := gotime.YearEnd()        // December 31st of current year

// For specific date
specificDate := time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC)
yearStartForDate := gotime.YearStart(specificDate)   // 2025-01-01 00:00:00

yearEndForDate := gotime.YearEnd(specificDate)       // 2025-12-31 23:59:59
```

## Business Calendar Utilities

### IsBusinessDay
```go
func IsBusinessDay(date time.Time, weekends []time.Weekday, holidays ...time.Time) bool
```
Returns true if the given date is a business day (not a weekend or holiday). Supports custom weekends and holidays.

**Example:**
```go
weekends := []time.Weekday{time.Saturday, time.Sunday}
holidays := []time.Time{
    time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Independence Day
}
gotime.IsBusinessDay(time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), weekends, holidays...) // true
```

### NextBusinessDay
```go
func NextBusinessDay(date time.Time, weekends []time.Weekday, holidays ...time.Time) time.Time
```
Returns the next business day after the given date, skipping weekends and holidays.

**Example:**
```go
gotime.NextBusinessDay(time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), weekends, holidays...) // 2025-07-07 (Monday)
```

### PrevBusinessDay
```go
func PrevBusinessDay(date time.Time, weekends []time.Weekday, holidays ...time.Time) time.Time
```
Returns the previous business day before the given date, skipping weekends and holidays.

**Example:**
```go
gotime.PrevBusinessDay(time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), weekends, holidays...) // 2025-07-03 (Thursday)
```

## Real-World Use Cases

### 1. User Interface Timestamps

```go
package ui

import (
    "github.com/maniartech/gotime"
    "time"
)

func formatTimestamp(t time.Time) string {
    now := time.Now()
    diff := now.Sub(t)

    // Use relative time for recent events
    if diff < 24*time.Hour {
        return gotime.TimeAgo(t)
    }

    // Use absolute time for older events
    if diff < 7*24*time.Hour {
        return gotime.Format(t, "wwww at h:ii aa")
    }

    // Use date only for very old events
    return gotime.Format(t, "mmmm dt, yyyy")
}

// Examples:
// 5 minutes ago: "5 minutes ago"
// Yesterday: "Yesterday"
// This week: "Monday at 2:30 PM"
// Older: "July 7th, 2025"
```

### 2. Meeting and Appointment Scheduling

```go
package scheduler

import (
    "github.com/maniartech/gotime"
    "time"
)

type Meeting struct {
    Title     string
    StartTime time.Time
    Duration  time.Duration
}

func (m *Meeting) EndTime() time.Time {
    minutes := int(m.Duration.Minutes())
    return gotime.Minutes(minutes, m.StartTime)
}

func (m *Meeting) ReminderTimes() []time.Time {
    return []time.Time{
        gotime.Days(-1, m.StartTime),    // 1 day before
        gotime.Hours(-2, m.StartTime),   // 2 hours before
        gotime.Minutes(-15, m.StartTime), // 15 minutes before
        gotime.Minutes(-5, m.StartTime),  // 5 minutes before
    }
}

func (m *Meeting) BufferTime() (start, end time.Time) {
    start = gotime.Minutes(-15, m.StartTime) // 15 min setup
    end = gotime.Minutes(15, m.EndTime())    // 15 min cleanup
    return
}

// Schedule back-to-back meetings with buffers
func scheduleConsecutiveMeetings(meetings []Meeting) {
    for i := 1; i < len(meetings); i++ {
        prevEnd := meetings[i-1].EndTime()
        buffer := gotime.Minutes(30, prevEnd) // 30-minute buffer

        if meetings[i].StartTime.Before(buffer) {
            meetings[i].StartTime = buffer
        }
    }
}
```

### 3. Content Scheduling

```go
package scheduler

import (
    "github.com/maniartech/gotime"
    "time"
)

type ContentScheduler struct {
    timezone *time.Location
}

func NewContentScheduler(tz *time.Location) *ContentScheduler {
    return &ContentScheduler{timezone: tz}
}

func (cs *ContentScheduler) SchedulePost(content string, when string) error {
    var scheduledTime time.Time
    var err error

    switch when {
    case "tomorrow":
        scheduledTime = gotime.Tomorrow()
    case "next_week":
        scheduledTime = gotime.NextWeek()
    case "next_month":
        scheduledTime = gotime.NextMonth()
    case "in_2_hours":
        scheduledTime = gotime.Hours(2)
    case "in_30_minutes":
        scheduledTime = gotime.Minutes(30)
    default:
        // Parse custom date
        scheduledTime, err = gotime.ParseInLocation(when, "yyyy-mm-dd", cs.timezone)
        if err != nil {
            return err
        }
    }

    // Schedule the content
    return cs.scheduleAt(content, scheduledTime)
}

func (cs *ContentScheduler) scheduleAt(content string, when time.Time) error {
    // Implementation details...
    return nil
}
```

### 4. Process Monitoring and Timeouts

```go
package monitoring

import (
    "github.com/maniartech/gotime"
    "time"
    "context"
)

type ProcessMonitor struct {
    processes map[string]*Process
}

type Process struct {
    ID        string
    StartTime time.Time
    Timeout   time.Duration
    Status    string
}

func (pm *ProcessMonitor) StartProcess(id string, timeout time.Duration) {
    now := time.Now()
    process := &Process{
        ID:        id,
        StartTime: now,
        Timeout:   timeout,
        Status:    "running",
    }

    pm.processes[id] = process

    // Set up timeout
    timeoutTime := gotime.Seconds(int(timeout.Seconds()), now)
    go pm.checkTimeout(id, timeoutTime)
}

func (pm *ProcessMonitor) checkTimeout(processID string, timeoutTime time.Time) {
    for {
        if time.Now().After(timeoutTime) {
            if process, exists := pm.processes[processID]; exists && process.Status == "running" {
                process.Status = "timeout"
            }
            break
        }
        time.Sleep(1 * time.Second)
    }
}

func (pm *ProcessMonitor) GetProcessInfo(id string) map[string]interface{} {
    process := pm.processes[id]
    if process == nil {
        return nil
    }

    runtime := time.Since(process.StartTime)
    timeoutAt := gotime.Seconds(int(process.Timeout.Seconds()), process.StartTime)

    return map[string]interface{}{
        "id":           process.ID,
        "status":       process.Status,
        "runtime":      gotime.TimeAgo(process.StartTime),
        "timeout_at":   gotime.TimeAgo(timeoutAt),
        "will_timeout": gotime.TimeAgo(timeoutAt),
    }
}
```

### 5. Report Generation

```go
package reports

import (
    "github.com/maniartech/gotime"
    "time"
)

type ReportPeriod struct {
    Start time.Time
    End   time.Time
    Name  string
}

func GetReportPeriods() []ReportPeriod {
    now := time.Now()

    return []ReportPeriod{
        {
            Start: gotime.SoD(),
            End:   gotime.EoD(),
            Name:  "Today",
        },
        {
            Start: gotime.SoD(gotime.Yesterday()),
            End:   gotime.EoD(gotime.Yesterday()),
            Name:  "Yesterday",
        },
        {
            Start: gotime.WeekStart(),
            End:   gotime.WeekEnd(),
            Name:  "This Week",
        },
        {
            Start: gotime.WeekStart(gotime.LastWeek()),
            End:   gotime.WeekEnd(gotime.LastWeek()),
            Name:  "Last Week",
        },
        {
            Start: gotime.MonthStart(),
            End:   gotime.MonthEnd(),
            Name:  "This Month",
        },
        {
            Start: gotime.MonthStart(gotime.LastMonth()),
            End:   gotime.MonthEnd(gotime.LastMonth()),
            Name:  "Last Month",
        },
        {
            Start: gotime.YearStart(),
            End:   gotime.YearEnd(),
            Name:  "This Year",
        },
    }
}

func GenerateReport(period ReportPeriod) *Report {
    // Generate report for the specified period
    return &Report{
        Period:    period.Name,
        StartDate: gotime.Format(period.Start, "yyyy-mm-dd"),
        EndDate:   gotime.Format(period.End, "yyyy-mm-dd"),
        Generated: gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss"),
    }
}
```

```go
package events

import (
    "github.com/maniartech/gotime"
    "time"
)

type Event struct {
    Name      string
    Date      time.Time
    Reminders []time.Time
}

func CreateEvent(name string, date time.Time) *Event {
    event := &Event{
        Name: name,
        Date: date,
    }

    // Set up automatic reminders
    event.Reminders = []time.Time{
        gotime.Days(-7, date),   // 1 week before
        gotime.Days(-1, date),   // 1 day before
        date.Add(-2 * time.Hour), // 2 hours before
    }

    return event
}

func (e *Event) TimeUntilEvent() string {
    return gotime.TimeAgo(e.Date)
}

func (e *Event) GetUpcomingReminders() []string {
    now := time.Now()
    var upcoming []string

    for _, reminder := range e.Reminders {
        if reminder.After(now) {
            timeUntil := gotime.TimeAgo(reminder)
            upcoming = append(upcoming, fmt.Sprintf("Reminder %s", timeUntil))
        }
    }

    return upcoming
}
```

### 5. Cache Expiration

```go
package cache

import (
    "github.com/maniartech/gotime"
    "time"
)

type CacheEntry struct {
    Data      interface{}
    ExpiresAt time.Time
    CreatedAt time.Time
}

func NewCacheEntry(data interface{}, ttl time.Duration) *CacheEntry {
    now := time.Now()
    return &CacheEntry{
        Data:      data,
        ExpiresAt: now.Add(ttl),
        CreatedAt: now,
    }
}

func (ce *CacheEntry) IsExpired() bool {
    return time.Now().After(ce.ExpiresAt)
}

func (ce *CacheEntry) TimeToExpiry() string {
    if ce.IsExpired() {
        return "Expired"
    }
    return gotime.TimeAgo(ce.ExpiresAt)
}

func (ce *CacheEntry) Age() string {
    return gotime.TimeAgo(ce.CreatedAt)
}

// Cache management
type Cache struct {
    entries map[string]*CacheEntry
}

func (c *Cache) CleanupExpired() int {
    cleaned := 0
    for key, entry := range c.entries {
        if entry.IsExpired() {
            delete(c.entries, key)
            cleaned++
        }
    }
    return cleaned
}

func (c *Cache) GetExpiryInfo() map[string]string {
    info := make(map[string]string)
    for key, entry := range c.entries {
        info[key] = entry.TimeToExpiry()
    }
    return info
}
```

### 6. Logging with Relative Times

```go
package logging

import (
    "fmt"
    "github.com/maniartech/gotime"
    "time"
)

type Logger struct {
    startTime time.Time
}

func NewLogger() *Logger {
    return &Logger{
        startTime: time.Now(),
    }
}

func (l *Logger) Log(level, message string) {
    now := time.Now()

    // Show both absolute and relative time
    absolute := gotime.Format(now, "hh:ii:ss.000")
    relative := gotime.TimeAgo(l.startTime, now)

    fmt.Printf("[%s] (%s since start) %s: %s\n",
        absolute, relative, level, message)
}

func (l *Logger) LogEvent(event string, eventTime time.Time) {
    now := time.Now()
    relative := gotime.TimeAgo(eventTime, now)

    fmt.Printf("Event: %s occurred %s\n", event, relative)
}
```

## Advanced Patterns

### 1. Flexible Date Calculation

```go
func calculateDate(base time.Time, offset string) (time.Time, error) {
    switch offset {
    case "yesterday":
        return gotime.Days(-1, base), nil
    case "tomorrow":
        return gotime.Days(1, base), nil
    case "last_week":
        return gotime.Weeks(-1, base), nil
    case "next_week":
        return gotime.Weeks(1, base), nil
    case "last_month":
        return gotime.Months(-1, base), nil
    case "next_month":
        return gotime.Months(1, base), nil
    case "last_year":
        return gotime.Years(-1, base), nil
    case "next_year":
        return gotime.Years(1, base), nil
    case "start_of_week":
        return gotime.WeekStart(base), nil
    case "end_of_week":
        return gotime.WeekEnd(base), nil
    case "start_of_month":
        return gotime.MonthStart(base), nil
    case "end_of_month":
        return gotime.MonthEnd(base), nil
    case "start_of_year":
        return gotime.YearStart(base), nil
    case "end_of_year":
        return gotime.YearEnd(base), nil
    default:
        return time.Time{}, fmt.Errorf("unknown offset: %s", offset)
    }
}
```

### 2. Smart Time Display

```go
func smartTimeDisplay(t time.Time) string {
    now := time.Now()
    diff := now.Sub(t)

    switch {
    case diff < time.Minute:
        return "Just now"
    case diff < time.Hour:
        return gotime.TimeAgo(t)
    case diff < 24*time.Hour:
        return gotime.TimeAgo(t)
    case diff < 7*24*time.Hour:
        return gotime.Format(t, "wwww")
    case t.Year() == now.Year():
        return gotime.Format(t, "mmmm dt")
    default:
        return gotime.Format(t, "mmmm dt, yyyy")
    }
}
```

### 3. Batch Relative Time Processing

```go
func processTimestamps(timestamps []time.Time) []string {
    results := make([]string, len(timestamps))
    now := time.Now()

    for i, ts := range timestamps {
        results[i] = gotime.TimeAgo(ts, now)
    }

    return results
}

func groupByRelativeTime(timestamps []time.Time) map[string][]time.Time {
    groups := make(map[string][]time.Time)
    now := time.Now()

    for _, ts := range timestamps {
        key := gotime.TimeAgo(ts, now)
        groups[key] = append(groups[key], ts)
    }

    return groups
}
```

## Performance Considerations

### 1. Caching Current Time

```go
// For batch operations, cache the current time
func processEvents(events []Event) []ProcessedEvent {
    now := time.Now() // Cache current time

    results := make([]ProcessedEvent, len(events))
    for i, event := range events {
        results[i] = ProcessedEvent{
            Event:       event,
            RelativeTime: gotime.TimeAgo(event.Timestamp, now),
        }
    }

    return results
}
```

### 2. Avoiding Repeated Calculations

```go
// Pre-calculate common relative dates
type RelativeDates struct {
    Yesterday    time.Time
    Tomorrow     time.Time
    LastWeek     time.Time
    NextWeek     time.Time
    MonthStart   time.Time
    MonthEnd     time.Time
}

func NewRelativeDates() *RelativeDates {
    return &RelativeDates{
        Yesterday:  gotime.Yesterday(),
        Tomorrow:   gotime.Tomorrow(),
        LastWeek:   gotime.LastWeek(),
        NextWeek:   gotime.NextWeek(),
        MonthStart: gotime.MonthStart(),
        MonthEnd:   gotime.MonthEnd(),
    }
}
```

---

## Summary

GoTime's relative time functions provide:

- **Human-readable time descriptions** with TimeAgo
- **Intuitive date arithmetic** with Days, Weeks, Months, Years
- **Convenient shortcuts** for common dates (Yesterday, Tomorrow, etc.)
- **Flexible period boundaries** (start/end of day, week, month, year)
- **Real-world utility** for UIs, scheduling, reporting, and logging

These functions make it easy to work with relative times in a natural, readable way, eliminating the need for complex manual calculations.

---

Next: [Time Calculations](time-calculations.md) | [Back to API Reference](../api-reference/)
