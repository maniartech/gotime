package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

func BenchmarkTimeAgo(b *testing.B) {
	// Benchmark for TimeAgo() - use fixed date to avoid time-dependent behavior
	baseTime := time.Date(2025, 7, 8, 12, 0, 0, 0, time.UTC)
	date := baseTime.AddDate(0, 0, 10)
	for i := 0; i < b.N; i++ {
		gotime.TimeAgo(date, baseTime)
	}
}

func TestTimeAgo(t *testing.T) {
	// Use fixed base time to prevent time-dependent test failures
	baseTime := time.Date(2025, 7, 8, 12, 0, 0, 0, time.UTC) // Fixed: July 8, 2025 12:00:00 UTC

	// Test with current time using base time for consistency
	timeAgoTestCaseWithBaseTime(t, "Just now", baseTime, baseTime)
	timeAgoTestCaseWithBaseTime(t, "Just now", baseTime.Add(time.Second*-9), baseTime)
	timeAgoTestCaseWithBaseTime(t, "In a few seconds", baseTime.Add(time.Second*9), baseTime)

	timeAgoTestCaseWithBaseTime(t, "A minute ago", baseTime.Add(time.Second*-30), baseTime)
	timeAgoTestCaseWithBaseTime(t, "In a minute", baseTime.Add(time.Second*30), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In a few minutes", baseTime.Add(time.Minute*2), baseTime)
	timeAgoTestCaseWithBaseTime(t, "Few minutes ago", baseTime.Add(time.Minute*-2), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In 2 hours", baseTime.Add(time.Hour*2), baseTime)
	timeAgoTestCaseWithBaseTime(t, "2 hours ago", baseTime.Add(time.Hour*-2), baseTime)

	timeAgoTestCaseWithBaseTime(t, "Tomorrow", baseTime.AddDate(0, 0, 1), baseTime)
	timeAgoTestCaseWithBaseTime(t, "Yesterday", baseTime.AddDate(0, 0, -1), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In 2 days", baseTime.AddDate(0, 0, 2), baseTime)
	timeAgoTestCaseWithBaseTime(t, "2 days ago", baseTime.AddDate(0, 0, -2), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In a week", baseTime.AddDate(0, 0, 8), baseTime)
	timeAgoTestCaseWithBaseTime(t, "Last week", baseTime.AddDate(0, 0, -8), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In 2 months", baseTime.AddDate(0, 2, 0), baseTime)
	timeAgoTestCaseWithBaseTime(t, "2 months ago", baseTime.AddDate(0, -2, 0), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In 2 years", baseTime.AddDate(2, 0, 0), baseTime)
	timeAgoTestCaseWithBaseTime(t, "2 years ago", baseTime.AddDate(-2, 0, -1), baseTime)

	timeAgoTestCaseWithBaseTime(t, "In 2 years", baseTime.AddDate(2, 2, 2), baseTime)
	timeAgoTestCaseWithBaseTime(t, "2 years ago", baseTime.AddDate(-2, -2, -2), baseTime)

	// Test very large time differences to reach 100% coverage (trigger final return path)
	timeAgoTestCaseWithBaseTime(t, "In 10 years", baseTime.AddDate(10, 0, 0), baseTime)
	timeAgoTestCaseWithBaseTime(t, "10 years ago", baseTime.AddDate(-10, 0, 0), baseTime)

	// Test case for TimeAgo() without base time (uses fixed base time for consistency)
	timeAgoTestCaseWithBaseTime(t, "Just now", baseTime, baseTime)

	// Test the branch where no baseTime is provided - this should hit the else branch
	// We'll test with a very recent time to ensure predictable results
	result := gotime.TimeAgo(time.Now().Add(-time.Second))
	if result != "Just now" && result != "A minute ago" {
		t.Errorf("Expected 'Just now' or 'A minute ago', got: %v", result)
	}
}

func timeAgoTestCaseWithBaseTime(t *testing.T, expected string, date time.Time, baseTime time.Time) {
	timeAgo := gotime.TimeAgo(date, baseTime)
	if timeAgo != expected {
		t.Errorf("Expected \"%v\", got, \"%v\"", expected, timeAgo)
	}
}
