[Home](../README.md) > [Examples](README.md) > Common Use Cases

# Common Use Cases

This guide demonstrates practical, real-world applications of GoTime across different domains and scenarios.

## Web Development

### 1. REST API Date Handling

```go
package api

import (
    "encoding/json"
    "github.com/maniartech/gotime"
    "net/http"
    "time"
)

// Standard API response format
type APIResponse struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    PublishAt string `json:"publish_at,omitempty"`
}

// Internal model with proper time.Time
type Article struct {
    ID        int
    Title     string
    CreatedAt time.Time
    UpdatedAt time.Time
    PublishAt *time.Time
}

const APIDateFormat = "yyyy-mm-dd hh:ii:ss"

func (a Article) ToAPIResponse() APIResponse {
    response := APIResponse{
        ID:        a.ID,
        Title:     a.Title,
        CreatedAt: gotime.Format(a.CreatedAt, APIDateFormat),
        UpdatedAt: gotime.Format(a.UpdatedAt, APIDateFormat),
    }

    if a.PublishAt != nil {
        response.PublishAt = gotime.Format(*a.PublishAt, APIDateFormat)
    }

    return response
}

func (a *Article) FromAPIRequest(data APIResponse) error {
    a.ID = data.ID
    a.Title = data.Title

    var err error
    a.CreatedAt, err = gotime.Parse(data.CreatedAt, APIDateFormat)
    if err != nil {
        return fmt.Errorf("invalid created_at: %v", err)
    }

    a.UpdatedAt, err = gotime.Parse(data.UpdatedAt, APIDateFormat)
    if err != nil {
        return fmt.Errorf("invalid updated_at: %v", err)
    }

    if data.PublishAt != "" {
        publishTime, err := gotime.Parse(data.PublishAt, APIDateFormat)
        if err != nil {
            return fmt.Errorf("invalid publish_at: %v", err)
        }
        a.PublishAt = &publishTime
    }

    return nil
}

// HTTP handlers
func createArticleHandler(w http.ResponseWriter, r *http.Request) {
    var req APIResponse
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    article := &Article{}
    if err := article.FromAPIRequest(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Set timestamps
    now := time.Now()
    article.CreatedAt = now
    article.UpdatedAt = now

    // Save to database...

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(article.ToAPIResponse())
}
```

### 2. User-Friendly Timestamps

```go
package frontend

import (
    "github.com/maniartech/gotime"
    "time"
)

type DisplayTimestamp struct {
    Absolute string `json:"absolute"`
    Relative string `json:"relative"`
    Tooltip  string `json:"tooltip"`
}

func FormatUserTimestamp(t time.Time) DisplayTimestamp {
    now := time.Now()
    diff := now.Sub(t)

    var absolute, relative, tooltip string

    switch {
    case diff < time.Minute:
        absolute = "Just now"
        relative = "Just now"
        tooltip = gotime.Format(t, "yyyy-mm-dd hh:ii:ss")

    case diff < time.Hour:
        absolute = gotime.TimeAgo(t)
        relative = gotime.TimeAgo(t)
        tooltip = gotime.Format(t, "yyyy-mm-dd hh:ii:ss")

    case diff < 24*time.Hour:
        absolute = gotime.Format(t, "h:ii aa")
        relative = gotime.TimeAgo(t)
        tooltip = gotime.Format(t, "wwww, mmmm dt, yyyy at h:ii aa")

    case diff < 7*24*time.Hour:
        absolute = gotime.Format(t, "wwww")
        relative = gotime.TimeAgo(t)
        tooltip = gotime.Format(t, "wwww, mmmm dt, yyyy at h:ii aa")

    case t.Year() == now.Year():
        absolute = gotime.Format(t, "mmmm dt")
        relative = gotime.TimeAgo(t)
        tooltip = gotime.Format(t, "wwww, mmmm dt, yyyy at h:ii aa")

    default:
        absolute = gotime.Format(t, "mmmm dt, yyyy")
        relative = gotime.TimeAgo(t)
        tooltip = gotime.Format(t, "wwww, mmmm dt, yyyy at h:ii aa")
    }

    return DisplayTimestamp{
        Absolute: absolute,
        Relative: relative,
        Tooltip:  tooltip,
    }
}

// Usage in templates:
// <span title="{{ .Tooltip }}">{{ .Absolute }}</span>
// Or dynamically switch between absolute/relative
```

