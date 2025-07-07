package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

var (
	benchDate       = time.Date(2025, 7, 15, 14, 30, 45, 0, time.UTC)
	benchResult     int
	benchBoolResult bool
)

func BenchmarkDayOfYear(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = gotime.DayOfYear(benchDate)
	}
	benchResult = result
}

func BenchmarkWeekOfMonth(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = gotime.WeekOfMonth(benchDate)
	}
	benchResult = result
}

func BenchmarkIsFirstDayOfMonth(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = gotime.IsFirstDayOfMonth(benchDate)
	}
	benchBoolResult = result
}

func BenchmarkIsLastDayOfMonth(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = gotime.IsLastDayOfMonth(benchDate)
	}
	benchBoolResult = result
}
