# TODOs

[✓] Intuitive Datetime Format (IDF), parsing and conversion function 💯

```go
// Parse a date
t, err := gotime.Parse("2023-11-25", "yyyy-mm-dd") // 2023-11-25 00:00:00 +0000 UTC

// Format a date
s := gotime.Format(time.Now(), "yyyy-mm-dd") // 2021-09-25

// Convert date string to different format
s, err := gotime.Convert("2012-12-01", "yyyy-mm-dd", "mmmm dt, yyyy") // December 1st, 2012
```

[✓] Relative date range 💯

```go
```

[✓] Time Ago 💯

```go
time.Now().Add(-1 * time.Minute).Ago() // 1 minute ago

time.Now().Add(1 * time.Hour).Ago() // 1 hour from now
```

[ ] General time related utility functions ⛔

```go
gotime.IsLeapYear(2020) // true
gotime.DaysInMonth(2020, 2) // 29
gotime.DaysInYear(2020) // 366
gotime.EoD(time.Now()) // 2021-09-25 23:59:59.999999999 +0000 UTC
gotime.EoM(time.Now()) // 2021-09-30 23:59:59.999999999 +0000 UTC
```

[ ] Relative time ⛔

```go

gotime.Yesterday(time.Now()) // 2021-09-24 00:00:00 +0000 UTC
gotime.Tomorrow(time.Now()) // 2021-09-26 00:00:00 +0000 UTC
```

[ ] Relative date range with custom start date ⛔

```go
start, end, err := gotime.RelativeRange("last-week", time.Now()) // 2021-09-19 00:00:00 +0000 UTC, 2021-09-25 00:00:00 +0000 UTC
```
