package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

func BenchmarkTimeAgo(b *testing.B) {
	// Benchmark for TimeAgo()
	date := time.Now().UTC().AddDate(0, 0, 10)
	for i := 0; i < b.N; i++ {
		gotime.TimeAgo(date)
	}
}

func TestTimeAgo(t *testing.T) {
	timeAgoTestCase(t, "Just now", time.Now().UTC())
	timeAgoTestCase(t, "Just now", time.Now().UTC().Add(time.Second*-9))
	timeAgoTestCase(t, "In a few seconds", time.Now().UTC().Add(time.Second*9))

	timeAgoTestCase(t, "A minute ago", time.Now().UTC().Add(time.Second*-30))
	timeAgoTestCase(t, "In a minute", time.Now().UTC().Add(time.Second*30))

	timeAgoTestCase(t, "In a few minutes", time.Now().UTC().Add(time.Minute*2))
	timeAgoTestCase(t, "Few minutes ago", time.Now().UTC().Add(time.Minute*-2))

	timeAgoTestCase(t, "In 2 hours", time.Now().UTC().Add(time.Hour*2))
	timeAgoTestCase(t, "2 hours ago", time.Now().UTC().Add(time.Hour*-2))

	timeAgoTestCase(t, "Tomorrow", time.Now().UTC().AddDate(0, 0, 1))
	timeAgoTestCase(t, "Yesterday", time.Now().UTC().AddDate(0, 0, -1))

	timeAgoTestCase(t, "In 2 days", time.Now().UTC().AddDate(0, 0, 2))
	timeAgoTestCase(t, "2 days ago", time.Now().UTC().AddDate(0, 0, -2))

	timeAgoTestCase(t, "In a week", time.Now().UTC().AddDate(0, 0, 8))
	timeAgoTestCase(t, "Last week", time.Now().UTC().AddDate(0, 0, -8))

	timeAgoTestCase(t, "In 2 months", time.Now().UTC().AddDate(0, 2, 0))
	timeAgoTestCase(t, "2 months ago", time.Now().UTC().AddDate(0, -2, 0))

	timeAgoTestCase(t, "In 2 years", time.Now().UTC().AddDate(2, 0, 0))
	timeAgoTestCase(t, "2 years ago", time.Now().UTC().AddDate(-2, 0, -1))

	timeAgoTestCase(t, "In 2 years", time.Now().UTC().AddDate(2, 2, 2))
	timeAgoTestCase(t, "2 years ago", time.Now().UTC().AddDate(-2, -2, -2))

	// Test case for TimeAgo() with base time
	timeAgoTestCaseWithBaseTime(t, "Just now", time.Now().UTC(), time.Now().UTC())
}

func timeAgoTestCase(t *testing.T, expected string, date time.Time) {
	timeAgo := gotime.TimeAgo(date)
	if timeAgo != expected {
		t.Errorf("Expected \"%v\", got, \"%v\"", expected, timeAgo)
	}
}

func timeAgoTestCaseWithBaseTime(t *testing.T, expected string, date time.Time, baseTime time.Time) {
	timeAgo := gotime.TimeAgo(date, baseTime)
	if timeAgo != expected {
		t.Errorf("Expected \"%v\", got, \"%v\"", expected, timeAgo)
	}
}
