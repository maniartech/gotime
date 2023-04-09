package temporal_test

import (
	"testing"

	"github.com/maniartech/temporal"
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
		if temporal.IsLeapYear(year) != leap {
			t.Errorf("Expected %v, got, %v", leap, temporal.IsLeapYear(year))
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
		if temporal.DaysInMonth(2019, month) != days {
			t.Errorf("Expected %v, got, %v", days, temporal.DaysInMonth(2019, month))
		}
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
		if temporal.DaysInYear(year) != days {
			t.Errorf("Expected %v, got, %v", days, temporal.DaysInYear(year))
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
		if temporal.DaysInQuarter(2019, quarter) != days {
			t.Errorf("Expected %v, got, %v", days, temporal.DaysInQuarter(2019, quarter))
		}
	}
}
