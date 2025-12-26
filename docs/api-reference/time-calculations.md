[Home](../README.md) > [API Reference](README.md) > Time Calculations

# Time Calculations

GoTime provides advanced functions for working with time differences, business days, date comparisons, and age calculations.

## Age Calculation Functions

### Age

Calculates the precise age in years, months, and days between a birth date and a reference date.

```go
func Age(birthDate time.Time, asOf ...time.Time) (years, months, days int)
```

**Parameters:**
- `birthDate`: The birth date
- `asOf`: Optional reference date (uses current time if not provided)

**Returns:**
- `years`: Complete years
- `months`: Additional months after years
- `days`: Additional days after years and months

**Examples:**

```go
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
asOf := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)

years, months, days := gotime.Age(birthDate, asOf)
fmt.Printf("Age: %d years, %d months, %d days", years, months, days)
// Output: Age: 35 years, 1 months, 22 days

// Use current time as reference
years, months, days = gotime.Age(birthDate)
// Calculates age as of now

// Handle leap year birthdays correctly
leapBirthDate := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
nonLeapRef := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
years, months, days = gotime.Age(leapBirthDate, nonLeapRef)
// Correctly handles Feb 29 -> Mar 1 in non-leap years
```

### YearsBetween

Calculates the precise number of years between two dates as a float64.

```go
func YearsBetween(start, end time.Time) float64
```

**Examples:**

```go
start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

years := gotime.YearsBetween(start, end)
fmt.Printf("%.1f years", years) // Output: 5.5 years

// Works with reverse order
years = gotime.YearsBetween(end, start) // Same result: 5.5 years

// Precise calculations
halfYear := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
years = gotime.YearsBetween(start, halfYear) // 0.5 years
```

### MonthsBetween

Calculates the precise number of months between two dates as a float64.

```go
func MonthsBetween(start, end time.Time) float64
```

**Examples:**

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC)

months := gotime.MonthsBetween(start, end)
fmt.Printf("%.1f months", months) // Output: 6.5 months

// Exact month calculations
oneMonth := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
months = gotime.MonthsBetween(start, oneMonth) // 1.0 months
```

### DaysBetween

Calculates the number of days between two dates.

```go
func DaysBetween(start, end time.Time) int
```

**Examples:**

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC)

days := gotime.DaysBetween(start, end)
fmt.Println(days) // Output: 7

// Handles leap years correctly
leapStart := time.Date(2020, 2, 28, 0, 0, 0, 0, time.UTC)
leapEnd := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
days = gotime.DaysBetween(leapStart, leapEnd) // 2 days (includes Feb 29)
```

### WeeksBetween

Calculates the precise number of weeks between two dates as a float64.

```go
func WeeksBetween(start, end time.Time) float64
```

**Examples:**

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 11, 0, 0, 0, 0, time.UTC)

weeks := gotime.WeeksBetween(start, end)
fmt.Printf("%.2f weeks", weeks) // Output: 1.43 weeks

// Exact week calculations
oneWeek := time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC)
weeks = gotime.WeeksBetween(start, oneWeek) // 1.0 weeks
```

### DurationInWords

Returns a human-readable representation of a duration.

```go
func DurationInWords(d time.Duration) string
```

**Examples:**

```go
// Simple durations
fmt.Println(gotime.DurationInWords(30 * time.Second))           // "30 seconds"
fmt.Println(gotime.DurationInWords(1 * time.Minute))           // "1 minute"
fmt.Println(gotime.DurationInWords(2 * time.Hour))             // "2 hours"
fmt.Println(gotime.DurationInWords(24 * time.Hour))            // "1 day"

// Combined durations (shows top 2 units)
duration := 2*time.Hour + 30*time.Minute
fmt.Println(gotime.DurationInWords(duration))                  // "2 hours 30 minutes"

duration = 25 * time.Hour
fmt.Println(gotime.DurationInWords(duration))                  // "1 day 1 hour"

duration = 2*24*time.Hour + 3*time.Hour + 45*time.Minute
fmt.Println(gotime.DurationInWords(duration))                  // "2 days 3 hours"

