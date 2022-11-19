package dateutils

import (
	"time"
)

var now time.Time = time.Now().UTC()
var todayMidnight time.Time = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

func Today() *time.Time {
	return &todayMidnight
}

func EoD() *time.Time {
	eod := todayMidnight.Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59)
	return &eod
}

func Yesterday() *time.Time {
	yesterday := todayMidnight.AddDate(0, 0, -1)
	return &yesterday
}

func Tomorrow() *time.Time {
	tomorrow := todayMidnight.AddDate(0, 0, 1)
	return &tomorrow
}

func LastWeek() *time.Time {
	lastWeek := todayMidnight.AddDate(0, 0, -7)
	return &lastWeek
}

func LastMonth() *time.Time {
	lastMonth := todayMidnight.AddDate(0, -1, 0)
	return &lastMonth
}

func LastYear() *time.Time {
	lastYear := todayMidnight.AddDate(-1, 0, 0)
	return &lastYear
}

func NextWeek() *time.Time {
	nextWeek := todayMidnight.AddDate(0, 0, 7)
	return &nextWeek
}

func NextMonth() *time.Time {
	nextMonth := todayMidnight.AddDate(0, 1, 0)
	return &nextMonth
}

func NextYear() *time.Time {
	nextYear := todayMidnight.AddDate(1, 0, 0)
	return &nextYear
}