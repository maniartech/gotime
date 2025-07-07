package gotime

import "time"

// DayOfYear returns the day number (1-366) within the year for the given date.
// This uses Go's built-in YearDay() method for optimal performance.
//
// Example:
//   gotime.DayOfYear(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))  // Returns 1
//   gotime.DayOfYear(time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)) // Returns 365
//   gotime.DayOfYear(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)) // Returns 366 (leap year)
func DayOfYear(t time.Time) int {
	return t.YearDay()
}

// WeekOfMonth returns the week number (1-5) within the month for the given date.
// The first week starts on the 1st of the month, regardless of the day of the week.
// This follows the common business calendar convention.
//
// Example:
//   gotime.WeekOfMonth(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))  // Returns 1
//   gotime.WeekOfMonth(time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC))  // Returns 2
//   gotime.WeekOfMonth(time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC)) // Returns 5
func WeekOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	firstWeekday := int(firstOfMonth.Weekday())
	return (t.Day()+firstWeekday-1)/7 + 1
}

// IsFirstDayOfMonth returns true if the date is the first day of its month.
// This is a high-performance check that simply compares the day component.
//
// Example:
//   gotime.IsFirstDayOfMonth(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))  // Returns true
//   gotime.IsFirstDayOfMonth(time.Date(2025, 7, 2, 0, 0, 0, 0, time.UTC))  // Returns false
func IsFirstDayOfMonth(t time.Time) bool {
	return t.Day() == 1
}

// IsLastDayOfMonth returns true if the date is the last day of its month.
// This efficiently checks by adding one day and seeing if the month changes.
// This approach correctly handles leap years and varying month lengths.
//
// Example:
//   gotime.IsLastDayOfMonth(time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC))  // Returns true (non-leap year)
//   gotime.IsLastDayOfMonth(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC))  // Returns true (leap year)
//   gotime.IsLastDayOfMonth(time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC))  // Returns true
//   gotime.IsLastDayOfMonth(time.Date(2025, 7, 30, 0, 0, 0, 0, time.UTC))  // Returns false
func IsLastDayOfMonth(t time.Time) bool {
	nextDay := t.AddDate(0, 0, 1)
	return nextDay.Month() != t.Month()
}
