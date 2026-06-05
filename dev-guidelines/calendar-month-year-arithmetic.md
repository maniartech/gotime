# Developer Note: Calendar Month/Year Arithmetic (Clamping & Differences)

Status: Implemented (pending review)
Affects: `relative_functions.go`, `age_calculation.go` (+ tests)
Author note date: 2026-06-05

This note documents a behavioral fix and two new functions related to
month/year arithmetic, why the change was made, the exact semantics now
guaranteed, and guidance on which function to use.

---

## 1. Background: fixed vs. calendar time

Time amounts fall into two categories that must not be conflated:

- **Fixed-length units** — millisecond, second, minute, hour, day, week. These
  have a constant length and are pure arithmetic on an instant.
- **Calendar (variable-length) units** — **month** (28–31 days) and **year**
  (365 or 366 days). These have *no fixed length* and *no agreed fixed
  conversion to days*. Operating on them requires looking at the specific date.

All of the behavior below concerns the calendar units only. Fixed units never
have these subtleties.

---

## 2. The bug that was fixed

### 2.1 Symptom

`Months` and `Years` were thin wrappers over the standard library's
`time.AddDate`:

```go
func Months(months int, dt ...time.Time) time.Time {
    // ...
    return t.AddDate(0, months, 0)   // OLD
}
```

`time.AddDate` **normalizes overflow** instead of clamping. Adding one month to
January 31 produces "February 31", which Go normalizes by rolling into March:

| Call | Old (overflow) result | Expected (clamp) |
|------|-----------------------|------------------|
| `Months(1, 2020-01-31)` | `2020-03-02` | `2020-02-29` |
| `Months(1, 2023-01-31)` | `2023-03-03` | `2023-02-28` |
| `Months(1, 2024-03-31)` | `2024-05-01` | `2024-04-30` |
| `Years(1, 2024-02-29)`  | `2025-03-01` | `2025-02-28` |

This contradicted:

1. **The functions' own doc comments**, which claimed
   `Months(1, 2020-01-31)` → `2020-02-29` ("handles month-end edge cases").
2. **Every other major date library** — `.NET AddMonths`, `java.time
   plusMonths`, Python `dateutil.relativedelta`, moment.js, and SQL `DATEADD`
   all **clamp** to the last day of the target month. Go's `AddDate` is the
   outlier.

### 2.2 Why the test never caught it

`TestMonth`/`TestYear` exercised `Months`/`Years` using `2022-01-01` (day 1,
which can never overflow) and asserted the result equalled
`fixedDate.AddDate(0, n, 0)` — i.e. they compared the function to the very
standard-library call it wrapped. The assertion was tautological and the
month-end edge case had zero coverage.

---

## 3. The fix

A calendar-aware helper performs end-of-month clamping; `Months` and `Years`
route through it.

```go
// addMonthsClamped adds months with end-of-month clamping: if the original
// day-of-month does not exist in the target month, the day is clamped to the
// last valid day. Time-of-day and location are preserved.
func addMonthsClamped(t time.Time, months int) time.Time {
    total := t.Year()*12 + (int(t.Month()) - 1) + months
    year := total / 12
    month := total%12 + 1
    if month < 1 { // normalize negative totals (Go truncated modulo)
        month += 12
        year--
    }
    day := t.Day()
    if maxDay := DaysInMonth(year, month); day > maxDay {
        day = maxDay
    }
    return time.Date(year, time.Month(month), day,
        t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
```

- `Months(n, t)` → `addMonthsClamped(t, n)`
- `Years(n, t)`  → `addMonthsClamped(t, n*12)`

`LastMonth`, `NextMonth`, `LastYear`, `NextYear` were changed to delegate to
`Months`/`Years` so they clamp consistently (previously they each called
`time.AddDate` directly and would overflow on a month-end "today").

### 3.1 Guarantees

- The day is clamped to `DaysInMonth(targetYear, targetMonth)` when the original
  day does not exist in the target month.
- Negative amounts clamp identically (`Months(-1, 2024-03-31)` → `2024-02-29`).
- `n == 0` is a no-op returning the original value unchanged.
- Time-of-day, nanoseconds, and `*time.Location` are preserved.

---

## 4. New functions: calendar-accurate differences

### 4.1 Why they are needed

The existing `MonthsBetween` / `YearsBetween` return a **fractional, nominal**
value computed as `elapsed_hours / average_unit_hours`, using the Gregorian
average of `365.2425` days/year (and `365.2425 / 12` days/month):

```go
gotime.MonthsBetween(2024-01-15, 2024-03-10) // 1.8070  (approximate elapsed)
gotime.YearsBetween(2020-06-01, 2024-05-01)  // 3.9152  (approximate elapsed)
```