### 3. Form Date Validation

```go
package forms

import (
    "github.com/maniartech/gotime"
    "time"
)

type DateValidator struct {
    AllowedFormats []string
    MinDate        *time.Time
    MaxDate        *time.Time
}

func NewDateValidator() *DateValidator {
    return &DateValidator{
        AllowedFormats: []string{
            "yyyy-mm-dd",
            "mm/dd/yyyy",
            "dd/mm/yyyy",
            "m/d/yyyy",
            "d/m/yyyy",
            "mmmm dt, yyyy",
        },
    }
}

func (dv *DateValidator) WithRange(min, max time.Time) *DateValidator {
    dv.MinDate = &min
    dv.MaxDate = &max
    return dv
}

func (dv *DateValidator) Validate(input string) (time.Time, error) {
    // Try parsing with each allowed format
    var parsed time.Time
    var err error

    for _, format := range dv.AllowedFormats {
        if parsed, err = gotime.Parse(input, format); err == nil {
            break
        }
    }

    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format: %s", input)
    }

    // Check range constraints
    if dv.MinDate != nil && parsed.Before(*dv.MinDate) {
        return time.Time{}, fmt.Errorf("date must be after %s",
            gotime.Format(*dv.MinDate, "mmmm dt, yyyy"))
    }

    if dv.MaxDate != nil && parsed.After(*dv.MaxDate) {
        return time.Time{}, fmt.Errorf("date must be before %s",
            gotime.Format(*dv.MaxDate, "mmmm dt, yyyy"))
    }

    return parsed, nil
}

// Usage in form handlers
func validateBirthDate(input string) (time.Time, error) {
    validator := NewDateValidator().WithRange(
        time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
        gotime.Yesterday(), // Must be in the past
    )
    return validator.Validate(input)
}

func validateEventDate(input string) (time.Time, error) {
    validator := NewDateValidator().WithRange(
        gotime.SoD(), // Must be today or future
        gotime.Years(5), // Within 5 years
    )
    return validator.Validate(input)
}
```

## Database Integration

### 1. Custom Database Types

```go
package models

import (
    "database/sql/driver"
    "github.com/maniartech/gotime"
    "time"
)

// Custom Date type for database interaction
type Date struct {
    time.Time
}

func (d *Date) Scan(value interface{}) error {
    switch v := value.(type) {
    case string:
        formats := []string{
            "yyyy-mm-dd hh:ii:ss",
            "yyyy-mm-dd",
            time.RFC3339,
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

func (d Date) MarshalJSON() ([]byte, error) {
    if d.Time.IsZero() {
        return []byte("null"), nil
    }
    formatted := gotime.Format(d.Time, "yyyy-mm-dd")
    return []byte(`"` + formatted + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        d.Time = time.Time{}
        return nil
    }

    // Remove quotes
    dateStr := string(data[1 : len(data)-1])

    parsed, err := gotime.Parse(dateStr, "yyyy-mm-dd")
    if err != nil {
        return err
    }

    d.Time = parsed
    return nil
}

// Usage in models
type User struct {
    ID        int    `json:"id" db:"id"`
    Name      string `json:"name" db:"name"`
    BirthDate Date   `json:"birth_date" db:"birth_date"`
    CreatedAt Date   `json:"created_at" db:"created_at"`
}
```

### 2. Query Builder with Date Filters

```go
package database

import (
    "github.com/maniartech/gotime"
    "strings"
    "time"
)