// Negative durations
fmt.Println(gotime.DurationInWords(-2 * time.Hour))            // "-2 hours"

// Very small durations
fmt.Println(gotime.DurationInWords(500 * time.Millisecond))    // "less than 1 second"
```

### IsValidAge

Checks if the given birth date results in a valid age (not negative, not unreasonably old).

```go
func IsValidAge(birthDate time.Time, asOf ...time.Time) bool
```

**Parameters:**
- `birthDate`: The birth date to validate
- `asOf`: Optional reference date (uses current time if not provided)

**Returns:**
- `true` if the age is valid (0-150 years)
- `false` if the birth date is in the future or results in an unreasonable age

**Examples:**

```go
// Valid ages
validBirth := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(validBirth))                     // true

newborn := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(newborn))                        // true

elderly := time.Date(1925, 1, 1, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(elderly))                        // true

// Invalid ages
futureBirth := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(futureBirth))                    // false

tooOld := time.Date(1850, 1, 1, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(tooOld))                         // false

// With specific reference date
refDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.IsValidAge(validBirth, refDate))            // true
```

## Time Arithmetic Functions

### Hours

Adds or subtracts hours from a time.

```go
func Hours(h int, times ...time.Time) time.Time
```

**Parameters:**
- `h`: Number of hours to add (negative to subtract)
- `times`: Optional time(s) to use as base (uses current time if not provided)

**Examples:**

```go
baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

// Add hours
result := gotime.Hours(5, baseTime)
fmt.Println(result) // 2025-07-07 17:00:00 +0000 UTC

// Subtract hours
result = gotime.Hours(-2, baseTime)
fmt.Println(result) // 2025-07-07 10:00:00 +0000 UTC

// Cross day boundaries
result = gotime.Hours(25, baseTime)
fmt.Println(result) // 2025-07-08 13:00:00 +0000 UTC

// Use current time if no base time provided
result = gotime.Hours(1) // 1 hour from now
```

### Minutes

Adds or subtracts minutes from a time.

```go
func Minutes(m int, times ...time.Time) time.Time
```

**Examples:**

```go
baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

// Add minutes
result := gotime.Minutes(30, baseTime)
fmt.Println(result) // 2025-07-07 12:30:00 +0000 UTC

// Subtract minutes
result = gotime.Minutes(-15, baseTime)
fmt.Println(result) // 2025-07-07 11:45:00 +0000 UTC

// Cross hour boundaries
result = gotime.Minutes(90, baseTime)
fmt.Println(result) // 2025-07-07 13:30:00 +0000 UTC

// Use current time if no base time provided
result = gotime.Minutes(30) // 30 minutes from now
```

### Seconds

Adds or subtracts seconds from a time.

```go
func Seconds(s int, times ...time.Time) time.Time
```

**Examples:**

```go
baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

// Add seconds
result := gotime.Seconds(45, baseTime)
fmt.Println(result) // 2025-07-07 12:00:45 +0000 UTC

// Subtract seconds
result = gotime.Seconds(-30, baseTime)
fmt.Println(result) // 2025-07-07 11:59:30 +0000 UTC

// Cross minute boundaries
result = gotime.Seconds(90, baseTime)
fmt.Println(result) // 2025-07-07 12:01:30 +0000 UTC

// Use current time if no base time provided
result = gotime.Seconds(10) // 10 seconds from now
```

## Quarter Functions

### QuarterStart

Returns the start of the quarter for a given time.

```go
func QuarterStart(times ...time.Time) time.Time
```

**Examples:**

```go
// Get start of current quarter
start := gotime.QuarterStart()

// Get start of specific quarter
date := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC) // Q3
start = gotime.QuarterStart(date)
fmt.Println(start) // 2025-07-01 00:00:00 +0000 UTC

// Quarter boundaries:
// Q1: January 1   - March 31
// Q2: April 1     - June 30
// Q3: July 1      - September 30
// Q4: October 1   - December 31
```

### QuarterEnd

Returns the end of the quarter for a given time.

```go
func QuarterEnd(times ...time.Time) time.Time
```

**Examples:**

```go
// Get end of current quarter
end := gotime.QuarterEnd()

