package gotime

import (
	"math"
	"time"
)

// Create a struct called Date
// type time.Time time.Time

// func Time() time.Time {
// 	return time.Time(d)
// }

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// WeekDay return the what day of the week it is from 0 to 6
// 0 is Sunday, 1 is Monday, 2 is Tuesday, 3 is Wednesday, 4 is Thursday, 5 is Friday, 6 is Saturday
// func WeekDay() int {
// 	return int(d.Weekday())
// }

// WeekDayName return the name of the day of the week
// func WeekDayName() string {
// 	return d.Weekday().String()
// }

// WeekDayShortName return the short name of the day of the week
func WeekDayShortName(d time.Time) string {
	return d.Weekday().String()[0:3]
}

// // Month return the month of the year from 1 to 12
// func Month() int {
// 	return int(d.Month())
// }

// // MonthName return the name of the month of the year
// func MonthName() string {
// 	return d.Month().String()
// }

// MonthShortName return the short name of the month of the year
func MonthShortName(d time.Time) string {
	return d.Month().String()[0:3]
}

// // Year return the year
// func Year() int {
// 	return d.Year()
// }

// Day return the day of the month from 1 to 31
// func Day() int {
// 	return d.Day()
// }

// // Hour return the hour of the day from 0 to 23
// func Hour() int {
// 	return d.Hour()
// }

// // Minute return the minute of the hour from 0 to 59
// func Minute() int {
// 	return d.Minute()
// }

// // Second return the second of the minute from 0 to 59
// func Second() int {
// 	return d.Second()
// }

// // Millisecond return the millisecond of the second from 0 to 999
// func Millisecond() int {
// 	return d.Nanosecond() / 1000000
// }

// // Microsecond return the microsecond of the second from 0 to 999999
// func Microsecond() int {
// 	return d.Nanosecond() / 1000
// }

// Date().Monday(weeks) returns the time.Time of the current week's Monday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Monday
func Monday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+1))

	return d
}

// Date().Tuesday(weeks) returns the time.Time of the current week's Tuesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Tuesday
func Tuesday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+2))

	return d
}

// Date().Wednesday(weeks) returns the time.Time of the current week's Wednesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Wednesday
func Wednesday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+3))

	return d
}

// Date().Thursday(weeks) returns the time.Time of the current week's Thursday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Thursday
func Thursday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+4))

	return d
}

// Date().Friday(weeks) returns the time.Time of the current week's Friday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Friday
func Friday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+5))

	return d
}

// Date().Saturday(weeks) returns the time.Time of the current week's Saturday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Saturday
func Saturday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+6))

	return d
}

// Date().Sunday(weeks) returns the time.Time of the current week's Sunday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Sunday
func Sunday(weeks ...int) time.Time {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d := (time.Now().AddDate(0, 0, -int(time.Now().Weekday())+w*7+7))

	return d
}
