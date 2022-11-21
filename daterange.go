package dateutils

import (
	"errors"
	"strconv"
	"time"
)

// RelativeRange returns the range of dates relative to the current date.
// It accepts r as string in the following format:
// today, yesterday, tomorrow, last-<n>days, next-<n>days,
// thisweek, lastweek, nextweek, last-<n>weeks, next-<n>weeks,
// thismonth, lastmonth, nextmonth, last-<n>months, next-<n>months,
// thisyear, lastyear, nextyear, last-<n>years, next-<n>years
// or absolute dates separated by commas
// 2018-01-01,2018-01-01T00:00:00Z, 2018-01-01T00:00:00.000Z
func RelativeRange(r string) (*time.Time, *time.Time, error) {
	if r == "" {
		return &time.Time{}, &time.Time{}, errors.New(ErrInvalidArgument)
	}

	now := time.Now().UTC()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// if today, return time from 00:00:00 to next day 00:00:00
	if r == "today" {
		nextday := todayMidnight.AddDate(0, 0, 1)
		return &todayMidnight, &nextday, nil
	}

	// yesterday, return time from yesterday 00:00:00 to today 00:00:00
	if r == "yesterday" {
		yesterday := todayMidnight.AddDate(0, 0, -1)
		return &yesterday, &todayMidnight, nil
	}

	// tomorrow, return time from tomorrow 00:00:00 to day after tomorrow 00:00:00
	if r == "tomorrow" {
		tomorrow := todayMidnight.AddDate(0, 0, 1)
		dayAfterTomorrow := tomorrow.AddDate(0, 0, 1)
		return &tomorrow, &dayAfterTomorrow, nil
	}

	// thisweek, return time from 00:00:00 to next week 00:00:00
	if r == "thisweek" {
		nextweek := todayMidnight.AddDate(0, 0, 7)
		return &todayMidnight, &nextweek, nil
	}

	// lastweek, return time from lastwek 00:00:00 to today 00:00:00
	if r == "lastweek" {
		lastweek := todayMidnight.AddDate(0, 0, -7)
		return &lastweek, &todayMidnight, nil
	}

	// nextweek, return time from next week 00:00:00 to a week later to that 00:00:00
	if r == "nextweek" {
		nextweek := todayMidnight.AddDate(0, 0, 7)
		secondweek := nextweek.AddDate(0, 0, 7)
		return &nextweek, &secondweek, nil
	}

	// thismonth, return time from 00:00:00 to next month 00:00:00
	if r == "thismonth" {
		nextmonth := todayMidnight.AddDate(0, 1, 0)
		return &todayMidnight, &nextmonth, nil
	}

	// lastmonth, return time from last month 00:00:00 to today 00:00:00
	if r == "lastmonth" {
		lastmonth := todayMidnight.AddDate(0, -1, 0)
		return &lastmonth, &todayMidnight, nil
	}

	// nextmonth, return time from next month 00:00:00 to a month later to that 00:00:00
	if r == "nextmonth" {
		nextmonth := todayMidnight.AddDate(0, 1, 0)
		secondmonth := nextmonth.AddDate(0, 1, 0)
		return &nextmonth, &secondmonth, nil
	}

	// thisyear, return time from 00:00:00 to next year 00:00:00
	if r == "thisyear" {
		nextyear := todayMidnight.AddDate(1, 0, 0)
		return &todayMidnight, &nextyear, nil
	}

	// lastyear, return time from last year 00:00:00 to today 00:00:00
	if r == "lastyear" {
		lastyear := todayMidnight.AddDate(-1, 0, 0)
		return &lastyear, &todayMidnight, nil
	}

	// nextyear, return time from next year 00:00:00 to a year later to that 00:00:00
	if r == "nextyear" {
		nextyear := todayMidnight.AddDate(1, 0, 0)
		secondyear := nextyear.AddDate(1, 0, 0)
		return &nextyear, &secondyear, nil
	}

	// last-<n>days, return time from n days/weeks/months/years ago 00:00:00 to end of day
	if r[:5] == "last-" {
		if r[len(r)-4:] == "days" {
			days, err := strconv.Atoi(r[5 : len(r)-4])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			start := todayMidnight.AddDate(0, 0, -days)
			end := todayMidnight.AddDate(0, 0, 1)
			return &start, &end, nil
		} else if r[len(r)-5:] == "weeks" {
			weeks, err := strconv.Atoi(r[5 : len(r)-5])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			start := todayMidnight.AddDate(0, 0, -7*weeks)
			return &start, &todayMidnight, nil
		} else if r[len(r)-6:] == "months" {
			months, err := strconv.Atoi(r[5 : len(r)-6])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			start := todayMidnight.AddDate(0, -months, 0)
			return &start, &todayMidnight, nil
		} else if r[len(r)-5:] == "years" {
			years, err := strconv.Atoi(r[5 : len(r)-5])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			start := todayMidnight.AddDate(-years, 0, 0)
			return &start, &todayMidnight, nil
		}
	}

	// next-<n>days, return time from midnight to next n days/weeks/months/years.
	// It should return time from today 00:00:00 to n days/weeks/months/years from the end of the day.
	if r[:5] == "next-" {
		if r[len(r)-4:] == "days" {
			days, err := strconv.Atoi(r[5 : len(r)-4])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			end := todayMidnight.AddDate(0, 0, days+1)
			return &todayMidnight, &end, nil
		} else if r[len(r)-5:] == "weeks" {
			weeks, err := strconv.Atoi(r[5 : len(r)-5])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			end := todayMidnight.AddDate(0, 0, 7*weeks)
			return &todayMidnight, &end, nil
		} else if r[len(r)-6:] == "months" {
			months, err := strconv.Atoi(r[5 : len(r)-6])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			end := todayMidnight.AddDate(0, months, 0)
			return &todayMidnight, &end, nil
		} else if r[len(r)-5:] == "years" {
			years, err := strconv.Atoi(r[5 : len(r)-5])
			if err != nil {
				return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
			}
			end := todayMidnight.AddDate(years, 0, 0)
			return &todayMidnight, &end, nil
		}
	}

	return &time.Time{}, &time.Time{}, errors.New(errs.ErrInvalidArgument)
}
