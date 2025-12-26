package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
)

// TestIsLeapYear tests IsLeapYear function.
func TestIsLeapYear(t *testing.T) {

	// Years Leap Map is a map of years and their leap status.
	// It contains random years to test the IsLeapYear function.
	yearsLeapMap := map[int]bool{
		1971: false,
		1980: true,
		1992: true,
		1995: false,
		2000: true,
		2001: false,
		2004: true,
		2009: false,
		2012: true,
		2015: false,
		2016: true,
		2020: true,
		2021: false,
		2024: true,
	}

	// Test Case for IsLeapYear function
	for year, leap := range yearsLeapMap {
		if gotime.IsLeapYear(year) != leap {
			t.Errorf("Expected %v, got, %v", leap, gotime.IsLeapYear(year))
		}
	}
}

// TestDaysInMonth tests DaysInMonth function.
func TestDaysInMonth(t *testing.T) {

	// MonthDaysMap is a map of months and their number of days.
	// It contains random months to test the DaysInMonth function.
	monthDaysMap := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	// Test Case for DaysInMonth function
	for month, days := range monthDaysMap {
		if gotime.DaysInMonth(2019, month) != days {
			t.Errorf("Expected %v, got, %v", days, gotime.DaysInMonth(2019, month))
		}
	}
	if gotime.DaysInMonth(2020, 2) != 29 {
		t.Errorf("Expected %v, got, %v", 29, gotime.DaysInMonth(2019, 2))
	}
}

// TestDaysInYear tests DaysInYear function.
func TestDaysInYear(t *testing.T) {

	// YearsDaysMap is a map of years and their number of days.
	// It contains random years to test the DaysInYear function.
	yearsDaysMap := map[int]int{
		1971: 365,
		1980: 366,
		1992: 366,
		1995: 365,
		2000: 366,
		2001: 365,
		2004: 366,
		2009: 365,
		2012: 366,
		2015: 365,
		2016: 366,
		2020: 366,
		2021: 365,
		2024: 366,
	}

	// Test Case for DaysInYear function
	for year, days := range yearsDaysMap {
		if gotime.DaysInYear(year) != days {
			t.Errorf("Expected %v, got, %v", days, gotime.DaysInYear(year))
		}
	}
}

// TestDaysInQuarter tests DaysInQuarter function.
func TestDaysInQuarter(t *testing.T) {

	// QuarterDaysMap is a map of quarters and their number of days.
	// It contains random quarters to test the DaysInQuarter function.
	quarterDaysMap := map[int]int{
		1: 90,
		2: 91,
		3: 92,
		4: 92,
	}

	// Test Case for DaysInQuarter function
	for quarter, days := range quarterDaysMap {
		if gotime.DaysInQuarter(2019, quarter) != days {
			t.Errorf("Expected %v, got, %v", days, gotime.DaysInQuarter(2019, quarter))
		}
	}
}

// TestDateCreate tests DateCreate function.
func TestDateCreate(t *testing.T) {
	date := gotime.NewDate(2021, 1, 1, time.Local)
	if date.Year() != 2021 || date.Month() != 1 || date.Day() != 1 {
		t.Errorf("Expected 2021-01-01, got, %v", date)
	}
}

// TestNewTime tests NewTime function.
func TestNewTime(t *testing.T) {
	time := gotime.NewTime(12, 0, 0, time.Local)
	if time.Hour() != 12 || time.Minute() != 0 || time.Second() != 0 {
		t.Errorf("Expected 12:00:00, got, %v", time)
	}
}

// TestEoD tests EoD function.
func TestEoD(t *testing.T) {
	date := gotime.EoD(gotime.NewDate(2021, 1, 1, time.Local))
	if date.Year() != 2021 || date.Month() != 1 || date.Day() != 1 || date.Hour() != 23 || date.Minute() != 59 || date.Second() != 59 {
		t.Errorf("Expected 2021-01-01 23:59:59, got, %v", date)
	}
}

// TestSoD tests SoD function.
func TestSoD(t *testing.T) {
	date := gotime.SoD(gotime.NewDate(2021, 1, 1, time.Local))
	if date.Year() != 2021 || date.Month() != 1 || date.Day() != 1 || date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 {
		t.Errorf("Expected 2021-01-01 00:00:00, got, %v", date)
	}
}

func TestReplaceDate(t *testing.T) {
	date := time.Date(2022, time.April, 15, 12, 30, 0, 0, time.UTC)
	newDate := gotime.ReplaceDate(date, 2023, 5, 20)
	expected := time.Date(2023, time.May, 20, 12, 30, 0, 0, time.UTC)

	if !newDate.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, newDate)
	}
}

func TestReplaceTime(t *testing.T) {
	date := time.Date(2022, time.April, 15, 12, 30, 0, 0, time.UTC)
	newTime := gotime.ReplaceTime(date, 15, 45, 0)
	expected := time.Date(2022, time.April, 15, 15, 45, 0, 0, time.UTC)

	if !newTime.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, newTime)
	}
}
