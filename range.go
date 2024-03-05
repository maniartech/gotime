package gotime

import "time"

// IsInRange checks if the given time is in the range of the start and end time.
// It performs inclusive comparison, i.e., the start and end time are included in
// the range.
func IsInRange(time, start, end time.Time) bool {
	return time.After(start) && time.Before(end) ||
		time.Equal(start) ||
		time.Equal(end)
}

// IsInDateRange checks if the given time is in the range of the start and end
// date. Before performing inclusive comparison, it sets the time to the start
// of the day for the start date and the end of the day for the end date.
func IsInDateRange(time, start, end time.Time) bool {
	start = SoD(start)
	end = EoD(end)

	return time.After(start) && time.Before(end) ||
		time.Equal(start) ||
		time.Equal(end)
}