type QueryBuilder struct {
    table      string
    conditions []string
    params     []interface{}
}

func NewQuery(table string) *QueryBuilder {
    return &QueryBuilder{table: table}
}

func (q *QueryBuilder) WhereDateEquals(column, date, format string) *QueryBuilder {
    parsed, err := gotime.Parse(date, format)
    if err != nil {
        return q // Skip invalid dates
    }

    q.conditions = append(q.conditions, fmt.Sprintf("%s = ?", column))
    q.params = append(q.params, gotime.Format(parsed, "yyyy-mm-dd"))
    return q
}

func (q *QueryBuilder) WhereDateRange(column, startDate, endDate, format string) *QueryBuilder {
    start, err1 := gotime.Parse(startDate, format)
    end, err2 := gotime.Parse(endDate, format)

    if err1 != nil || err2 != nil {
        return q // Skip invalid dates
    }

    q.conditions = append(q.conditions, fmt.Sprintf("%s BETWEEN ? AND ?", column))
    q.params = append(q.params,
        gotime.Format(start, "yyyy-mm-dd"),
        gotime.Format(end, "yyyy-mm-dd"))
    return q
}

func (q *QueryBuilder) WhereRelativeDate(column string, offset string) *QueryBuilder {
    var targetDate time.Time

    switch offset {
    case "today":
        targetDate = gotime.SoD()
    case "yesterday":
        targetDate = gotime.Yesterday()
    case "last_week":
        targetDate = gotime.LastWeek()
    case "this_month":
        targetDate = gotime.MonthStart()
    case "last_month":
        targetDate = gotime.MonthStart(gotime.LastMonth())
    default:
        return q // Unknown offset
    }

    q.conditions = append(q.conditions, fmt.Sprintf("%s >= ?", column))
    q.params = append(q.params, gotime.Format(targetDate, "yyyy-mm-dd"))
    return q
}

func (q *QueryBuilder) Build() (string, []interface{}) {
    query := fmt.Sprintf("SELECT * FROM %s", q.table)

    if len(q.conditions) > 0 {
        query += " WHERE " + strings.Join(q.conditions, " AND ")
    }

    return query, q.params
}

// Usage
query, params := NewQuery("orders").
    WhereDateRange("created_at", "2025-01-01", "2025-12-31", "yyyy-mm-dd").
    WhereRelativeDate("updated_at", "last_week").
    Build()
```

## Business Applications

### 1. Invoice Due Date Calculation

```go
package billing

import (
    "github.com/maniartech/gotime"
    "time"
)

type PaymentTerms int

const (
    Net15 PaymentTerms = iota
    Net30
    Net45
    Net60
    Net90
)

type Invoice struct {
    ID          string
    Amount      float64
    IssueDate   time.Time
    DueDate     time.Time
    Terms       PaymentTerms
    PaidDate    *time.Time
}

func NewInvoice(id string, amount float64, terms PaymentTerms) *Invoice {
    issueDate := time.Now()

    var dueDate time.Time
    switch terms {
    case Net15:
        dueDate = gotime.Days(15, issueDate)
    case Net30:
        dueDate = gotime.Days(30, issueDate)
    case Net45:
        dueDate = gotime.Days(45, issueDate)
    case Net60:
        dueDate = gotime.Days(60, issueDate)
    case Net90:
        dueDate = gotime.Days(90, issueDate)
    }

    return &Invoice{
        ID:        id,
        Amount:    amount,
        IssueDate: issueDate,
        DueDate:   dueDate,
        Terms:     terms,
    }
}

func (inv *Invoice) IsOverdue() bool {
    return inv.PaidDate == nil && time.Now().After(inv.DueDate)
}

func (inv *Invoice) DaysUntilDue() int {
    if inv.PaidDate != nil {
        return 0 // Already paid
    }

    now := time.Now()
    if now.After(inv.DueDate) {
        return -int(now.Sub(inv.DueDate).Hours() / 24) // Negative for overdue
    }

    return int(inv.DueDate.Sub(now).Hours() / 24)
}

