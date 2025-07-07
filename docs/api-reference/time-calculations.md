[Home](../README.md) > [API Reference](README.md) > Time Calculations

# Time Calculations

GoTime provides advanced functions for working with time differences, business days, and date comparisons.

## Date Comparison Functions

### Latest

Returns the latest (most recent) time from a list of times.

```go
func Latest(t1, t2 time.Time, tn ...time.Time) time.Time
```

**Examples:**

```go
date1 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
date3 := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

latest := gotime.Latest(date1, date2, date3)    // Returns date3 (2025-12-31)
latest = gotime.Latest(date1, date2)            // Returns date2 (2025-06-15)

// Works with any number of dates
now := time.Now()
latest = gotime.Latest(date1, date2, date3, now)
```

### Earliest

Returns the earliest time from a list of times.

```go
func Earliest(t1, t2 time.Time, tn ...time.Time) time.Time
```

**Examples:**

```go
date1 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
date3 := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

earliest := gotime.Earliest(date1, date2, date3)  // Returns date1 (2025-01-01)
earliest = gotime.Earliest(date2, date3)          // Returns date2 (2025-06-15)
```

## Time Difference Calculations

### Diff

Calculates the difference between two times in the specified unit.

```go
func Diff(t1, t2 time.Time, unit time.Duration, rounded ...bool) float64
```

**Parameters:**
- `t1`: First time
- `t2`: Second time
- `unit`: Unit for the difference (time.Second, time.Minute, time.Hour, etc.)
- `rounded`: Optional boolean to round the result (default: false)

**Examples:**

```go
start := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)

// Difference in hours
hours := gotime.Diff(end, start, time.Hour)         // 2.5
hoursRounded := gotime.Diff(end, start, time.Hour, true)  // 3.0

// Difference in minutes
minutes := gotime.Diff(end, start, time.Minute)     // 150.0

// Difference in days
nextDay := time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC)
days := gotime.Diff(nextDay, start, 24*time.Hour)   // 1.0
```

## Business Day Functions

### WorkDay

Calculates a future date after adding the specified number of working days.

```go
func WorkDay(startDate time.Time, days int, workingDays [7]bool, holidays ...time.Time) (time.Time, error)
```

**Parameters:**
- `startDate`: Starting date
- `days`: Number of working days to add (must be positive)
- `workingDays`: Array defining which days are working days [Sunday, Monday, ..., Saturday]
- `holidays`: Optional list of holiday dates to exclude

**Examples:**

```go
// Define working days (Monday to Friday)
workingDays := [7]bool{false, true, true, true, true, true, false}
//                      Sun   Mon  Tue  Wed  Thu  Fri  Sat

startDate := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday

// 5 business days from Monday = next Monday
result, err := gotime.WorkDay(startDate, 5, workingDays)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 2025-07-14 (next Monday)

// With holidays
holidays := []time.Time{
    time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Independence Day
}

result, err = gotime.WorkDay(startDate, 5, workingDays, holidays...)
// Will skip July 4th and calculate accordingly
```

### PrevWorkDay

Calculates a past date by subtracting the specified number of working days.

```go
func PrevWorkDay(startDate time.Time, days int, workingDays [7]bool, holidays ...time.Time) (time.Time, error)
```

**Examples:**

```go
workingDays := [7]bool{false, true, true, true, true, true, false}
startDate := time.Date(2025, 7, 11, 0, 0, 0, 0, time.UTC) // Friday

// 5 business days before Friday = previous Friday
result, err := gotime.PrevWorkDay(startDate, 5, workingDays)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 2025-07-04 (previous Friday)
```

### NetWorkDays

Counts the number of working days between two dates.

```go
func NetWorkDays(startDate, endDate time.Time, workingDays [7]bool, holidays ...time.Time) (int, error)
```

**Examples:**

