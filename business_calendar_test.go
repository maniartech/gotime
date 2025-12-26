package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
)

func TestIsBusinessDay(t *testing.T) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Independence Day
	}
	tests := []struct {
		name string
		date time.Time
		want bool
	}{
		{"Weekday, not holiday", time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), true}, // Monday
		{"Saturday", time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC), false},
		{"Sunday", time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC), false},
		{"Holiday", time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gotime.IsBusinessDay(tc.date, weekends, holidays...)
			if got != tc.want {
				t.Errorf("IsBusinessDay(%v) = %v, want %v", tc.date, got, tc.want)
			}
		})
	}
}

func TestNextBusinessDay(t *testing.T) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Friday
	}
	tests := []struct {
		name  string
		start time.Time
		want  time.Time
	}{
		{"Friday (holiday)", time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)}, // Monday
		{"Saturday", time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)},
		{"Sunday", time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)},
		{"Monday", time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 8, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gotime.NextBusinessDay(tc.start, weekends, holidays...)
			if !got.Equal(tc.want) {
				t.Errorf("NextBusinessDay(%v) = %v, want %v", tc.start, got, tc.want)
			}
		})
	}
}

func TestPrevBusinessDay(t *testing.T) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Friday
	}
	tests := []struct {
		name  string
		start time.Time
		want  time.Time
	}{
		{"Monday", time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC)}, // Thursday
		{"Saturday", time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC)},
		{"Sunday", time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC)},
		{"Friday (holiday)", time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gotime.PrevBusinessDay(tc.start, weekends, holidays...)
			if !got.Equal(tc.want) {
				t.Errorf("PrevBusinessDay(%v) = %v, want %v", tc.start, got, tc.want)
			}
		})
	}
}
