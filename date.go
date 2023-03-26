package temporal

import (
	"time"
)

// Return today's DateTime at 00:00:00
func DayStart(dt ...time.Time) DateTime {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return DateTime(start)
}

// Returns today's DateTime at 11:59 PM
func DayEnd(dt ...time.Time) DateTime {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return DateTime(end)
}

// Today is an alias for DayStart
func Today(dt ...time.Time) DateTime {
	return Now(dt...)
}

// Return last week's DateTime at 00:00:00
func LastWeek(dt ...time.Time) DateTime {
	lastWeek := Now(dt...).AddDate(0, 0, -7)
	return lastWeek
}

// Return last month's DateTime at 00:00:00
func LastMonth(dt ...time.Time) DateTime {
	lastMonth := Now(dt...).AddDate(0, -1, 0)
	return lastMonth
}

// Return last year's DateTime at 00:00:00
func LastYear(dt ...time.Time) DateTime {
	lastYear := Now(dt...).AddDate(-1, 0, 0)
	return lastYear
}

// Return next week's DateTime at 00:00:00
func NextWeek(dt ...time.Time) DateTime {
	nextWeek := Now(dt...).AddDate(0, 0, 7)
	return nextWeek
}

// Return next month's DateTime at 00:00:00
func NextMonth(dt ...time.Time) DateTime {
	nextMonth := Now(dt...).AddDate(0, 1, 0)
	return nextMonth
}

// Return next year's DateTime at 00:00:00
func NextYear(dt ...time.Time) DateTime {
	nextYear := Now(dt...).AddDate(1, 0, 0)
	return nextYear
}