```go
workingDays := [7]bool{false, true, true, true, true, true, false}

start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)  // Monday
end := time.Date(2025, 7, 18, 0, 0, 0, 0, time.UTC)   // Friday (next week)

count, err := gotime.NetWorkDays(start, end, workingDays)
if err != nil {
    log.Fatal(err)
}
fmt.Println(count) // 10 working days

// With holidays
holidays := []time.Time{
    time.Date(2025, 7, 14, 0, 0, 0, 0, time.UTC), // Holiday on Monday
}

count, err = gotime.NetWorkDays(start, end, workingDays, holidays...)
fmt.Println(count) // 9 working days (excluding holiday)
```

## Date Manipulation Functions

### DateValue

Returns a serial number representing the date (days since 1900-01-01).

```go
func DateValue(date time.Time) int
```

**Examples:**

```go
date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
serial := gotime.DateValue(date)
fmt.Println(serial) // 45252 (days since 1900-01-01)

// Useful for Excel compatibility or date calculations
baseDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
baseSerial := gotime.DateValue(baseDate)
fmt.Println(baseSerial) // 2
```

### TruncateTime

Removes the time portion from a date, keeping only the date part.

```go
func TruncateTime(date time.Time) time.Time
```

**Examples:**

```go
datetime := time.Date(2025, 7, 7, 14, 30, 45, 0, time.UTC)
dateOnly := gotime.TruncateTime(datetime)
fmt.Println(dateOnly) // 2025-07-07 00:00:00 +0000 UTC

// Useful for date-only comparisons
today := gotime.TruncateTime(time.Now())
targetDate := gotime.TruncateTime(someDateTime)
if today.Equal(targetDate) {
    fmt.Println("Same date!")
}
```

## Real-World Use Cases

### 1. Project Timeline Calculator

```go
package timeline

import (
    "github.com/maniartech/gotime"
    "time"
)

type ProjectCalculator struct {
    workingDays [7]bool
    holidays    []time.Time
}

func NewProjectCalculator() *ProjectCalculator {
    // Monday to Friday working days
    workingDays := [7]bool{false, true, true, true, true, true, false}

    return &ProjectCalculator{
        workingDays: workingDays,
        holidays:    []time.Time{},
    }
}

func (pc *ProjectCalculator) AddHoliday(date time.Time) {
    pc.holidays = append(pc.holidays, gotime.TruncateTime(date))
}

func (pc *ProjectCalculator) CalculateProjectEnd(startDate time.Time, workingDays int) (time.Time, error) {
    return gotime.WorkDay(startDate, workingDays, pc.workingDays, pc.holidays...)
}

func (pc *ProjectCalculator) CalculateProjectDuration(startDate, endDate time.Time) (int, error) {
    return gotime.NetWorkDays(startDate, endDate, pc.workingDays, pc.holidays...)
}

func (pc *ProjectCalculator) GetMilestones(startDate time.Time, totalDays int) (map[string]time.Time, error) {
    milestones := make(map[string]time.Time)

    quarter, err := gotime.WorkDay(startDate, totalDays/4, pc.workingDays, pc.holidays...)
    if err != nil {
        return nil, err
    }
    milestones["25%"] = quarter

    half, err := gotime.WorkDay(startDate, totalDays/2, pc.workingDays, pc.holidays...)
    if err != nil {
        return nil, err
    }
    milestones["50%"] = half

    threeQuarter, err := gotime.WorkDay(startDate, totalDays*3/4, pc.workingDays, pc.holidays...)
    if err != nil {
        return nil, err
    }
    milestones["75%"] = threeQuarter

    end, err := gotime.WorkDay(startDate, totalDays, pc.workingDays, pc.holidays...)
    if err != nil {
        return nil, err
    }
    milestones["100%"] = end

    return milestones, nil
}
```

### 2. Performance Benchmarking

