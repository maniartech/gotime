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
	seconds := int(timeSince.Seconds())
	// If timeSince is negative, then the date is in the future
	if timeSince < 0 {
		timeSince = timeSince * -1
		future = true
	}
	// Checking through all the time scales from seconds to years
	if future {
		if seconds < 10 {
			return "In a few seconds"
		} else if seconds < 60 {
			return "In a minute"
		} else if seconds < 3600 {
			return "In a few minutes"
		}
	}

	//Past
	if seconds < 10 {
		return "Just now"
	} else if seconds < 60 {
		return "A Minute ago"
	} else if seconds < 3600 {
		return "A few minutes ago"
	}

	// Going through all the time scales and finding the closest one
	hours := timeSince.Hours()
	for i, val := range hoursList[:4] {
		if hours >= val && hours < hoursList[i+1] {
			timeVal = int(math.Round(hours / val))
			timeScaleVal = timeScale[i]
			break
		}
	}
	if timeScaleVal == "" {
		timeVal = int(math.Round(hours / hoursList[4]))
		timeScaleVal = timeScale[4]
	}

	// Returning the time difference if the date is in the future
	if future {
		// If the date is after tomorrow midnight and before day after tomorrow midnight then print "Tomorrow"
		if timeScaleVal == "days" || timeScaleVal == "hours" {
			now := time.Now().UTC()
			nowYear, nowMonth, nowDay := now.Date()

			dayAfterTomorrowMidnight := Date(nowYear, int(nowMonth), nowDay+2)
			tomorrowMidnight := Date(nowYear, int(nowMonth), nowDay+1)

			//If the date is after tomorrow midnight and before day after tomorrow midnight then print "Tomorrow"
			if date.After(tomorrowMidnight) && date.Before(dayAfterTomorrowMidnight) {
				return "Tomorrow"
			}
		}

		// If the time difference is 1, then we need to print singular form of the time scale
		if timeVal == 1 {
			return fmt.Sprintf("In a %s", timeScaleVal[:len(timeScaleVal)-1])
		}
		return fmt.Sprintf("In %d %s", timeVal, timeScaleVal)
	}
	// Past
	// If the date is before yesterday midnight and after day before yesterday midnight then print "Yesterday"
	if timeScaleVal == "days" || timeScaleVal == "hours" {

		now := time.Now().UTC()
		nowYear, nowMonth, nowDay := now.Date()

		// Calculating the midnight of the day after tomorrow, tomorrow, yesterday and day before yesterday
		yesterdayMidnight := Date(nowYear, int(nowMonth), nowDay)
		dayBeforeYesterdayMidnight := Date(nowYear, int(nowMonth), nowDay-1)

		if date.After(dayBeforeYesterdayMidnight) && date.Before(yesterdayMidnight) {

			return "Yesterday"
		}
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
