package dateutils

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	today := *Today()
	if today != todayMidnight {
		t.Error("")
	}
}

func TestEoD(t *testing.T) {
	eod := *EoD()
	if eod != todayMidnight.Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
		t.Error("")
	}
}

func TestYesterday(t *testing.T) {
	yesterday := *Yesterday()
	if yesterday != todayMidnight.AddDate(0, 0, -1) {
		t.Error("")
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := *Tomorrow()
	if tomorrow != todayMidnight.AddDate(0, 0, 1) {
		t.Error("")
	}
}

func LastWeekTest(t *testing.T) {
	lastWeek := *LastWeek()
	if lastWeek != todayMidnight.AddDate(0, 0, -1) {
		t.Error("")
	}
}

func LastMonthTest(t *testing.T) {
	lastMonth := *LastMonth()
	if lastMonth != todayMidnight.AddDate(0, -1, 0) {
		t.Error("")
	}
}

func LastYearTest(t *testing.T) {
	lastYear := *LastYear()
	if lastYear != todayMidnight.AddDate(-1, 0, 0) {
		t.Error("")
	}
}

func NextWeekTest(t *testing.T) {
	nextWeek := *NextWeek()
	if nextWeek != todayMidnight.AddDate(0, 0, 1) {
		t.Error("")
	}
}

func NextMonthTest(t *testing.T) {
	nextMonth := *NextMonth()
	if nextMonth != todayMidnight.AddDate(0, 1, 0) {
		t.Error("")
	}
}

func NextYearTest(t *testing.T) {
	nextYear := *NextYear()
	if nextYear != todayMidnight.AddDate(1, 0, 0) {
		t.Error("")
	}
}
