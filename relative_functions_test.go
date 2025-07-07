package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
	"github.com/maniartech/gotime/internal/utils"
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

	// LastYear
	expectedDate = time.Now().AddDate(-1, 0, 0)
	functionDate = gotime.LastYear()

	utils.AssertEqual(t, expectedDate, functionDate)

	// NextYear
	expectedDate = time.Now().AddDate(1, 0, 0)
	functionDate = gotime.NextYear()

	utils.AssertEqual(t, expectedDate, functionDate)

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

	// LastMonth
	expectedDate = time.Now().AddDate(0, -1, 0)
	functionDate = gotime.LastMonth()

	utils.AssertEqual(t, expectedDate, functionDate)

	// NextMonth
	expectedDate = time.Now().AddDate(0, 1, 0)
	functionDate = gotime.NextMonth()

	utils.AssertEqual(t, expectedDate, functionDate)

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
