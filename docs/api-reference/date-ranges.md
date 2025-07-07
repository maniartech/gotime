[Home](../README.md) > [API Reference](README.md) > Date Ranges

# Date Range Operations

GoTime provides powerful functions for working with date ranges, checking if dates fall within specific periods, and performing range-based validations.

## Core Range Functions

### IsBetween

Checks if a time falls between two other times (inclusive).

```go
func IsBetween(t1, t2, t3 time.Time) bool
```

**Parameters:**
- `t1`: The time to check
- `t2`: Start of range
- `t3`: End of range (order doesn't matter - function will swap if needed)

**Returns:**
- `bool`: True if t1 is between t2 and t3 (inclusive)

**Examples:**

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
check := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

// Check if date is in range
inRange := gotime.IsBetween(check, start, end)  // true

// Order doesn't matter
inRange = gotime.IsBetween(check, end, start)   // true (automatically swapped)

// Exact boundaries
onStart := gotime.IsBetween(start, start, end)  // true (inclusive)
onEnd := gotime.IsBetween(end, start, end)      // true (inclusive)

// Outside range
before := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
isInRange := gotime.IsBetween(before, start, end)  // false
```

### IsBetweenDates

Checks if a date falls between two other dates, ignoring time components.

```go
func IsBetweenDates(t1, t2, t3 time.Time) bool
```

**Key Difference from IsBetween:**
- Compares only the date portion (year, month, day)
- Automatically sets start time to beginning of day (00:00:00)
- Automatically sets end time to end of day (23:59:59)

**Examples:**

```go
// Different times, same dates
start := time.Date(2025, 7, 1, 14, 30, 0, 0, time.UTC)      // July 1, 2:30 PM
end := time.Date(2025, 7, 31, 9, 15, 0, 0, time.UTC)        // July 31, 9:15 AM
check := time.Date(2025, 7, 15, 23, 45, 0, 0, time.UTC)     // July 15, 11:45 PM

// IsBetween considers exact times
exactTime := gotime.IsBetween(check, start, end)            // true

// IsBetweenDates only considers dates
dateOnly := gotime.IsBetweenDates(check, start, end)        // true

// Edge case: same date, different times
sameDay1 := time.Date(2025, 7, 15, 1, 0, 0, 0, time.UTC)   // 1:00 AM
sameDay2 := time.Date(2025, 7, 15, 10, 0, 0, 0, time.UTC)  // 10:00 AM
sameDay3 := time.Date(2025, 7, 15, 20, 0, 0, 0, time.UTC)  // 8:00 PM

// All return true because it's the same date
result1 := gotime.IsBetweenDates(sameDay2, sameDay1, sameDay3)  // true
result2 := gotime.IsBetweenDates(sameDay1, sameDay2, sameDay3)  // true
```

## Practical Range Applications

### 1. Event Scheduling Validation

```go
package events

import (
    "github.com/maniartech/gotime"
    "time"
)

type Event struct {
    ID        string
    Name      string
    StartTime time.Time
    EndTime   time.Time
    Location  string
}

type EventScheduler struct {
    events []Event
}

func NewEventScheduler() *EventScheduler {
    return &EventScheduler{
        events: make([]Event, 0),
    }
}

func (es *EventScheduler) IsTimeSlotAvailable(start, end time.Time) bool {
    for _, event := range es.events {
        // Check if new event overlaps with existing event
        if gotime.IsBetween(start, event.StartTime, event.EndTime) ||
           gotime.IsBetween(end, event.StartTime, event.EndTime) ||
           gotime.IsBetween(event.StartTime, start, end) {
            return false
        }
    }
    return true
}

func (es *EventScheduler) GetEventsInRange(start, end time.Time) []Event {
    var eventsInRange []Event

    for _, event := range es.events {
        // Event overlaps with the query range
        if gotime.IsBetween(event.StartTime, start, end) ||
           gotime.IsBetween(event.EndTime, start, end) ||
           gotime.IsBetween(start, event.StartTime, event.EndTime) {
            eventsInRange = append(eventsInRange, event)
        }
    }

    return eventsInRange
}

func (es *EventScheduler) GetEventsOnDate(date time.Time) []Event {
    var eventsOnDate []Event

    dayStart := gotime.SoD(date)
    dayEnd := gotime.EoD(date)

    for _, event := range es.events {
        if gotime.IsBetweenDates(event.StartTime, dayStart, dayEnd) ||
           gotime.IsBetweenDates(event.EndTime, dayStart, dayEnd) {
            eventsOnDate = append(eventsOnDate, event)
        }
    }

    return eventsOnDate
}

func (es *EventScheduler) ScheduleEvent(event Event) error {
    if !es.IsTimeSlotAvailable(event.StartTime, event.EndTime) {
        return fmt.Errorf("time slot not available")
    }

    if event.EndTime.Before(event.StartTime) {
        return fmt.Errorf("end time must be after start time")
    }

    es.events = append(es.events, event)
    return nil
}

// Usage
func ExampleEventScheduling() {
    scheduler := NewEventScheduler()

    meeting := Event{
        ID:        "1",
        Name:      "Team Meeting",
        StartTime: time.Date(2025, 7, 7, 10, 0, 0, 0, time.UTC),
        EndTime:   time.Date(2025, 7, 7, 11, 0, 0, 0, time.UTC),
        Location:  "Conference Room A",
    }

    err := scheduler.ScheduleEvent(meeting)
    if err != nil {
        fmt.Printf("Failed to schedule: %v\n", err)
    }

    // Check availability
    newStart := time.Date(2025, 7, 7, 10, 30, 0, 0, time.UTC)
    newEnd := time.Date(2025, 7, 7, 11, 30, 0, 0, time.UTC)

    if scheduler.IsTimeSlotAvailable(newStart, newEnd) {
        fmt.Println("Time slot is available")
    } else {
        fmt.Println("Time slot conflicts with existing event")
    }
}
```

### 2. Business Hours Validation

```go
package business

import (
    "github.com/maniartech/gotime"
    "time"
)

type BusinessHours struct {
    Open  time.Time
    Close time.Time
}

type BusinessSchedule struct {
    Hours   map[time.Weekday]BusinessHours
    Closed  map[time.Weekday]bool
    TimeZone *time.Location
}

func NewBusinessSchedule(tz *time.Location) *BusinessSchedule {
    return &BusinessSchedule{
        Hours:    make(map[time.Weekday]BusinessHours),
        Closed:   make(map[time.Weekday]bool),
        TimeZone: tz,
    }
}

func (bs *BusinessSchedule) SetHours(day time.Weekday, open, close string) error {
    openTime, err := gotime.Parse(open, "hh:ii")
    if err != nil {
        return fmt.Errorf("invalid open time: %v", err)
    }

    closeTime, err := gotime.Parse(close, "hh:ii")
    if err != nil {
        return fmt.Errorf("invalid close time: %v", err)
    }

    bs.Hours[day] = BusinessHours{
        Open:  openTime,
        Close: closeTime,
    }
    bs.Closed[day] = false

    return nil
}

func (bs *BusinessSchedule) SetClosed(day time.Weekday) {
    bs.Closed[day] = true
    delete(bs.Hours, day)
}

func (bs *BusinessSchedule) IsBusinessHours(t time.Time) bool {
    // Convert to business timezone
    localTime := t.In(bs.TimeZone)
    weekday := localTime.Weekday()

    // Check if closed
    if bs.Closed[weekday] {
        return false
    }

    // Check if hours defined
    hours, exists := bs.Hours[weekday]
    if !exists {
        return false
    }

    // Create times for the same date
    dayStart := time.Date(localTime.Year(), localTime.Month(), localTime.Day(),
        hours.Open.Hour(), hours.Open.Minute(), 0, 0, bs.TimeZone)
    dayEnd := time.Date(localTime.Year(), localTime.Month(), localTime.Day(),
        hours.Close.Hour(), hours.Close.Minute(), 0, 0, bs.TimeZone)

    return gotime.IsBetween(localTime, dayStart, dayEnd)
}

func (bs *BusinessSchedule) GetNextBusinessHour(from time.Time) time.Time {
    localTime := from.In(bs.TimeZone)

    // Check next 7 days
    for days := 0; days < 7; days++ {
        checkDate := localTime.AddDate(0, 0, days)
        weekday := checkDate.Weekday()

        // Skip if closed
        if bs.Closed[weekday] {
            continue
        }

        // Check if hours defined
        hours, exists := bs.Hours[weekday]
        if !exists {
            continue
        }

        openTime := time.Date(checkDate.Year(), checkDate.Month(), checkDate.Day(),
            hours.Open.Hour(), hours.Open.Minute(), 0, 0, bs.TimeZone)

        // If it's the same day and still business hours
        if days == 0 && bs.IsBusinessHours(localTime) {
            return localTime
        }

        // Return opening time
        if days > 0 || openTime.After(localTime) {
            return openTime
        }
    }

    return time.Time{} // No business hours found in next 7 days
}

// Usage
func ExampleBusinessHours() {
    est, _ := time.LoadLocation("America/New_York")
    schedule := NewBusinessSchedule(est)

    // Set business hours
    schedule.SetHours(time.Monday, "09:00", "17:00")
    schedule.SetHours(time.Tuesday, "09:00", "17:00")
    schedule.SetHours(time.Wednesday, "09:00", "17:00")
    schedule.SetHours(time.Thursday, "09:00", "17:00")
    schedule.SetHours(time.Friday, "09:00", "17:00")
    schedule.SetClosed(time.Saturday)
    schedule.SetClosed(time.Sunday)

    // Check if current time is business hours
    now := time.Now()
    if schedule.IsBusinessHours(now) {
        fmt.Println("Currently business hours")
    } else {
        next := schedule.GetNextBusinessHour(now)
        fmt.Printf("Next business hour: %s\n",
            gotime.Format(next, "wwww, yyyy-mm-dd hh:ii"))
    }
}
```

### 3. Date Range Analytics

```go
package analytics

import (
    "github.com/maniartech/gotime"
    "time"
)

type DateRangeAnalyzer struct {
    data map[time.Time]float64
}

func NewDateRangeAnalyzer() *DateRangeAnalyzer {
    return &DateRangeAnalyzer{
        data: make(map[time.Time]float64),
    }
}

func (dra *DateRangeAnalyzer) AddDataPoint(date time.Time, value float64) {
    // Normalize to date only
    dateOnly := gotime.TruncateTime(date)
    dra.data[dateOnly] = value
}

func (dra *DateRangeAnalyzer) GetDataInRange(start, end time.Time) map[time.Time]float64 {
    result := make(map[time.Time]float64)

    startDate := gotime.TruncateTime(start)
    endDate := gotime.TruncateTime(end)

    for date, value := range dra.data {
        if gotime.IsBetweenDates(date, startDate, endDate) {
            result[date] = value
        }
    }

    return result
}

func (dra *DateRangeAnalyzer) GetQuarterlyData(year int, quarter int) map[time.Time]float64 {
    var startMonth, endMonth int

    switch quarter {
    case 1:
        startMonth, endMonth = 1, 3
    case 2:
        startMonth, endMonth = 4, 6
    case 3:
        startMonth, endMonth = 7, 9
    case 4:
        startMonth, endMonth = 10, 12
    default:
        return make(map[time.Time]float64)
    }

    start := time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)
    end := gotime.MonthEnd(time.Date(year, time.Month(endMonth), 1, 0, 0, 0, 0, time.UTC))

    return dra.GetDataInRange(start, end)
}

