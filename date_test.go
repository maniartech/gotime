package temporal

import (
	"testing"
	"time"
)

const (
	ErrInvalidFormat = "Invalid format"
)

func TestDateMonday(t *testing.T) {
	d := Date().Monday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+1)
	if d.Start != e {
		t.Error("Monday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Monday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+1+14)
	if d.Start != e {
		t.Error("Date is not a Monday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}
}

func TestDateTuesday(t *testing.T) {
	d := Date().Tuesday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+2)
	if d.Start != e {
		t.Error("Tuesday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Tuesday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+2+14)
	if d.Start != e {
		t.Error("Date is not a Tuesday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}
}

func TestDateWednesday(t *testing.T) {
	d := Date().Wednesday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+3)
	if d.Start != e {
		t.Error("Wednesday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Wednesday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+3+14)
	if d.Start != e {
		t.Error("Date is not a Wednesday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}
}

func TestDateThursday(t *testing.T) {
	d := Date().Thursday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+4)
	if d.Start != e {
		t.Error("Thursday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Thursday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+4+14)
	if d.Start != e {
		t.Error("Date is not a Thursday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

}

func TestDateFriday(t *testing.T) {
	d := Date().Friday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+5)
	if d.Start != e {
		t.Error("Friday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Friday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+5+14)
	if d.Start != e {
		t.Error("Date is not a Friday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

}

func TestDateSaturday(t *testing.T) {
	d := Date().Saturday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+6)
	if d.Start != e {
		t.Error("Saturday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Saturday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+6+14)
	if d.Start != e {
		t.Error("Date is not a Saturday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

}

func TestDateSunday(t *testing.T) {
	d := Date().Sunday()
	e := DayStart().AddDate(0, 0, -int(DayStart().Weekday())+7)
	if d.Start != e {
		t.Error("Sunday is not the start of the week")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

	d = Date().Sunday(2)
	e = DayStart().AddDate(0, 0, -int(DayStart().Weekday())+7+14)
	if d.Start != e {
		t.Error("Date is not a Sunday")
	}
	if d.End != DayEnd(e) {
		t.Error("Not the correct end to the day")
	}

}

func TestDate(t *testing.T) {
	ti := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	d := Date(ti)

	if d.Start != DayStart(ti) {
		t.Error("Date is not the start of the day")
	}
	if d.End != DayEnd(ti) {
		t.Error("Date is not the end of the day")
	}
}

func TestToday(t *testing.T) {
	today := DayStart()
	if !today.Equal(DayStart()) {
		t.Error(ErrInvalidFormat)
	}

	today = Today()
	if !today.Equal(DayStart()) {
		t.Error(ErrInvalidFormat)
	}
}

func TestEoD(t *testing.T) {
	eod := DayEnd()
	if eod != DayStart().Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
		t.Error(ErrInvalidFormat)
	}
}

func TestYesterday(t *testing.T) {
	yesterday := Yesterday()
	if yesterday != DayStart().AddDate(0, 0, -1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := Tomorrow()
	if tomorrow != DayStart().AddDate(0, 0, 1) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastWeek(t *testing.T) {
	lastWeek := LastWeek()
	if lastWeek != DayStart().AddDate(0, 0, -7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastMonth(t *testing.T) {
	lastMonth := LastMonth()
	if lastMonth != DayStart().AddDate(0, -1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestLastYear(t *testing.T) {
	lastYear := LastYear()
	if lastYear != DayStart().AddDate(-1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextWeek(t *testing.T) {
	nextWeek := NextWeek()
	if nextWeek != DayStart().AddDate(0, 0, 7) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextMonth(t *testing.T) {
	nextMonth := NextMonth()
	if nextMonth != DayStart().AddDate(0, 1, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func TestNextYear(t *testing.T) {
	nextYear := NextYear()
	if nextYear != DayStart().AddDate(1, 0, 0) {
		t.Error(ErrInvalidFormat)
	}
}

func BenchmarkToday(b *testing.B) {
	// Benchmarking for Today function
	for i := 0; i < b.N; i++ {
		DayStart()
	}
	for i := 0; i < b.N; i++ {
		Today()
	}
}
func BenchmarkEoD(b *testing.B) {
	// Benchmarking for EoD function
	for i := 0; i < b.N; i++ {
		DayEnd()
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
