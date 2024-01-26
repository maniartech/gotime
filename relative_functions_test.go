package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	// YearStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := temporal.YearStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	functionDate = temporal.YearStart()
	assert.Equal(t, expectedDate, functionDate)

	// YearEnd
	expectedDate = time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.YearEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.YearEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastYear
	expectedDate = time.Now().AddDate(-1, 0, 0)
	functionDate = temporal.LastYear()

	assert.Equal(t, expectedDate, functionDate)

	// NextYear
	expectedDate = time.Now().AddDate(1, 0, 0)
	functionDate = temporal.NextYear()

	assert.Equal(t, expectedDate, functionDate)

	// Years
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, year := range testRange {
		expectedDate := fixedDate.AddDate(year, 0, 0)
		functionDate := temporal.Years(year, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateMillisecond(now.AddDate(1, 0, 0))
	functionDate = truccateMillisecond(temporal.Years(1))
	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Years(0, fixedDate) })
}

func TestMonth(t *testing.T) {
	// MonthStart - write a test that loops through all the months and checks that the first day of the month is correct
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	totalDaysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		functionDate := temporal.MonthStart(time.Date(2023, time.Month(month), 1, 11, 2, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}
	now := time.Now()
	expectedDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	functionDate := temporal.MonthStart()
	assert.Equal(t, expectedDate, functionDate)
	// MonthEnd
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), totalDaysInMonth[month-1], 23, 59, 59, 999999999, time.UTC)
		functionDate := temporal.MonthEnd(time.Date(2023, time.Month(month), 1, 11, 22, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}
	expectedDate = time.Date(now.Year(), now.Month(), totalDaysInMonth[now.Month()-1], 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.MonthEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastMonth
	expectedDate = time.Now().AddDate(0, -1, 0)
	functionDate = temporal.LastMonth()

	assert.Equal(t, expectedDate, functionDate)

	// NextMonth
	expectedDate = time.Now().AddDate(0, 1, 0)
	functionDate = temporal.NextMonth()

	assert.Equal(t, expectedDate, functionDate)

	// Months
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, month := range testRange {
		expectedDate := fixedDate.AddDate(0, month, 0)
		functionDate := temporal.Months(month, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateMillisecond(now.AddDate(0, 1, 0))
	functionDate = truccateMillisecond(temporal.Months(1))
	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Months(0, fixedDate) })

}

func truccateMillisecond(t time.Time) time.Time {
	return t.Truncate(time.Millisecond)
}

func TestWeek(t *testing.T) {
	// WeekStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := temporal.WeekStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = truccateMillisecond(time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday()), 0, 0, 0, 0, time.Local))
	functionDate = truccateMillisecond(temporal.WeekStart())
	assert.Equal(t, expectedDate, functionDate)

	// WeekEnd
	expectedDate = time.Date(2023, 1, 7, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.WeekEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+6, 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.WeekEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastWeek
	expectedDate = time.Now().AddDate(0, 0, -7)
	functionDate = temporal.LastWeek()

	assert.Equal(t, expectedDate, functionDate)

	// NextWeek
	expectedDate = time.Now().AddDate(0, 0, 7)
	functionDate = temporal.NextWeek()

	assert.Equal(t, expectedDate, functionDate)

	// Weeks
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, week := range testRange {
		expectedDate := fixedDate.AddDate(0, 0, week*7)
		functionDate := temporal.Weeks(week, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateMillisecond(now.AddDate(0, 0, 7))
	functionDate = truccateMillisecond(temporal.Weeks(1))

	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Weeks(0, fixedDate) })

}
