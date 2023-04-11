package temp

import (
	"time"
)

type dateRange struct {
	// For is the date to use as the base date for the date range. For example,
	// if you want to get the date range for yesterday, you can use the current
	// date as the base date.
	For time.Time

	// From is the start of the date range
	From time.Time

	// To is the end of the date range
	To time.Time
}

// DateRange returns a struct that contains functions to create date ranges.
// The functions are chainable and con
func DateRange(dt ...time.Time) *dateRange {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return &dateRange{For: t}
}

func (d *dateRange) Today() *dateRange {
	d.From = DayStart(d.For)
	d.To = DayEnd(d.For)
	return d
}

func (d *dateRange) Yesterday() *dateRange {
	d.From = time.Time(Yesterday())
	d.To = DayEnd(d.For)
	return d
}

func (d *dateRange) Tomorrow() *dateRange {
	d.From = DayStart(d.For)
	d.To = time.Time(Tomorrow())
	return d
}

// Days returns the date range based of specified number of days. If the number of days is negative,
// it returns the from date as exactly the specified number of days ago from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of days is -2, the from date will be 2023-02-05. If the number of days is
// positive, it returns the from date as exactly the specified number of days from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of days is 2, the from date will be 2023-02-07. The base date is always included
// in the date range.
func (d *dateRange) Days(days int) *dateRange {
	if days < 0 {

		d.From = DayStart(d.For).AddDate(0, 0, days+1) // +1 because the base date is included
		d.To = DayEnd(d.For)
	} else {
		d.From = DayStart(d.For)
		d.To = DayEnd(d.For).AddDate(0, 0, days-1) // -1 because the base date is included
	}
	return d
}

// ThisWeek returns the current week's date range. It returns the from date
// as the start of the current week until the end of the current week.
func (d *dateRange) ThisWeek() *dateRange {
	d.From = DayStart(d.For)
	// Iterating through the days of the week to get the sunday of the week
	for d.From.Weekday() != time.Sunday {
		d.From = d.From.AddDate(0, 0, -1)
	}
	// To is the end of the week
	d.To = d.From.AddDate(0, 0, 6)
	return d
}

func (d *dateRange) LastWeek() *dateRange {
	d.From = DayStart().AddDate(0, 0, -7)
	// Iterating through the days of the week to get the sunday of the week
	for d.From.Weekday() != time.Sunday {
		d.From = d.From.AddDate(0, 0, -1)
	}
	d.To = d.From.AddDate(0, 0, 6)
	return d
}

// NextWeek returns dateRange from the start of the next week until the end of the next week.
func (d *dateRange) NextWeek() *dateRange {
	d.From = NextWeek(d.For)

	// Iterating through the days of the week to get the sunday of the week
	for d.From.Weekday() != time.Sunday {
		d.From = d.From.AddDate(0, 0, -1)
	}
	// To is the end of the week
	d.To = d.From.AddDate(0, 0, 6)
	return d
}

// Weeks returns the date range based of specified number of weeks. If the number of weeks is negative,
// it returns the from date as exactly the specified number of weeks ago from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of weeks is -2, the from date will be 2023-01-23. If the number of weeks is
// positive, it returns the from date as exactly the specified number of weeks from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of weeks is 2, the from date will be 2023-02-19. The base date is always included
// in the date range. If the number of weeks is 0, it returns the current week's date range.
func (d *dateRange) Weeks(weeks int) *dateRange {
	if weeks == 0 {
		return d.ThisWeek()
	}

	// From is the start of the week
	d.From = DayStart(d.For).AddDate(0, 0, weeks*7)

	// To is the end of the week
	d.To = DayEnd(d.For).AddDate(0, 0, weeks*7+6)
	return d
}

