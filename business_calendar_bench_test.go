package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

func BenchmarkIsBusinessDay(b *testing.B) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
	}
	dates := []time.Time{
		time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC), // Thursday
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC), // Friday (holiday)
		time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC), // Saturday
		time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC), // Sunday
		time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC), // Monday
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gotime.IsBusinessDay(dates[i%len(dates)], weekends, holidays...)
	}
}

func BenchmarkNextBusinessDay(b *testing.B) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
	}
	dates := []time.Time{
		time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gotime.NextBusinessDay(dates[i%len(dates)], weekends, holidays...)
	}
}

func BenchmarkPrevBusinessDay(b *testing.B) {
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holidays := []time.Time{
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
	}
	dates := []time.Time{
		time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gotime.PrevBusinessDay(dates[i%len(dates)], weekends, holidays...)
	}
}
