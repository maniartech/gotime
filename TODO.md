# TODOs

[✓] Intuitive Datetime Format (IDF), parsing and conversion function 💯

```go
// Parse a date
t, err := temporal.Parse("2023-11-25", "yyyy-mm-dd") // 2023-11-25 00:00:00 +0000 UTC

// Format a date
s := temporal.Format(time.Now(), "yyyy-mm-dd") // 2021-09-25

// Convert date string to different format
s, err := temporal.Convert("2012-12-01", "yyyy-mm-dd", "mmmm dt, yyyy") // December 1st, 2012
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
temporal.IsLeapYear(2020) // true
temporal.DaysInMonth(2020, 2) // 29
temporal.DaysInYear(2020) // 366
temporal.EoD(time.Now()) // 2021-09-25 23:59:59.999999999 +0000 UTC
temporal.EoM(time.Now()) // 2021-09-30 23:59:59.999999999 +0000 UTC
```

[ ] Relative time ⛔

```go

temporal.Yesterday(time.Now()) // 2021-09-24 00:00:00 +0000 UTC
temporal.Tomorrow(time.Now()) // 2021-09-26 00:00:00 +0000 UTC
```

[ ] Relative date range with custom start date ⛔

```go
start, end, err := temporal.RelativeRange("last-week", time.Now()) // 2021-09-19 00:00:00 +0000 UTC, 2021-09-25 00:00:00 +0000 UTC
```
