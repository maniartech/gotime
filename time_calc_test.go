package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
	"github.com/maniartech/gotime/v2/internal/utils"
)

func TestDiff(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	twoHoursAgo := now.Add(-2 * time.Hour)
	fiveMinutesAgo := now.Add(-5 * time.Minute)

	// Test different units
	t.Run("Test difference in hours", func(t *testing.T) {
		result := gotime.Diff(now, yesterday, time.Hour)
		expected := 24.0 // Adjusted for potential time zone differences
		if result != expected {
			t.Errorf("Expected difference in hours: %f, got: %f", expected, result)
		}
	})

	t.Run("Test difference in minutes", func(t *testing.T) {
		result := gotime.Diff(now, twoHoursAgo, time.Minute)
		expected := 120.0
		if result != expected {
			t.Errorf("Expected difference in minutes: %f, got: %f", expected, result)
		}
	})

	// Test rounding
	t.Run("Test rounding to nearest second", func(t *testing.T) {
		result := gotime.Diff(now, fiveMinutesAgo, time.Second, true)
		expected := 300.0
		if result != expected {
			t.Errorf("Expected rounded difference in seconds: %f, got: %f", expected, result)
		}
	})

	// Test zero difference
	t.Run("Test zero difference", func(t *testing.T) {
		result := gotime.Diff(now, now, time.Second)
		expected := 0.0
		if result != expected {
			t.Errorf("Expected zero difference: %f, got: %f", expected, result)
		}
	})

	// Test negative difference
	t.Run("Test negative difference", func(t *testing.T) {
		result := gotime.Diff(yesterday, now, time.Hour)
		expected := -24.0
		if result != expected {
			t.Errorf("Expected negative difference in hours: %f, got: %f", expected, result)
		}
	})
}

func TestLatest(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test with multiple times
	result := trunccateSecond(gotime.Latest(now, yesterday, tomorrow))
	expected := trunccateSecond(tomorrow)
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = trunccateSecond(gotime.Latest(now, yesterday, tomorrow))
	expected = trunccateSecond(tomorrow)
	utils.AssertEqual(t, expected, result)
}

func TestEarliest(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test with multiple times
	result := trunccateSecond(gotime.Earliest(now, yesterday, tomorrow))
	expected := trunccateSecond(yesterday)
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = trunccateSecond(gotime.Earliest(now, tomorrow, yesterday))
	expected = trunccateSecond(yesterday)
	utils.AssertEqual(t, expected, result)

}

func TestTruncateTime(t *testing.T) {
	now := time.Now()
	expected := trunccateSecond(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	result := trunccateSecond(gotime.TruncateTime(now))
	utils.AssertEqual(t, expected, result)
}

func TestWorkDay(t *testing.T) {
	startDay := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	workingDays := [7]bool{false, true, true, true, true, true, false}
	holidays := []time.Time{
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC),
	}
	days := 7

	expectedDate := time.Date(2024, 1, 9, 0, 0, 0, 0, time.UTC)
	functionDate, err := gotime.WorkDay(startDay, days, workingDays)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 11, 0, 0, 0, 0, time.UTC)
	functionDate, err = gotime.WorkDay(startDay, days, workingDays, holidays...)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDate, functionDate)

	// Test error cases
	// Test negative days
	_, err = gotime.WorkDay(startDay, -1, workingDays)
	if err == nil {
		t.Error("Expected error for negative days")
	}

	// Test zero date
	_, err = gotime.WorkDay(time.Time{}, days, workingDays)
	if err == nil {
		t.Error("Expected error for zero date")
	}

	// Test no working days
	noWorkingDays := [7]bool{false, false, false, false, false, false, false}
	_, err = gotime.WorkDay(startDay, days, noWorkingDays)
	if err == nil {
		t.Error("Expected error for no working days")
	}
}

func TestWorkDayWithUnsortedHolidays(t *testing.T) {
	startDay := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	workingDays := [7]bool{false, true, true, true, true, true, false}

	// Deliberately unsorted holidays
	holidays := []time.Time{
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
	}

	days := 5

	expectedDate := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	functionDate, err := gotime.WorkDay(startDay, days, workingDays, holidays...)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDate, functionDate)
}