func (dra *DateRangeAnalyzer) GetMonthlyData(year int, month int) map[time.Time]float64 {
    start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    end := gotime.MonthEnd(start)

    return dra.GetDataInRange(start, end)
}

func (dra *DateRangeAnalyzer) GetWeeklyData(date time.Time) map[time.Time]float64 {
    start := gotime.WeekStart(date)
    end := gotime.WeekEnd(date)

    return dra.GetDataInRange(start, end)
}

func (dra *DateRangeAnalyzer) GetYearToDateData(year int) map[time.Time]float64 {
    start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
    end := time.Now()

    return dra.GetDataInRange(start, end)
}

func (dra *DateRangeAnalyzer) GenerateRangeReport(start, end time.Time) map[string]interface{} {
    data := dra.GetDataInRange(start, end)

    if len(data) == 0 {
        return map[string]interface{}{
            "error": "No data in specified range",
            "range": fmt.Sprintf("%s to %s",
                gotime.Format(start, "yyyy-mm-dd"),
                gotime.Format(end, "yyyy-mm-dd")),
        }
    }

    var sum, min, max float64
    var count int
    first := true

    for _, value := range data {
        sum += value
        count++

        if first {
            min = value
            max = value
            first = false
        } else {
            if value < min {
                min = value
            }
            if value > max {
                max = value
            }
        }
    }

    average := sum / float64(count)

    return map[string]interface{}{
        "range": map[string]string{
            "start": gotime.Format(start, "yyyy-mm-dd"),
            "end":   gotime.Format(end, "yyyy-mm-dd"),
        },
        "statistics": map[string]float64{
            "count":   float64(count),
            "sum":     sum,
            "average": average,
            "min":     min,
            "max":     max,
        },
        "generated_at": gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss"),
    }
}
```

### 4. Subscription Management

```go
package subscription

