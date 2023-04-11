package temporal

import "time"

// IsLeapYear returns true if the year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in the month.
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
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// DaysInQuarter returns the number of days in the quarter.
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

// DateCreate creates a time.Time object from the given year, month and day.
func DateCreate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// TimeCreate creates a time.Time object from the given hour, minute and second.
func TimeCreate(hour, minute, second int) time.Time {
	return time.Date(0, 0, 0, hour, minute, second, 0, time.UTC)
}
