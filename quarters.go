package gotime

import "time"

// QuarterStart returns the first day of the quarter for the given time.
// If no time is provided, it uses the current time.
//
// Example:
//   q1Start := gotime.QuarterStart(time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC))
//   // Returns: 2025-01-01 00:00:00 +0000 UTC
func QuarterStart(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}

	quarter := (int(t.Month()) - 1) / 3
	startMonth := quarter*3 + 1
	return time.Date(t.Year(), time.Month(startMonth), 1, 0, 0, 0, 0, t.Location())
}

// QuarterEnd returns the last day and last second of the quarter for the given time.
// If no time is provided, it uses the current time.
//
// Example:
//   q1End := gotime.QuarterEnd(time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC))
//   // Returns: 2025-03-31 23:59:59.999999999 +0000 UTC
func QuarterEnd(dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}

	quarter := (int(t.Month()) - 1) / 3
	endMonth := quarter*3 + 3

	// Get the last day of the quarter's last month
	lastDay := DaysInMonth(t.Year(), endMonth)
	return time.Date(t.Year(), time.Month(endMonth), lastDay, 23, 59, 59, 999999999, t.Location())
}

// LastQuarter returns the same date/time in the previous quarter.
// This is equivalent to Quarters(-1).
//
// Example:
//   lastQ := gotime.LastQuarter() // If current is 2025-07-15, returns 2025-04-15
func LastQuarter() time.Time {
	return Quarters(-1)
}

// NextQuarter returns the same date/time in the next quarter.
// This is equivalent to Quarters(1).
//
// Example:
//   nextQ := gotime.NextQuarter() // If current is 2025-07-15, returns 2025-10-15
func NextQuarter() time.Time {
	return Quarters(1)
}

// Quarters returns the time after adding the specified number of quarters to the given time.
// If no time is provided, it uses the current time. Negative values subtract quarters.
//
// Example:
//   futureQuarter := gotime.Quarters(2, someTime)  // 2 quarters (6 months) after someTime
//   pastQuarter := gotime.Quarters(-1, someTime)   // 1 quarter (3 months) before someTime
//   noChange := gotime.Quarters(0)                 // same as time.Now()
func Quarters(quarters int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}

	// Each quarter is 3 months
	return t.AddDate(0, quarters*3, 0)
}

// QuarterOfYear returns the quarter number (1-4) for the given time.
// If no time is provided, it uses the current time.
//
// Example:
//   quarter := gotime.QuarterOfYear(time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC))
//   // Returns: 3 (July is in Q3)
func QuarterOfYear(dt ...time.Time) int {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}

	return (int(t.Month())-1)/3 + 1
}