func (inv *Invoice) StatusDescription() string {
    if inv.PaidDate != nil {
        paidWhen := gotime.TimeAgo(*inv.PaidDate)
        return fmt.Sprintf("Paid %s", paidWhen)
    }

    if inv.IsOverdue() {
        overdueFor := gotime.TimeAgo(inv.DueDate)
        return fmt.Sprintf("Overdue %s", overdueFor)
    }

    dueWhen := gotime.TimeAgo(inv.DueDate)
    return fmt.Sprintf("Due %s", dueWhen)
}

func (inv *Invoice) FormatForDisplay() map[string]string {
    return map[string]string{
        "issue_date":    gotime.Format(inv.IssueDate, "mmmm dt, yyyy"),
        "due_date":      gotime.Format(inv.DueDate, "mmmm dt, yyyy"),
        "status":        inv.StatusDescription(),
        "days_until":    fmt.Sprintf("%d days", inv.DaysUntilDue()),
    }
}
```

### 2. Employee Schedule Management

```go
package schedule

import (
    "github.com/maniartech/gotime"
    "time"
)

type Shift struct {
    EmployeeID int
    Date       time.Time
    StartTime  time.Time
    EndTime    time.Time
    Position   string
}

type ScheduleManager struct {
    workingDays [7]bool // Sunday = 0, Monday = 1, etc.
    holidays    []time.Time
}

func NewScheduleManager() *ScheduleManager {
    // Default: Monday to Friday
    workingDays := [7]bool{false, true, true, true, true, true, false}

    return &ScheduleManager{
        workingDays: workingDays,
        holidays:    []time.Time{},
    }
}

func (sm *ScheduleManager) AddHoliday(date time.Time) {
    sm.holidays = append(sm.holidays, gotime.SoD(date))
}

func (sm *ScheduleManager) IsWorkingDay(date time.Time) bool {
    // Check if it's a configured working day
    if !sm.workingDays[date.Weekday()] {
        return false
    }

    // Check if it's a holiday
    dateOnly := gotime.SoD(date)
    for _, holiday := range sm.holidays {
        if dateOnly.Equal(holiday) {
            return false
        }
    }

    return true
}

func (sm *ScheduleManager) GenerateWeeklySchedule(startDate time.Time, employees []int) []Shift {
    var shifts []Shift
    weekStart := gotime.WeekStart(startDate)

    for day := 0; day < 7; day++ {
        currentDate := gotime.Days(day, weekStart)

        if !sm.IsWorkingDay(currentDate) {
            continue
        }

        // Create shifts for each employee
        for _, empID := range employees {
            shift := Shift{
                EmployeeID: empID,
                Date:       currentDate,
                StartTime:  time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 9, 0, 0, 0, currentDate.Location()),
                EndTime:    time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 17, 0, 0, 0, currentDate.Location()),
                Position:   "Regular",
            }
            shifts = append(shifts, shift)
        }
    }

    return shifts
}

func (sm *ScheduleManager) GetNextWorkingDay(from time.Time) time.Time {
    nextDay := gotime.Days(1, from)

    for !sm.IsWorkingDay(nextDay) {
        nextDay = gotime.Days(1, nextDay)
    }

    return nextDay
}

func (sm *ScheduleManager) GetWorkingDaysInMonth(year, month int) []time.Time {
    var workingDays []time.Time

    monthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    monthEnd := gotime.MonthEnd(monthStart)

    current := monthStart
    for !current.After(monthEnd) {
        if sm.IsWorkingDay(current) {
            workingDays = append(workingDays, current)
        }
        current = gotime.Days(1, current)
    }

    return workingDays
}
```

### 3. Project Timeline Management

```go
package projects

import (
    "github.com/maniartech/gotime"
    "time"
)

