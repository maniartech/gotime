# Utility Functions

GoTime provides a comprehensive set of utility functions for common date and time operations, data manipulation, and helper methods that simplify everyday programming tasks.

## Time Manipulation Utilities

### TruncateTime

Removes the time component from a datetime, keeping only the date portion.

```go
func TruncateTime(t time.Time) time.Time
```

**Purpose:** Normalizes datetime values to date-only for comparison and grouping operations.

**Examples:**

```go
// Remove time component
datetime := time.Date(2025, 7, 7, 14, 30, 45, 123456789, time.UTC)
dateOnly := gotime.TruncateTime(datetime)
// Result: 2025-07-07 00:00:00 +0000 UTC

// Useful for date comparisons
date1 := gotime.TruncateTime(time.Now())
date2 := gotime.TruncateTime(time.Date(2025, 7, 7, 23, 59, 59, 0, time.UTC))
isSameDate := date1.Equal(date2)

// Group data by date
type Transaction struct {
    ID     string
    Amount float64
    Date   time.Time
}

func GroupTransactionsByDate(transactions []Transaction) map[time.Time][]Transaction {
    groups := make(map[time.Time][]Transaction)

    for _, tx := range transactions {
        dateKey := gotime.TruncateTime(tx.Date)
        groups[dateKey] = append(groups[dateKey], tx)
    }

    return groups
}
```

### DateValue

Converts various date representations (string, timestamp, etc.) into a standardized time.Time value.

```go
func DateValue(input interface{}) (time.Time, error)
```

**Supported Input Types:**
- `time.Time` - Returns as-is
- `string` - Attempts to parse using common formats
- `int64` - Unix timestamp
- `float64` - Unix timestamp with fractional seconds

**Examples:**

```go
// From string
date1, err := gotime.DateValue("2025-07-07")
date2, err := gotime.DateValue("07/07/2025")
date3, err := gotime.DateValue("July 7, 2025")

// From timestamp
date4, err := gotime.DateValue(1720310400) // Unix timestamp

// From time.Time (no conversion)
now := time.Now()
date5, err := gotime.DateValue(now)

// Error handling
if err != nil {
    fmt.Printf("Invalid date format: %v\n", err)
}

// Data import utility
func ImportDateColumn(values []interface{}) ([]time.Time, error) {
    var dates []time.Time

    for i, value := range values {
        date, err := gotime.DateValue(value)
        if err != nil {
            return nil, fmt.Errorf("row %d: invalid date '%v': %v", i+1, value, err)
        }
        dates = append(dates, date)
    }

    return dates, nil
}
```

## Time Period Boundaries

### SoD (Start of Day)

Returns the beginning of the day (00:00:00) for a given date.

```go
func SoD(t time.Time) time.Time
```

### EoD (End of Day)

Returns the end of the day (23:59:59.999999999) for a given date.

```go
func EoD(t time.Time) time.Time
```

### WeekStart / WeekEnd

Returns the beginning and end of the week for a given date.

```go
func WeekStart(t time.Time) time.Time
func WeekEnd(t time.Time) time.Time
```

### MonthStart / MonthEnd

Returns the beginning and end of the month for a given date.

```go
func MonthStart(t time.Time) time.Time
func MonthEnd(t time.Time) time.Time
```

**Examples:**

```go
date := time.Date(2025, 7, 15, 14, 30, 45, 0, time.UTC) // Tuesday, July 15

// Day boundaries
dayStart := gotime.SoD(date)     // 2025-07-15 00:00:00
dayEnd := gotime.EoD(date)       // 2025-07-15 23:59:59.999999999

// Week boundaries (assuming Monday = start of week)
weekStart := gotime.WeekStart(date)  // 2025-07-14 00:00:00 (Monday)
weekEnd := gotime.WeekEnd(date)      // 2025-07-20 23:59:59 (Sunday)

// Month boundaries
monthStart := gotime.MonthStart(date)  // 2025-07-01 00:00:00
monthEnd := gotime.MonthEnd(date)      // 2025-07-31 23:59:59

// Generate report for current month
func GenerateMonthlyReport() Report {
    now := time.Now()
    start := gotime.MonthStart(now)
    end := gotime.MonthEnd(now)

    return Report{
        Period: fmt.Sprintf("%s to %s",
            gotime.Format(start, "yyyy-mm-dd"),
            gotime.Format(end, "yyyy-mm-dd")),
        GeneratedAt: now,
        Data: getDataInRange(start, end),
    }
}
```

