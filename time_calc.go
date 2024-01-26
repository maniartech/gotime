package temporal

import (
	"math"
	"time"
)


// Diff returns the difference between the given time.Time and the current time.Time in the given unit
func Diff(d, t time.Time, unit time.Duration, rounded ...bool) float64 {
	isRounded := false
	if len(rounded) > 0 {
		isRounded = rounded[0]
	}

	if isRounded {
		return math.Round(float64(d.Sub(t) / unit))
	}
	return float64(d.Sub(t) / unit)
}

// Latest returns the latest time from the given time.Time
// panic if no time is given
func Latest(dt ...time.Time) time.Time {

	if len(dt) == 0 {
		panic("No time given")
	}

	timeStamps := make([]int64, 0)
	for _, t := range dt {
		timeStamps = append(timeStamps, t.UnixMilli())
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
func Earliest(dt ...time.Time) time.Time {

	if len(dt) == 0 {
		panic("No time given")
	}

	timeStamps := make([]int64, 0)
	for _, t := range dt {
		timeStamps = append(timeStamps, t.UnixMilli())
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
