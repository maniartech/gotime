package gotime

import "time"

//-----------------Year Functions-----------------
// YearStart returns the first day of the year for the given date.
// If no date is provided, it uses the current time.
//
// Example:
//	start := YearStart() // First day of current year
//	// start: 2025-01-01 00:00:00
//
//	someDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
//	start = YearStart(someDate)
//	// start: 2024-01-01 00:00:00
func YearStart(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()

	}
	start := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	return start
}

// YearEnd returns the last day and last second of the year for the given date.
// If no date is provided, it uses the current time.
//
// Example:
//	end := YearEnd() // Last day of current year
//	// end: 2025-12-31 23:59:59.999999999
//
//	someDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
//	end = YearEnd(someDate)
//	// end: 2024-12-31 23:59:59.999999999
func YearEnd(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()

	}
	end := time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
	return end
}

// Years returns the date after adding the specified number of years to the given date.
// If no date is provided, it uses the current time.
// If years is 0, it returns the original date unchanged.
//
// Example:
//	future := Years(2) // 2 years from now
//	// future: 2027-07-08 (current time + 2 years)
//
//	someDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
//	result := Years(5, someDate)
//	// result: 2025-01-01
func Years(years int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	if years == 0 {
		return t
	}
	return t.AddDate(years, 0, 0)
}

// LastYear returns the date one year ago from the current time.
//
// Example:
//	lastYear := LastYear()
//	// lastYear: 2024-07-08 (if current time is 2025-07-08)
func LastYear() time.Time {
	return time.Now().AddDate(-1, 0, 0)
}

// NextYear returns the date one year from the current time.
//
// Example:
//	nextYear := NextYear()
//	// nextYear: 2026-07-08 (if current time is 2025-07-08)
func NextYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

//-----------------Month Functions-----------------

// MonthStart returns the first day of the month for the given date.
// If no date is provided, it uses the current time.
//
// Example:
//	start := MonthStart() // First day of current month
//	// start: 2025-07-01 00:00:00
//
//	someDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
//	start = MonthStart(someDate)
//	// start: 2024-06-01 00:00:00
func MonthStart(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return start
}

// MonthEnd returns the last day and last second of the month for the given date.
// If no date is provided, it uses the current time.
//
// Example:
//	end := MonthEnd() // Last day of current month
//	// end: 2025-07-31 23:59:59.999999999
//
//	someDate := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
//	end = MonthEnd(someDate)
//	// end: 2024-02-29 23:59:59.999999999 (leap year)
func MonthEnd(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
	return end
}

// LastMonth returns the date one month ago from the current time.
//
// Example:
//	lastMonth := LastMonth()
//	// lastMonth: 2025-06-08 (if current time is 2025-07-08)
func LastMonth() time.Time {
	return time.Now().AddDate(0, -1, 0)
}

// NextMonth returns the date one month from the current time.
//
// Example:
//	nextMonth := NextMonth()
//	// nextMonth: 2025-08-08 (if current time is 2025-07-08)
func NextMonth() time.Time {
	return time.Now().AddDate(0, 1, 0)
}

// Months returns the date after adding the specified number of months to the given date.
// If no date is provided, it uses the current time.
// If months is 0, it returns the original date unchanged.
//
// Example:
//	future := Months(3) // 3 months from now
//	// future: 2025-10-08 (current time + 3 months)
//
//	someDate := time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC)
//	result := Months(1, someDate)
//	// result: 2020-02-29 (handles month-end edge cases)
func Months(months int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	if months == 0 {
		return t
	}
	return t.AddDate(0, months, 0)
}

//-----------------Week Functions-----------------

// WeekStart returns the first day of the week (Sunday) for the given date.
// If no date is provided, it uses the current time.
//
// Example:
//	start := WeekStart() // Sunday of current week
//	// start: 2025-07-06 00:00:00 (if current is 2025-07-08)
//
//	someDate := time.Date(2025, 7, 10, 15, 30, 0, 0, time.UTC) // Thursday
//	start = WeekStart(someDate)
//	// start: 2025-07-06 00:00:00 (Sunday of that week)
func WeekStart(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := t.AddDate(0, 0, -int(t.Weekday()))
	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
}

