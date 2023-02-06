package temporal

import (
	"time"
)

// Return today's date at 00:00:00
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

// Returns today's date at 11:59 PM
func DayEnd(dt ...time.Time) time.Time {
	start := DayStart(dt...)
	eod := start.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	return eod
}

// Today is an alias for DayStart
func Today(dt ...time.Time) time.Time {
	return DayStart(dt...)
}

// Return yesterday's date at 00:00:00
func Yesterday(dt ...time.Time) time.Time {
	yesterday := DayStart(dt...).AddDate(0, 0, -1)
	return yesterday
}

// Return tomorrow's date at 00:00:00
func Tomorrow(dt ...time.Time) time.Time {
	tomorrow := DayStart(dt...).AddDate(0, 0, 1)
	return tomorrow
}

// Return last week's date at 00:00:00
func LastWeek(dt ...time.Time) time.Time {
	lastWeek := DayStart(dt...).AddDate(0, 0, -7)
	return lastWeek
}

// Return last month's date at 00:00:00
func LastMonth(dt ...time.Time) time.Time {
	lastMonth := DayStart(dt...).AddDate(0, -1, 0)
	return lastMonth
}

// Return last year's date at 00:00:00
func LastYear(dt ...time.Time) time.Time {
	lastYear := DayStart(dt...).AddDate(-1, 0, 0)
	return lastYear
}

// Return next week's date at 00:00:00
func NextWeek(dt ...time.Time) time.Time {
	nextWeek := DayStart(dt...).AddDate(0, 0, 7)
	return nextWeek
}

// Return next month's date at 00:00:00
func NextMonth(dt ...time.Time) time.Time {
	nextMonth := DayStart(dt...).AddDate(0, 1, 0)
	return nextMonth
}

// Return next year's date at 00:00:00
func NextYear(dt ...time.Time) time.Time {
	nextYear := DayStart(dt...).AddDate(1, 0, 0)
	return nextYear
}