type Task struct {
    ID           string
    Name         string
    StartDate    time.Time
    EndDate      time.Time
    Dependencies []string
    Status       string
}

type Project struct {
    ID         string
    Name       string
    StartDate  time.Time
    EndDate    time.Time
    Tasks      []Task
}

func (p *Project) GetMilestones() []string {
    var milestones []string

    now := time.Now()

    // Project milestones
    milestones = append(milestones, fmt.Sprintf("Project started %s",
        gotime.TimeAgo(p.StartDate)))

    quarterPoints := []float64{0.25, 0.5, 0.75}
    duration := p.EndDate.Sub(p.StartDate)

    for _, point := range quarterPoints {
        milestoneDate := p.StartDate.Add(time.Duration(float64(duration) * point))
        if milestoneDate.Before(now) {
            milestones = append(milestones, fmt.Sprintf("%.0f%% milestone reached %s",
                point*100, gotime.TimeAgo(milestoneDate)))
        } else {
            milestones = append(milestones, fmt.Sprintf("%.0f%% milestone %s",
                point*100, gotime.TimeAgo(milestoneDate)))
        }
    }

    if p.EndDate.After(now) {
        milestones = append(milestones, fmt.Sprintf("Project completion %s",
            gotime.TimeAgo(p.EndDate)))
    } else {
        milestones = append(milestones, fmt.Sprintf("Project completed %s",
            gotime.TimeAgo(p.EndDate)))
    }

    return milestones
}

func (p *Project) GetTasksByStatus() map[string][]Task {
    taskMap := make(map[string][]Task)

    for _, task := range p.Tasks {
        taskMap[task.Status] = append(taskMap[task.Status], task)
    }

    return taskMap
}

func (p *Project) GetUpcomingDeadlines(days int) []Task {
    var upcoming []Task
    cutoff := gotime.Days(days)

    for _, task := range p.Tasks {
        if task.Status != "completed" && task.EndDate.Before(cutoff) {
            upcoming = append(upcoming, task)
        }
    }

    return upcoming
}

func (p *Project) GenerateStatusReport() map[string]interface{} {
    now := time.Now()

    totalTasks := len(p.Tasks)
    completedTasks := len(p.GetTasksByStatus()["completed"])

    progress := float64(completedTasks) / float64(totalTasks) * 100

    return map[string]interface{}{
        "project_name":     p.Name,
        "start_date":       gotime.Format(p.StartDate, "mmmm dt, yyyy"),
        "end_date":         gotime.Format(p.EndDate, "mmmm dt, yyyy"),
        "duration":         gotime.TimeAgo(p.StartDate, p.EndDate),
        "time_remaining":   gotime.TimeAgo(p.EndDate),
        "progress":         fmt.Sprintf("%.1f%%", progress),
        "total_tasks":      totalTasks,
        "completed_tasks":  completedTasks,
        "upcoming_deadlines": len(p.GetUpcomingDeadlines(7)),
        "status":           p.getProjectStatus(),
    }
}

func (p *Project) getProjectStatus() string {
    now := time.Now()

    if now.After(p.EndDate) {
        return "Completed"
    }

    if now.Before(p.StartDate) {
        return "Not Started"
    }

    return "In Progress"
}
```

## Analytics and Reporting

### 1. Time-Series Data Processing

```go
package analytics

import (
    "github.com/maniartech/gotime"
    "time"
)

type DataPoint struct {
    Timestamp time.Time
    Value     float64
    Category  string
}

type TimeSeriesAnalyzer struct {
    data []DataPoint
}

func NewTimeSeriesAnalyzer(data []DataPoint) *TimeSeriesAnalyzer {
    return &TimeSeriesAnalyzer{data: data}
}

