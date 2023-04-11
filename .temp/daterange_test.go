package temp_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
	"github.com/maniartech/temporal/temp"
)

func TestRange(t *testing.T) {
	ti := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	d := temp.DateRange().Range(ti)
	if d.From != temporal.SoD(ti) {
		t.Error("From date is not today's start date")
	}
	if d.To != temporal.EoD() {
		t.Error("To date is not today's end date")
	}

	ti = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	d = temp.DateRange().Range(ti)
	if d.From != temporal.SoD() {
		t.Error("From date is not today's start date")
	}
	ti = temporal.EoD(ti)
	if d.To != ti {
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.For
func TestDateRangeFor(t *testing.T) {
	d := temp.DateRange()

	if time.Until(d.For) > time.Second*1 { // time.Until returns the duration until t (time.Now() in
		t.Error("For date is not now")
	}

	baseTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	d = temp.DateRange(baseTime)
	if d.For != baseTime {
		t.Error("For date is not base time")
	}
}

// Tests for DateRange.Today
func TestDateRangeToday(t *testing.T) {
	d := temp.DateRange().Today()

	if d.From != temporal.SoD() {
		t.Error("From date is not today's start date")
	}
	if d.To != temporal.EoD() {
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.Yesterday
func TestDateRangeYesterday(t *testing.T) {
	d := temp.DateRange().Yesterday()

	if temporal.DateTime(d.From) != temporal.Yesterday() {
		t.Error("From date is not yesterday's date")
	}
	if d.To != temporal.EoD() {
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.Tomorrow
func TestDateRangeTomorrow(t *testing.T) {
	d := temp.DateRange().Tomorrow()

	if d.From != temporal.SoD() {
		t.Error("From date is not today's start date")
	}
	if temporal.DateTime(d.To) != temporal.Tomorrow() {
		t.Error("To date is not tomorrow's date")
	}
}

// Tests for DateRange.Days
func TestDateRangeDays(t *testing.T) {
	d := temp.DateRange().Days(-2)

	if d.From != temporal.SoD().AddDate(0, 0, -1) { // -1 because the base date is included
		t.Error("From date is not 2 days ago")
	}
	if d.To != temporal.EoD() {
		t.Error("To date is not today's end date")
	}

	d = temp.DateRange().Days(2)

	if d.From != temporal.SoD() {
		t.Error("From date is not 2 days from now")
	}
	if d.To != temporal.EoD().AddDate(0, 0, 1) { // +1 because the base date is included
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.LastWeek
func TestDateRangeLastWeek(t *testing.T) {
	d := temp.DateRange().LastWeek()
	e := temporal.LastWeek()
	for e.Weekday() != time.Sunday {
		e = e.AddDate(0, 0, -1)
	}

	if d.From != e {
		t.Error("From date is not last week's start date")
	}
	if d.To != e.AddDate(0, 0, 6) {
		t.Error("To date is not last week's end date")
	}
}

// Tests for DateRange.ThisWeek
func TestDateRangeThisWeek(t *testing.T) {
	d := temp.DateRange().ThisWeek()
	e := temporal.SoD()
	for e.Weekday() != time.Sunday {
		e = e.AddDate(0, 0, -1)
	}
	if d.From != e {
		t.Error("From date is not this week's start date")
	}
	if d.To != e.AddDate(0, 0, 6) {
		t.Error("To date is not this week's end date")
	}
}

// Tests for DateRange.NextWeek
func TestDateRangeNextWeek(t *testing.T) {
	d := temp.DateRange().NextWeek()
	e := temporal.NextWeek()
	for e.Weekday() != time.Sunday {
		e = e.AddDate(0, 0, -1)
	}
	if d.From != e {
		t.Error("From date is not this week's start date")
	}
	if d.To != e.AddDate(0, 0, 6) {
		t.Error("To date is not this week's end date")
	}
}

// Tests for DateRange.Weeks
func TestDateRangeWeeks(t *testing.T) {
	d := temp.DateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)).Weeks(-2)
	e := time.Date(2018, 12, 18, 0, 0, 0, 0, time.UTC)
	if d.From != e {
		t.Error("From date is not 2 weeks ago")
	}
	e = temporal.EoD(time.Date(2018, 12, 24, 0, 0, 0, 0, time.UTC))
	if d.To != e {
		t.Error("To date is not today's end date")
	}

	d = temp.DateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)).Weeks(0)
	e = time.Date(2018, 12, 30, 0, 0, 0, 0, time.UTC)
	if d.From != e {
		t.Error("From date is not 0 weeks from now")
	}

	e = time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)
	if d.To != e {
		t.Error("To date is not today's end date")
	}

}

// Tests for DateRange.LastMonth
func TestDateRangeLastMonth(t *testing.T) {
	d := temp.DateRange().LastMonth()
	e := temporal.SoD().AddDate(0, -1, 0)
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	// e = e.AddDate(0, 0, -1)
	if d.From != e {
		t.Error("From date is not last month's start date")
	}

	e = e.AddDate(0, 1, 0)
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	e = e.AddDate(0, 0, -1)
	if d.To != e {
		t.Error("To date is not last month's end date")
	}
}

// Tests for DateRange.ThisMonth
func TestDateRangeThisMonth(t *testing.T) {
	d := temp.DateRange().ThisMonth()
	e := temporal.SoD()
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	if d.From != e {
		t.Error("From date is not this month's start date")
	}

	e = e.AddDate(0, 1, 0)
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	e = e.AddDate(0, 0, -1)
	if d.To != e {
		t.Error("To date is not this month's end date")
	}
}

// Tests for DateRange.NextMonth
func TestDateRangeNextMonth(t *testing.T) {
	d := temp.DateRange().NextMonth()
	e := temporal.SoD().AddDate(0, 1, 0)
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	if d.From != e {
		t.Error("From date is not next month's start date")
	}

	e = e.AddDate(0, 1, 0)
	for e.Day() != 1 {
		e = e.AddDate(0, 0, -1)
	}
	e = e.AddDate(0, 0, -1)
	if d.To != e {
		t.Error("To date is not next month's end date")
	}
}

// Tests for DateRange.Months
func TestDateRangeMonths(t *testing.T) {
	d := temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).Months(-1)
	e := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if d.From != e { // -7 because the base date is included
		t.Error("From date is not correct")
	}

	// d = temp.DateRange(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)).Months(2)
	e = time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	if d.To != e {
		t.Error("From date is not correct")
	}

	d = temp.DateRange(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)).Months(0)
	e = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if d.From != e {
		t.Error("From date is not correct")
	}

	e = time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	if d.To != e {
		t.Error("From date is not correct")
	}

}

// Tests for DateRange.LastYear
func TestDateRangeLastYear(t *testing.T) {
	d := temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).LastYear()
	e := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	if d.From != e {
		t.Error("From date is not last year's start date")
	}

	e = e.AddDate(1, 0, -1)
	if d.To != e {
		t.Error("To date is not last year's end date")
	}
}