import (
    "github.com/maniartech/gotime"
    "time"
)

type SubscriptionStatus string

const (
    Active   SubscriptionStatus = "active"
    Expired  SubscriptionStatus = "expired"
    Canceled SubscriptionStatus = "canceled"
    Trial    SubscriptionStatus = "trial"
)

type Subscription struct {
    ID        string
    UserID    string
    StartDate time.Time
    EndDate   time.Time
    Status    SubscriptionStatus
    PlanType  string
}

type SubscriptionManager struct {
    subscriptions []Subscription
}

func NewSubscriptionManager() *SubscriptionManager {
    return &SubscriptionManager{
        subscriptions: make([]Subscription, 0),
    }
}

func (sm *SubscriptionManager) IsActiveSubscription(userID string, checkDate time.Time) bool {
    for _, sub := range sm.subscriptions {
        if sub.UserID == userID && sub.Status == Active {
            if gotime.IsBetween(checkDate, sub.StartDate, sub.EndDate) {
                return true
            }
        }
    }
    return false
}

func (sm *SubscriptionManager) GetActiveSubscriptions(date time.Time) []Subscription {
    var active []Subscription

    for _, sub := range sm.subscriptions {
        if sub.Status == Active && gotime.IsBetween(date, sub.StartDate, sub.EndDate) {
            active = append(active, sub)
        }
    }

    return active
}

