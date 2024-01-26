package temporal

import "time"

//-----------------Year Functions-----------------
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

func LastYear() time.Time {
	return time.Now().AddDate(-1, 0, 0)
}

func NextYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

//-----------------Month Functions-----------------

// MonthStart returns the first day of the month.
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
// If the date is not provided, it will return the date of the given number of months from the current date.
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
// For example, WeekStartOn(time.Sunday) returns the first day of the week (Sunday).
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
// The default value is 1 week, that is the date of the next week from the specified date.
// The weeks parameter can't be zero. If it is zero, it will panic.
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

// DayStart returns the first second of the day.
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
// If the date is not provided, it will return the date of the given number of days from the current date.
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