func (tsa *TimeSeriesAnalyzer) GroupByPeriod(period string) map[string][]DataPoint {
    groups := make(map[string][]DataPoint)

    for _, point := range tsa.data {
        var key string

        switch period {
        case "hour":
            key = gotime.Format(point.Timestamp, "yyyy-mm-dd hh")
        case "day":
            key = gotime.Format(point.Timestamp, "yyyy-mm-dd")
        case "week":
            weekStart := gotime.WeekStart(point.Timestamp)
            key = gotime.Format(weekStart, "yyyy-mm-dd")
        case "month":
            key = gotime.Format(point.Timestamp, "yyyy-mm")
        case "quarter":
            quarter := (int(point.Timestamp.Month()) - 1) / 3 + 1
            key = fmt.Sprintf("%d-Q%d", point.Timestamp.Year(), quarter)
        case "year":
            key = gotime.Format(point.Timestamp, "yyyy")
        default:
            key = gotime.Format(point.Timestamp, "yyyy-mm-dd")
        }

        groups[key] = append(groups[key], point)
    }

    return groups
}

func (tsa *TimeSeriesAnalyzer) GetPeriodSummary(period string) map[string]float64 {
    groups := tsa.GroupByPeriod(period)
    summary := make(map[string]float64)

    for key, points := range groups {
        var sum float64
        for _, point := range points {
            sum += point.Value
        }
        summary[key] = sum
    }

    return summary
}

func (tsa *TimeSeriesAnalyzer) GetTrends() map[string]interface{} {
    if len(tsa.data) < 2 {
        return map[string]interface{}{"error": "insufficient data"}
    }

    // Sort by timestamp
    data := make([]DataPoint, len(tsa.data))
    copy(data, tsa.data)

    earliest := data[0].Timestamp
    latest := data[len(data)-1].Timestamp

    for i := 1; i < len(data); i++ {
        if data[i].Timestamp.Before(earliest) {
            earliest = data[i].Timestamp
        }
        if data[i].Timestamp.After(latest) {
            latest = data[i].Timestamp
        }
    }

    return map[string]interface{}{
        "period_start":    gotime.Format(earliest, "yyyy-mm-dd"),
        "period_end":      gotime.Format(latest, "yyyy-mm-dd"),
        "duration":        gotime.TimeAgo(earliest, latest),
        "data_points":     len(data),
        "average_per_day": tsa.calculateDailyAverage(),
    }
}

func (tsa *TimeSeriesAnalyzer) calculateDailyAverage() float64 {
    dailyGroups := tsa.GroupByPeriod("day")

    var totalDays int
    var totalValue float64

    for _, points := range dailyGroups {
        totalDays++
        for _, point := range points {
            totalValue += point.Value
        }
    }

    if totalDays == 0 {
        return 0
    }

    return totalValue / float64(totalDays)
}
```

### 2. Performance Monitoring

```go
package monitoring

import (
    "github.com/maniartech/gotime"
    "sync"
    "time"
)

type Metric struct {
    Name      string
    Value     float64
    Timestamp time.Time
    Tags      map[string]string
}

type MetricsCollector struct {
    metrics []Metric
    mu      sync.RWMutex
}

func NewMetricsCollector() *MetricsCollector {
    return &MetricsCollector{
        metrics: make([]Metric, 0),
    }
}

func (mc *MetricsCollector) Record(name string, value float64, tags map[string]string) {
    mc.mu.Lock()
    defer mc.mu.Unlock()

    metric := Metric{
        Name:      name,
        Value:     value,
        Timestamp: time.Now(),
        Tags:      tags,
    }

    mc.metrics = append(mc.metrics, metric)
}

func (mc *MetricsCollector) GetRecentMetrics(duration time.Duration) []Metric {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    cutoff := time.Now().Add(-duration)
    var recent []Metric

    for _, metric := range mc.metrics {
        if metric.Timestamp.After(cutoff) {
            recent = append(recent, metric)
        }
    }

    return recent
}