// Get end of specific quarter
date := time.Date(2025, 5, 15, 12, 30, 45, 0, time.UTC) // Q2
end = gotime.QuarterEnd(date)
fmt.Println(end) // 2025-06-30 23:59:59.999999999 +0000 UTC
```

### Quarters

Adds or subtracts quarters from a time.

```go
func Quarters(q int, times ...time.Time) time.Time
```

**Examples:**

```go
baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)

// Add quarters
future := gotime.Quarters(2, baseTime)
fmt.Println(future) // 2026-01-15 12:30:45 +0000 UTC

// Subtract quarters
past := gotime.Quarters(-1, baseTime)
fmt.Println(past) // 2025-04-15 12:30:45 +0000 UTC

// Handle end-of-month edge cases correctly
endOfMonth := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)
result := gotime.Quarters(1, endOfMonth) // November has only 30 days
fmt.Println(result) // 2025-11-30 00:00:00 +0000 UTC
```

### LastQuarter

Returns the same time in the previous quarter.

```go
func LastQuarter(times ...time.Time) time.Time
```

**Examples:**

```go
current := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC) // Q3
last := gotime.LastQuarter(current)
fmt.Println(last) // 2025-04-15 12:30:45 +0000 UTC (Q2)
```

### NextQuarter

Returns the same time in the next quarter.

```go
func NextQuarter(times ...time.Time) time.Time
```

**Examples:**

```go
current := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC) // Q3
next := gotime.NextQuarter(current)
fmt.Println(next) // 2025-10-15 12:30:45 +0000 UTC (Q4)
```

### QuarterOfYear

Returns the quarter number (1-4) for a given time.

```go
func QuarterOfYear(times ...time.Time) int
```

**Examples:**

```go
// January-March = Q1
q1 := time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.QuarterOfYear(q1)) // 1

// April-June = Q2
q2 := time.Date(2025, 5, 15, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.QuarterOfYear(q2)) // 2

// July-September = Q3
q3 := time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.QuarterOfYear(q3)) // 3

// October-December = Q4
q4 := time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC)
fmt.Println(gotime.QuarterOfYear(q4)) // 4

// Current quarter
fmt.Println(gotime.QuarterOfYear()) // Quarter of current time
```

## Time Arithmetic Use Cases

### 1. Appointment Scheduling System

```go
package scheduling

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type AppointmentScheduler struct {
    workingHours struct {
        start int // 9 for 9 AM
        end   int // 17 for 5 PM
    }
    appointmentDuration int // minutes
}

func NewScheduler() *AppointmentScheduler {
    return &AppointmentScheduler{
        workingHours: struct {
            start int
            end   int
        }{start: 9, end: 17},
        appointmentDuration: 30,
    }
}

func (s *AppointmentScheduler) NextAvailableSlot(after time.Time) time.Time {
    // Start looking from the next working hour
    candidate := s.nextWorkingHour(after)

    // For this example, assume all slots are available
    return candidate
}

func (s *AppointmentScheduler) nextWorkingHour(t time.Time) time.Time {
    hour := t.Hour()

    if hour < s.workingHours.start {
        // Before working hours - move to start of working day
        return time.Date(t.Year(), t.Month(), t.Day(), s.workingHours.start, 0, 0, 0, t.Location())
    } else if hour >= s.workingHours.end {
        // After working hours - move to next working day
        nextDay := gotime.Days(1, t)
        return time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), s.workingHours.start, 0, 0, 0, t.Location())
    }

    // During working hours - round up to next hour
    return gotime.Hours(1, time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, t.Location()))
}

func (s *AppointmentScheduler) ScheduleAppointment(preferredTime time.Time) time.Time {
    // Find next available slot
    slot := s.NextAvailableSlot(preferredTime)

    // Make sure it's during working hours
    hour := slot.Hour()
    if hour < s.workingHours.start || hour >= s.workingHours.end {
        slot = s.nextWorkingHour(slot)
    }

    return slot
}

func (s *AppointmentScheduler) GetAppointmentEnd(start time.Time) time.Time {
    return gotime.Minutes(s.appointmentDuration, start)
}