## Data Conversion Utilities

### ToISO8601

Converts a time.Time to ISO 8601 format string.

```go
func ToISO8601(t time.Time) string
```

### FromISO8601

Parses an ISO 8601 format string to time.Time.

```go
func FromISO8601(s string) (time.Time, error)
```

**Examples:**

```go
now := time.Now()

// Convert to ISO 8601
iso := gotime.ToISO8601(now)
// Result: "2025-07-07T14:30:45Z"

// Parse from ISO 8601
parsed, err := gotime.FromISO8601("2025-07-07T14:30:45Z")
if err != nil {
    fmt.Printf("Parse error: %v\n", err)
}

// API integration
type APIResponse struct {
    Timestamp string `json:"timestamp"`
    Data      interface{} `json:"data"`
}

func ParseAPIResponse(response APIResponse) (time.Time, interface{}, error) {
    timestamp, err := gotime.FromISO8601(response.Timestamp)
    if err != nil {
        return time.Time{}, nil, fmt.Errorf("invalid timestamp: %v", err)
    }

    return timestamp, response.Data, nil
}
```

## Helper Functions

### IsLeapYear

Checks if a given year is a leap year.

```go
func IsLeapYear(year int) bool
```

### DaysInMonth

Returns the number of days in a specific month and year.

```go
func DaysInMonth(year int, month time.Month) int
```

### IsWeekend

Checks if a given date falls on a weekend (Saturday or Sunday).

```go
func IsWeekend(t time.Time) bool
```

### IsWeekday

Checks if a given date falls on a weekday (Monday through Friday).

```go
func IsWeekday(t time.Time) bool
```

**Examples:**

```go
// Leap year calculation
year := 2024
if gotime.IsLeapYear(year) {
    fmt.Printf("%d is a leap year\n", year)  // true
}

// Days in month
days := gotime.DaysInMonth(2024, time.February)  // 29 (leap year)
days = gotime.DaysInMonth(2025, time.February)   // 28 (regular year)

// Weekend/weekday checking
date := time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC)  // Saturday

if gotime.IsWeekend(date) {
    fmt.Println("It's the weekend!")
}

if gotime.IsWeekday(date) {
    fmt.Println("It's a weekday")
} else {
    fmt.Println("It's not a weekday")
}

// Business logic example
func CalculateBusinessDays(start, end time.Time) int {
    count := 0
    current := gotime.TruncateTime(start)
    endDate := gotime.TruncateTime(end)

    for !current.After(endDate) {
        if gotime.IsWeekday(current) {
            count++
        }
        current = gotime.Days(1, current)
    }

    return count
}
```

## Time Zone Utilities

### ConvertTimezone

Converts a time from one timezone to another.

```go
func ConvertTimezone(t time.Time, fromTZ, toTZ *time.Location) time.Time
```

### GetTimezoneOffset

Gets the timezone offset in hours for a specific time and location.

```go
func GetTimezoneOffset(t time.Time, loc *time.Location) float64
```

**Examples:**

```go
// Load timezones
utc, _ := time.LoadLocation("UTC")
est, _ := time.LoadLocation("America/New_York")
pst, _ := time.LoadLocation("America/Los_Angeles")
tokyo, _ := time.LoadLocation("Asia/Tokyo")

utcTime := time.Date(2025, 7, 7, 12, 0, 0, 0, utc)

// Convert between timezones
estTime := gotime.ConvertTimezone(utcTime, utc, est)
pstTime := gotime.ConvertTimezone(utcTime, utc, pst)
tokyoTime := gotime.ConvertTimezone(utcTime, utc, tokyo)

// Get timezone offsets
estOffset := gotime.GetTimezoneOffset(utcTime, est)     // -4.0 (EDT)
pstOffset := gotime.GetTimezoneOffset(utcTime, pst)     // -7.0 (PDT)
tokyoOffset := gotime.GetTimezoneOffset(utcTime, tokyo) // 9.0

// Multi-timezone scheduling
type GlobalMeeting struct {
    Title     string
    UTCTime   time.Time
    Timezones []*time.Location
}

func (gm *GlobalMeeting) GetLocalTimes() map[string]string {
    times := make(map[string]string)

    for _, tz := range gm.Timezones {
        localTime := gm.UTCTime.In(tz)
        times[tz.String()] = gotime.Format(localTime, "yyyy-mm-dd hh:ii:ss")
    }

    return times
}
```

## Advanced Utility Patterns

### 1. Batch Date Processing

