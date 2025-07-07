package gotime

import (
	"testing"
	"time"
)

func TestIsWeekdayPresentInRange_SinglePresent(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC)  // Sunday
	if !IsWeekdayPresentInRange(start, end, time.Wednesday) {
		t.Errorf("Expected Wednesday to be present in range")
	}
}

func TestIsWeekdayPresentInRange_SingleAbsent(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)   // Monday
	if IsWeekdayPresentInRange(start, end, time.Tuesday) {
		t.Errorf("Expected Tuesday to be absent in range")
	}
}

func TestIsWeekdayPresentInRange_MultipleAnyPresent(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 9, 0, 0, 0, 0, time.UTC)   // Wednesday
	if !IsWeekdayPresentInRange(start, end, time.Friday, time.Wednesday) {
		t.Errorf("Expected Wednesday to be present in range")
	}
}

func TestIsWeekdayPresentInRange_MultipleNonePresent(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 9, 0, 0, 0, 0, time.UTC)   // Wednesday
	if IsWeekdayPresentInRange(start, end, time.Friday, time.Sunday) {
		t.Errorf("Expected Friday and Sunday to be absent in range")
	}
}

func TestIsWeekdayPresentInRange_ReverseOrder(t *testing.T) {
	end := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)   // Monday
	start := time.Date(2025, 7, 9, 0, 0, 0, 0, time.UTC) // Wednesday
	if !IsWeekdayPresentInRange(start, end, time.Tuesday) {
		t.Errorf("Expected Tuesday to be present in range (reverse order)")
	}
}

func TestIsWeekdayPresentInRange_EmptyWeekdays(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC)
	if IsWeekdayPresentInRange(start, end) {
		t.Errorf("Expected false when no weekdays provided")
	}
}
