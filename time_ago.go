package temporal

import (
	"fmt"
	"math"
	"time"
)

const (
	hoursInYear  = 8760
	hoursInMonth = 730.001
	hoursInWeek  = 168
	hoursInDay   = 24
	hoursInHour  = 1
)

var hoursList = [5]float64{hoursInHour, hoursInDay, hoursInWeek, hoursInMonth, hoursInYear}
var timeScale = [5]string{"hours", "days", "weeks", "months", "years"}

// Calculates the relative time difference since time.Now()
func TimeAgo(date time.Time) string {
	future := false
	timeSince := time.Since(date)

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
	val := yesterdayOrTomorrow(date, future)
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

func DateCreate(Year, Month, Day int) time.Time {
	return time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC)
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