```go
package batch

import (
    "github.com/maniartech/gotime"
    "time"
)

type DateProcessor struct {
    BatchSize int
    Timezone  *time.Location
}

func NewDateProcessor(batchSize int, tz *time.Location) *DateProcessor {
    return &DateProcessor{
        BatchSize: batchSize,
        Timezone:  tz,
    }
}

func (dp *DateProcessor) ProcessDates(dates []string, format string) ([]time.Time, []error) {
    var results []time.Time
    var errors []error

    for i := 0; i < len(dates); i += dp.BatchSize {
        end := i + dp.BatchSize
        if end > len(dates) {
            end = len(dates)
        }

        batch := dates[i:end]
        batchResults, batchErrors := dp.processBatch(batch, format)

        results = append(results, batchResults...)
        errors = append(errors, batchErrors...)
    }

    return results, errors
}

func (dp *DateProcessor) processBatch(dates []string, format string) ([]time.Time, []error) {
    var results []time.Time
    var errors []error

    for _, dateStr := range dates {
        parsed, err := gotime.ParseInLocation(dateStr, format, dp.Timezone)
        if err != nil {
            errors = append(errors, fmt.Errorf("failed to parse '%s': %v", dateStr, err))
            results = append(results, time.Time{})
        } else {
            results = append(results, parsed)
            errors = append(errors, nil)
        }
    }

    return results, errors
}

func (dp *DateProcessor) NormalizeDates(dates []time.Time) []time.Time {
    normalized := make([]time.Time, len(dates))

    for i, date := range dates {
        // Convert to target timezone and truncate time
        local := date.In(dp.Timezone)
        normalized[i] = gotime.TruncateTime(local)
    }

    return normalized
}
```

### 2. Date Range Generator

```go
package generator

type DateRange struct {
    Start time.Time
    End   time.Time
    Step  time.Duration
}

func NewDateRange(start, end time.Time, step time.Duration) *DateRange {
    return &DateRange{
        Start: start,
        End:   end,
        Step:  step,
    }
}

func (dr *DateRange) GenerateDates() []time.Time {
    var dates []time.Time
    current := dr.Start

    for !current.After(dr.End) {
        dates = append(dates, current)
        current = current.Add(dr.Step)
    }

    return dates
}

func (dr *DateRange) GenerateBusinessDays() []time.Time {
    var businessDays []time.Time
    current := gotime.TruncateTime(dr.Start)
    end := gotime.TruncateTime(dr.End)

    for !current.After(end) {
        if gotime.IsWeekday(current) {
            businessDays = append(businessDays, current)
        }
        current = gotime.Days(1, current)
    }

    return businessDays
}

func (dr *DateRange) GenerateMonthlyBoundaries() []time.Time {
    var boundaries []time.Time
    current := gotime.MonthStart(dr.Start)
    end := gotime.MonthEnd(dr.End)

    for !current.After(end) {
        boundaries = append(boundaries, current)
        // Move to next month
        if current.Month() == 12 {
            current = time.Date(current.Year()+1, 1, 1, 0, 0, 0, 0, current.Location())
        } else {
            current = time.Date(current.Year(), current.Month()+1, 1, 0, 0, 0, 0, current.Location())
        }
    }

    return boundaries
}

// Usage examples
func ExampleDateGeneration() {
    start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

    // Generate all days in year
    daily := NewDateRange(start, end, 24*time.Hour)
    allDays := daily.GenerateDates()
    fmt.Printf("Days in 2025: %d\n", len(allDays))

    // Generate business days only
    businessDays := daily.GenerateBusinessDays()
    fmt.Printf("Business days in 2025: %d\n", len(businessDays))

    // Generate month boundaries
    monthBoundaries := daily.GenerateMonthlyBoundaries()
    fmt.Printf("Months in 2025: %d\n", len(monthBoundaries))
}
```

### 3. Time Series Helper

