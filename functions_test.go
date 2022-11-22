package dateutils

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	today := *Today()
	if !today.Equal(*Today()) {
		t.Error(ErrInvalidFormat)
	}
}

func TestEoD(t *testing.T) {
	eod := *EoD()
	if eod != TodayStart().Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
		t.Error(ErrInvalidFormat)
	}
}

func TestYesterday(t *testing.T) {
	yesterday := *Yesterday()
	if yesterday != TodayStart().AddDate(0, 0, -1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := *Tomorrow()
	if tomorrow != TodayStart().AddDate(0, 0, 1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastWeek(t *testing.T) {
	lastWeek := *LastWeek()
	if lastWeek != TodayStart().AddDate(0, 0, -7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastMonth(t *testing.T) {
	lastMonth := *LastMonth()
	if lastMonth != TodayStart().AddDate(0, -1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastYear(t *testing.T) {
	lastYear := *LastYear()
	if lastYear != TodayStart().AddDate(-1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextWeek(t *testing.T) {
	nextWeek := *NextWeek()
	if nextWeek != TodayStart().AddDate(0, 0, 7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextMonth(t *testing.T) {
	nextMonth := *NextMonth()
	if nextMonth != TodayStart().AddDate(0, 1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextYear(t *testing.T) {
	nextYear := *NextYear()
	if nextYear != TodayStart().AddDate(1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func BenchmarkToday(b *testing.B) {
	// Benchmarking for Today function
	for i := 0; i < b.N; i++ {
		Today()
	}
}
func BenchmarkEoD(b *testing.B) {
	// Benchmarking for EoD function
	for i := 0; i < b.N; i++ {
		EoD()
	}
}

func BenchmarkYesterday(b *testing.B) {
	// Benchmarking for Yesterday function
	for i := 0; i < b.N; i++ {
		Yesterday()
	}
}

func BenchmarkTomorrow(b *testing.B) {
	// Benchmarking for Tomorrow function
	for i := 0; i < b.N; i++ {
		Tomorrow()
	}
}

func BenchmarkLastWeek(b *testing.B) {
	// Benchmarking for LastWeek function
	for i := 0; i < b.N; i++ {
		LastWeek()
	}
}

func BenchmarkLastMonth(b *testing.B) {
	// Benchmarking for LastMonth function
	for i := 0; i < b.N; i++ {
		LastMonth()
	}
}

func BenchmarkLastYear(b *testing.B) {
	// Benchmarking for LastYear function
	for i := 0; i < b.N; i++ {
		LastYear()
	}
}

func BenchmarkNextWeek(b *testing.B) {
	// Benchmarking for NextWeek function
	for i := 0; i < b.N; i++ {
		NextWeek()
	}
}

func BenchmarkNextMonth(b *testing.B) {
	// Benchmarking for NextMonth function
	for i := 0; i < b.N; i++ {
		NextMonth()
	}
}

func BenchmarkNextYear(b *testing.B) {
	// Benchmarking for NextYear function
	for i := 0; i < b.N; i++ {
		NextYear()
	}
}
