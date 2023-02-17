package temporal

import (
	"time"
)

// Create a struct called Date
type date struct {
	Start time.Time
	End   time.Time
}

func Date(dt ...time.Time) date {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return date{Start: DayStart(t), End: DayEnd(t)}
}

// Date().Monday(weeks) returns the date of the current week's Monday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Monday
func (d date) Monday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+1)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Tuesday(weeks) returns the date of the current week's Tuesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Tuesday
func (d date) Tuesday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+2)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Wednesday(weeks) returns the date of the current week's Wednesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Wednesday
func (d date) Wednesday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+3)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Thursday(weeks) returns the date of the current week's Thursday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Thursday
func (d date) Thursday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+4)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Friday(weeks) returns the date of the current week's Friday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Friday
func (d date) Friday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+5)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Saturday(weeks) returns the date of the current week's Saturday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Saturday
func (d date) Saturday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7+6)
	d.End = DayEnd(d.Start)

	return d
}

// Date().Sunday(weeks) returns the date of the current week's Sunday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Sunday
func (d date) Sunday(weeks ...int) date {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d.Start = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+w*7)
	d.End = DayEnd(d.Start)

	return d
}

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
