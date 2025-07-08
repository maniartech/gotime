package gotime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime"
	"github.com/maniartech/gotime/internal/utils"
)

func TestQuarterStart(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 - January",
			input:    time.Date(2025, 1, 15, 12, 30, 45, 0, time.UTC),
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 - February",
			input:    time.Date(2025, 2, 28, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 - March",
			input:    time.Date(2025, 3, 31, 18, 45, 30, 0, time.UTC),
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - April",
			input:    time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - May",
			input:    time.Date(2025, 5, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 - June",
			input:    time.Date(2025, 6, 30, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - July",
			input:    time.Date(2025, 7, 7, 14, 30, 0, 0, time.UTC),
			expected: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - August",
			input:    time.Date(2025, 8, 15, 9, 15, 30, 0, time.UTC),
			expected: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 - September",
			input:    time.Date(2025, 9, 30, 21, 45, 0, 0, time.UTC),
			expected: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - October",
			input:    time.Date(2025, 10, 1, 6, 30, 0, 0, time.UTC),
			expected: time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - November",
			input:    time.Date(2025, 11, 15, 16, 20, 10, 0, time.UTC),
			expected: time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 - December",
			input:    time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Different timezone",
			input:    time.Date(2025, 5, 15, 12, 0, 0, 0, time.FixedZone("EST", -5*3600)),
			expected: time.Date(2025, 4, 1, 0, 0, 0, 0, time.FixedZone("EST", -5*3600)),
		},
		{
			name:     "Leap year Q1",
			input:    time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gotime.QuarterStart(tt.input)
			utils.AssertEqual(t, tt.expected, result)
		})
	}

	// Test with no arguments (uses current time)
	t.Run("No arguments uses current time", func(t *testing.T) {
		before := time.Now()
		result := gotime.QuarterStart()
		after := time.Now()

		// Verify the result is a valid quarter start
		expectedQuarter := gotime.QuarterOfYear(before)
		resultQuarter := gotime.QuarterOfYear(result)
		utils.AssertEqual(t, expectedQuarter, resultQuarter)

		// Verify it's the start of quarter (day 1, time 00:00:00)
		if result.Day() != 1 || result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
			t.Errorf("QuarterStart() should return start of quarter, got %v", result)
		}

		// Verify it's within reasonable time bounds
		if result.After(after) {
			t.Errorf("QuarterStart() returned future time")
		}
	})
}

func TestQuarterEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 - January (non-leap year)",
			input:    time.Date(2025, 1, 15, 12, 30, 45, 0, time.UTC),
			expected: time.Date(2025, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q1 - February (leap year)",
			input:    time.Date(2024, 2, 15, 10, 20, 30, 0, time.UTC),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q1 - March",
			input:    time.Date(2025, 3, 31, 18, 45, 30, 0, time.UTC),
			expected: time.Date(2025, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - April",
			input:    time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - May",
			input:    time.Date(2025, 5, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q2 - June",
			input:    time.Date(2025, 6, 30, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2025, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - July",
			input:    time.Date(2025, 7, 7, 14, 30, 0, 0, time.UTC),
			expected: time.Date(2025, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - August",
			input:    time.Date(2025, 8, 15, 9, 15, 30, 0, time.UTC),
			expected: time.Date(2025, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q3 - September",
			input:    time.Date(2025, 9, 30, 21, 45, 0, 0, time.UTC),
			expected: time.Date(2025, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - October",
			input:    time.Date(2025, 10, 1, 6, 30, 0, 0, time.UTC),
			expected: time.Date(2025, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - November",
			input:    time.Date(2025, 11, 15, 16, 20, 10, 0, time.UTC),
			expected: time.Date(2025, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Q4 - December",
			input:    time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2025, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "Different timezone",
			input:    time.Date(2025, 5, 15, 12, 0, 0, 0, time.FixedZone("EST", -5*3600)),
			expected: time.Date(2025, 6, 30, 23, 59, 59, 999999999, time.FixedZone("EST", -5*3600)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gotime.QuarterEnd(tt.input)
			utils.AssertEqual(t, tt.expected, result)
		})
	}

	// Test QuarterEnd without arguments (uses current time)
	t.Run("No arguments uses current time", func(t *testing.T) {
		result := gotime.QuarterEnd()
		now := time.Now()

		// Should return a valid quarter end
		if result.Before(now) {
			t.Errorf("QuarterEnd() without arguments should return a time >= now")
		}

		// Should be at end of day (23:59:59.999999999)
		if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
			t.Errorf("QuarterEnd() should return end of day time, got %v", result)
		}
	})
}

func TestQuarters(t *testing.T) {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		quarters int
		baseTime *time.Time
		expected time.Time
	}{
		{
			name:     "Add one quarter",
			quarters: 1,
			baseTime: &baseTime,
			expected: time.Date(2025, 10, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Add two quarters",
			quarters: 2,
			baseTime: &baseTime,
			expected: time.Date(2026, 1, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Subtract one quarter",
			quarters: -1,
			baseTime: &baseTime,
			expected: time.Date(2025, 4, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Subtract two quarters",
			quarters: -2,
			baseTime: &baseTime,
			expected: time.Date(2025, 1, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Zero quarters",
			quarters: 0,
			baseTime: &baseTime,
			expected: baseTime,
		},
		{
			name:     "Add four quarters (one year)",
			quarters: 4,
			baseTime: &baseTime,
			expected: time.Date(2026, 7, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Cross year boundary backwards",
			quarters: -3,
			baseTime: &baseTime,
			expected: time.Date(2024, 10, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Handle end of month edge case",
			quarters: 1,
			baseTime: func() *time.Time { t := time.Date(2025, 5, 31, 12, 0, 0, 0, time.UTC); return &t }(),
			expected: time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "Use current time when no base time provided",
			quarters: 0,
			baseTime: nil,
			expected: time.Time{}, // Will be set in test
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			if tt.baseTime != nil {
				result = gotime.Quarters(tt.quarters, *tt.baseTime)
				utils.AssertEqual(t, tt.expected, result)
			} else {
				// Test with current time
				before := time.Now()
				result = gotime.Quarters(tt.quarters)
				after := time.Now()

				// Result should be close to current time (within test execution time)
				if result.Before(before) || result.After(after.Add(time.Second)) {
					t.Errorf("Quarters() with current time not within expected range")
				}
			}
		})
	}
}

func TestLastQuarter(t *testing.T) {
	// Test that LastQuarter() is equivalent to Quarters(-1)
	expected := gotime.Quarters(-1)
	result := gotime.LastQuarter()

	// Allow for small time differences due to execution time
	diff := result.Sub(expected)
	if diff > time.Second || diff < -time.Second {
		t.Errorf("LastQuarter() = %v, want %v (diff: %v)", result, expected, diff)
	}
}

func TestNextQuarter(t *testing.T) {
	// Test that NextQuarter() is equivalent to Quarters(1)
	expected := gotime.Quarters(1)
	result := gotime.NextQuarter()

	// Allow for small time differences due to execution time
	diff := result.Sub(expected)
	if diff > time.Second || diff < -time.Second {
		t.Errorf("NextQuarter() = %v, want %v (diff: %v)", result, expected, diff)
	}
}

func TestQuarterOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected int
	}{
		{
			name:     "January - Q1",
			input:    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "February - Q1",
			input:    time.Date(2025, 2, 15, 12, 30, 45, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "March - Q1",
			input:    time.Date(2025, 3, 31, 23, 59, 59, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "April - Q2",
			input:    time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "May - Q2",
			input:    time.Date(2025, 5, 15, 12, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "June - Q2",
			input:    time.Date(2025, 6, 30, 23, 59, 59, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "July - Q3",
			input:    time.Date(2025, 7, 7, 14, 30, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "August - Q3",
			input:    time.Date(2025, 8, 15, 9, 15, 30, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "September - Q3",
			input:    time.Date(2025, 9, 30, 21, 45, 0, 0, time.UTC),
			expected: 3,
		},
		{
			name:     "October - Q4",
			input:    time.Date(2025, 10, 1, 6, 30, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "November - Q4",
			input:    time.Date(2025, 11, 15, 16, 20, 10, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "December - Q4",
			input:    time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gotime.QuarterOfYear(tt.input)
			utils.AssertEqual(t, tt.expected, result)
		})
	}

	// Test with no arguments (uses current time)
	t.Run("No arguments uses current time", func(t *testing.T) {
		now := time.Now()
		result := gotime.QuarterOfYear()
		expected := (int(now.Month())-1)/3 + 1

		// Allow for edge case where time might cross quarter boundary during test
		if result != expected {
			// Check if we're at quarter boundary
			nowQuarter := (int(now.Month())-1)/3 + 1
			if result == nowQuarter {
				return // Valid result due to timing
			}
			t.Errorf("QuarterOfYear() = %d, want %d", result, expected)
		}
	})
}

// Benchmark tests
func BenchmarkQuarterStart(b *testing.B) {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.QuarterStart(baseTime)
	}
}

func BenchmarkQuarterEnd(b *testing.B) {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.QuarterEnd(baseTime)
	}
}

func BenchmarkQuarters(b *testing.B) {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.Quarters(2, baseTime)
	}
}

func BenchmarkQuarterOfYear(b *testing.B) {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.QuarterOfYear(baseTime)
	}
}

// Example tests for documentation
func ExampleQuarterStart() {
	// Get the start of Q3 2025
	t := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	start := gotime.QuarterStart(t)
	fmt.Println(start)
	// Output: 2025-07-01 00:00:00 +0000 UTC
}

func ExampleQuarterEnd() {
	// Get the end of Q2 2025
	t := time.Date(2025, 5, 15, 12, 30, 45, 0, time.UTC)
	end := gotime.QuarterEnd(t)
	fmt.Println(end)
	// Output: 2025-06-30 23:59:59.999999999 +0000 UTC
}

func ExampleQuarters() {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)

	// Add 2 quarters (6 months)
	future := gotime.Quarters(2, baseTime)
	fmt.Println(future)
	// Output: 2026-01-15 12:30:45 +0000 UTC
}

func ExampleQuarters_negative() {
	baseTime := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)

	// Subtract 1 quarter (3 months)
	past := gotime.Quarters(-1, baseTime)
	fmt.Println(past)
	// Output: 2025-04-15 12:30:45 +0000 UTC
}

func ExampleQuarterOfYear() {
	// July is in Q3
	t := time.Date(2025, 7, 15, 12, 30, 45, 0, time.UTC)
	quarter := gotime.QuarterOfYear(t)
	fmt.Println(quarter)
	// Output: 3
}
