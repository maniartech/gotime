package gotime

import (
	"testing"
	"time"
)

func TestCountWeekdaysInRange_SingleWeek(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 13, 0, 0, 0, 0, time.UTC)  // Sunday
	counts := CountWeekdaysInRange(start, end)
	if counts.Monday != 1 || counts.Tuesday != 1 || counts.Wednesday != 1 || counts.Thursday != 1 || counts.Friday != 1 || counts.Saturday != 1 || counts.Sunday != 1 {
		t.Errorf("Expected 1 of each weekday, got: %+v", counts)
	}
}

func TestCountWeekdaysInRange_MultipleWeeks(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 20, 0, 0, 0, 0, time.UTC)  // Sunday (2 weeks)
	counts := CountWeekdaysInRange(start, end)
	if counts.Monday != 2 || counts.Tuesday != 2 || counts.Wednesday != 2 || counts.Thursday != 2 || counts.Friday != 2 || counts.Saturday != 2 || counts.Sunday != 2 {
		t.Errorf("Expected 2 of each weekday, got: %+v", counts)
	}
}

func TestCountWeekdaysInRange_PartialWeek(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 9, 0, 0, 0, 0, time.UTC)   // Wednesday
	counts := CountWeekdaysInRange(start, end)
	if counts.Monday != 1 || counts.Tuesday != 1 || counts.Wednesday != 1 {
		t.Errorf("Expected 1 each for Mon/Tue/Wed, got: %+v", counts)
	}
	if counts.Thursday != 0 || counts.Friday != 0 || counts.Saturday != 0 || counts.Sunday != 0 {
		t.Errorf("Expected 0 for Thu-Sun, got: %+v", counts)
	}
}

func TestCountWeekdaysInRange_ReverseOrder(t *testing.T) {
	end := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)   // Monday
	start := time.Date(2025, 7, 9, 0, 0, 0, 0, time.UTC) // Wednesday
	counts := CountWeekdaysInRange(start, end)
	if counts.Monday != 1 || counts.Tuesday != 1 || counts.Wednesday != 1 {
		t.Errorf("Expected 1 each for Mon/Tue/Wed, got: %+v", counts)
	}
}

func TestCountWeekdaysInRange_SingleDay(t *testing.T) {
	start := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
	counts := CountWeekdaysInRange(start, end)
	if counts.Monday != 1 {
		t.Errorf("Expected 1 Monday, got: %+v", counts)
	}
	if counts.Tuesday != 0 || counts.Wednesday != 0 || counts.Thursday != 0 || counts.Friday != 0 || counts.Saturday != 0 || counts.Sunday != 0 {
		t.Errorf("Expected 0 for other days, got: %+v", counts)
	}
}
