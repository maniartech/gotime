package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

func TestDayOfYear(t *testing.T) {
	tests := []struct {
		date time.Time
		want int
	}{
		{time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), 188},
		{time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 366}, // Leap year
	}
	for _, tc := range tests {
		if got := gotime.DayOfYear(tc.date); got != tc.want {
			t.Errorf("DayOfYear(%v) = %d, want %d", tc.date, got, tc.want)
		}
	}
}

func TestWeekOfMonth(t *testing.T) {
	tests := []struct {
		date time.Time
		want int
	}{
		{time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC), 3},
		{time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC), 5},
	}
	for _, tc := range tests {
		if got := gotime.WeekOfMonth(tc.date); got != tc.want {
			t.Errorf("WeekOfMonth(%v) = %d, want %d", tc.date, got, tc.want)
		}
	}
}

func TestIsFirstDayOfMonth(t *testing.T) {
	tests := []struct {
		date time.Time
		want bool
	}{
		{time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2025, 7, 2, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tc := range tests {
		if got := gotime.IsFirstDayOfMonth(tc.date); got != tc.want {
			t.Errorf("IsFirstDayOfMonth(%v) = %v, want %v", tc.date, got, tc.want)
		}
	}
}

func TestIsLastDayOfMonth(t *testing.T) {
	tests := []struct {
		date time.Time
		want bool
	}{
		{time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC), true},  // Last day of Feb in non-leap year
		{time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC), false}, // Not last day of Feb
		{time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), true},  // Last day of Feb in leap year
		{time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC), false}, // Not last day of Feb in leap year
		{time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC), true},  // Last day of July
		{time.Date(2025, 7, 30, 0, 0, 0, 0, time.UTC), false}, // Not last day of July
	}
	for _, tc := range tests {
		if got := gotime.IsLastDayOfMonth(tc.date); got != tc.want {
			t.Errorf("IsLastDayOfMonth(%v) = %v, want %v", tc.date, got, tc.want)
		}
	}
}
