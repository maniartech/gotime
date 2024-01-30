package temporal

import "time"

// IsLeapYear returns true if the year is a leap year.
//
// A leap year is a year that is evenly divisible by 4, but not by 100 unless it is also divisible by 400.
//
// Example:
//   isLeap := IsLeapYear(2024)
//   // isLeap == true
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in the month.
//
// For February, the number of days depends on whether the year is a leap year or not.
// For April, June, September, and November, the number of days is 30.
// For all other months, the number of days is 31.
//
// Example:
//   days := DaysInMonth(2022, 2)
//   // days == 28
func DaysInMonth(year, month int) int {
	switch month {
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

// DaysInYear returns the number of days in the year.
//
// If the year is a leap year, it returns 366, otherwise it returns 365.
//
// Example:
//   days := DaysInYear(2022)
//   // days == 365
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// DaysInQuarter returns the number of days in the quarter.
//
// The number of days in each quarter depends on the number of days in each month.
// Q1: Jan (31) + Feb (28 or 29) + Mar (31) = 90 or 91 days
// Q2: Apr (30) + May (31) + Jun (30) = 91 days
// Q3: Jul (31) + Aug (31) + Sep (30) = 92 days
// Q4: Oct (31) + Nov (30) + Dec (31) = 92 days
//
// Example:
//   days := DaysInQuarter(2022, 1)
//   // days == 90
func DaysInQuarter(year, quarter int) int {
	switch quarter {
	case 1:
		return DaysInMonth(year, 1) + DaysInMonth(year, 2) + DaysInMonth(year, 3)
	case 2:
		return DaysInMonth(year, 4) + DaysInMonth(year, 5) + DaysInMonth(year, 6)
	case 3:
		return DaysInMonth(year, 7) + DaysInMonth(year, 8) + DaysInMonth(year, 9)
	default:
		return DaysInMonth(year, 10) + DaysInMonth(year, 11) + DaysInMonth(year, 12)
	}
}

// NewDate creates a time.Time object from the given year, month and day.
//
// It constructs a new time.Time object with the given year, month, and day, and sets the hour, minute,
// second, and nanosecond fields to 0. It also sets the location field to time.UTC.
//
// Example:
//   date := NewDate(2022, 4, 15)
//   // date == time.Date(2022, time.April, 15, 0, 0, 0, 0, time.UTC)
func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// NewTime creates a time.Time object from the given hour, minute and second.
//
// It constructs a new time.Time object with the given hour, minute, and second, and sets the year, month,
// and day fields to 0. It also sets the location field to time.UTC.
//
// Example:
//   timeObj := NewTime(12, 30, 0)
//   // timeObj == time.Date(0, 1, 1, 12, 30, 0, 0, time.UTC)
func NewTime(hour, minute, second int) time.Time {
	return time.Date(0, 0, 0, hour, minute, second, 0, time.UTC)
}
