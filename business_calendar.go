package gotime

import "time"

// IsBusinessDay returns true if t is a business day (not a weekend or holiday).
// weekends: a slice of time.Weekday considered weekends (e.g., []time.Weekday{time.Saturday, time.Sunday})
// holidays: a slice of time.Time representing holidays (date only, time ignored)
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

// NextBusinessDay returns the next business day after t (skipping weekends and holidays).
func NextBusinessDay(t time.Time, weekends []time.Weekday, holidays ...time.Time) time.Time {
	next := t.AddDate(0, 0, 1)
	for !IsBusinessDay(next, weekends, holidays...) {
		next = next.AddDate(0, 0, 1)
	}
	return next
}

// PrevBusinessDay returns the previous business day before t (skipping weekends and holidays).
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
