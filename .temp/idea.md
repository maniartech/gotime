# Contains, the library structural and functional research, trial and error notes!

```go
gotime.DateRange().Today() // 2019-01-01 00:00:00 +0000 UTC - 2019-01-01 23:59:59.999999999 +0000 UTC
gotime.DateRange().Yesterday()

// DateRange Functions

gotime.DateRange().Weeks(-2) // From exactly 2 weeks ago to now
gotime.DateRange(gotime.FromTime(t)).Weeks(-2) // Yesterday as a base date, from exactly 2 weeks ago to yesterday

gotime.DateRange().Days(3)
gotime.DateRanae().Months(-1) // From exactly 1 month ago to now

gotime.DateRange(baseTime).Months() // From the base time to the end of the month
gotime.DateRange()
  .Months(1) // From the current time to one month from now
  .Months(-1) // From exactly one month ago to now
  .Months(-1, true) // true for previous month. Previous months 1st day at 00:00:00 to the end of the month
  .Months(1, true) // true for next month. Next months 1st day at 00:00:00 to the end of the month

  .Weeks(1) // From the current time to one week from now
  .Weeks(-1) // From exactly one week ago to now
  .Weeks(-1, true) // true for previous week. Previous weeks 1st day at 00:00:00 to the end of the week
  .Weeks(1, true) // true for next week. Next weeks 1st day at 00:00:00 to the end of the week

  .Days(1) // From the 00:00:00 of the current day to the 23:59:59 of the next day
  .Days(-1) // From the 00:00:00 of the previous day to the 23:59:59 of the current day

  .From()
  .To()

// Date Functions
gotime.Yesterday()
gotime.Now() // 2019-01-01 13:37:00 +0000 UTC, the current time
gotime.Today() // 2019-01-01 00:00:00 +0000 UTC, the current date at 00:00:00 without time
gotime.Tomorrow()
gotime.Today() // 2019-01-01 00:00:00 +0000 UTC
gotime.Parse()
gotime.DateFromTS()
gotime.DateFromTime()
gotime.DateFromString()

gotime.DateTime
  .Yesterday()
  .DayStart() // 2019-01-01 00:00:00 +0000 UTC
  .DayEnd() // 2019-01-01 23:59:59.999999999 +0000 UTC

  .Local() // 2019-01-01 05:30:00 +0530 IST
  .UTC() // 2019-01-01 00:00:00 +0000 UTC

  // WeekDays
  .Mondays() // Accepts a number of weeks, if not provided, it will return the Monday of the current week. If negative, it will return the Monday of the previous weeks. If positive, it will return the Monday of the next weeks.
  .Tuesdays(1)
  .Wednesdays(1)
  .Thursdays(1)
  .Fridays(1)
  .Saturdays(1)
  .Sundays(1)

  .WeekDay() // 0-6
  .WeekDayName() // Monday
  .WeekDayNameShort() // Mon

  .Month() // 1-12
  .MonthName() // January
  .MonthNameShort() // Jan

  .Year() // 2019

  .Day() // 1-31
  .Hour() // 0-23
  .Minute() // 0-59
  .Second() // 0-59
  .Millisecond() // 0-999
  .Microsecond() // 0-999999

  // Add
  .Add(1, gotime.Days) // Days, Weeks, Months, Years, Hours, Minutes, Seconds, Milliseconds, Microseconds

  // Diff
  .Diff(time.Now(), gotime.Days) // Days, Weeks, Months, Years
  .Diff(time.Now(), gotime.Weeks, true) // true for rounded value

  // Range - Range function always returns the earliest date first
  .Range() // Returns a range of dates from the current date to the end of the month
  .Range(gotime.Yesterday()) // Returns a range of dates from the current date to yesterday

  // Format

  .ToTime()
  .ToTS()
  .ToString()
  .Format("dd-mm-yyyy") // 01-01-2019
  .Format(gotime.ISODate) // 2019-01-01 00:00:00
  .Format(gotime.ISODateTime) // 2019-01-01 00:00:00 +0000 UTC
  .Format(gotime.ISODateTimeShort) // 2019-01-01 00:00:00
  .Format(gotime.ISODateTimeMicro) // 2019-01-01 00:00:00.000000 +0000 UTC
  .Format(gotime.JSONDate) // 2019-01-01T00:00:00.000Z

// Factory Functions - Creates a new gotime.DateTime object
gotime.FromTS(4324234234)
gotime.FromTime(time.Now())
gotime.FromString("2019-01-01", "yyyy-mm-dd")

// Format functions

// Time Ago
gotime.TryParse("2019-01-01")                  // time.Time, errors.Error
gotime.Parse("2019-01-01", "yyyy-mm-dd")    // time.Time, errors.Error

gotime.ParseAndFormat("2019-01-01", "dd-mm-yyyy") // Automatically detects the format, and formats it to dd-mm-yyyy
gotime.ParseAndFormat("2019-01-01", "dd-mm-yyyy", "yyyy-mm-dd") // Format is provided as the third argument


gotime.TimeAgo(
  gotime.Catch(gotime.Parse("2019-01-01", "yyyy-mm-dd"), time.Now()),
)

gotime.Weeks(1, time.Now())
gotime.Months(2)
gotime.Years(2)

gotime.EoD(
  gotime.Monday(-1) // Last Monday
) // Last Monday at 23:59:59.999999999


// First Monday of the month
gotime.Monday(
  1, gotime.MonthStart(), true // inclusive true
                                 // inclusive means consider the current date
                                 // as well
)

// Second Monday of the month
gotime.Monday(
  2, gotime.MonthStart(), true // inclusive true
                                 // inclusive means consider the current date
                                 // as well
)

// Last Monday of the month
gotime.Monday(
  -1, gotime.MonthEnd(), true // inclusive true
                                // inclusive means consider the current date
                                // as well
)

// Monday from today, inclusive
gotime.WeekDay(gotime.Monday)

// Count of Mondays in the given month
gotime.TotalWeekdays(gotime.Monday, gotime.MonthStart(), gotime.MonthEnd())

Options
  .CacheFormats() // Default true
  .DisableCache()
  .EnableCache()

  .DateTimeFormat() // Default "YYYY-MM-DD HH:mm:ss"
  .DateFormat() // Default "YYYY-MM-DD"
  .TimeFormat() // Default "HH:mm:ss"

  .Formatter()

  // Future
  .Timezone() // Default "UTC"
  .Locale() // Default "en"
  .SetWeekDays()


```
