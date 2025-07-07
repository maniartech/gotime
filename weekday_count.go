package gotime

import "time"

type WeekdayCounts struct {
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
	Saturday  int
	Sunday    int
}

// CountWeekdaysInRange returns a struct with the count of each weekday in the given date range (inclusive).
func CountWeekdaysInRange(start, end time.Time) *WeekdayCounts {
	if end.Before(start) {
		start, end = end, start
	}
	counts := &WeekdayCounts{}
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
