package temporal

import "time"

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

//-----------------Week Functions-----------------

// WeekStart returns the first day of the week.
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

// LastWeek returns the last week's time.Time corresponding to the current time.
func LastWeek() time.Time {
	return time.Now().AddDate(0, 0, -7)
}

// NextWeek returns the next week's time.Time corresponding to the current time.
func NextWeek() time.Time {
	return time.Now().AddDate(0, 0, 7)
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


