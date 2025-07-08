package gotime

import (
	"fmt"
	"math"
	"time"
)

const (

	// The actual average number of hours in a month is 730.001. This number is obtained by
	// taking the average number of days in a year (365.2425) and dividing it by 12, which
	// gives 30.436875. Multiplying this by 24 gives 730.001 hours.
	hoursInMonth = 730.001
	hoursInYear  = hoursInMonth * 12
	hoursInWeek  = 168 // 7 * 24
	hoursInDay   = 24  // 24
	hoursInHour  = 1   // 1
)

var hoursList = [5]float64{hoursInHour, hoursInDay, hoursInWeek, hoursInMonth, hoursInYear}
var timeScale = [5]string{"hours", "days", "weeks", "months", "years"}

// TimeAgo returns a human-readable string describing the relative time difference
// between t and the current time (or baseTime if provided).
//
// The function returns phrases like "Just now", "5 minutes ago", "Tomorrow",
// "Last week", etc. If the given time is in the future, it returns phrases
// like "In a few seconds", "In 5 minutes", "Tomorrow", etc.
//
// Example:
//
//	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)
//	result := TimeAgo(oneWeekAgo)
//	// result: "Last week"
//
//	threeHoursAgo := time.Now().Add(-3 * time.Hour)
//	result = TimeAgo(threeHoursAgo)
//	// result: "3 hours ago"
//
//	futureTime := time.Now().Add(30 * time.Minute)
//	result = TimeAgo(futureTime)
//	// result: "In a few minutes"
func TimeAgo(t time.Time, baseTime ...time.Time) string {
	future := false
	var timeSince time.Duration
	if len(baseTime) > 0 {
		timeSince = baseTime[0].Sub(t)
	} else {
		timeSince = time.Since(t)
	}

	// If timeSince is negative, then the date is in the future
	if timeSince < 0 {
		timeSince = timeSince * -1
		future = true
	}
	seconds := int(timeSince.Seconds())
	// Checking through all the time scales from seconds to years
	if seconds < 10 {
		return justNow(future)
	} else if seconds < 60 {
		return minuteAgo(future)
	} else if seconds < 3600 {
		return fewMinutesAgo(future)
	}

	//Checking if the date is yesterday or tomorrow
	val := yesterdayOrTomorrow(t, future)
	if val != "" {
		return val
	}

	hours := timeSince.Hours()
	// Going through all the time scales and finding the closest one
	for i := 0; i < len(hoursList)-1; i++ {
		if hours > hoursList[i] && hours < hoursList[i+1] {
			return calculateTimeVal(i, hours, future)
		}
	}
	return calculateTimeVal(len(hoursList)-1, hours, future)
}

// calculateTimeVal calculates the appropriate time value and returns a human-readable
// string describing the time difference.
func calculateTimeVal(scale int, hours float64, future bool) string {
	timeVal := int(math.Round(hours / hoursList[scale]))
	timeScaleVal := timeScale[scale]

	if timeVal == 1 {
		return lastOrNextSingular(timeScaleVal, future)
	}
	return lastOrNext(timeScaleVal, timeVal, future)
}

// justNow returns "Just now" for past time or "In a few seconds" for future time.
func justNow(future bool) string {
	if future {
		return "In a few seconds"
	}
	return "Just now"
}

// minuteAgo returns "A minute ago" for past time or "In a minute" for future time.
func minuteAgo(future bool) string {
	if future {
		return "In a minute"
	}
	return "A minute ago"
}

// fewMinutesAgo returns "Few minutes ago" for past time or "In a few minutes" for future time.
func fewMinutesAgo(future bool) string {
	if future {
		return "In a few minutes"
	}
	return "Few minutes ago"
}

// yesterdayOrTomorrow returns "Yesterday" for past time or "Tomorrow" for future time
// if the date falls within these ranges, otherwise returns an empty string.
func yesterdayOrTomorrow(date time.Time, future bool) string {
	now := time.Now().In(date.Location())
	nowYear, nowMonth, nowDay := now.Date()
	if future {
		dayAfterTomorrowMidnight := NewDate(nowYear, int(nowMonth), nowDay+2, date.Location())
		tomorrowMidnight := NewDate(nowYear, int(nowMonth), nowDay+1, date.Location())

		//If the date is after tomorrow midnight and before day after tomorrow midnight then print "Tomorrow"
		if date.After(tomorrowMidnight) && date.Before(dayAfterTomorrowMidnight) {
			return "Tomorrow"
		}
		return ""
	}

	// Past
	// Calculating the midnight of the day after tomorrow, tomorrow, yesterday and day before yesterday
	yesterdayMidnight := NewDate(nowYear, int(nowMonth), nowDay, date.Location())
	dayBeforeYesterdayMidnight := NewDate(nowYear, int(nowMonth), nowDay-1, date.Location())

	if date.After(dayBeforeYesterdayMidnight) && date.Before(yesterdayMidnight) {
		return "Yesterday"
	}
	return ""
}

// lastOrNextSingular returns singular forms like "Last week" for past time
// or "In a week" for future time.
func lastOrNextSingular(timeScaleVal string, future bool) string {
	if future {
		return fmt.Sprintf("In a %s", timeScaleVal[:len(timeScaleVal)-1])
	}
	return fmt.Sprintf("Last %s", timeScaleVal[:len(timeScaleVal)-1])
}

// lastOrNext returns plural forms like "5 days ago" for past time
// or "In 5 days" for future time.
func lastOrNext(timeScaleVal string, timeVal int, future bool) string {
	if future {
		return fmt.Sprintf("In %d %s", timeVal, timeScaleVal)
	}
	return fmt.Sprintf("%d %s ago", timeVal, timeScaleVal)
}
