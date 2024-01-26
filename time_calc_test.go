package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	twoHoursAgo := now.Add(-2 * time.Hour)
	fiveMinutesAgo := now.Add(-5 * time.Minute)

	// Test different units
	t.Run("Test difference in hours", func(t *testing.T) {
		result := temporal.Diff(now, yesterday, time.Hour)
		expected := 24.0 // Adjusted for potential time zone differences
		if result != expected {
			t.Errorf("Expected difference in hours: %f, got: %f", expected, result)
		}
	})

	t.Run("Test difference in minutes", func(t *testing.T) {
		result := temporal.Diff(now, twoHoursAgo, time.Minute)
		expected := 120.0
		if result != expected {
			t.Errorf("Expected difference in minutes: %f, got: %f", expected, result)
		}
	})

	// Test rounding
	t.Run("Test rounding to nearest second", func(t *testing.T) {
		result := temporal.Diff(now, fiveMinutesAgo, time.Second, true)
		expected := 300.0
		if result != expected {
			t.Errorf("Expected rounded difference in seconds: %f, got: %f", expected, result)
		}
	})

	// Test zero difference
	t.Run("Test zero difference", func(t *testing.T) {
		result := temporal.Diff(now, now, time.Second)
		expected := 0.0
		if result != expected {
			t.Errorf("Expected zero difference: %f, got: %f", expected, result)
		}
	})

	// Test negative difference
	t.Run("Test negative difference", func(t *testing.T) {
		result := temporal.Diff(yesterday, now, time.Hour)
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
	result := trunccateSecond(temporal.Latest(now, yesterday, tomorrow))
	expected := trunccateSecond(tomorrow)
	assert.Equal(t, expected, result)

	// Test with a single time
	result = trunccateSecond(temporal.Latest(now))
	expected = trunccateSecond(now)
	assert.Equal(t, expected, result)

	// Test panic for empty input
	defer func() {
		if r := recover(); r != "No time given" {
			t.Errorf("Expected panic for empty input, got: %v", r)
		}
	}()
	temporal.Latest()
}

func TestEarliest(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test with multiple times
	result := trunccateSecond(temporal.Earliest(now, yesterday, tomorrow))
	expected := trunccateSecond(yesterday)
	assert.Equal(t, expected, result)

	// Test with a single time
	result = trunccateSecond(temporal.Earliest(now))
	expected = trunccateSecond(now)
	assert.Equal(t, expected, result)

	// Test panic for empty input
	defer func() {
		if r := recover(); r != "No time given" {
			t.Errorf("Expected panic for empty input, got: %v", r)
		}
	}()
	temporal.Earliest()
}