func (s *AppointmentScheduler) GetReminders(appointmentTime time.Time) map[string]time.Time {
    return map[string]time.Time{
        "24_hours":  gotime.Hours(-24, appointmentTime),
        "2_hours":   gotime.Hours(-2, appointmentTime),
        "15_minutes": gotime.Minutes(-15, appointmentTime),
    }
}
```

### 2. Shift Management System

```go
package shifts

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type Shift struct {
    ID        string
    StartTime time.Time
    Duration  int // hours
    BreakTime int // minutes
}

func (s *Shift) GetEndTime() time.Time {
    return gotime.Hours(s.Duration, s.StartTime)
}

func (s *Shift) GetBreakStart() time.Time {
    // Break starts halfway through shift
    return gotime.Hours(s.Duration/2, s.StartTime)
}

func (s *Shift) GetBreakEnd() time.Time {
    breakStart := s.GetBreakStart()
    return gotime.Minutes(s.BreakTime, breakStart)
}

func (s *Shift) GetActualWorkTime() time.Duration {
    totalHours := time.Duration(s.Duration) * time.Hour
    breakDuration := time.Duration(s.BreakTime) * time.Minute
    return totalHours - breakDuration
}

func (s *Shift) IsOvertime() bool {
    return s.Duration > 8 // More than 8 hours
}

func (s *Shift) GetOvertimeHours() int {
    if s.Duration > 8 {
        return s.Duration - 8
    }
    return 0
}

type ShiftManager struct {
    shifts []Shift
}

func (sm *ShiftManager) ScheduleShift(start time.Time, duration int) Shift {
    shift := Shift{
        ID:        generateShiftID(),
        StartTime: start,
        Duration:  duration,
        BreakTime: 30, // 30 minute break
    }

    sm.shifts = append(sm.shifts, shift)
    return shift
}

func (sm *ShiftManager) GetWeeklySchedule(weekStart time.Time) []Shift {
    var weeklyShifts []Shift
    weekEnd := gotime.Days(7, weekStart)

    for _, shift := range sm.shifts {
        if shift.StartTime.After(weekStart) && shift.StartTime.Before(weekEnd) {
            weeklyShifts = append(weeklyShifts, shift)
        }
    }

    return weeklyShifts
}

func generateShiftID() string {
    return gotime.Format(time.Now(), "yyyymmdd-hhiiss")
}
```

### 3. Financial Quarter Reporting

```go
package reporting

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type QuarterlyReport struct {
    Year     int
    Quarter  int
    Revenue  float64
    Expenses float64
}

func (qr *QuarterlyReport) GetProfit() float64 {
    return qr.Revenue - qr.Expenses
}

func (qr *QuarterlyReport) GetPeriodStart() time.Time {
    return gotime.QuarterStart(time.Date(qr.Year, time.Month((qr.Quarter-1)*3+1), 1, 0, 0, 0, 0, time.UTC))
}

func (qr *QuarterlyReport) GetPeriodEnd() time.Time {
    return gotime.QuarterEnd(time.Date(qr.Year, time.Month((qr.Quarter-1)*3+1), 1, 0, 0, 0, 0, time.UTC))
}

type FinancialReporting struct {
    reports []QuarterlyReport
}

func (fr *FinancialReporting) GetCurrentQuarter() int {
    return gotime.QuarterOfYear()
}

func (fr *FinancialReporting) GetQuarterStartDate(year, quarter int) time.Time {
    firstMonthOfQuarter := (quarter-1)*3 + 1
    someTimeInQuarter := time.Date(year, time.Month(firstMonthOfQuarter), 15, 0, 0, 0, 0, time.UTC)
    return gotime.QuarterStart(someTimeInQuarter)
}

func (fr *FinancialReporting) GetQuarterEndDate(year, quarter int) time.Time {
    firstMonthOfQuarter := (quarter-1)*3 + 1
    someTimeInQuarter := time.Date(year, time.Month(firstMonthOfQuarter), 15, 0, 0, 0, 0, time.UTC)
    return gotime.QuarterEnd(someTimeInQuarter)
}

func (fr *FinancialReporting) GetNextQuarterStart() time.Time {
    nextQuarterTime := gotime.NextQuarter()
    return gotime.QuarterStart(nextQuarterTime)
}