// WeekStartOn returns the first occurrence of the specified weekday for the given date's week.
// If no date is provided, it uses the current time.
//
// Example:
//	start := WeekStartOn(time.Monday) // Monday of current week
//	// start: 2025-07-07 00:00:00 (if current is 2025-07-08)
//
//	someDate := time.Date(2025, 7, 10, 0, 0, 0, 0, time.UTC) // Thursday
//	start = WeekStartOn(time.Wednesday, someDate)
//	// start: 2025-07-09 00:00:00 (Wednesday of that week)
func WeekStartOn(day time.Weekday, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := t.AddDate(0, 0, -int(t.Weekday())+int(day))
	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
}

// WeekEnd returns the last day and the last second of the week.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the last day of the week.
//
// # Note
//
// If the date is not provided, it will return the last day of the week from the current date.
func WeekEnd(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := t.AddDate(0, 0, 6-int(t.Weekday()))
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, end.Location())
}

// WeekEndOn returns the last day and the last second of the week on the given day.
// For example, WeekEndOn(time.Sunday) returns the last day of the week (Sunday).
//
// # Arguments
//
// day: (time.Weekday) The day to be used to calculate the last day of the week.
//
// dt: (time.Time) The date to be used to calculate the last day of the week.
//
// # Note
//
// If the date is not provided, it will return the last day of the week from the current date.
func WeekEndOn(day time.Weekday, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := t.AddDate(0, 0, 6-int(t.Weekday())+int(day))
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, end.Location())
}

// LastWeek returns the last week's time.Time corresponding to the current time.
func LastWeek() time.Time {
	return time.Now().AddDate(0, 0, -7)
}

// NextWeek returns the next week's time.Time corresponding to the current time.
func NextWeek() time.Time {
	return time.Now().AddDate(0, 0, 7)
}

// Weeks returns the date of the given number of weeks from the current date.
// If the value is negative, it will return the date of the previous week.
//
// # Arguments
//
// weeks: (int) The number of weeks to be added to the date.
//
// dt: (time.Time) The date to be used to calculate the date of the given number of weeks. (Only takes the first date if multiple dates are provided)
//
// # Note
//
// If the weeks parameter is 0 it will panic.
func Weeks(weeks int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	if weeks == 0 {
		return t
	}
	return t.AddDate(0, 0, weeks*7)
}

//-----------------Day Functions-----------------

// EoD returns the end of the day for the given time.
//
// It calculates the start of the day for the given time using the SoD function, and adds 24 hours to it.
// To get the end of the day, it subtracts one nanosecond from the resulting time.Time value,
// since the SoD function returns the first nanosecond of the day.
//
// Example:
//   t := time.Date(2022, time.December, 30, 16, 30, 0, 0, time.UTC)
//   endOfDay := EoD(t)
//   // endOfDay == time.Date(2022, time.December, 30, 23, 59, 59, 999999999, time.UTC)
func EoD(t ...time.Time) time.Time {
	var dt time.Time
	if len(t) > 0 {
		dt = t[0]
	} else {
		dt = time.Now()
	}

	return time.Date(
		dt.Year(), dt.Month(), dt.Day(),
		23, 59, 59, 999999999,
		dt.Location(),
	)
}

// SoD returns the start of the day for the given time.
//
// It constructs a new time.Time value using the year, month, and day of the given time.Time value,
// and setting the hour, minute, second, and nanosecond fields to 0. It also sets the location field to the
// same location as the input time.Time value. The resulting time.Time value represents the start of the day
// for the given time.
//
// Example:
//   t := time.Date(2022, time.December, 30, 16, 30, 0, 0, time.UTC)
//   startOfDay := SoD(t)
//   // startOfDay == time.Date(2022, time.December, 30, 0, 0, 0, 0, time.UTC)
func SoD(t ...time.Time) time.Time {
	if len(t) > 0 {
		return time.Date(t[0].Year(), t[0].Month(), t[0].Day(), 0, 0, 0, 0, t[0].Location())
	}

	dt := time.Now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}

// Yesterday returns the yesterday's time.Time corresponding to the current time.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// Tomorrow returns the tomorrow's time.Time corresponding to the current time.
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// Days returns the date of the given number of days from the date provided,
//
// # Arguments
//
// days: (int) The number of days to be added to the date.
//
// dt: (time.Time) The date to be used to calculate the date of the given number of days. (Only takes the first date if multiple dates are provided)
//
// # Note
//
// If the days parameter is 0 it will panic.
func Days(days int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	if days == 0 {
		return t
	}
	return t.AddDate(0, 0, days)
}
