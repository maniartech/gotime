package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
)

const (
	ErrInvalidFormat = "Invalid format"
)

func TestDateMonday(t *testing.T) {
	d := temporal.Date().Monday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+1)
	if time.Time(d) != e {
		t.Error("Monday is not the start of the week")
	}

	d = temporal.Date().Monday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+1+14)
	if time.Time(d) != e {
		t.Error("Date is not a Monday")
	}
}

func TestDateTuesday(t *testing.T) {
	d := temporal.Date().Tuesday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+2)
	if time.Time(d) != e {
		t.Error("Tuesday is not the start of the week")
	}

	d = temporal.Date().Tuesday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+2+14)
	if time.Time(d) != e {
		t.Error("Date is not a Tuesday")
	}
}

func TestDateWednesday(t *testing.T) {
	d := temporal.Date().Wednesday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+3)
	if time.Time(d) != e {
		t.Error("Wednesday is not the start of the week")
	}

	d = temporal.Date().Wednesday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+3+14)
	if time.Time(d) != e {
		t.Error("Date is not a Wednesday")
	}
}

func TestDateThursday(t *testing.T) {
	d := temporal.Date().Thursday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+4)
	if time.Time(d) != e {
		t.Error("Thursday is not the start of the week")
	}

	d = temporal.Date().Thursday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+4+14)
	if time.Time(d) != e {
		t.Error("Date is not a Thursday")
	}

}

func TestDateFriday(t *testing.T) {
	d := temporal.Date().Friday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+5)
	if time.Time(d) != e {
		t.Error("Friday is not the start of the week")
	}

	d = temporal.Date().Friday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+5+14)
	if time.Time(d) != e {
		t.Error("Date is not a Friday")
	}

}

func TestDateSaturday(t *testing.T) {
	d := temporal.Date().Saturday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+6)
	if time.Time(d) != e {
		t.Error("Saturday is not the start of the week")
	}

	d = temporal.Date().Saturday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+6+14)
	if time.Time(d) != e {
		t.Error("Date is not a Saturday")
	}

}

func TestDateSunday(t *testing.T) {
	d := temporal.Date().Sunday()
	e := temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday()))
	if time.Time(d) != e {
		t.Error("Sunday is not the start of the week")
	}

	d = temporal.Date().Sunday(2)
	e = temporal.DayStart().AddDate(0, 0, -int(temporal.DayStart().Weekday())+7+7)
	if time.Time(d) != e {
		t.Error("Date is not a Sunday")
	}
}

func TestDate(t *testing.T) {
	ti := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	d := temporal.Date(ti)

	if time.Time(d) != temporal.DayStart(ti) {
		t.Error("Date is not the start of the day")
	}
	if time.Time(d) != temporal.DayEnd(ti) {
		t.Error("Date is not the end of the day")
	}
}

func TestToday(t *testing.T) {
	today := temporal.DayStart()
	if !today.Equal(temporal.DayStart()) {
		t.Error(ErrInvalidFormat)
	}

	today = temporal.Today()
	if !today.Equal(temporal.DayStart()) {
		t.Error(ErrInvalidFormat)
	}
}

func TestEoD(t *testing.T) {
	eod := temporal.DayEnd()
	if eod != temporal.DayStart().Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
		t.Error(ErrInvalidFormat)
	}
}

func TestYesterday(t *testing.T) {
	yesterday := temporal.Yesterday()
	if time.Time(yesterday) != temporal.DayStart().AddDate(0, 0, -1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := temporal.Tomorrow()
	if time.Time(tomorrow) != temporal.DayStart().AddDate(0, 0, 1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastWeek(t *testing.T) {
	lastWeek := temporal.LastWeek()
	if lastWeek != temporal.DayStart().AddDate(0, 0, -7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastMonth(t *testing.T) {
	lastMonth := temporal.LastMonth()
	if lastMonth != temporal.DayStart().AddDate(0, -1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastYear(t *testing.T) {
	lastYear := temporal.LastYear()
	if lastYear != temporal.DayStart().AddDate(-1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextWeek(t *testing.T) {
	nextWeek := temporal.NextWeek()
	if nextWeek != temporal.DayStart().AddDate(0, 0, 7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextMonth(t *testing.T) {
	nextMonth := temporal.NextMonth()
	if nextMonth != temporal.DayStart().AddDate(0, 1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextYear(t *testing.T) {
	nextYear := temporal.NextYear()
	if nextYear != temporal.DayStart().AddDate(1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func BenchmarkToday(b *testing.B) {
	// Benchmarking for Today function
	for i := 0; i < b.N; i++ {
		temporal.DayStart()
	}
	for i := 0; i < b.N; i++ {
		temporal.Today()
	}
}
func BenchmarkEoD(b *testing.B) {
	// Benchmarking for EoD function
	for i := 0; i < b.N; i++ {
		temporal.DayEnd()
	}
}

func BenchmarkYesterday(b *testing.B) {
	// Benchmarking for Yesterday function
	for i := 0; i < b.N; i++ {
		temporal.Yesterday()
	}
}

func BenchmarkTomorrow(b *testing.B) {
	// Benchmarking for Tomorrow function
	for i := 0; i < b.N; i++ {
		temporal.Tomorrow()
	}
}

func BenchmarkLastWeek(b *testing.B) {
	// Benchmarking for temporal.LastWeek function
	for i := 0; i < b.N; i++ {
		temporal.LastWeek()
	}
}

func BenchmarkLastMonth(b *testing.B) {
	// Benchmarking for temporal.LastMonth function
	for i := 0; i < b.N; i++ {
		temporal.LastMonth()
	}
}

func BenchmarkLastYear(b *testing.B) {
	// Benchmarking for temporal.LastYear function
	for i := 0; i < b.N; i++ {
		temporal.LastYear()
	}
}

func BenchmarkNextWeek(b *testing.B) {
	// Benchmarking for temporal.NextWeek function
	for i := 0; i < b.N; i++ {
		temporal.NextWeek()
	}
}

func BenchmarkNextMonth(b *testing.B) {
	// Benchmarking for temporal.NextMonth function
	for i := 0; i < b.N; i++ {
		temporal.NextMonth()
	}
}

func BenchmarkNextYear(b *testing.B) {
	// Benchmarking for temporal.NextYear function
	for i := 0; i < b.N; i++ {
		temporal.NextYear()
	}
}