func (fr *FinancialReporting) GetLastQuarterReport() *QuarterlyReport {
    lastQuarterTime := gotime.LastQuarter()
    targetYear := lastQuarterTime.Year()
    targetQuarter := gotime.QuarterOfYear(lastQuarterTime)

    for _, report := range fr.reports {
        if report.Year == targetYear && report.Quarter == targetQuarter {
            return &report
        }
    }
    return nil
}

func (fr *FinancialReporting) GenerateQuarterComparison(year, quarter int) map[string]interface{} {
    current := fr.getQuarterReport(year, quarter)
    if current == nil {
        return map[string]interface{}{"error": "Current quarter data not found"}
    }

    // Get same quarter from previous year
    previous := fr.getQuarterReport(year-1, quarter)

    comparison := map[string]interface{}{
        "current_quarter": map[string]interface{}{
            "year":     current.Year,
            "quarter":  current.Quarter,
            "revenue":  current.Revenue,
            "expenses": current.Expenses,
            "profit":   current.GetProfit(),
            "start":    gotime.Format(current.GetPeriodStart(), "yyyy-mm-dd"),
            "end":      gotime.Format(current.GetPeriodEnd(), "yyyy-mm-dd"),
        },
    }

    if previous != nil {
        revenueGrowth := ((current.Revenue - previous.Revenue) / previous.Revenue) * 100
        profitGrowth := ((current.GetProfit() - previous.GetProfit()) / previous.GetProfit()) * 100

        comparison["previous_year"] = map[string]interface{}{
            "revenue":        previous.Revenue,
            "profit":         previous.GetProfit(),
            "revenue_growth": revenueGrowth,
            "profit_growth":  profitGrowth,
        }
    }

    return comparison
}

func (fr *FinancialReporting) getQuarterReport(year, quarter int) *QuarterlyReport {
    for _, report := range fr.reports {
        if report.Year == year && report.Quarter == quarter {
            return &report
        }
    }
    return nil
}
```

fmt.Println(gotime.IsValidAge(validBirth, refDate))            // true
```

## Age Calculation Use Cases

### 1. User Registration System

```go
package registration

import (
    "errors"
    "github.com/maniartech/gotime/v2"
    "time"
)

type User struct {
    Name      string
    BirthDate time.Time
    Email     string
}

func (u *User) GetAge() (int, int, int) {
    return gotime.Age(u.BirthDate)
}

func (u *User) GetAgeInYears() float64 {
    return gotime.YearsBetween(u.BirthDate, time.Now())
}

func (u *User) IsAdult() bool {
    years, _, _ := u.GetAge()
    return years >= 18
}

func (u *User) IsEligibleForSeniorDiscount() bool {
    years, _, _ := u.GetAge()
    return years >= 65
}

func ValidateUser(u *User) error {
    if !gotime.IsValidAge(u.BirthDate) {
        return errors.New("invalid birth date")
    }

    if !u.IsAdult() {
        return errors.New("user must be at least 18 years old")
    }

    return nil
}

func GetAgeGroup(birthDate time.Time) string {
    years, _, _ := gotime.Age(birthDate)

    switch {
    case years < 13:
        return "child"
    case years < 20:
        return "teenager"
    case years < 30:
        return "young_adult"
    case years < 50:
        return "adult"
    case years < 65:
        return "middle_aged"
    default:
        return "senior"
    }
}
```

### 2. Insurance Premium Calculator

