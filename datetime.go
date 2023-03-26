package temporal

import (
	"math"
	"time"
)

// Create a struct called Date
type DateTime time.Time

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) UTC() DateTime {
	// return DateTime((*time.Time)(&d).UTC())
	return DateTime(time.Time(d).UTC())
}

// WeekDay return the what day of the week it is from 0 to 6
// 0 is Sunday, 1 is Monday, 2 is Tuesday, 3 is Wednesday, 4 is Thursday, 5 is Friday, 6 is Saturday
func (d DateTime) WeekDay() int {
	return int(d.Time().Weekday())
}

// WeekDayName return the name of the day of the week
func (d DateTime) WeekDayName() string {
	return d.Time().Weekday().String()
}

// WeekDayShortName return the short name of the day of the week
func (d DateTime) WeekDayShortName() string {
	return d.Time().Weekday().String()[0:3]
}

// Month return the month of the year from 1 to 12
func (d DateTime) Month() int {
	return int(d.Time().Month())
}

// MonthName return the name of the month of the year
func (d DateTime) MonthName() string {
	return d.Time().Month().String()
}

// MonthShortName return the short name of the month of the year
func (d DateTime) MonthShortName() string {
	return d.Time().Month().String()[0:3]
}

// Year return the year
func (d DateTime) Year() int {
	return d.Time().Year()
}

// Day return the day of the month from 1 to 31
func (d DateTime) Day() int {
	return d.Time().Day()
}

// Hour return the hour of the day from 0 to 23
func (d DateTime) Hour() int {
	return d.Time().Hour()
}

// Minute return the minute of the hour from 0 to 59
func (d DateTime) Minute() int {
	return d.Time().Minute()
}

// Second return the second of the minute from 0 to 59
func (d DateTime) Second() int {
	return d.Time().Second()
}

// Millisecond return the millisecond of the second from 0 to 999
func (d DateTime) Millisecond() int {
	return d.Time().Nanosecond() / 1000000
}

// Microsecond return the microsecond of the second from 0 to 999999
func (d DateTime) Microsecond() int {
	return d.Time().Nanosecond() / 1000
}

// Add adds the given duration to the DateTime
func (d DateTime) Add(duration time.Duration) DateTime {
	return DateTime(d.Time().Add(duration))
}

// AddDate adds the given number of years, months, and days to the DateTime
func (d DateTime) AddDate(years, months, days int) DateTime {
	return DateTime(d.Time().AddDate(years, months, days))
}

// Diff returns the difference between the given DateTime and the current DateTime in the given unit
func (d DateTime) Diff(t DateTime, unit time.Duration, rounded ...bool) float64 {
	isRounded := false
	if len(rounded) > 0 {
		isRounded = rounded[0]
	}
	if isRounded {
		return math.Round(float64(d.Time().Sub(t.Time()) / unit))
	}
	return float64(d.Time().Sub(t.Time()) / unit)
}

// Local returns the local time
func (d DateTime) Local() DateTime {
	return DateTime(d.Time().Local())
}

// In returns the DateTime in the given location
func (d DateTime) In(loc *time.Location) DateTime {
	return DateTime(d.Time().In(loc))
}

// Round returns the DateTime rounded to the nearest unit
func (d DateTime) Round(unit time.Duration) DateTime {
	return DateTime(d.Time().Round(unit))
}

// Date().Monday(weeks) returns the DateTime of the current week's Monday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Monday
func (d DateTime) Monday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+1))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+1)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Tuesday(weeks) returns the DateTime of the current week's Tuesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Tuesday
func (d DateTime) Tuesday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+2))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+2)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Wednesday(weeks) returns the DateTime of the current week's Wednesday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Wednesday
func (d DateTime) Wednesday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+3))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+3)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Thursday(weeks) returns the DateTime of the current week's Thursday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Thursday
func (d DateTime) Thursday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+4))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+4)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Friday(weeks) returns the DateTime of the current week's Friday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Friday
func (d DateTime) Friday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+5))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+5)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Saturday(weeks) returns the DateTime of the current week's Saturday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Saturday
func (d DateTime) Saturday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+6))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+6)
	// d.End = DayEnd(d.Start)

	return d
}

// Date().Sunday(weeks) returns the DateTime of the current week's Sunday
// weeks is the number of weeks to add to the current week
// If the value is negative, it will return the previous week's Sunday
func (d DateTime) Sunday(weeks ...int) DateTime {
	w := 0
	if len(weeks) > 0 {
		w = weeks[0]
	}
	d = (Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7+7))
	// d.Start = Now().AddDate(0, 0, -int(Now().Time().Weekday())+w*7)
	// d.End = DayEnd(d.Start)

	return d
}