This answers "approximately how much time elapsed" — it is **not** a count of
*complete calendar* months/years, and it intentionally uses an averaged month
length. For "how many whole calendar months/years are between two dates"
(age, tenure, billing periods), a different, calendar-accurate function is
required.

### 4.2 The functions

```go
func CalendarMonthsBetween(start, end time.Time) int
func CalendarYearsBetween(start, end time.Time) int
```

Semantics:

- Returns the **signed** count of **complete** calendar units, **truncated
  toward zero**. Negative when `end` precedes `start`.
- Defined as the largest `k ≥ 0` such that `addMonths(start, k) ≤ end`
  (respectively `addYears`). This makes the difference the **inverse of the add
  operation**, so the end-of-month clamp applies consistently.
- No nominal/average constant is used — the result is exact against the
  calendar.

Implementation sketch:

```go
func CalendarMonthsBetween(start, end time.Time) int {
    if start.Equal(end) { return 0 }
    sign, a, b := 1, start, end
    if a.After(b) { sign, a, b = -1, end, start }
    months := (b.Year()-a.Year())*12 + int(b.Month()) - int(a.Month())
    if addMonthsClamped(a, months).After(b) { months-- }
    return sign * months
}
```

### 4.3 Behavior table

| Call | Result | Reason |
|------|--------|--------|
| `CalendarMonthsBetween(2024-01-15, 2024-03-10)` | `1` | second month incomplete |
| `CalendarMonthsBetween(2024-01-15, 2024-03-15)` | `2` | exactly two months |
| `CalendarMonthsBetween(2024-01-31, 2024-02-29)` | `1` | clamped target reached |
| `CalendarMonthsBetween(2024-01-31, 2024-02-28)` | `0` | clamped target not reached |
| `CalendarMonthsBetween(2024-03-10, 2024-01-15)` | `-1` | reversed → negative |
| `CalendarYearsBetween(2020-06-01, 2024-05-01)`  | `3` | fourth year incomplete |
| `CalendarYearsBetween(2024-02-29, 2025-02-28)`  | `1` | clamped anniversary reached |

### 4.4 Edge-case note (clamp consistency)

Because the difference is defined as the inverse of the clamped add, at
month-end clamp points it can differ from a naïve "compare raw day numbers"
approach (such as `java.time`'s packed-day algorithm). Example:
`CalendarYearsBetween(2024-02-29, 2025-02-28)` returns `1` here (the clamped
anniversary `addYears(2024-02-29, 1) = 2025-02-28` is reached), whereas a strict
day-compare would return `0`. The inverse-of-add definition was chosen so that
`addMonths`/`addYears` and the diff functions are mutually consistent.

---

## 5. Which function should I use?

| I want… | Use |
|---------|-----|
| Shift a date by whole months/years (calendar, clamped) | `Months`, `Years` |
| Shift a date by a fixed amount (days/hours/…) | standard `Add` / `AddDate(0,0,n)` |
| **Whole calendar** months/years between two dates (age, tenure) | `CalendarMonthsBetween`, `CalendarYearsBetween` |
| **Approximate elapsed** months/years (a fractional estimate) | `MonthsBetween`, `YearsBetween` |
| A precise (years, months, days) breakdown | `Age` |

Rule of thumb: `Calendar*Between` is the calendar-correct integer answer;
`*Between` (without `Calendar`) is a nominal fractional estimate. They are
different questions — pick deliberately.

---

## 6. Tests

- `TestMonthYearClamp` (in `relative_functions_test.go`) — leap/non-leap
  clamps, spans > 12 months, negative shifts, year leap→non-leap, no-clamp
  cases, zero no-op, and time-of-day/nanosecond preservation.
- `TestCalendarBetween` (in `age_calculation_test.go`) — complete vs incomplete
  units, clamp-edge cases, reversed (negative) ranges, equal dates, multi-year
  spans, for both months and years.
- Existing `TestMonth` / `TestYear` continue to pass unchanged (their day-1
  fixtures are unaffected by clamping).

---

## 7. Compatibility / impact

- **Behavioral change:** code relying on the previous *overflow* behavior of
  `Months`/`Years` (and the four `Last*`/`Next*` helpers) on month-end inputs
  will now see clamped results. This aligns the functions with their documented
  behavior and with other ecosystems, so it is a correctness fix, but it is an
  observable change for those specific inputs.
- `MonthsBetween` / `YearsBetween` are **unchanged**; their nominal-fractional
  semantics are preserved. The new behavior is additive via the `Calendar*`
  functions.