```go
package timeseries

type TimeSeriesData struct {
    Timestamp time.Time
    Value     float64
    Tags      map[string]string
}

type TimeSeriesProcessor struct {
    Data     []TimeSeriesData
    Timezone *time.Location
}

func NewTimeSeriesProcessor(tz *time.Location) *TimeSeriesProcessor {
    return &TimeSeriesProcessor{
        Data:     make([]TimeSeriesData, 0),
        Timezone: tz,
    }
}

func (tsp *TimeSeriesProcessor) AddDataPoint(timestamp time.Time, value float64, tags map[string]string) {
    tsp.Data = append(tsp.Data, TimeSeriesData{
        Timestamp: timestamp.In(tsp.Timezone),
        Value:     value,
        Tags:      tags,
    })
}

func (tsp *TimeSeriesProcessor) GroupByPeriod(period string) map[time.Time][]TimeSeriesData {
    groups := make(map[time.Time][]TimeSeriesData)

    for _, data := range tsp.Data {
        var key time.Time

        switch period {
        case "day":
            key = gotime.TruncateTime(data.Timestamp)
        case "week":
            key = gotime.WeekStart(data.Timestamp)
        case "month":
            key = gotime.MonthStart(data.Timestamp)
        case "hour":
            key = data.Timestamp.Truncate(time.Hour)
        default:
            key = gotime.TruncateTime(data.Timestamp)
        }

        groups[key] = append(groups[key], data)
    }

    return groups
}

func (tsp *TimeSeriesProcessor) CalculateAggregates(period string) map[time.Time]map[string]float64 {
    groups := tsp.GroupByPeriod(period)
    aggregates := make(map[time.Time]map[string]float64)

    for timestamp, dataPoints := range groups {
        if len(dataPoints) == 0 {
            continue
        }

        var sum, min, max float64
        first := true

        for _, point := range dataPoints {
            sum += point.Value

            if first {
                min = point.Value
                max = point.Value
                first = false
            } else {
                if point.Value < min {
                    min = point.Value
                }
                if point.Value > max {
                    max = point.Value
                }
            }
        }

        count := float64(len(dataPoints))
        avg := sum / count

        aggregates[timestamp] = map[string]float64{
            "count":   count,
            "sum":     sum,
            "average": avg,
            "min":     min,
            "max":     max,
        }
    }

    return aggregates
}

func (tsp *TimeSeriesProcessor) ExportToCSV(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write header
    header := []string{"timestamp", "value"}

    // Add tag columns
    tagKeys := make(map[string]bool)
    for _, data := range tsp.Data {
        for key := range data.Tags {
            tagKeys[key] = true
        }
    }

    for key := range tagKeys {
        header = append(header, key)
    }

    writer.Write(header)

    // Write data
    for _, data := range tsp.Data {
        record := []string{
            gotime.Format(data.Timestamp, "yyyy-mm-dd hh:ii:ss"),
            fmt.Sprintf("%.2f", data.Value),
        }

        for key := range tagKeys {
            value := data.Tags[key]
            record = append(record, value)
        }

        writer.Write(record)
    }

    return nil
}
```

## Performance Optimization

### 1. Time Caching

```go
type TimeCache struct {
    cache map[string]time.Time
    mutex sync.RWMutex
}

func NewTimeCache() *TimeCache {
    return &TimeCache{
        cache: make(map[string]time.Time),
    }
}

func (tc *TimeCache) GetOrParse(dateStr, format string) (time.Time, error) {
    key := fmt.Sprintf("%s|%s", dateStr, format)

    tc.mutex.RLock()
    if cached, exists := tc.cache[key]; exists {
        tc.mutex.RUnlock()
        return cached, nil
    }
    tc.mutex.RUnlock()

    parsed, err := gotime.Parse(dateStr, format)
    if err != nil {
        return time.Time{}, err
    }

    tc.mutex.Lock()
    tc.cache[key] = parsed
    tc.mutex.Unlock()

    return parsed, nil
}
```

### 2. Bulk Operations

```go
func BulkTruncateTime(times []time.Time) []time.Time {
    results := make([]time.Time, len(times))
    for i, t := range times {
        results[i] = gotime.TruncateTime(t)
    }
    return results
}

func BulkConvertTimezone(times []time.Time, fromTZ, toTZ *time.Location) []time.Time {
    results := make([]time.Time, len(times))
    for i, t := range times {
        results[i] = gotime.ConvertTimezone(t, fromTZ, toTZ)
    }
    return results
}
```

---

## Summary

GoTime's utility functions provide:

- **Time manipulation** - TruncateTime, DateValue, period boundaries
- **Data conversion** - ISO 8601, timezone handling, format conversion
- **Helper functions** - Leap year, days in month, weekend detection
- **Advanced patterns** - Batch processing, date generation, time series
- **Performance optimization** - Caching, bulk operations

These utilities are essential for:
- Data preprocessing and normalization
- Report generation with proper time boundaries
- Multi-timezone application development
- Time series data analysis
- Performance-critical date operations

---

[Back to API Reference](../api-reference/) | [Getting Started](../getting-started/)
