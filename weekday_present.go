package gotime

import "time"

// IsWeekdayPresentInRange returns true if any of the specified weekdays are present in the date range (inclusive).
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
