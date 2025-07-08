package gotime

import "time"

// IsBetween reports whether t1 falls within the time range defined by t2 and t3 (inclusive).
// The order of t2 and t3 doesn't matter - they will be automatically ordered.
//
// Example:
//	target := time.Date(2025, 7, 5, 12, 0, 0, 0, time.UTC)
//	start := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 7, 10, 0, 0, 0, 0, time.UTC)
//	result := gotime.IsBetween(target, start, end)
//	// result: true (July 5th is between July 1st and 10th)
func IsBetween(t1, t2, t3 time.Time) bool {
	t1Unix := t1.UnixMilli()
	t2Unix := t2.UnixMilli()
	t3Unix := t3.UnixMilli()

	// Swapping the values if t2Unix is greater than t3Unix
	if t2Unix > t3Unix {
		t2Unix, t3Unix = t3Unix, t2Unix
	}

	return t1Unix >= t2Unix && t1Unix <= t3Unix
}

// IsBetweenDates reports whether t1 falls within the date range defined by t2 and t3 (inclusive).
// Unlike IsBetween, this function compares only the date portion by normalizing t2 to
// start-of-day and t3 to end-of-day before comparison.
//
// Example:
//	target := time.Date(2025, 7, 5, 23, 30, 0, 0, time.UTC)
//	start := time.Date(2025, 7, 5, 10, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 7, 5, 14, 0, 0, 0, time.UTC)
//	result := gotime.IsBetweenDates(target, start, end)
//	// result: true (all dates are July 5th, regardless of time)
func IsBetweenDates(t1, t2, t3 time.Time) bool {
	sec1 := SoD(t1).Unix()
	sec2 := SoD(t2).Unix()
	sec3 := EoD(t3).Unix()

	// Swapping the values if sec2 is greater than sec3
	if sec2 > sec3 {
		sec2, sec3 = sec3, sec2
	}

	return sec1 >= sec2 && sec1 <= sec3
}