```go
package benchmark

import (
    "github.com/maniartech/gotime"
    "time"
)

type Benchmark struct {
    Name      string
    StartTime time.Time
    EndTime   time.Time
    Metrics   map[string]float64
}

func NewBenchmark(name string) *Benchmark {
    return &Benchmark{
        Name:      name,
        StartTime: time.Now(),
        Metrics:   make(map[string]float64),
    }
}

func (b *Benchmark) Stop() {
    b.EndTime = time.Now()
}

func (b *Benchmark) GetDuration(unit time.Duration) float64 {
    if b.EndTime.IsZero() {
        b.Stop()
    }
    return gotime.Diff(b.EndTime, b.StartTime, unit)
}

func (b *Benchmark) GetDurationRounded(unit time.Duration) float64 {
    if b.EndTime.IsZero() {
        b.Stop()
    }
    return gotime.Diff(b.EndTime, b.StartTime, unit, true)
}

func (b *Benchmark) AddMetric(name string, value float64) {
    b.Metrics[name] = value
}

func (b *Benchmark) GenerateReport() map[string]interface{} {
    if b.EndTime.IsZero() {
        b.Stop()
    }

    return map[string]interface{}{
        "name":         b.Name,
        "start_time":   gotime.Format(b.StartTime, "yyyy-mm-dd hh:ii:ss.000"),
        "end_time":     gotime.Format(b.EndTime, "yyyy-mm-dd hh:ii:ss.000"),
        "duration_ms":  b.GetDuration(time.Millisecond),
        "duration_sec": b.GetDurationRounded(time.Second),
        "metrics":      b.Metrics,
    }
}

// Usage
func BenchmarkFunction() {
    bench := NewBenchmark("Database Query")

    // ... perform operation ...

    bench.Stop()

    fmt.Printf("Operation took %.2f ms\n", bench.GetDuration(time.Millisecond))
    fmt.Printf("Rounded: %.0f seconds\n", bench.GetDurationRounded(time.Second))
}
```

### 3. SLA (Service Level Agreement) Monitor

```go
package sla

import (
    "github.com/maniartech/gotime"
    "time"
)

type SLAMonitor struct {
    workingDays [7]bool
    holidays    []time.Time
}

func NewSLAMonitor() *SLAMonitor {
    // Business hours: Monday to Friday
    workingDays := [7]bool{false, true, true, true, true, true, false}

    return &SLAMonitor{
        workingDays: workingDays,
        holidays:    []time.Time{},
    }
}

func (sla *SLAMonitor) CalculateResponseDeadline(requestTime time.Time, responseSLA int) (time.Time, error) {
    // Calculate deadline based on business days
    return gotime.WorkDay(requestTime, responseSLA, sla.workingDays, sla.holidays...)
}

func (sla *SLAMonitor) CalculateResolutionDeadline(requestTime time.Time, resolutionSLA int) (time.Time, error) {
    return gotime.WorkDay(requestTime, resolutionSLA, sla.workingDays, sla.holidays...)
}

func (sla *SLAMonitor) CheckSLABreach(requestTime, responseTime time.Time, slaHours int) bool {
    deadline, err := sla.CalculateResponseDeadline(requestTime, slaHours)
    if err != nil {
        return false
    }

    return responseTime.After(deadline)
}

func (sla *SLAMonitor) GetSLAStatus(requestTime time.Time, slaHours int) map[string]interface{} {
    deadline, err := sla.CalculateResponseDeadline(requestTime, slaHours)
    if err != nil {
        return map[string]interface{}{"error": err.Error()}
    }

    now := time.Now()

    status := map[string]interface{}{
        "request_time":    gotime.Format(requestTime, "yyyy-mm-dd hh:ii:ss"),
        "deadline":        gotime.Format(deadline, "yyyy-mm-dd hh:ii:ss"),
        "sla_hours":       slaHours,
    }

    if now.After(deadline) {
        overdue := gotime.Diff(now, deadline, time.Hour, true)
        status["status"] = "BREACHED"
        status["overdue_hours"] = overdue
    } else {
        remaining := gotime.Diff(deadline, now, time.Hour, true)
        status["status"] = "ON_TIME"
        status["remaining_hours"] = remaining
        status["time_until_deadline"] = gotime.TimeAgo(deadline)
    }

    return status
}
```

### 4. Data Retention Manager

