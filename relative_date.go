package temporal

import (
	"errors"
	"time"
)

type RelativeDateType string

const (
	Yesterday   RelativeDateType = "yesterday"
	Today       RelativeDateType = "today"
	Tomorrow    RelativeDateType = "tomorrow"
	NextWeek    RelativeDateType = "next-week"
	WeekAgo     RelativeDateType = "week-ago"
	PrevWeek    RelativeDateType = "prev-week"
	WeekFromNow RelativeDateType = "week-from-now"
	NextMonth   RelativeDateType = "next-month"
	LastMonth   RelativeDateType = "last-month"
	NextYear    RelativeDateType = "next-year"
	LastYear    RelativeDateType = "last-year"
)

// DateR is a utility function accepts a relative date-time string such as "now", "today", "tomorrow", "yesterday", "next-week", "last month", etc. and returns a time.Time value and an error if the parsing fails.
//
// Example usage:
//
//	parsedDate, err := DateR(temporal.NextWeek)
//	if err != nil {
//	    panic(err)
//	}
//	fmt.Println(parsedDate) // 2021-01-10 00:00:00 +0000 UTC
func DateR(value RelativeDateType, relatedTo ...time.Time) (time.Time, error) {
	now := time.Now()
	if len(relatedTo) > 0 {
		now = relatedTo[0]
	}

	switch value {
	case Yesterday:
		return now.AddDate(0, 0, -1), nil
	case Today:
		return now, nil
	case Tomorrow:
		return now.AddDate(0, 0, 1), nil
	case NextWeek:
		return now.AddDate(0, 0, 7), nil
	case WeekAgo:
		return now.AddDate(0, 0, -7), nil
	case PrevWeek: // Previous week starting from Monday
		return now.AddDate(0, 0, -7-now.Weekday()), nil

	}

	return time.Time{}, errors.New("invalid relative date type")
}
