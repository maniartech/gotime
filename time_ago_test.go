package datetime

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkTimeAgo(b *testing.B) {
	// Benchmark for TimeAgo()
	date := time.Now().UTC().AddDate(0, 0, -1)
	for i := 0; i < b.N; i++ {
		TimeAgo(date)
	}
}

func TestTimeAgo(t *testing.T) {
	// Test case for TimeAgo()

	fmt.Println(TimeAgo(time.Date(2022, 1, 1, 19, 40, 0, 0, time.Local)))

	timeAgoTestCase(t, "Just now", time.Now().UTC())
	timeAgoTestCase(t, "Just now", time.Now().UTC().Add(time.Second*-10))
	timeAgoTestCase(t, "In a few seconds", time.Now().UTC().Add(time.Second*10))

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

}

func timeAgoTestCase(t *testing.T, expected string, date time.Time) {
	timeAgo := TimeAgo(date)
	if timeAgo != expected {
		t.Errorf("Expected \"%v\", got, \"%v\"", expected, timeAgo)
	}
}
