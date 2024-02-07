package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
	"github.com/maniartech/gotime/internal/utils"
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

func TestIsBetween(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test with multiple times
	result := gotime.IsBetween(now, yesterday, tomorrow)
	expected := true
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetween(now, tomorrow, tomorrow)
	expected = false
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetween(now, now, now)
	expected = true
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetween(now, tomorrow, yesterday)
	expected = true
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

	expectedDate := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	functionDate := gotime.WorkDay(startDay, days, workingDays)
	utils.AssertEqual(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 12, 0, 0, 0, 0, time.UTC)
	functionDate = gotime.WorkDay(startDay, days, workingDays, holidays...)
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
	functionDays := gotime.NetWorkDays(startDay, endDay, workingDays)
	utils.AssertEqual(t, expectedDays, functionDays)

	expectedDays = 6
	functionDays = gotime.NetWorkDays(startDay, endDay, workingDays, holidays...)
	utils.AssertEqual(t, expectedDays, functionDays)
}
