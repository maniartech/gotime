package gotime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime"
	"github.com/maniartech/gotime/internal/utils"
)

func TestHours(t *testing.T) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		hours    int
		baseTime *time.Time
		expected time.Time
	}{
		{
			name:     "Add positive hours",
			hours:    5,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 17, 30, 45, 0, time.UTC),
		},
		{
			name:     "Add negative hours",
			hours:    -3,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 9, 30, 45, 0, time.UTC),
		},
		{
			name:     "Add zero hours",
			hours:    0,
			baseTime: &baseTime,
			expected: baseTime,
		},
		{
			name:     "Add hours crossing day boundary",
			hours:    15,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 8, 3, 30, 45, 0, time.UTC),
		},
		{
			name:     "Subtract hours crossing day boundary",
			hours:    -15,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 6, 21, 30, 45, 0, time.UTC),
		},
		{
			name:     "Add 24 hours (exactly one day)",
			hours:    24,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 8, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Use current time when no base time provided",
			hours:    0,
			baseTime: nil,
			expected: time.Time{}, // Will be set in test
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			if tt.baseTime != nil {
				result = gotime.Hours(tt.hours, *tt.baseTime)
				utils.AssertEqual(t, tt.expected, result)
			} else {
				// Test with current time
				before := time.Now()
				result = gotime.Hours(tt.hours)
				after := time.Now()

				// Result should be close to current time (within test execution time)
				if result.Before(before) || result.After(after.Add(time.Second)) {
					t.Errorf("Hours() with current time not within expected range")
				}
			}
		})
	}
}

func TestMinutes(t *testing.T) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		minutes  int
		baseTime *time.Time
		expected time.Time
	}{
		{
			name:     "Add positive minutes",
			minutes:  30,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 13, 0, 45, 0, time.UTC),
		},
		{
			name:     "Add negative minutes",
			minutes:  -15,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 12, 15, 45, 0, time.UTC),
		},
		{
			name:     "Add zero minutes",
			minutes:  0,
			baseTime: &baseTime,
			expected: baseTime,
		},
		{
			name:     "Add minutes crossing hour boundary",
			minutes:  45,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 13, 15, 45, 0, time.UTC),
		},
		{
			name:     "Subtract minutes crossing hour boundary",
			minutes:  -45,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 11, 45, 45, 0, time.UTC),
		},
		{
			name:     "Add 1440 minutes (exactly one day)",
			minutes:  1440,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 8, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Use current time when no base time provided",
			minutes:  0,
			baseTime: nil,
			expected: time.Time{}, // Will be set in test
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			if tt.baseTime != nil {
				result = gotime.Minutes(tt.minutes, *tt.baseTime)
				utils.AssertEqual(t, tt.expected, result)
			} else {
				// Test with current time
				before := time.Now()
				result = gotime.Minutes(tt.minutes)
				after := time.Now()

				// Result should be close to current time (within test execution time)
				if result.Before(before) || result.After(after.Add(time.Second)) {
					t.Errorf("Minutes() with current time not within expected range")
				}
			}
		})
	}
}

func TestSeconds(t *testing.T) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		seconds  int
		baseTime *time.Time
		expected time.Time
	}{
		{
			name:     "Add positive seconds",
			seconds:  30,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 12, 31, 15, 0, time.UTC),
		},
		{
			name:     "Add negative seconds",
			seconds:  -15,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 12, 30, 30, 0, time.UTC),
		},
		{
			name:     "Add zero seconds",
			seconds:  0,
			baseTime: &baseTime,
			expected: baseTime,
		},
		{
			name:     "Add seconds crossing minute boundary",
			seconds:  30,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 12, 31, 15, 0, time.UTC),
		},
		{
			name:     "Subtract seconds crossing minute boundary",
			seconds:  -60,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 7, 12, 29, 45, 0, time.UTC),
		},
		{
			name:     "Add 86400 seconds (exactly one day)",
			seconds:  86400,
			baseTime: &baseTime,
			expected: time.Date(2025, 7, 8, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "Use current time when no base time provided",
			seconds:  0,
			baseTime: nil,
			expected: time.Time{}, // Will be set in test
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			if tt.baseTime != nil {
				result = gotime.Seconds(tt.seconds, *tt.baseTime)
				utils.AssertEqual(t, tt.expected, result)
			} else {
				// Test with current time
				before := time.Now()
				result = gotime.Seconds(tt.seconds)
				after := time.Now()

				// Result should be close to current time (within test execution time)
				if result.Before(before) || result.After(after.Add(time.Second)) {
					t.Errorf("Seconds() with current time not within expected range")
				}
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkHours(b *testing.B) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.Hours(5, baseTime)
	}
}

func BenchmarkMinutes(b *testing.B) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.Minutes(30, baseTime)
	}
}

func BenchmarkSeconds(b *testing.B) {
	baseTime := time.Date(2025, 7, 7, 12, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		gotime.Seconds(45, baseTime)
	}
}

// Example tests for documentation
func ExampleHours() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Add 5 hours
	result := gotime.Hours(5, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 17:00:00 +0000 UTC
}

func ExampleHours_negative() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Subtract 2 hours
	result := gotime.Hours(-2, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 10:00:00 +0000 UTC
}

func ExampleMinutes() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Add 30 minutes
	result := gotime.Minutes(30, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 12:30:00 +0000 UTC
}

func ExampleMinutes_negative() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Subtract 15 minutes
	result := gotime.Minutes(-15, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 11:45:00 +0000 UTC
}

func ExampleSeconds() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Add 45 seconds
	result := gotime.Seconds(45, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 12:00:45 +0000 UTC
}

func ExampleSeconds_negative() {
	baseTime := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)

	// Subtract 30 seconds
	result := gotime.Seconds(-30, baseTime)
	fmt.Println(result)
	// Output: 2025-07-07 11:59:30 +0000 UTC
}
