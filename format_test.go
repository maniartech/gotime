package datetime

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// Converting time.Time to Go format.
	date1 := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
	date1Formatted, _ := Format(date1, "yyyy/mm/dd")
	if date1Formatted != "0018/06/04" {
		t.Errorf("Expected 0018/06/04, got, %s", date1Formatted)
	}

	// Converting string to Go format.
	date2Formatted, _ := Format("2001-01-01T15:04:05Z", "yyyy/mm/dd")
	if date2Formatted != "2001/01/01" {
		t.Errorf("Expected 2001/01/01, got, %s", date2Formatted)
	}

	// Converting int64 to Go format.
	date3 := time.Date(12, 11, 1999, 0, 0, 0, 0, time.UTC)
	unixTime := date3.Unix()
	date3Formatted, _ := Format(unixTime, "mm/dd/yyyy")
	if date3Formatted != "04/22/0018" {
		t.Errorf("Expected 04/22/0018, got %s, ", date3Formatted)
	}

	// Converting uint64 to Go format.
	date4 := time.Date(12, 11, 1999, 0, 0, 0, 0, time.UTC)
	unsignedUnixTime := uint64(date4.Unix())
	date4Formatted, _ := Format(unsignedUnixTime, "mm/dd/yyyy")
	if date4Formatted != "04/22/0018" {
		t.Errorf("Expected 04/22/0018, got %s, ", date4Formatted)
	}

	// Returns empty string and error.
	date5, err := Format("2001-01-01T15:04:05Z07:00", "yyyy/mm/dd")
	if date5 != "" && err == nil {
		t.Errorf("Expected %v, got nil", err)
	}

	// Returns empty string and error.
	date6, err := Format("2001-01-01T15:04:05Z07:00", "yyyy/mm/dd")
	if date6 != "" && err == nil {
		t.Errorf("Expected %v, got nil", err)
	}

	// Returns empty string and error.
	date7, err := Format("2001-01-01T15:04:05Z07:00", "yyyy/mm/dd")
	if date7 != "" && err == nil {
		t.Errorf("Expected %v, got nil", err)
	}

	// Supplying any other type then time.Time, string, int64, uint64
	// should return empty string and error.
	date8, err := Format(123, "yyyy/mm/dd")
	if date8 != "" && err == nil {
		t.Errorf("Expected %v, got nil", err)
	}
}

func BenchmarkFormat(b *testing.B) {
	// Benchmarking for Format function
	for i := 0; i < b.N; i++ {
		date := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
		Format(date, "yyyy/mm/dd")
	}
}
