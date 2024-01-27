package temporal

import "time"

//-----------------Year Functions-----------------
// YearStart returns the first day of the year.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the first day of the year.
//
// # Note
//
// If the date is not provided, it will return the first day of the year from the current date.
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

// YearEnd returns the last day and the last second of the year.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the last day of the year.
//
// # Note
//
// If the date is not provided, it will return the last day of the year from the current date.
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

// Years returns the date of the given number of years from the date provided,
// If the date is not provided, it will return the date of the given number of years from the current date.
//
// # Arguments
//
// years: (int) The number of years to be added to the date.
//
// dt: (time.Time) The date to be used to calculate the date of the given number of year. (Only takes the first date if multiple dates are provided)
//
// # Note
//
// If the years parameter is 0 it will panic.
func Years(years int, dt ...time.Time) time.Time {
	if years == 0 {
		panic("Years parameter can't be zero")
	}

	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()

	}
	return t.AddDate(years, 0, 0)
}

// LastYear returns the last year's time.Time corresponding to the current time.
func LastYear() time.Time {
	return time.Now().AddDate(-1, 0, 0)
}

// NextYear returns the next year's time.Time corresponding to the current time.
func NextYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

//-----------------Month Functions-----------------

// MonthStart returns the first day of the month.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the first day of the month.
//
// # Note
//
// If the date is not provided, it will return the first day of the month from the current date.
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

// MonthEnd returns the last day and the last second of the month.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the last day of the month.
//
// # Note
//
// If the date is not provided, it will return the last day of the month from the current date.
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

// LastMonth returns the last month's time.Time corresponding to the current time.
func LastMonth() time.Time {
	return time.Now().AddDate(0, -1, 0)
}

// NextMonth returns the next month's time.Time corresponding to the current time.
func NextMonth() time.Time {
	return time.Now().AddDate(0, 1, 0)
}

// Months returns the date of the given number of months from the date provided,
//
// # Arguments
//
// months: (int) The number of months to be added to the date.
//
// dt: (time.Time) The date to be used to calculate the date of the given number of months. (Only takes the first date if multiple dates are provided)
//
// # Note
//
// If the months parameter is 0 it will panic.
func Months(months int, dt ...time.Time) time.Time {
	if months == 0 {
		panic("Months parameter can't be zero")
	}

	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return t.AddDate(0, months, 0)
}

//-----------------Week Functions-----------------

// WeekStart returns the first day of the week (Monday).
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the first day of the week.
//
// # Note
//
// If the date is not provided, it will return the first day of the week from the current date.
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

// WeekStartOn returns the first day of the week on the given day.
//
// # Arguments
//
// day: (time.Weekday) The day to be used to calculate the first day of the week.
//
// dt: (time.Time) The date to be used to calculate the first day of the week.
//
// # Note
//
// If the date is not provided, it will return the first day of the week from the current date.
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
	if weeks == 0 {
		panic("Weeks parameter can't be zero")
	}

	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}

	return t.AddDate(0, 0, weeks*7)
}

//-----------------Day Functions-----------------

// DayStart returns the start of the day.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the start of the day.
//
// # Note
//
// If the date is not provided, it will return the start of the day from the current date.
func DayStart(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return start
}

// DayEnd returns the last second of the day.
//
// # Arguments
//
// dt: (time.Time) The date to be used to calculate the last second of the day.
//
// # Note
//
// If the date is not provided, it will return the last second of the day from the current date.
func DayEnd(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return end
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
	if days == 0 {
		panic("Days parameter can't be zero")
	}

	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return t.AddDate(0, 0, days)
}
