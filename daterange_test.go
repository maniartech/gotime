package dateutils_test

import (
	"testing"
	"time"

	"github.com/maniartech/dateutils"
)

// Covers all test cases for the function RelativeRange in daterange.go
func TestRelativeRange(t *testing.T) {

	now := time.Now().UTC()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	t.Run("today", func(t *testing.T) {
		// Test case for today
		d1, d2, err := dateutils.RelativeRange("today")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 1))
	})

	t.Run("yesterday", func(t *testing.T) {
		// Test case for yesterday
		d1, d2, err := dateutils.RelativeRange("yesterday")
		if err != nil {
			t.Error(err)
		}
		yesterday := todayMidnight.AddDate(0, 0, -1)
		isEqual(t, d1, yesterday)
		isEqual(t, d2, todayMidnight)
	})

	t.Run("tomorrow", func(t *testing.T) {
		// Test case for tomorrow
		d1, d2, err := dateutils.RelativeRange("tomorrow")
		if err != nil {
			t.Error(err)
		}
		tomorrow := todayMidnight.AddDate(0, 0, 1)
		isEqual(t, d1, tomorrow)
		isEqual(t, d2, tomorrow.AddDate(0, 0, 1))
	})

	t.Run("last-<n>days", func(t *testing.T) {
		// Test case for last-<n>days
		d1, d2, err := dateutils.RelativeRange("last-5days")
		if err != nil {
			t.Error(err)
		}
		last5days := todayMidnight.AddDate(0, 0, -5)
		isEqual(t, d1, last5days)
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 1))
	})

	t.Run("next-<n>days", func(t *testing.T) {
		// Test case for next-<n>days
		d1, d2, err := dateutils.RelativeRange("next-5days")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 6))
	})

	t.Run("thisweek", func(t *testing.T) {
		// Test case for thisweek
		d1, d2, err := dateutils.RelativeRange("thisweek")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 7))
	})

	t.Run("lastweek", func(t *testing.T) {
		// Test case for lastweek
		d1, d2, err := dateutils.RelativeRange("lastweek")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, 0, -7))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("nextweek", func(t *testing.T) {
		// Test case for nextweek
		d1, d2, err := dateutils.RelativeRange("nextweek")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, 0, 7))
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 14))
	})

	t.Run("last-<n>weeks", func(t *testing.T) {
		// Test case for last-<n>weeks
		d1, d2, err := dateutils.RelativeRange("last-2weeks")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, 0, -14))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("next-<n>weeks", func(t *testing.T) {
		// Test case for next-<n>weeks
		d1, d2, err := dateutils.RelativeRange("next-2weeks")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 0, 14))
	})

	t.Run("thismonth", func(t *testing.T) {
		// Test case for thismonth
		d1, d2, err := dateutils.RelativeRange("thismonth")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 1, 0))
	})


	t.Run("lastmonth", func(t *testing.T) {
		// Test case for lastmonth
		d1, d2, err := dateutils.RelativeRange("lastmonth")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, -1, 0))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("nextmonth", func(t *testing.T) {
		// Test case for nextmonth
		d1, d2, err := dateutils.RelativeRange("nextmonth")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, 1, 0))
		isEqual(t, d2, todayMidnight.AddDate(0, 2, 0))
	})

	t.Run("last-<n>months", func(t *testing.T) {
		// Test case for last-<n>months
		d1, d2, err := dateutils.RelativeRange("last-2months")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(0, -2, 0))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("next-<n>months", func(t *testing.T) {
		// Test case for next-<n>months
		d1, d2, err := dateutils.RelativeRange("next-2months")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(0, 2, 0))
	})

	t.Run("thisyear", func(t *testing.T) {
		// Test case for thisyear
		d1, d2, err := dateutils.RelativeRange("thisyear")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(1, 0, 0))
	})


	t.Run("lastyear", func(t *testing.T) {
		// Test case for lastyear
		d1, d2, err := dateutils.RelativeRange("lastyear")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(-1, 0, 0))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("nextyear", func(t *testing.T) {
		// Test case for nextyear
		d1, d2, err := dateutils.RelativeRange("nextyear")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(1, 0, 0))
		isEqual(t, d2, todayMidnight.AddDate(2, 0, 0))
	})

	t.Run("last-<n>years", func(t *testing.T) {
		// Test case for last-<n>years
		d1, d2, err := dateutils.RelativeRange("last-2years")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight.AddDate(-2, 0, 0))
		isEqual(t, d2, todayMidnight)
	})

	t.Run("next-<n>years", func(t *testing.T) {
		// Test case for next-<n>years
		d1, d2, err := dateutils.RelativeRange("next-2years")
		if err != nil {
			t.Error(err)
		}
		isEqual(t, d1, todayMidnight)
		isEqual(t, d2, todayMidnight.AddDate(2, 0, 0))
	})

}

// isEqual compares two datetime values and returns true if they are equal
func isEqual(t *testing.T, d1 time.Time, d2 time.Time) {
	if d1.IsZero() || d2.IsZero() {
		t.Error("Expected non-zero time")
	}
	if d1 != d2 {
		t.Error("Expected equal time", d1, d2)
	}
}
