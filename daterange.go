package dateutils

import (
	"errors"
	"strconv"
	"time"

	"github.com/maniartech/dateutils/errs"
)

// RelativeRange returns the range of dates relative to the current date.
// It accepts r as string in the following format:
// today, yesterday, tomorrow, last-<n>days, next-<n>days,
// thisweek, lastweek, nextweek, last-<n>weeks, next-<n>weeks,
// thismonth, lastmonth, nextmonth, last-<n>months, next-<n>months,
// thisyear, lastyear, nextyear, last-<n>years, next-<n>years
// or absolute dates separated by commas
// 2018-01-01,2018-01-01T00:00:00Z, 2018-01-01T00:00:00.000Z
func RelativeRange(r string) (time.Time, time.Time, error) {
	if r == "" {
		return time.Time{}, time.Time{}, errors.New(errs.ErrInvalidArgument)
	}

	now := time.Now().UTC()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// if today, return time from 00:00:00 to next day 00:00:00
	if r == "today" {
		return todayMidnight, todayMidnight.AddDate(0, 0, 1), nil
	}

	// yesterday, return time from yesterday 00:00:00 to today 00:00:00
	if r == "yesterday" {
		yesterday := todayMidnight.AddDate(0, 0, -1)
		return yesterday, todayMidnight, nil
	}

	// tomorrow, return time from tomorrow 00:00:00 to day after tomorrow 00:00:00
	if r == "tomorrow" {
		tomorrow := todayMidnight.AddDate(0, 0, 1)
		return tomorrow, tomorrow.AddDate(0, 0, 1), nil
	}

	// last-<n>days, return time from n days ago 00:00:00 to end of day
	if r[:4] == "last" {
		days, err := strconv.Atoi(r[5 : len(r)-4])
		if err != nil {
			return time.Time{}, time.Time{}, errors.New(errs.ErrInvalidArgument)
		}
		start := todayMidnight.AddDate(0, 0, -days)
		return start, todayMidnight.AddDate(0, 0, 1), nil
	}

	// next-<n>days, return time from midnight to next n days. It should
	// return time from today 00:00:00 to n days from the end of the day
	if r[:4] == "next" {
		days, err := strconv.Atoi(r[5 : len(r)-4])
		if err != nil {
			return time.Time{}, time.Time{}, errors.New(errs.ErrInvalidArgument)
		}
		return todayMidnight, todayMidnight.AddDate(0, 0, days+1), nil
	}

	return time.Time{}, time.Time{}, errors.New(errs.ErrInvalidArgument)
}
