package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
)

// Tests for DateRange.For
func TestDateRangeFor(t *testing.T) {
	d := temporal.DateRange()

	if time.Until(d.For) > time.Second*1 { // time.Until returns the duration until t (time.Now() in
		t.Error("For date is not now")
	}

	baseTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	d = temporal.DateRange(baseTime)
	if d.For != baseTime {
		t.Error("For date is not base time")
	}
}

// Tests for DateRange.Today
func TestDateRangeToday(t *testing.T) {
	d := temporal.DateRange().Today()

	if d.From != temporal.DayStart() {
		t.Error("From date is not today's start date")
	}
	if d.To != temporal.DayEnd() {
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.Yesterday
func TestDateRangeYesterday(t *testing.T) {
	d := temporal.DateRange().Yesterday()

	if d.From != temporal.Yesterday() {
		t.Error("From date is not yesterday's date")
	}
	if d.To != temporal.DayEnd() {
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.Tomorrow
func TestDateRangeTomorrow(t *testing.T) {
	d := temporal.DateRange().Tomorrow()

	if d.From != temporal.DayStart() {
		t.Error("From date is not today's start date")
	}
	if d.To != temporal.Tomorrow() {
		t.Error("To date is not tomorrow's date")
	}
}

// Tests for DateRange.Days
func TestDateRangeDays(t *testing.T) {
	d := temporal.DateRange().Days(-2)

	if d.From != temporal.DayStart().AddDate(0, 0, -1) { // -1 because the base date is included
		t.Error("From date is not 2 days ago")
	}
	if d.To != temporal.DayEnd() {
		t.Error("To date is not today's end date")
	}

	d = temporal.DateRange().Days(2)

	if d.From != temporal.DayStart() {
		t.Error("From date is not 2 days from now")
	}
	if d.To != temporal.DayEnd().AddDate(0, 0, 1) { // +1 because the base date is included
		t.Error("To date is not today's end date")
	}
}

// Tests for DateRange.LastWeek
func TestDateRangeLastWeek(t *testing.T) {
	d := temporal.DateRange().LastWeek()
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
	d := temporal.DateRange().ThisWeek()
	e := temporal.DayStart()
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
	d := temporal.DateRange().NextWeek()
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
