package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
	"github.com/maniartech/gotime/v2/internal/utils"
)

func TestYear(t *testing.T) {
	// YearStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := gotime.YearStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	functionDate = gotime.YearStart()
	utils.AssertEqual(t, expectedDate, functionDate)

	// YearEnd
	expectedDate = time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)
	functionDate = gotime.YearEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, time.Local)
	functionDate = gotime.YearEnd()
	utils.AssertEqual(t, expectedDate, functionDate)

	// LastYear / NextYear delegate to the (clamped) Years function.
	utils.AssertEqual(t, trunccateSecond(gotime.Years(-1)), trunccateSecond(gotime.LastYear()))
	utils.AssertEqual(t, trunccateSecond(gotime.Years(1)), trunccateSecond(gotime.NextYear()))

	// Years
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, year := range testRange {
		expectedDate := fixedDate.AddDate(year, 0, 0)
		functionDate := gotime.Years(year, fixedDate)

		utils.AssertEqual(t, expectedDate, functionDate)
	}

	expectedDate = trunccateSecond(now.AddDate(1, 0, 0))
	functionDate = trunccateSecond(gotime.Years(1))
	utils.AssertEqual(t, expectedDate, functionDate)

	// Zero is a no-op, should return the same date
	expectedDate = fixedDate
	functionDate = gotime.Years(0, fixedDate)
	utils.AssertEqual(t, expectedDate, functionDate)
}

func TestMonth(t *testing.T) {
	// MonthStart - write a test that loops through all the months and checks that the first day of the month is correct
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	totalDaysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		functionDate := gotime.MonthStart(time.Date(2023, time.Month(month), 1, 11, 2, 10, 0, time.UTC))

		utils.AssertEqual(t, expectedDate, functionDate)
	}
	now := time.Now()
	expectedDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	functionDate := gotime.MonthStart()
	utils.AssertEqual(t, expectedDate, functionDate)
	// MonthEnd
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), totalDaysInMonth[month-1], 23, 59, 59, 999999999, time.UTC)
		functionDate := gotime.MonthEnd(time.Date(2023, time.Month(month), 1, 11, 22, 10, 0, time.UTC))

		utils.AssertEqual(t, expectedDate, functionDate)
	}

	expectedDate = time.Date(now.Year(), now.Month(), gotime.DaysInMonth(now.Year(), int(now.Month())), 23, 59, 59, 999999999, time.Local)
	functionDate = gotime.MonthEnd()
	utils.AssertEqual(t, expectedDate, functionDate)

	// LastMonth / NextMonth delegate to the (clamped) Months function.
	utils.AssertEqual(t, trunccateSecond(gotime.Months(-1)), trunccateSecond(gotime.LastMonth()))
	utils.AssertEqual(t, trunccateSecond(gotime.Months(1)), trunccateSecond(gotime.NextMonth()))

	// Months
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, month := range testRange {
		expectedDate := fixedDate.AddDate(0, month, 0)
		functionDate := gotime.Months(month, fixedDate)

		utils.AssertEqual(t, expectedDate, functionDate)
	}

	expectedDate = trunccateSecond(now.AddDate(0, 1, 0))
	functionDate = trunccateSecond(gotime.Months(1))
	utils.AssertEqual(t, expectedDate, functionDate)

	// Zero is a no-op, should return the same date
	expectedDate = fixedDate
	functionDate = gotime.Months(0, fixedDate)
	utils.AssertEqual(t, expectedDate, functionDate)

}

func trunccateSecond(t time.Time) time.Time {
	return t.Truncate(time.Second)
}

// TestMonthYearClamp verifies that Months and Years apply calendar-aware
// end-of-month clamping instead of the standard library's overflow
// normalization. For example, Jan 31 + 1 month must yield Feb 28/29, not
// Mar 2/3. This matches .NET AddMonths, java.time plusMonths, and SQL DATEADD.
func TestMonthYearClamp(t *testing.T) {
	d := func(y, m, day int) time.Time {
		return time.Date(y, time.Month(m), day, 0, 0, 0, 0, time.UTC)
	}

	monthCases := []struct {
		in   time.Time
		n    int
		want time.Time
	}{
		{d(2020, 1, 31), 1, d(2020, 2, 29)},  // leap year -> Feb 29
		{d(2023, 1, 31), 1, d(2023, 2, 28)},  // non-leap -> Feb 28
		{d(2024, 3, 31), 1, d(2024, 4, 30)},  // 31 -> 30-day month
		{d(2024, 5, 31), 1, d(2024, 6, 30)},  // 31 -> 30-day month
		{d(2024, 1, 31), 13, d(2025, 2, 28)}, // span >12 months
		{d(2024, 3, 31), -1, d(2024, 2, 29)}, // negative shift clamps too
		{d(2024, 1, 15), 1, d(2024, 2, 15)},  // no clamp needed
		{d(2024, 1, 31), 0, d(2024, 1, 31)},  // zero is a no-op
	}
	for _, c := range monthCases {
		got := gotime.Months(c.n, c.in)
		utils.AssertEqual(t, c.want, got)
	}

	yearCases := []struct {
		in   time.Time
		n    int
		want time.Time
	}{
		{d(2024, 2, 29), 1, d(2025, 2, 28)},  // leap -> non-leap clamps
		{d(2024, 2, 29), 4, d(2028, 2, 29)},  // leap -> leap stays
		{d(2024, 2, 29), -1, d(2023, 2, 28)}, // negative shift clamps
		{d(2024, 6, 15), 2, d(2026, 6, 15)},  // no clamp needed
	}
	for _, c := range yearCases {
		got := gotime.Years(c.n, c.in)
		utils.AssertEqual(t, c.want, got)
	}

	// Time-of-day and sub-second components must be preserved.
	withTime := time.Date(2024, 1, 31, 13, 45, 30, 123456789, time.UTC)
	got := gotime.Months(1, withTime)
	want := time.Date(2024, 2, 29, 13, 45, 30, 123456789, time.UTC)
	utils.AssertEqual(t, want, got)
}

