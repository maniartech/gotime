package temporal

import "time"

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
