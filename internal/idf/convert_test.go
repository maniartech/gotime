package idf_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/temporal/internal/idf"
)

func TestConvertLayoutA(t *testing.T) {
	c := idf.Format(time.Now(), time.Kitchen)
	fmt.Println(c)
	// temporal.TimeAgo(time.Now())
}

func TestConvert(t *testing.T) {
	date, err := idf.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss.000000")
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if date != "2012-06-03 12:00:AM:00.123000" {
		t.Error("Expected 2012-06-03 12:00:AM:00.123000, got ", date)
	}
}

// func TestConvertWithCache(t *testing.T) {
// 	// Enable cache
// 	temporal.Options.EnableCache()

// 	// Convert date
// 	f := "yyyy-mmm-dd hhh:ii:ss.999"
// 	f1 := temporal.convertLayout(f)
// 	f2 := temporal.convertLayout(f)

// 	if f1 != f2 {
// 		t.Error("Expected ", f1, ", got ", f2)
// 	}
// }

// func TestconvertLayout(t *testing.T) {
// 	// Convert format dd-mm-yyyy to to Go format.
// 	converted := temporal.convertLayout("dd-mm-yyyy")[0]
// 	if converted != "02-01-2006" {
// 		t.Error("Expected 02-01-2006, got ", converted)
// 	}

// 	// Convert format w-ww to to Go format.
// 	converted = temporal.convertLayout("www-wwww")[0]
// 	if converted != "Mon-Monday" {
// 		t.Error("Expected Mon-Monday, got ", converted)
// 	}

// 	// Convert format hhh-hh-h to to Go format.
// 	converted = temporal.convertLayout("hhh-hh-h")[0]
// 	if converted != "15-03-3" {
// 		t.Error("Expected 15-03-3, got ", converted)
// 	}

// 	// Convert format ww-MM-yyyy to to Go format.
// 	converted = temporal.convertLayout("wwww-mmmm-yyyy")[0]
// 	if converted != "Monday-January-2006" {
// 		t.Error("Expected Monday-January-2006, got ", converted)
// 	}

// 	// Convert format i-s to Go format.
// 	converted = temporal.convertLayout("i-s")[0]
// 	if converted != "4-5" {
// 		t.Error("Expected 4-5, got ", converted)
// 	}

// 	// Convert format ii-ss to Go format.
// 	converted = temporal.convertLayout("ss-ii")[0]
// 	if converted != "05-04" {
// 		t.Error("Expected 05-04, got ", converted)
// 	}

// 	// Convert format ss-u to Go format.
// 	converted = temporal.convertLayout("ss-000000")[0]
// 	if converted != "05-000000" {
// 		t.Error("Expected 05-000000, got ", converted)
// 	}

// 	// Convert format dd-z-hh to Go format.
// 	converted = temporal.convertLayout("dd-z-hh")[0]
// 	if converted != "02-±07:00-03" {
// 		t.Error("Expected 02-±0700-03, got ", converted)
// 	}

// 	// Convert format zh-d to Go format.
// 	converted = temporal.convertLayout("zh-d")[0]
// 	if converted != "±07-2" {
// 		t.Error("Expected ±07-2, got ", converted)
// 	}

// 	// Convert format w-zzz to Go format.
// 	converted = temporal.convertLayout("www-zzz")[0]
// 	if converted != "Mon-MST" {
// 		t.Error("Expected Mon-MST, got ", converted)
// 	}

// 	// Convert format yy-zz to Go format.
// 	converted = temporal.convertLayout("yy-z")[0]
// 	if converted != "06-±07:00" {
// 		t.Error("Expected 06-±07:00, got ", converted)
// 	}

// 	// Convert format zzzz-ss to Go format.
// 	converted = temporal.convertLayout("zzzz-ss")[0]
// 	if converted != "GMT-07:00-05" {
// 		t.Error("Expected GMT-07:00-05, got ", converted)
// 	}

// 	// Convert format M-ddd-yy to Go format.
// 	converted = temporal.convertLayout("mmm-ddd-yy")[0]
// 	if converted != "Jan-002-06" {
// 		t.Error("Expected Jan-002-06, got ", converted)
// 	}

// 	// Convert format a-h to Go format.
// 	converted = temporal.convertLayout("a-h")[0]
// 	if converted != "pm-3" {
// 		t.Error("Expected pm-3, got ", converted)
// 	}

// 	// Convert format w-A to Go format.
// 	converted = temporal.convertLayout("www-aa")[0]
// 	if converted != "Mon-PM" {
// 		t.Error("Expected Mon-PM, got ", converted)
// 	}

// 	// Convert format M-yy-m to Go format.
// 	converted = temporal.convertLayout("mmm-yy")[0]
// 	if converted != "Jan-06" {
// 		t.Error("Expected Jan-06, got ", converted)
// 	}

// 	// Convert format time.Layout format to Go format.
// 	converted = temporal.convertLayout(time.Layout)[0]
// 	if converted != "01/02 03:04:05PM '06 -0700" {
// 		t.Error("Expected 01/02 03:04:05PM '06 -0700, got ", converted)
// 	}

// 	// Convert format time.UnixDate format to Go format.
// 	converted = temporal.convertLayout(time.UnixDate)[0]
// 	if converted != "Mon Jan _2 15:04:05 MST 2006" {
// 		t.Error("Expected Mon Jan _2 15:04:05 MST 2006, got ", converted)
// 	}

// 	// Convert 12-01-01 from format yyyy-mm-dd to dd/mm/yyyy format.
// 	date, _ := temporal.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")
// 	if date != "01/01/2012" {
// 		t.Error("Expected 01/01/2012, got ", date)
// 	}

// 	// Convert 01/24/1984 from format mm/dd/yyyy to dd-mm-yyyy format.
// 	date, _ = temporal.Convert("01/24/1984", "mm/dd/yyyy", "dd-mm-yyyy")
// 	if date != "24-01-1984" {
// 		t.Error("Expected 24-01-1984, got ", date)
// 	}

// 	// Convert 17/09/1991 from format dd/mm/yyyy to dd/mm/yyyy format.
// 	date, _ = temporal.Convert("17/09/1991", "dd/mm/yyyy", "dd/mm/yyyy")
// 	if date != "17/09/1991" {
// 		t.Error("Expected 19/09/1991, got ", date)
// 	}

// 	// Returns empty string and error.
// 	date, err := temporal.Convert(time.RFC3339, time.RFC3339, "dd-mm-yyyy")
// 	if date != "" && err == nil {
// 		t.Errorf("Expected %v, got nil", err)
// 	}

// 	// Test Microseconds
// 	date, err = temporal.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss.000000")
// 	if err != nil {
// 		t.Error("Expected no error, got ", err)
// 	}

// 	if date != "2012-06-03 12:00:AM:00.123000" {
// 		t.Error("Expected 2012-06-03 12:00:AM:00.123000, got ", date)
// 	}

// }

// func BenchmarkconvertLayout(b *testing.B) {
// 	// Benchmarking for temporal.convertLayout function
// 	for i := 0; i < b.N; i++ {
// 		temporal.convertLayout("dd-mm-yyyy")
// 	}
// }

func BenchmarkConvert(b *testing.B) {
	// Benchmarking for Convert function
	for i := 0; i < b.N; i++ {
		idf.Convert("01/24/1984", "dd/mm/yyyy", "yyyy/dd/mm")
	}
}
