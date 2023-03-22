package temporal

import (
	"time"
)

func Yesterday() DateTime {
	t := time.Now()
	return DateTime(DayStart(t).AddDate(0, 0, -1))
}
func Tomorrow() DateTime {
	t := time.Now()
	return DateTime(DayStart(t).AddDate(0, 0, 1))
}
func Now() DateTime {
	t := time.Now()
	return DateTime(DayStart(t))
}

func DateFromTS(ts int64) DateTime {
	t := time.Unix(ts, 0)
	return DateTime(DayStart(t))
}

func DateFromTime(t time.Time) DateTime {
	return DateTime(t)
}

// Unix returns a DateTime with the given Unix timestamp
func Unix(ts int64) DateTime {
	return DateTime(time.Unix(ts, 0))
}

// UnixMilli returns a DateTime with the given Unix timestamp in milliseconds
func UnixMilli(ts int64) DateTime {
	return DateTime(time.UnixMilli(ts))
}

// UnixMicro returns a DateTime with the given Unix timestamp in microsecond
func UnixNano(ts int64) DateTime {
	return DateTime(time.UnixMicro(ts))
}

func Date(dt ...time.Time) DateTime {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return DateTime(t)
}

// Return today's DateTime at 00:00:00
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

// Returns today's DateTime at 11:59 PM
func DayEnd(dt ...time.Time) time.Time {
	start := DayStart(dt...)
	eod := start.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	return eod
}

// Today is an alias for DayStart
func Today(dt ...time.Time) time.Time {
	return DayStart(dt...)
}

// Return last week's DateTime at 00:00:00
func LastWeek(dt ...time.Time) time.Time {
	lastWeek := DayStart(dt...).AddDate(0, 0, -7)
	return lastWeek
}

// Return last month's DateTime at 00:00:00
func LastMonth(dt ...time.Time) time.Time {
	lastMonth := DayStart(dt...).AddDate(0, -1, 0)
	return lastMonth
}

// Return last year's DateTime at 00:00:00
func LastYear(dt ...time.Time) time.Time {
	lastYear := DayStart(dt...).AddDate(-1, 0, 0)
	return lastYear
}

// Return next week's DateTime at 00:00:00
func NextWeek(dt ...time.Time) time.Time {
	nextWeek := DayStart(dt...).AddDate(0, 0, 7)
	return nextWeek
}

// Return next month's DateTime at 00:00:00
func NextMonth(dt ...time.Time) time.Time {
	nextMonth := DayStart(dt...).AddDate(0, 1, 0)
	return nextMonth
}

// Return next year's DateTime at 00:00:00
func NextYear(dt ...time.Time) time.Time {
	nextYear := DayStart(dt...).AddDate(1, 0, 0)
	return nextYear
}
