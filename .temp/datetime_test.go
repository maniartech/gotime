package gotime_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/maniartech/gotime"
// )

// const (
// 	ErrInvalidFormat = "Invalid format"
// )

// func TestDateMonday(t *testing.T) {
// 	d := gotime.Date().Monday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+1)
// 	if d != e {
// 		t.Error("Monday is not the start of the week")
// 	}

// 	d = gotime.Date().Monday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+1+14)
// 	if d != e {
// 		t.Error("Date is not a Monday")
// 	}
// }

// func TestDateTuesday(t *testing.T) {
// 	d := gotime.Date().Tuesday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+2)
// 	if d != e {
// 		t.Error("Tuesday is not the start of the week")
// 	}

// 	d = gotime.Date().Tuesday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+2+14)
// 	if d != e {
// 		t.Error("Date is not a Tuesday")
// 	}
// }

// func TestDateWednesday(t *testing.T) {
// 	d := gotime.Date().Wednesday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+3)
// 	if d != e {
// 		t.Error("Wednesday is not the start of the week")
// 	}

// 	d = gotime.Date().Wednesday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+3+14)
// 	if d != e {
// 		t.Error("Date is not a Wednesday")
// 	}
// }

// func TestDateThursday(t *testing.T) {
// 	d := gotime.Date().Thursday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+4)
// 	if d != e {
// 		t.Error("Thursday is not the start of the week")
// 	}

// 	d = gotime.Date().Thursday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+4+14)
// 	if d != e {
// 		t.Error("Date is not a Thursday")
// 	}

// }

// func TestDateFriday(t *testing.T) {
// 	d := gotime.Date().Friday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+5)
// 	if d != e {
// 		t.Error("Friday is not the start of the week")
// 	}

// 	d = gotime.Date().Friday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+5+14)
// 	if d != e {
// 		t.Error("Date is not a Friday")
// 	}

// }

// func TestDateSaturday(t *testing.T) {
// 	d := gotime.Date().Saturday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+6)
// 	if d != e {
// 		t.Error("Saturday is not the start of the week")
// 	}

// 	d = gotime.Date().Saturday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+6+14)
// 	if d != e {
// 		t.Error("Date is not a Saturday")
// 	}

// }

// func TestDateSunday(t *testing.T) {
// 	d := gotime.Date().Sunday()
// 	e := gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday()))
// 	if d != e {
// 		t.Error("Sunday is not the start of the week")
// 	}

// 	d = gotime.Date().Sunday(2)
// 	e = gotime.DayStart().AddDate(0, 0, -int(gotime.DayStart().Time().Weekday())+7+7)
// 	if d != e {
// 		t.Error("Date is not a Sunday")
// 	}
// }

// func TestDate(t *testing.T) {
// 	ti := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
// 	d := gotime.Date(ti)

// 	if d != gotime.DayStart(ti) {
// 		t.Error("Date is not the start of the day")
// 	}
// 	if d != gotime.DayEnd(ti) {
// 		t.Error("Date is not the end of the day")
// 	}
// }

// func TestToday(t *testing.T) {
// 	today := gotime.DayStart()
// 	if today != (gotime.DayStart()) {
// 		t.Error(ErrInvalidFormat)
// 	}

// 	today = gotime.Today()
// 	if today != (gotime.DayStart()) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestEoD(t *testing.T) {
// 	eod := gotime.DayEnd()
// 	if eod != gotime.DayStart().Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestYesterday(t *testing.T) {
// 	yesterday := gotime.Yesterday()
// 	if yesterday != gotime.DayStart().AddDate(0, 0, -1) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestTomorrow(t *testing.T) {
// 	tomorrow := gotime.Tomorrow()
// 	if tomorrow != gotime.DayStart().AddDate(0, 0, 1) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestLastWeek(t *testing.T) {
// 	lastWeek := gotime.LastWeek()
// 	if lastWeek != gotime.DayStart().AddDate(0, 0, -7) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestLastMonth(t *testing.T) {
// 	lastMonth := gotime.LastMonth()
// 	if lastMonth != gotime.DayStart().AddDate(0, -1, 0) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestLastYear(t *testing.T) {
// 	lastYear := gotime.LastYear()
// 	if lastYear != gotime.DayStart().AddDate(-1, 0, 0) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestNextWeek(t *testing.T) {
// 	nextWeek := gotime.NextWeek()
// 	if nextWeek != gotime.DayStart().AddDate(0, 0, 7) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestNextMonth(t *testing.T) {
// 	nextMonth := gotime.NextMonth()
// 	if nextMonth != gotime.DayStart().AddDate(0, 1, 0) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func TestNextYear(t *testing.T) {
// 	nextYear := gotime.NextYear()
// 	if nextYear != gotime.DayStart().AddDate(1, 0, 0) {
// 		t.Error(ErrInvalidFormat)
// 	}
// }

// func BenchmarkToday(b *testing.B) {
// 	// Benchmarking for Today function
// 	for i := 0; i < b.N; i++ {
// 		gotime.DayStart()
// 	}
// 	for i := 0; i < b.N; i++ {
// 		gotime.Today()
// 	}
// }
// func BenchmarkEoD(b *testing.B) {
// 	// Benchmarking for EoD function
// 	for i := 0; i < b.N; i++ {
// 		gotime.DayEnd()
// 	}
// }

// func BenchmarkYesterday(b *testing.B) {
// 	// Benchmarking for Yesterday function
// 	for i := 0; i < b.N; i++ {
// 		gotime.Yesterday()
// 	}
// }

// func BenchmarkTomorrow(b *testing.B) {
// 	// Benchmarking for Tomorrow function
// 	for i := 0; i < b.N; i++ {
// 		gotime.Tomorrow()
// 	}
// }

// func BenchmarkLastWeek(b *testing.B) {
// 	// Benchmarking for gotime.LastWeek function
// 	for i := 0; i < b.N; i++ {
// 		gotime.LastWeek()
// 	}
// }

// func BenchmarkLastMonth(b *testing.B) {
// 	// Benchmarking for gotime.LastMonth function
// 	for i := 0; i < b.N; i++ {
// 		gotime.LastMonth()
// 	}
// }

// func BenchmarkLastYear(b *testing.B) {
// 	// Benchmarking for gotime.LastYear function
// 	for i := 0; i < b.N; i++ {
// 		gotime.LastYear()
// 	}
// }

// func BenchmarkNextWeek(b *testing.B) {
// 	// Benchmarking for gotime.NextWeek function
// 	for i := 0; i < b.N; i++ {
// 		gotime.NextWeek()
// 	}
// }

// func BenchmarkNextMonth(b *testing.B) {
// 	// Benchmarking for gotime.NextMonth function
// 	for i := 0; i < b.N; i++ {
// 		gotime.NextMonth()
// 	}
// }

// func BenchmarkNextYear(b *testing.B) {
// 	// Benchmarking for gotime.NextYear function
// 	for i := 0; i < b.N; i++ {
// 		gotime.NextYear()
// 	}
// }