```go
package retention

import (
    "github.com/maniartech/gotime"
    "time"
)

type RetentionPolicy struct {
    Name            string
    RetentionPeriod time.Duration
    ArchivePeriod   time.Duration
}

type DataManager struct {
    policies map[string]RetentionPolicy
}

func NewDataManager() *DataManager {
    return &DataManager{
        policies: make(map[string]RetentionPolicy),
    }
}

func (dm *DataManager) AddPolicy(name string, retention, archive time.Duration) {
    dm.policies[name] = RetentionPolicy{
        Name:            name,
        RetentionPeriod: retention,
        ArchivePeriod:   archive,
    }
}

func (dm *DataManager) ShouldArchive(dataCreated time.Time, policyName string) bool {
    policy, exists := dm.policies[policyName]
    if !exists {
        return false
    }

    archiveDate := dataCreated.Add(policy.RetentionPeriod)
    return time.Now().After(archiveDate)
}

func (dm *DataManager) ShouldDelete(dataCreated time.Time, policyName string) bool {
    policy, exists := dm.policies[policyName]
    if !exists {
        return false
    }

    deleteDate := dataCreated.Add(policy.RetentionPeriod + policy.ArchivePeriod)
    return time.Now().After(deleteDate)
}

func (dm *DataManager) GetRetentionStatus(dataCreated time.Time, policyName string) map[string]interface{} {
    policy, exists := dm.policies[policyName]
    if !exists {
        return map[string]interface{}{"error": "Policy not found"}
    }

    now := time.Now()
    archiveDate := dataCreated.Add(policy.RetentionPeriod)
    deleteDate := dataCreated.Add(policy.RetentionPeriod + policy.ArchivePeriod)

    age := gotime.Diff(now, dataCreated, 24*time.Hour, true)

    status := map[string]interface{}{
        "policy":       policyName,
        "created":      gotime.Format(dataCreated, "yyyy-mm-dd"),
        "age_days":     age,
        "archive_date": gotime.Format(archiveDate, "yyyy-mm-dd"),
        "delete_date":  gotime.Format(deleteDate, "yyyy-mm-dd"),
    }

    switch {
    case now.After(deleteDate):
        status["action"] = "DELETE"
        status["overdue_days"] = gotime.Diff(now, deleteDate, 24*time.Hour, true)
    case now.After(archiveDate):
        status["action"] = "ARCHIVE"
        status["archive_overdue_days"] = gotime.Diff(now, archiveDate, 24*time.Hour, true)
        status["days_until_deletion"] = gotime.Diff(deleteDate, now, 24*time.Hour, true)
    default:
        status["action"] = "RETAIN"
        status["days_until_archive"] = gotime.Diff(archiveDate, now, 24*time.Hour, true)
    }

    return status
}

// Usage
func SetupRetentionPolicies() *DataManager {
    dm := NewDataManager()

    // Different retention policies
    dm.AddPolicy("logs", 30*24*time.Hour, 365*24*time.Hour)      // 30 days active, 1 year archive
    dm.AddPolicy("analytics", 90*24*time.Hour, 2*365*24*time.Hour) // 90 days active, 2 years archive
    dm.AddPolicy("backups", 7*24*time.Hour, 30*24*time.Hour)     // 7 days active, 30 days archive

    return dm
}
```

## Error Handling

All business day functions return errors for invalid inputs:

```go
workingDays := [7]bool{false, false, false, false, false, false, false} // No working days

_, err := gotime.WorkDay(time.Now(), 5, workingDays)
if err != nil {
    fmt.Println(err) // "at least one working day must be specified"
}

_, err = gotime.WorkDay(time.Now(), -5, workingDays)
if err != nil {
    fmt.Println(err) // "number of days cannot be negative"
}
```

## Performance Tips

1. **Pre-define working day arrays** as constants when possible
2. **Cache holiday lists** rather than recreating them
3. **Use batch operations** for multiple calculations
4. **Consider timezone implications** for business day calculations

---

## Summary

GoTime's time calculation functions provide:

- **Flexible date comparison** with Latest and Earliest
- **Precise time differences** with customizable units and rounding
- **Business day calculations** with holiday support
- **Data manipulation utilities** for common operations
- **Real-world applications** for project management, SLA monitoring, and data retention

These functions are essential for building robust business applications that need to handle complex time-based calculations.

---

Next: [Date Range Operations](date-ranges.md) | [Back to API Reference](../api-reference/)
