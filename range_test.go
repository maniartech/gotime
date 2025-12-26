package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
	"github.com/maniartech/gotime/v2/internal/utils"
)

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

func TestIsBetweenDates(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test with multiple times
	result := gotime.IsBetweenDates(now, yesterday, tomorrow)
	expected := true
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetweenDates(now, tomorrow, tomorrow)
	expected = false
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetweenDates(now, now, now)
	expected = true
	utils.AssertEqual(t, expected, result)

	// Test with a single time
	result = gotime.IsBetweenDates(now, tomorrow, yesterday)
	expected = true
	utils.AssertEqual(t, expected, result)
}
