package temporal

import (
	"time"
)

func Yesterday() DateTime {
	t := time.Now()
	return DateTime(t.AddDate(0, 0, -1))
}
func Tomorrow() DateTime {
	t := time.Now()
	return DateTime(t.AddDate(0, 0, 1))
}
func Now(dt ...time.Time) DateTime {
	if len(dt) > 0 {
		t := time.Now()
		return DateTime(time.Date(dt[0].Year(), dt[0].Month(), dt[0].Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), dt[0].Location()))
	}
	return DateTime(time.Now())
}

func DateFromTS(ts int64) DateTime {
	t := time.Unix(ts, 0)
	return DateTime(t)
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