func (sm *SubscriptionManager) GetExpiringSubscriptions(days int) []Subscription {
    var expiring []Subscription
    now := time.Now()
    futureDate := gotime.Days(days, now)

    for _, sub := range sm.subscriptions {
        if sub.Status == Active && gotime.IsBetween(sub.EndDate, now, futureDate) {
            expiring = append(expiring, sub)
        }
    }

    return expiring
}

func (sm *SubscriptionManager) GetSubscriptionsInRange(start, end time.Time) []Subscription {
    var inRange []Subscription

    for _, sub := range sm.subscriptions {
        // Check if subscription overlaps with the range
        if gotime.IsBetween(sub.StartDate, start, end) ||
           gotime.IsBetween(sub.EndDate, start, end) ||
           gotime.IsBetween(start, sub.StartDate, sub.EndDate) {
            inRange = append(inRange, sub)
        }
    }

    return inRange
}

func (sm *SubscriptionManager) GenerateSubscriptionReport(reportDate time.Time) map[string]interface{} {
    active := sm.GetActiveSubscriptions(reportDate)
    expiring7 := sm.GetExpiringSubscriptions(7)
    expiring30 := sm.GetExpiringSubscriptions(30)

    // Count by plan type
    planCounts := make(map[string]int)
    for _, sub := range active {
        planCounts[sub.PlanType]++
    }

    return map[string]interface{}{
        "report_date": gotime.Format(reportDate, "yyyy-mm-dd"),
        "active_subscriptions": len(active),
        "expiring_7_days": len(expiring7),
        "expiring_30_days": len(expiring30),
        "plans": planCounts,
        "generated_at": gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss"),
    }
}

