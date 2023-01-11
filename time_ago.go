package datetime

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
	timeScaleVal := ""
	timeVal := 0
	future := false
	timeSince := time.Since(date)
	// If timeSince is negative, then the date is in the future
	if timeSince < 0 {
		timeSince = timeSince * -1
		future = true
	}

	// Checking the timeScale from the smallest level {Just Now, Few Minutes ago, ... }
	if timeSince.Hours() < 1 {
		if timeSince.Minutes() < 1 {
			if future {
				return "In a few seconds"
			}
			return "Just now"
		} else {
			if future {
				return "In a few minutes"
			}
			return "Few minutes ago"
		}
	}

	// Going through all the time scales and finding the closest one
	hours := timeSince.Hours()
	for i, val := range hoursList[:4] {

		if hours >= val && hours < hoursList[i+1] {
			timeVal = int(math.Round(hours / val))
			timeScaleVal = timeScale[i]
		}
	}
	if timeScaleVal == "" {
		timeVal = int(math.Round(hours / hoursList[4]))
		timeScaleVal = timeScale[4]
	}

	// If the time scale is in days and the date is in the future,
	// then we need to print "Tomorrow" and "Yesterday" instead of "In a day" and "Last day"
	if timeScaleVal == "days" || timeScaleVal == "hours" {

		now := time.Now().UTC()
		nowYear, nowMonth, nowDay := now.Date()

		// Calculating the midnight of the day after tomorrow, tomorrow, yesterday and day before yesterday
		dayAfterTomorrowMidnight := Date(nowYear, int(nowMonth), nowDay+2)
		tomorrowMidnight := Date(nowYear, int(nowMonth), nowDay+1)
		yesterdayMidnight := Date(nowYear, int(nowMonth), nowDay)
		dayBeforeYesterdayMidnight := Date(nowYear, int(nowMonth), nowDay-1)

		//If the date is after tomorrow midnight and before day after tomorrow midnight then print "Tomorrow"
		if date.After(tomorrowMidnight) && date.Before(dayAfterTomorrowMidnight) {

			return "Tomorrow"
		}
		if date.After(dayBeforeYesterdayMidnight) && date.Before(yesterdayMidnight) {

			return "Yesterday"
		}
	}

	// Printing the time difference
	// If the time difference is 1, then we need to print singular form of the time scale

	// Returning the time difference if the date is in the future
	if future {
		if timeVal == 1 {
			return fmt.Sprintf("In a %s", timeScaleVal[:len(timeScaleVal)-1])
		}
		return fmt.Sprintf("In %d %s", timeVal, timeScaleVal)
	}

	// Returning the time difference if the date is in the future
	if timeVal == 1 {
		return fmt.Sprintf("Last %s", timeScaleVal[:len(timeScaleVal)-1])
	}
	return fmt.Sprintf("%d %s ago", timeVal, timeScaleVal)
}

func Date(Year, Month, Day int) time.Time {
	return time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC)
}