// This Month returns the current month's date range. It returns the from date
// as the start of the current month until the end of the current month.
func (d *dateRange) ThisMonth() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the month
	d.From = time.Date(dayStart.Year(), dayStart.Month(), 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the month is the first day of the next month minus 1 day
	d.To = time.Date(dayStart.Year(), dayStart.Month()+1, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// LastMonth returns the previous month's date range. It returns the from date
// as the start of the previous month until the end of the previous month.
func (d *dateRange) LastMonth() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the month
	d.From = time.Date(dayStart.Year(), dayStart.Month()-1, 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the month is the first day of the next month minus 1 day
	d.To = time.Date(dayStart.Year(), dayStart.Month(), 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// NextMonth returns the next month's date range. It returns the from date
// as the start of the next month until the end of the next month.
func (d *dateRange) NextMonth() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the month
	d.From = time.Date(dayStart.Year(), dayStart.Month()+1, 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the month is the first day of the next month minus 1 day
	d.To = time.Date(dayStart.Year(), dayStart.Month()+2, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// Months returns the date range based of specified number of months. If the number of months is negative,
// it returns the from date as exactly the specified number of months ago from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of months is -2, the from date will be 2022-12-29. If the number of months is
// positive, it returns the from date as exactly the specified number of months from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of months is 2, the from date will be 2023-04-05. The base date is always included
// in the date range.
func (d *dateRange) Months(months int) *dateRange {
	if months == 0 {
		return d.ThisMonth()
	}

	dayStart := DayStart(d.For)
	d.From = time.Date(dayStart.Year(), dayStart.Month(), 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, months, 0)
	d.To = time.Date(dayStart.Year(), dayStart.Month()+1, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, months, -1)

	return d
}

// This Year returns the current year's date range. It returns the from date
// as the start of the current year until the end of the current year.
func (d *dateRange) ThisYear() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the year
	d.From = time.Date(dayStart.Year(), 1, 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the year is the first day of the next year minus 1 day
	d.To = time.Date(dayStart.Year()+1, 1, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// LastYear returns the previous year's date range. It returns the from date
// as the start of the previous year until the end of the previous year.
func (d *dateRange) LastYear() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the year
	d.From = time.Date(dayStart.Year()-1, 1, 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the year is the first day of the next year minus 1 day
	d.To = time.Date(dayStart.Year(), 1, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// NextYear returns the next year's date range. It returns the from date
// as the start of the next year until the end of the next year.
func (d *dateRange) NextYear() *dateRange {
	dayStart := DayStart(d.For)
	// first day of the year
	d.From = time.Date(dayStart.Year()+1, 1, 1, 0, 0, 0, 0, dayStart.Location())

	// last day of the year is the first day of the next year minus 1 day
	d.To = time.Date(dayStart.Year()+2, 1, 1, 0, 0, 0, 0, dayStart.Location()).AddDate(0, 0, -1)
	return d
}

// Years returns the date range based of specified number of years. If the number of years is negative,
// it returns the from date as exactly the specified number of years ago from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of years is -2, the from date will be 2021-02-06. If the number of years is
// positive, it returns the from date as exactly the specified number of years from the base date until the base date. For example,
// if the base date is 2023-02-06 and the number of years is 2, the from date will be 2025-02-06. The base date is always included
// in the date range.
func (d *dateRange) Years(years int) *dateRange {
	if years == 0 {
		return d.ThisYear()
	}

	dayStart := DayStart(d.For)
	d.From = time.Date(dayStart.Year(), dayStart.Month(), dayStart.Day(), 0, 0, 0, 0, dayStart.Location()).AddDate(years, 0, 0)
	d.To = time.Date(dayStart.Year()+1, dayStart.Month(), dayStart.Day(), 0, 0, 0, 0, dayStart.Location()).AddDate(years, 0, -1)

	return d
}

// Range returns the date range from the specified time.Time. If time.Time is
// from the past, the range will be from the past date to the current date.
// If time.Time is from the future, the range will be from the current date
// to the future date.
func (d *dateRange) Range(dt time.Time) *dateRange {
	if dt.Before(d.For) {
		d.From = DayStart(dt)
		d.To = DayEnd(d.For)
	} else {
		d.From = DayStart(d.For)
		d.To = DayEnd(dt)
	}
	return d
}
