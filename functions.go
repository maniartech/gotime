package dateutils

import (
	"time"
)

func TodayStart() *time.Time {
	var now time.Time = time.Now().UTC()
	var todayMidnight time.Time = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return &todayMidnight
}

func Today() *time.Time {
	start := *TodayStart()
	return &start
}

func EoD() *time.Time {
	start := *TodayStart()
	eod := start.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	return &eod
}

func Yesterday() *time.Time {
	yesterday := TodayStart().AddDate(0, 0, -1)
	return &yesterday
}

func Tomorrow() *time.Time {
	tomorrow := TodayStart().AddDate(0, 0, 1)
	return &tomorrow
}

func LastWeek() *time.Time {
	lastWeek := TodayStart().AddDate(0, 0, -7)
	return &lastWeek
}

func LastMonth() *time.Time {
	lastMonth := TodayStart().AddDate(0, -1, 0)
	return &lastMonth
}

func LastYear() *time.Time {
	lastYear := TodayStart().AddDate(-1, 0, 0)
	return &lastYear
}

func NextWeek() *time.Time {
	nextWeek := TodayStart().AddDate(0, 0, 7)
	return &nextWeek
}

func NextMonth() *time.Time {
	nextMonth := TodayStart().AddDate(0, 1, 0)
	return &nextMonth
}

func NextYear() *time.Time {
	nextYear := TodayStart().AddDate(1, 0, 0)
	return &nextYear
}
