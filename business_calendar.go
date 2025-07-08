package gotime

import "time"

// IsBusinessDay reports whether t is a business day (not a weekend or holiday).
//
// The weekends parameter specifies which weekdays are considered weekends
// (e.g., []time.Weekday{time.Saturday, time.Sunday}).
// The holidays parameter is a slice of time.Time representing holidays
// (only the date portion is considered, time is ignored).
//
// Example:
//	weekends := []time.Weekday{time.Saturday, time.Sunday}
//	holidays := []time.Time{time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC)}
//	isBusiness := gotime.IsBusinessDay(time.Now(), weekends, holidays...)
//	// isBusiness: true if today is not a weekend or holiday
func IsBusinessDay(t time.Time, weekends []time.Weekday, holidays ...time.Time) bool {
	for _, w := range weekends {
		if t.Weekday() == w {
			return false
		}
	}
	for _, h := range holidays {
		if sameDay(t, h) {
			return false
		}
	}
	return true
}

// NextBusinessDay returns the next business day after t, skipping weekends and holidays.
//
// Example:
//	weekends := []time.Weekday{time.Saturday, time.Sunday}
//	holidays := []time.Time{time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC)}
//	nextBiz := gotime.NextBusinessDay(time.Now(), weekends, holidays...)
//	// nextBiz: the next date that is not a weekend or holiday
func NextBusinessDay(t time.Time, weekends []time.Weekday, holidays ...time.Time) time.Time {
	next := t.AddDate(0, 0, 1)
	for !IsBusinessDay(next, weekends, holidays...) {
		next = next.AddDate(0, 0, 1)
	}
	return next
}

// PrevBusinessDay returns the previous business day before t, skipping weekends and holidays.
//
// Example:
//	weekends := []time.Weekday{time.Saturday, time.Sunday}
//	holidays := []time.Time{time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC)}
//	prevBiz := gotime.PrevBusinessDay(time.Now(), weekends, holidays...)
//	// prevBiz: the previous date that is not a weekend or holiday
func PrevBusinessDay(t time.Time, weekends []time.Weekday, holidays ...time.Time) time.Time {
	prev := t.AddDate(0, 0, -1)
	for !IsBusinessDay(prev, weekends, holidays...) {
		prev = prev.AddDate(0, 0, -1)
	}
	return prev
}

// sameDay checks if two times are on the same year, month, and day.
func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
