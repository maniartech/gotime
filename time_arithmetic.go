package gotime

import "time"

// Hours returns the time after adding the specified number of hours to the given time.
// If no time is provided, it uses the current time. Negative values subtract hours.
//
// Example:
//   futureTime := gotime.Hours(5)           // 5 hours from now
//   pastTime := gotime.Hours(-2, someTime)  // 2 hours before someTime
//   noChange := gotime.Hours(0)             // same as time.Now()
func Hours(hours int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return t.Add(time.Duration(hours) * time.Hour)
}

// Minutes returns the time after adding the specified number of minutes to the given time.
// If no time is provided, it uses the current time. Negative values subtract minutes.
//
// Example:
//   futureTime := gotime.Minutes(30)           // 30 minutes from now
//   pastTime := gotime.Minutes(-15, someTime)  // 15 minutes before someTime
//   noChange := gotime.Minutes(0)              // same as time.Now()
func Minutes(minutes int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return t.Add(time.Duration(minutes) * time.Minute)
}

// Seconds returns the time after adding the specified number of seconds to the given time.
// If no time is provided, it uses the current time. Negative values subtract seconds.
//
// Example:
//   futureTime := gotime.Seconds(45)           // 45 seconds from now
//   pastTime := gotime.Seconds(-30, someTime)  // 30 seconds before someTime
//   noChange := gotime.Seconds(0)              // same as time.Now()
func Seconds(seconds int, dt ...time.Time) time.Time {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	return t.Add(time.Duration(seconds) * time.Second)
}
