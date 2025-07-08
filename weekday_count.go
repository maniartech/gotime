package gotime

import "time"

// WeekdayCounts represents the count of each weekday in a date range.
type WeekdayCounts struct {
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
	Saturday  int
	Sunday    int
}

// CountWeekdaysInRange returns the count of each weekday within the specified
// date range (inclusive). If end is before start, the dates are swapped.
//
// Example:
//	start := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
//	counts := CountWeekdaysInRange(start, end)
//	// counts.Monday: 1, counts.Tuesday: 1, etc.
func CountWeekdaysInRange(start, end time.Time) WeekdayCounts {
	if end.Before(start) {
		start, end = end, start
	}
	counts := WeekdayCounts{}
	cur := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	for !cur.After(end) {
		switch cur.Weekday() {
		case time.Monday:
			counts.Monday++
		case time.Tuesday:
			counts.Tuesday++
		case time.Wednesday:
			counts.Wednesday++
		case time.Thursday:
			counts.Thursday++
		case time.Friday:
			counts.Friday++
		case time.Saturday:
			counts.Saturday++
		case time.Sunday:
			counts.Sunday++
		}
		cur = cur.AddDate(0, 0, 1)
	}
	return counts
}
