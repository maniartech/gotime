package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime"
)

func TestFormat(t *testing.T) {
	// Converting time.Time to Go format.
	date1 := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
	date1Formatted := gotime.Format(date1, "yyyy/mm/dd")
	if date1Formatted != "0018/06/04" {
		t.Errorf("Expected 0018/06/04, got, %s", date1Formatted)
	}

	// Converting string to Go format.
	date2Formatted, err := gotime.Convert("2001-01-01T15:04:05Z", `2006-01-02T15:04:05\Z`, "yyyy/mm/dd")
	if err != nil {
		t.Errorf("Expected no error, got, %s", err)
	}

	if date2Formatted != "2001/01/01" {
		t.Errorf("Expected 2001/01/01, got, %s", date2Formatted)
	}

	// Converting int64 to Go format.
	date3 := time.Date(12, 11, 1999, 0, 0, 0, 0, time.UTC)
	unixTime := date3.Unix()
	date3Formatted := gotime.FormatTimestamp(unixTime, "mm/dd/yyyy")
	if date3Formatted != "04/22/0018" {
		t.Errorf("Expected 04/22/0018, got %s, ", date3Formatted)
	}

	// Converting Layout to Go format.
	date4Formatted, _ := gotime.Parse("dd/mm/yy", "02/02/02")
	if date4Formatted != time.Date(2002, 2, 2, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Expected 2012/06/03, got %s, ", date4Formatted)
	}

	// Converting Unix timestamp to Go format.
	date5Formatted := gotime.FormatUnix(1234567890, 0, "yyyy/mm/dd")
	if date5Formatted != "2009/02/14" {
		t.Errorf("Expected 2009/02/14, got %s, ", date5Formatted)
	}
}

func BenchmarkFormat(b *testing.B) {
	// Benchmarking for gotime.Format function
	for i := 0; i < b.N; i++ {
		date := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
		gotime.Format(date, "yyyy/mm/dd")
	}
}
