package datetime

import (
	"time"
)

// Return today's date at 00:00:00
func TodayStart() *time.Time {
	var now time.Time = time.Now().UTC()
	var todayMidnight time.Time = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return &todayMidnight
}

// Return today's date at 00:00:00
func Today() *time.Time {
	start := *TodayStart()
	return &start
}

// Returns today's date at 11:59 PM
func EoD() *time.Time {
	start := *TodayStart()
	eod := start.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	return &eod
}

// Return yesterday's date at 00:00:00
func Yesterday() *time.Time {
	yesterday := TodayStart().AddDate(0, 0, -1)
	return &yesterday
}

// Return tomorrow's date at 00:00:00
func Tomorrow() *time.Time {
	tomorrow := TodayStart().AddDate(0, 0, 1)
	return &tomorrow
}

// Return last week's date at 00:00:00
func LastWeek() *time.Time {
	lastWeek := TodayStart().AddDate(0, 0, -7)
	return &lastWeek
}

// Return last month's date at 00:00:00
func LastMonth() *time.Time {
	lastMonth := TodayStart().AddDate(0, -1, 0)
	return &lastMonth
}

// Return last year's date at 00:00:00
func LastYear() *time.Time {
	lastYear := TodayStart().AddDate(-1, 0, 0)
	return &lastYear
}

// Return next week's date at 00:00:00
func NextWeek() *time.Time {
	nextWeek := TodayStart().AddDate(0, 0, 7)
	return &nextWeek
}

// Return next month's date at 00:00:00
func NextMonth() *time.Time {
	nextMonth := TodayStart().AddDate(0, 1, 0)
	return &nextMonth
}

// Return next year's date at 00:00:00
func NextYear() *time.Time {
	nextYear := TodayStart().AddDate(1, 0, 0)
	return &nextYear
}