func (mc *MetricsCollector) GenerateReport() map[string]interface{} {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    if len(mc.metrics) == 0 {
        return map[string]interface{}{"error": "no metrics collected"}
    }

    now := time.Now()

    // Find time range
    earliest := mc.metrics[0].Timestamp
    latest := mc.metrics[0].Timestamp

    for _, metric := range mc.metrics {
        if metric.Timestamp.Before(earliest) {
            earliest = metric.Timestamp
        }
        if metric.Timestamp.After(latest) {
            latest = metric.Timestamp
        }
    }

    // Group by time periods
    last5min := mc.getMetricsInRange(now.Add(-5*time.Minute), now)
    lastHour := mc.getMetricsInRange(now.Add(-time.Hour), now)
    today := mc.getMetricsInRange(gotime.SoD(), now)

    return map[string]interface{}{
        "collection_period": map[string]string{
            "start":    gotime.Format(earliest, "yyyy-mm-dd hh:ii:ss"),
            "end":      gotime.Format(latest, "yyyy-mm-dd hh:ii:ss"),
            "duration": gotime.TimeAgo(earliest, latest),
        },
        "metrics_count": map[string]int{
            "total":        len(mc.metrics),
            "last_5_min":   len(last5min),
            "last_hour":    len(lastHour),
            "today":        len(today),
        },
        "averages": map[string]float64{
            "last_5_min": mc.calculateAverage(last5min),
            "last_hour":  mc.calculateAverage(lastHour),
            "today":      mc.calculateAverage(today),
        },
        "generated_at": gotime.Format(now, "yyyy-mm-dd hh:ii:ss"),
    }
}

func (mc *MetricsCollector) getMetricsInRange(start, end time.Time) []Metric {
    var filtered []Metric

    for _, metric := range mc.metrics {
        if metric.Timestamp.After(start) && metric.Timestamp.Before(end) {
            filtered = append(filtered, metric)
        }
    }

    return filtered
}

func (mc *MetricsCollector) calculateAverage(metrics []Metric) float64 {
    if len(metrics) == 0 {
        return 0
    }

    var sum float64
    for _, metric := range metrics {
        sum += metric.Value
    }

    return sum / float64(len(metrics))
}
```

## Content Management

### 1. Content Scheduling System

```go
package cms

import (
    "github.com/maniartech/gotime"
    "time"
)

type ContentStatus string

const (
    Draft     ContentStatus = "draft"
    Scheduled ContentStatus = "scheduled"
    Published ContentStatus = "published"
    Archived  ContentStatus = "archived"
)

type Content struct {
    ID          string
    Title       string
    Body        string
    Author      string
    Status      ContentStatus
    CreatedAt   time.Time
    UpdatedAt   time.Time
    ScheduledAt *time.Time
    PublishedAt *time.Time
    ExpiresAt   *time.Time
}

type ContentManager struct {
    content []Content
}

func NewContentManager() *ContentManager {
    return &ContentManager{
        content: make([]Content, 0),
    }
}

func (cm *ContentManager) ScheduleContent(content Content, when string) error {
    var scheduledTime time.Time
    var err error

    // Parse flexible scheduling options
    switch when {
    case "now":
        scheduledTime = time.Now()
    case "in_1_hour":
        scheduledTime = time.Now().Add(time.Hour)
    case "tomorrow":
        scheduledTime = gotime.SoD(gotime.Tomorrow())
    case "next_week":
        scheduledTime = gotime.SoD(gotime.NextWeek())
    case "next_month":
        scheduledTime = gotime.SoD(gotime.NextMonth())
    default:
        // Try parsing as date
        formats := []string{
            "yyyy-mm-dd",
            "yyyy-mm-dd hh:ii",
            "mm/dd/yyyy",
            "mm/dd/yyyy h:ii aa",
        }

        for _, format := range formats {
            if scheduledTime, err = gotime.Parse(when, format); err == nil {
                break
            }
        }

        if err != nil {
            return fmt.Errorf("invalid schedule time: %s", when)
        }
    }

    content.ScheduledAt = &scheduledTime
    content.Status = Scheduled
    content.UpdatedAt = time.Now()

    cm.content = append(cm.content, content)
    return nil
}