func TestPrevWorkDay(t *testing.T) {
	// Define the working days (Monday to Friday)
	workingDays := [7]bool{false, true, true, true, true, true, false}

	// Define some holidays
	holidays := []time.Time{
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),   // New Year's Day
		time.Date(2022, time.December, 25, 0, 0, 0, 0, time.UTC), // Christmas Day
	}

	// Define some test cases
	testCases := []struct {
		name       string
		startDate  time.Time
		days       int
		expectDate time.Time
	}{
		{
			name:       "Subtract 1 working day, no holiday",
			startDate:  time.Date(2022, time.January, 5, 0, 0, 0, 0, time.UTC), // Wednesday
			days:       1,
			expectDate: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC), // Tuesday
		},
		{
			name:       "Subtract 1 working day, with holiday",
			startDate:  time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC), // Monday
			days:       1,
			expectDate: time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC), // Friday
		},
		{
			name:       "Subtract 5 working days, with holiday",
			startDate:  time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC), // Monday
			days:       5,
			expectDate: time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC), // Monday
		},
		{
			name:       "Subtract 20 working days, with holiday",
			startDate:  time.Date(2022, time.January, 31, 0, 0, 0, 0, time.UTC), // Monday
			days:       20,
			expectDate: time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC), // Monday
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotDate, err := gotime.PrevWorkDay(tc.startDate, tc.days, workingDays, holidays...)
			utils.AssertNoError(t, err)
			if !gotDate.Equal(tc.expectDate) {
				t.Errorf("got %v, want %v", gotDate, tc.expectDate)
			}
		})
	}

	// Test error cases
	// Test negative days
	_, err := gotime.PrevWorkDay(time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC), -1, workingDays)
	if err == nil {
		t.Error("Expected error for negative days")
	}

	// Test zero date
	_, err = gotime.PrevWorkDay(time.Time{}, 1, workingDays)
	if err == nil {
		t.Error("Expected error for zero date")
	}

	// Test no working days
	noWorkingDays := [7]bool{false, false, false, false, false, false, false}
	_, err = gotime.PrevWorkDay(time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC), 1, noWorkingDays)
	if err == nil {
		t.Error("Expected error for no working days")
	}
}

func TestPrevWorkDayWithUnsortedHolidays(t *testing.T) {
	startDay := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	workingDays := [7]bool{false, true, true, true, true, true, false}

	// Deliberately unsorted holidays
	holidays := []time.Time{
		time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
	}

	days := 5

	expectedDate := time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC)
	functionDate, err := gotime.PrevWorkDay(startDay, days, workingDays, holidays...)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDate, functionDate)
}

func TestNetWorkdays(t *testing.T) {
	startDay := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDay := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	workingDays := [7]bool{false, true, true, true, true, true, false}
	holidays := []time.Time{
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC),
	}

	expectedDays := 8
	functionDays, err := gotime.NetWorkDays(endDay, startDay, workingDays)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDays, functionDays)

	expectedDays = 6
	functionDays, err = gotime.NetWorkDays(startDay, endDay, workingDays, holidays...)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedDays, functionDays)

	// Test error cases
	// Test zero start date
	_, err = gotime.NetWorkDays(time.Time{}, endDay, workingDays)
	if err == nil {
		t.Error("Expected error for zero start date")
	}

	// Test zero end date
	_, err = gotime.NetWorkDays(startDay, time.Time{}, workingDays)
	if err == nil {
		t.Error("Expected error for zero end date")
	}

	// Test no working days
	noWorkingDays := [7]bool{false, false, false, false, false, false, false}
	_, err = gotime.NetWorkDays(startDay, endDay, noWorkingDays)
	if err == nil {
		t.Error("Expected error for no working days")
	}
}

func TestDateValue(t *testing.T) {
	testCases := []struct {
		date     time.Time
		expected int
	}{
		{time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(1900, 1, 2, 0, 0, 0, 0, time.UTC), 3},
		{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 45252},
	}

	for _, tc := range testCases {
		t.Run(tc.date.String(), func(t *testing.T) {
			result := gotime.DateValue(tc.date)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
