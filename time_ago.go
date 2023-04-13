package temporal

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

// TimeAgo calculates the relative time difference between a given timestamp and the current time, and returns a string
// that describes the time difference in human-readable terms. The function takes a time.Time object representing the
// timestamp, and an optional baseTime parameter, which can be used to specify a different base time to use instead of the
// current time. The function returns a string that describes the time difference in human-readable terms, such as "2 weeks
// ago" or "In a few minutes".
//
// Example usage:
//
//	// Create a time object representing one week ago
//	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)
//
//	// Calculate the relative time difference between the timestamp and the current time
//	timeAgo := TimeAgo(oneWeekAgo)
//
//	fmt.Println("One week ago:", timeAgo)
//	// Output: One week ago: Last week
//
//	// Create a time object representing 3 hours and 45 minutes ago
//	threeHoursAgo := time.Now().Add(-3 * time.Hour).Add(-45 * time.Minute)
//
//	// Calculate the relative time difference between the timestamp and the current time
//	timeAgo = TimeAgo(threeHoursAgo)
//
//	fmt.Println("Three hours ago:", timeAgo)
//	// Output: Three hours ago: 3 hours ago
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

func calculateTimeVal(scale int, hours float64, future bool) string {
	timeVal := int(math.Round(hours / hoursList[scale]))
	timeScaleVal := timeScale[scale]

	if timeVal == 1 {
		return lastOrNextSingular(timeScaleVal, future)
	}
	return lastOrNext(timeScaleVal, timeVal, future)
}

// Returns the string "Just now" or "In a few seconds" depending on the future bool
func justNow(future bool) string {
	if future {
		return "In a few seconds"
	}
	return "Just now"
}

// Returns the string "A minute ago" or "In a minute" depending on the future bool
func minuteAgo(future bool) string {
	if future {
		return "In a minute"
	}
	return "A minute ago"
}

// Returns the string "Few minutes ago" or "In a few minutes" depending on the future bool
func fewMinutesAgo(future bool) string {
	if future {
		return "In a few minutes"
	}
	return "Few minutes ago"
}

// Returns the string "Yesterday" or "Tomorrow" depending on the future bool
func yesterdayOrTomorrow(date time.Time, future bool) string {
	now := time.Now().UTC()
	nowYear, nowMonth, nowDay := now.Date()
	if future {
		dayAfterTomorrowMidnight := DateCreate(nowYear, int(nowMonth), nowDay+2)
		tomorrowMidnight := DateCreate(nowYear, int(nowMonth), nowDay+1)

		//If the date is after tomorrow midnight and before day after tomorrow midnight then print "Tomorrow"
		if date.After(tomorrowMidnight) && date.Before(dayAfterTomorrowMidnight) {
			return "Tomorrow"
		}
		return ""
	}

	// Past
	// Calculating the midnight of the day after tomorrow, tomorrow, yesterday and day before yesterday
	yesterdayMidnight := DateCreate(nowYear, int(nowMonth), nowDay)
	dayBeforeYesterdayMidnight := DateCreate(nowYear, int(nowMonth), nowDay-1)

	if date.After(dayBeforeYesterdayMidnight) && date.Before(yesterdayMidnight) {
		return "Yesterday"
	}
	return ""
}

// Returns the string "Last <timeScaleVal>" or "In a <timeScaleVal>" depending on the future bool
func lastOrNextSingular(timeScaleVal string, future bool) string {
	if future {
		return fmt.Sprintf("In a %s", timeScaleVal[:len(timeScaleVal)-1])
	}
	return fmt.Sprintf("Last %s", timeScaleVal[:len(timeScaleVal)-1])
}

func lastOrNext(timeScaleVal string, timeVal int, future bool) string {
	if future {
		return fmt.Sprintf("In %d %s", timeVal, timeScaleVal)
	}
	return fmt.Sprintf("%d %s ago", timeVal, timeScaleVal)
}