func (cm *ContentManager) GetContentByStatus(status ContentStatus) []Content {
    var filtered []Content

    for _, c := range cm.content {
        if c.Status == status {
            filtered = append(filtered, c)
        }
    }

    return filtered
}

func (cm *ContentManager) GetUpcomingContent() []Content {
    var upcoming []Content
    now := time.Now()

    for _, c := range cm.content {
        if c.Status == Scheduled && c.ScheduledAt != nil && c.ScheduledAt.After(now) {
            upcoming = append(upcoming, c)
        }
    }

    return upcoming
}

func (cm *ContentManager) GetExpiredContent() []Content {
    var expired []Content
    now := time.Now()

    for _, c := range cm.content {
        if c.ExpiresAt != nil && c.ExpiresAt.Before(now) {
            expired = append(expired, c)
        }
    }

    return expired
}

func (cm *ContentManager) GetContentDashboard() map[string]interface{} {
    now := time.Now()

    var totalContent, publishedContent, scheduledContent, expiredContent int
    var recentlyPublished []Content

    for _, c := range cm.content {
        totalContent++

        switch c.Status {
        case Published:
            publishedContent++
            if c.PublishedAt != nil && time.Since(*c.PublishedAt) < 24*time.Hour {
                recentlyPublished = append(recentlyPublished, c)
            }
        case Scheduled:
            scheduledContent++
        }

        if c.ExpiresAt != nil && c.ExpiresAt.Before(now) {
            expiredContent++
        }
    }

    upcoming := cm.GetUpcomingContent()

    return map[string]interface{}{
        "summary": map[string]int{
            "total":     totalContent,
            "published": publishedContent,
            "scheduled": scheduledContent,
            "expired":   expiredContent,
        },
        "recent_activity": map[string]interface{}{
            "published_today": len(recentlyPublished),
            "upcoming_count":  len(upcoming),
        },
        "next_publication": cm.getNextPublicationInfo(upcoming),
        "generated_at":     gotime.Format(now, "yyyy-mm-dd hh:ii:ss"),
    }
}

func (cm *ContentManager) getNextPublicationInfo(upcoming []Content) map[string]string {
    if len(upcoming) == 0 {
        return map[string]string{"message": "No upcoming publications"}
    }

    // Find earliest scheduled content
    earliest := upcoming[0]
    for _, c := range upcoming[1:] {
        if c.ScheduledAt.Before(*earliest.ScheduledAt) {
            earliest = c
        }
    }

    return map[string]string{
        "title":     earliest.Title,
        "scheduled": gotime.Format(*earliest.ScheduledAt, "yyyy-mm-dd hh:ii"),
        "relative":  gotime.TimeAgo(*earliest.ScheduledAt),
    }
}
```

---

## Summary

These common use cases demonstrate GoTime's versatility across different domains:

1. **Web Development**: API handling, user timestamps, form validation
2. **Database Integration**: Custom types, query builders, data migration
3. **Business Applications**: Invoicing, scheduling, project management
4. **Analytics**: Time-series analysis, performance monitoring
5. **Content Management**: Publishing schedules, content lifecycle

**Key Benefits Demonstrated:**
- **Intuitive formatting** reduces development time
- **Flexible parsing** handles diverse input formats
- **Rich relative time functions** simplify common calculations
- **Business logic integration** streamlines real-world applications
- **Consistent API** across all use cases

**Next Steps:**
- Explore [Advanced Usage](../advanced/) for optimization techniques
- Check out [Best Practices](../advanced/best-practices.md) for production guidance
- See [Real-world Examples](real-world.md) for complete application examples

---

Continue to: [Real-world Examples](real-world.md) | [Advanced Usage](../advanced/)
