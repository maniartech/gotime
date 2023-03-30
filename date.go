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
	return DayStart(dt...)
}