func TestWeek(t *testing.T) {
	// WeekStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := gotime.WeekStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = trunccateSecond(time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday()), 0, 0, 0, 0, time.Local))
	functionDate = trunccateSecond(gotime.WeekStart())
	utils.AssertEqual(t, expectedDate, functionDate)

	// WeekEnd
	expectedDate = time.Date(2023, 1, 7, 23, 59, 59, 999999999, time.UTC)
	functionDate = gotime.WeekEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+6, 23, 59, 59, 999999999, time.Local)
	functionDate = gotime.WeekEnd()
	utils.AssertEqual(t, expectedDate, functionDate)

	// LastWeek
	expectedDate = time.Now().AddDate(0, 0, -7)
	functionDate = gotime.LastWeek()

	utils.AssertEqual(t, expectedDate, functionDate)

	// NextWeek
	expectedDate = time.Now().AddDate(0, 0, 7)
	functionDate = gotime.NextWeek()

	utils.AssertEqual(t, expectedDate, functionDate)

	// Weeks
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, week := range testRange {
		expectedDate := fixedDate.AddDate(0, 0, week*7)
		functionDate := gotime.Weeks(week, fixedDate)

		utils.AssertEqual(t, expectedDate, functionDate)
	}

	expectedDate = trunccateSecond(now.AddDate(0, 0, 7))
	functionDate = trunccateSecond(gotime.Weeks(1))

	utils.AssertEqual(t, expectedDate, functionDate)

	// Zero is a no-op, should return the same date
	expectedDate = fixedDate
	functionDate = gotime.Weeks(0, fixedDate)
	utils.AssertEqual(t, expectedDate, functionDate)

	// WeekStartOn
	expectedDate = time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	functionDate = gotime.WeekStartOn(time.Sunday, time.Date(2024, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate = gotime.WeekStartOn(time.Monday, time.Date(2024, 1, 2, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	functionDate = gotime.WeekStartOn(time.Tuesday, time.Date(2024, 1, 3, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	now = time.Now()

	expectedDate = trunccateSecond(time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+1, 0, 0, 0, 0, time.Local))
	functionDate = trunccateSecond(gotime.WeekStartOn(time.Monday))

	utils.AssertEqual(t, expectedDate, functionDate)

	// WeekEndOn
	expectedDate = time.Date(2024, 1, 6, 23, 59, 59, 999999999, time.UTC)
	functionDate = gotime.WeekEndOn(time.Sunday, time.Date(2024, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 7, 23, 59, 59, 999999999, time.UTC)
	functionDate = gotime.WeekEndOn(time.Monday, time.Date(2024, 1, 2, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	now = time.Now()

	expectedDate = time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+7, 23, 59, 59, 999999999, time.Local)
	functionDate = gotime.WeekEndOn(time.Monday)

	utils.AssertEqual(t, expectedDate, functionDate)
}

func TestDay(t *testing.T) {
	// SoD
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := gotime.SoD(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	functionDate = gotime.SoD()
	utils.AssertEqual(t, expectedDate, functionDate)

	// EoD
	expectedDate = time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC)
	functionDate = gotime.EoD(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, time.Local)
	functionDate = gotime.EoD()
	utils.AssertEqual(t, expectedDate, functionDate)

	// Yesterday
	expectedDate = time.Now().AddDate(0, 0, -1)
	functionDate = gotime.Yesterday()

	utils.AssertEqual(t, expectedDate, functionDate)

	// Tomorrow
	expectedDate = time.Now().AddDate(0, 0, 1)
	functionDate = gotime.Tomorrow()

	utils.AssertEqual(t, expectedDate, functionDate)

	// Days
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, day := range testRange {
		expectedDate := fixedDate.AddDate(0, 0, day)
		functionDate := gotime.Days(day, fixedDate)

		utils.AssertEqual(t, expectedDate, functionDate)
	}

	expectedDate = trunccateSecond(now.AddDate(0, 0, 1))
	functionDate = trunccateSecond(gotime.Days(1))

	utils.AssertEqual(t, expectedDate, functionDate)

	// Zero is a no-op, should return the same date
	expectedDate = fixedDate
	functionDate = gotime.Days(0, fixedDate)
	utils.AssertEqual(t, expectedDate, functionDate)

}
