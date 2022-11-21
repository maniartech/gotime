package dateutils

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
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		date := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
		Format(date, "yyyy/mm/dd")
	}
}
