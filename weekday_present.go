package gotime

import "time"

// IsWeekdayPresentInRange reports whether any of the specified weekdays
// occur within the date range (inclusive). If end is before start, the dates are swapped.
//
// Example:
//	start := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC) // Tuesday
//	end := time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC)   // Saturday
//	present := gotime.IsWeekdayPresentInRange(start, end, time.Monday, time.Friday)
//	// present: true (Friday occurs on July 4th)
func IsWeekdayPresentInRange(start, end time.Time, weekdays ...time.Weekday) bool {
	if len(weekdays) == 0 {
		return false
	}
	// Normalize order
	if end.Before(start) {
		start, end = end, start
	}
	// Build a set for quick lookup
	weekdaySet := make(map[time.Weekday]struct{}, len(weekdays))
	for _, wd := range weekdays {
		weekdaySet[wd] = struct{}{}
	}
	cur := start
	for !cur.After(end) {
		if _, ok := weekdaySet[cur.Weekday()]; ok {
			return true
		}
		cur = cur.AddDate(0, 0, 1)
	}
	return false
}