```go
package insurance

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type InsuranceCalculator struct {
    BasePremium float64
}

func NewInsuranceCalculator(basePremium float64) *InsuranceCalculator {
    return &InsuranceCalculator{BasePremium: basePremium}
}

func (ic *InsuranceCalculator) CalculatePremium(birthDate time.Time) float64 {
    years, _, _ := gotime.Age(birthDate)
    premium := ic.BasePremium

    // Age-based multipliers
    switch {
    case years < 25:
        premium *= 1.5 // Higher risk for young drivers
    case years < 35:
        premium *= 1.0 // Standard rate
    case years < 55:
        premium *= 0.9 // Lower rate for experienced drivers
    case years < 70:
        premium *= 1.1 // Slightly higher for older drivers
    default:
        premium *= 1.8 // Much higher for elderly drivers
    }

    return premium
}

func (ic *InsuranceCalculator) GetRiskCategory(birthDate time.Time) string {
    years, _, _ := gotime.Age(birthDate)

    switch {
    case years < 25:
        return "high_risk"
    case years < 55:
        return "standard_risk"
    case years < 70:
        return "experienced"
    default:
        return "senior"
    }
}

func (ic *InsuranceCalculator) IsEligible(birthDate time.Time) bool {
    if !gotime.IsValidAge(birthDate) {
        return false
    }

    years, _, _ := gotime.Age(birthDate)
    return years >= 18 && years <= 80
}

func (ic *InsuranceCalculator) GetQuote(birthDate time.Time) map[string]interface{} {
    if !ic.IsEligible(birthDate) {
        return map[string]interface{}{
            "eligible": false,
            "reason":   "Age requirements not met",
        }
    }

    years, months, days := gotime.Age(birthDate)
    premium := ic.CalculatePremium(birthDate)
    category := ic.GetRiskCategory(birthDate)

    return map[string]interface{}{
        "eligible":     true,
        "age":          map[string]int{"years": years, "months": months, "days": days},
        "age_years":    gotime.YearsBetween(birthDate, time.Now()),
        "premium":      premium,
        "risk_category": category,
        "quote_date":   gotime.Format(time.Now(), "yyyy-mm-dd"),
    }
}
```

### 3. Employee Benefits System

```go
package benefits

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type Employee struct {
    ID          string
    Name        string
    BirthDate   time.Time
    HireDate    time.Time
    Salary      float64
}

func (e *Employee) GetAge() (int, int, int) {
    return gotime.Age(e.BirthDate)
}

func (e *Employee) GetServiceYears() float64 {
    return gotime.YearsBetween(e.HireDate, time.Now())
}

func (e *Employee) GetVacationDays() int {
    serviceYears := e.GetServiceYears()

    switch {
    case serviceYears < 1:
        return 10 // 10 days for new employees
    case serviceYears < 5:
        return 15 // 15 days after 1 year
    case serviceYears < 10:
        return 20 // 20 days after 5 years
    default:
        return 25 // 25 days after 10 years
    }
}

func (e *Employee) IsEligibleForRetirement() bool {
    age, _, _ := e.GetAge()
    serviceYears := e.GetServiceYears()

    // Rule of 85: age + service years >= 85, minimum age 55
    return age >= 55 && (float64(age) + serviceYears) >= 85
}

func (e *Employee) GetHealthInsuranceRate() float64 {
    age, _, _ := e.GetAge()

    baseRate := 150.0 // Base monthly premium

    switch {
    case age < 30:
        return baseRate * 0.8
    case age < 40:
        return baseRate * 1.0
    case age < 50:
        return baseRate * 1.2
    case age < 60:
        return baseRate * 1.5
    default:
        return baseRate * 1.8
    }
}

func (e *Employee) GetLifeInsuranceMultiplier() float64 {
    age, _, _ := e.GetAge()

    switch {
    case age < 30:
        return 3.0 // 3x salary
    case age < 40:
        return 2.5 // 2.5x salary
    case age < 50:
        return 2.0 // 2x salary
    case age < 60:
        return 1.5 // 1.5x salary
    default:
        return 1.0 // 1x salary
    }
}

func (e *Employee) GetBenefitsSummary() map[string]interface{} {
    years, months, days := e.GetAge()
    serviceYears := e.GetServiceYears()

    return map[string]interface{}{
        "employee_id":       e.ID,
        "name":             e.Name,
        "age":              map[string]int{"years": years, "months": months, "days": days},
        "service_years":    serviceYears,
        "vacation_days":    e.GetVacationDays(),
        "retirement_eligible": e.IsEligibleForRetirement(),
        "health_premium":   e.GetHealthInsuranceRate(),
        "life_insurance":   e.GetLifeInsuranceMultiplier() * e.Salary,
        "calculated_date":  gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss"),
    }
}
```

