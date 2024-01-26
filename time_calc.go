package temporal

import (
	"math"
	"time"
)

func DateValue(date time.Time) float64 {
	var val time.Time = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	val = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	// Converting the time.Time form to a serial number starting from 1/1/1900
	diff := val.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, date.Location())).Hours()
	return float64(diff/24) + 2
}

// Diff returns the difference between the given time.Time and the current time.Time in the given unit
func Diff(t1, t2 time.Time, unit time.Duration, rounded ...bool) float64 {
	isRounded := false
	if len(rounded) > 0 {
		isRounded = rounded[0]
	}

	if isRounded {
		return math.Round(float64(t1.Sub(t2) / unit))
	}
	return float64(t1.Sub(t2) / unit)
}

// Latest returns the latest time from the given time.Time
func Latest(t1, t2 time.Time, tn ...time.Time) time.Time {

	timeStamps := make([]int64, len(tn)+2)
	timeStamps[0] = t1.UnixMilli()
	timeStamps[1] = t2.UnixMilli()
	for i, t := range tn {
		timeStamps[i+2] = t.UnixMilli()
	}

	// Returing the largest unix timestamp
	max := timeStamps[0]
	for _, t := range timeStamps {
		if t > max {
			max = t
		}
	}

	return time.UnixMilli(max)
}

// Earliest returns the earliest time from the given time.Time
func Earliest(t1, t2 time.Time, tn ...time.Time) time.Time {

	timeStamps := make([]int64, len(tn)+2)
	timeStamps[0] = t1.UnixMilli()
	timeStamps[1] = t2.UnixMilli()
	for i, t := range tn {
		timeStamps[i+2] = t.UnixMilli()
	}

	// Returing the smallest unix timestamp
	min := timeStamps[0]
	for _, t := range timeStamps {
		if t < min {
			min = t
		}
	}

	return time.UnixMilli(min)
}

// TruncateTime truncates the time part of the given date. It returns
// the tructed date.
func TruncateTime(date time.Time) time.Time {
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		0, 0, 0, 0,
		date.Location(),
	)
}

func WorkDay(startDate time.Time, days float64, workingDays [7]bool, holidays ...time.Time) time.Time {
	finalDateSerial := DateValue(startDate)
	weekDay := startDate.Weekday()

	// truncating the days to get the integer part
	days = float64(int(days))

	holidaysSerial := make([]float64, len(holidays))
	for i, holiday := range holidays {
		holidaysSerial[i] = DateValue(holiday)
	}

	for days > 0 {
		finalDateSerial++
		weekDay = (weekDay + 1) % 7
		if !workingDays[weekDay] {
			continue
		}
		found := false
		for _, holiday := range holidaysSerial {
			if finalDateSerial == holiday {
				holidaysSerial = holidaysSerial[1:]
				found = true
			}
			if found {
				break
			}
		}
		if !found {
			days--
		}
	}

	return time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(finalDateSerial-2) * 24 * time.Hour)
}
