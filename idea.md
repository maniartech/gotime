# Contains, the library structural and functional trial and error notes

```go

temporal.Toady()

temporal.DateRange().Today() // 2019-01-01 00:00:00 +0000 UTC - 2019-01-01 23:59:59.999999999 +0000 UTC
temporal.DateRange().Yesterday()

temporal.DateRange().LastWees(1) // 2019-01-01 00:00:00 +0000 UTC - 2019-01-07 23:59:59.999999999 +0000 UTC

// DateRange Functions
temporal.DateRange().NextWeek(2)
temporal.DateRange().SinceWeeks(3)
temporal.DateRanae().Months(-1) // From exactly 1 month ago to now
temporal.DateRange().LastMonth() // From first day of last month to last day of last month
temporal.DateRange().SinceMonths(3).AddDays(1)

// Date Functions
temporal.Yesterday()
temporal.Tomorrow()
temporal.Today()

temporal.Date().Monday(1).Range()
temporal.Date().Monday(1).Time()
temporal.Parse().Time() // return time.Time
tempora.Parse("2023-1-1T12:30", "yyyy-m-dThh:mm").TruncateTime().Time()

// Root functions
temporal.Format(time.Now(), "YYYY-MM-DD HH:mm:ss") // 2019-01-01 00:00:00
temproal.FormatTS(1231232, "dd mmmm, yyyy") // 21 January, 2023

temporal.Parse("2019-01-01", "yyyy-mm-dd").TimeAgo()



temporal.ParseToDate("2019-01-01", "yyyy-mm-dd").Format("dd-mm-yyyy")
temporal.ParseToDate("2019-01-01", "yyyy-mm-dd").Monday()



temporal.Last().Weeks(1)
temporal.Last().Months(2)
temporal.Last().Years(2)

// Format Functions
temporal.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
temporal.FormatStr("2019-01-01 00:00:00 +0000", "YYYY-MM-DD")
temporal.FormatStr("2022-02", "DD-MM-YYYY", "YYYY-MM")
temporal.FormatTS(1231232, "dd mmmm, yyyy") // 21 January, 2023


// Time ago
```