func (sm *SubscriptionManager) AddSubscription(sub Subscription) error {
    if sub.EndDate.Before(sub.StartDate) {
        return fmt.Errorf("end date must be after start date")
    }

    sm.subscriptions = append(sm.subscriptions, sub)
    return nil
}
```

## Range Validation Patterns

### 1. Input Validation

```go
func ValidateDateRange(start, end string, format string) error {
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

    // Check for reasonable range (e.g., not more than 10 years)
    maxRange := gotime.Years(10, startDate)
    if endDate.After(maxRange) {
        return fmt.Errorf("date range too large (maximum 10 years)")
    }

    return nil
}
```

### 2. Range Overlap Detection

```go
func HasOverlap(start1, end1, start2, end2 time.Time) bool {
    return gotime.IsBetween(start1, start2, end2) ||
           gotime.IsBetween(end1, start2, end2) ||
           gotime.IsBetween(start2, start1, end1) ||
           gotime.IsBetween(end2, start1, end1)
}
```

### 3. Multiple Range Checking

```go
type DateRange struct {
    Start time.Time
    End   time.Time
    Name  string
}

func FindOverlappingRanges(ranges []DateRange) [][]DateRange {
    var overlapping [][]DateRange

    for i := 0; i < len(ranges); i++ {
        var group []DateRange
        group = append(group, ranges[i])

        for j := i + 1; j < len(ranges); j++ {
            if HasOverlap(ranges[i].Start, ranges[i].End, ranges[j].Start, ranges[j].End) {
                group = append(group, ranges[j])
            }
        }

        if len(group) > 1 {
            overlapping = append(overlapping, group)
        }
    }

    return overlapping
}
```

## Performance Considerations

### 1. Large Dataset Range Queries

```go
// For large datasets, consider indexing by date
type IndexedDateRange struct {
    data  map[string]interface{}
    index map[time.Time][]string // Date -> list of IDs
}

func (idr *IndexedDateRange) GetItemsInRange(start, end time.Time) []interface{} {
    var items []interface{}

    current := gotime.TruncateTime(start)
    endDate := gotime.TruncateTime(end)

    for !current.After(endDate) {
        if ids, exists := idr.index[current]; exists {
            for _, id := range ids {
                if item, exists := idr.data[id]; exists {
                    items = append(items, item)
                }
            }
        }
        current = gotime.Days(1, current)
    }

    return items
}
```

### 2. Caching Range Results

```go
type CachedRangeQuery struct {
    cache map[string][]interface{}
}

func (crq *CachedRangeQuery) GetRangeKey(start, end time.Time) string {
    return fmt.Sprintf("%s_%s",
        gotime.Format(start, "yyyy-mm-dd"),
        gotime.Format(end, "yyyy-mm-dd"))
}

func (crq *CachedRangeQuery) GetCachedRange(start, end time.Time) ([]interface{}, bool) {
    key := crq.GetRangeKey(start, end)
    result, exists := crq.cache[key]
    return result, exists
}
```

---

## Summary

GoTime's date range operations provide:

- **Flexible range checking** with IsBetween and IsBetweenDates
- **Automatic boundary handling** (inclusive ranges, order independence)
- **Date vs datetime distinction** for appropriate comparison logic
- **Real-world applications** for scheduling, analytics, and business logic
- **Performance patterns** for large-scale range operations

These functions are essential for building applications that need to:
- Validate date ranges in user input
- Check for schedule conflicts
- Analyze time-series data
- Manage subscriptions and renewals
- Implement business rules based on date ranges

---

Next: [Utility Functions](utilities.md) | [Back to API Reference](../api-reference/)
