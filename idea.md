# Contains, the library structural and functional trial and error notes

```go

temporal.Toady()

temporal.DateRange().Today() // 2019-01-01 00:00:00 +0000 UTC - 2019-01-01 23:59:59.999999999 +0000 UTC
temporal.DateRange().Yesterday()

temporal.DateRange().ThisWeek() // 2019-01-01 00:00:00 +0000 UTC - 2019-01-07 23:59:59.999999999 +0000 UTC
temporal.DateRange().LastWeek()

temporal.DateRange().SinceWeeks(3)
temporal.DateRanae().SinceMonth(s) // From exactly 1 month ago to now
temporal.DateRange().LastMonth() // From first day of last month to last day of last month

temporal.DateRange().SinceMonths(3).AddDays(1)


temporal.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
```