### 4. Medical Appointment System

```go
package medical

import (
    "github.com/maniartech/gotime/v2"
    "time"
)

type Patient struct {
    ID          string
    Name        string
    BirthDate   time.Time
    LastVisit   *time.Time
}

func (p *Patient) GetAge() (int, int, int) {
    return gotime.Age(p.BirthDate)
}

func (p *Patient) GetAgeInMonths() float64 {
    return gotime.MonthsBetween(p.BirthDate, time.Now())
}

func (p *Patient) GetTimeSinceLastVisit() string {
    if p.LastVisit == nil {
        return "No previous visits"
    }

    duration := time.Since(*p.LastVisit)
    return gotime.DurationInWords(duration)
}

func (p *Patient) GetRecommendedScreenings() []string {
    age, _, _ := p.GetAge()
    var screenings []string

    if age >= 40 {
        screenings = append(screenings, "Annual Physical")
        screenings = append(screenings, "Blood Pressure Check")
    }

    if age >= 45 {
        screenings = append(screenings, "Cholesterol Test")
        screenings = append(screenings, "Diabetes Screening")
    }

    if age >= 50 {
        screenings = append(screenings, "Colonoscopy")
        screenings = append(screenings, "Mammogram") // for applicable patients
    }

    if age >= 65 {
        screenings = append(screenings, "Bone Density Test")
        screenings = append(screenings, "Vision Test")
        screenings = append(screenings, "Hearing Test")
    }

    return screenings
}

func (p *Patient) IsOverdue() bool {
    if p.LastVisit == nil {
        return true // Never visited
    }

    age, _, _ := p.GetAge()
    var maxInterval time.Duration

    switch {
    case age < 18:
        maxInterval = 6 * 30 * 24 * time.Hour // 6 months for children
    case age < 40:
        maxInterval = 2 * 365 * 24 * time.Hour // 2 years for young adults
    case age < 65:
        maxInterval = 1 * 365 * 24 * time.Hour // 1 year for middle-aged
    default:
        maxInterval = 6 * 30 * 24 * time.Hour // 6 months for seniors
    }

    return time.Since(*p.LastVisit) > maxInterval
}

func (p *Patient) GetVaccineSchedule() map[string]interface{} {
    ageInMonths := p.GetAgeInMonths()
    age, _, _ := p.GetAge()

    schedule := make(map[string]interface{})

    // Pediatric vaccines (age in months)
    if ageInMonths <= 18 {
        if ageInMonths >= 2 {
            schedule["DTaP"] = "Due"
            schedule["Hib"] = "Due"
            schedule["IPV"] = "Due"
            schedule["PCV13"] = "Due"
            schedule["RV"] = "Due"
        }
        if ageInMonths >= 12 {
            schedule["MMR"] = "Due"
            schedule["Varicella"] = "Due"
            schedule["Hepatitis A"] = "Due"
        }
    }

    // Adult vaccines (age in years)
    if age >= 18 {
        schedule["Tdap"] = "Every 10 years"
        schedule["Influenza"] = "Annual"
    }

    if age >= 50 {
        schedule["Shingles"] = "One-time dose"
    }

    if age >= 65 {
        schedule["Pneumococcal"] = "One-time dose"
    }

    return schedule
}

func (p *Patient) GetMedicalSummary() map[string]interface{} {
    years, months, days := p.GetAge()

    summary := map[string]interface{}{
        "patient_id":         p.ID,
        "name":              p.Name,
        "age":               map[string]int{"years": years, "months": months, "days": days},
        "age_in_months":     p.GetAgeInMonths(),
        "is_overdue":        p.IsOverdue(),
        "time_since_visit":  p.GetTimeSinceLastVisit(),
        "recommended_screenings": p.GetRecommendedScreenings(),
        "vaccine_schedule":  p.GetVaccineSchedule(),
        "summary_date":      gotime.Format(time.Now(), "yyyy-mm-dd"),
    }

    return summary
}
```

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
    "github.com/maniartech/gotime/v2"
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
    "github.com/maniartech/gotime/v2"
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
    "github.com/maniartech/gotime/v2"
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
    "github.com/maniartech/gotime/v2"
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