// Tests for DateRange.ThisYear
func TestDateRangeThisYear(t *testing.T) {
	d := temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).ThisYear()
	e := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	if d.From != e {
		t.Error("From date is not this year's start date")
	}

	e = e.AddDate(1, 0, -1)
	if d.To != e {
		t.Error("To date is not this year's end date")
	}
}

// Tests for DateRange.NextYear
func TestDateRangeNextYear(t *testing.T) {
	d := temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).NextYear()
	e := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	if d.From != e {
		t.Error("From date is not next year's start date")
	}

	e = e.AddDate(1, 0, -1)
	if d.To != e {
		t.Error("To date is not next year's end date")
	}
}

// Tests for DateRange.Years
func TestDateRangeYears(t *testing.T) {
	d := temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).Years(-1)
	e := time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC)
	if d.From != e { // -7 because the base date is included
		t.Error("From date is not 2 years ago")
	}

	d = temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).Years(2)
	e = time.Date(2026, 2, 13, 0, 0, 0, 0, time.UTC)
	if d.To != e {
		t.Error("From date is not 2 years from now")
	}

	d = temp.DateRange(time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC)).Years(0)
	e = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if d.From != e {
		t.Error("From date is not this year's start date")
	}

	e = e.AddDate(1, 0, -1)
	if d.To != e {
		t.Error("To date is not this year's end date")
	}
}
