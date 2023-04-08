package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
)

func TestConvert(t *testing.T) {
	date, err := temporal.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss.000000")
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if date != "2012-06-03 12:00:AM:00.123000" {
		t.Error("Expected 2012-06-03 12:00:AM:00.123000, got ", date)
	}
}

func TestConvertWithCache(t *testing.T) {
	// Enable cache
	temporal.Options.EnableCache()

	// Convert date
	f := "yyyy-mmm-dd hhh:ii:ss.999"
	f1 := temporal.ConvertFormat(f)
	f2 := temporal.ConvertFormat(f)

	if f1 != f2 {
		t.Error("Expected ", f1, ", got ", f2)
	}
}

func TestConvertFormat(t *testing.T) {
	// Convert format dd-mm-yyyy to to Go format.
	converted := temporal.ConvertFormat("dd-mm-yyyy")
	if converted != "02-01-2006" {
		t.Error("Expected 02-01-2006, got ", converted)
	}

	// Convert format w-ww to to Go format.
	converted = temporal.ConvertFormat("www-wwww")
	if converted != "Mon-Monday" {
		t.Error("Expected Mon-Monday, got ", converted)
	}

	// Convert format hhh-hh-h to to Go format.
	converted = temporal.ConvertFormat("hhh-hh-h")
	if converted != "15-03-3" {
		t.Error("Expected 15-03-3, got ", converted)
	}

	// Convert format ww-MM-yyyy to to Go format.
	converted = temporal.ConvertFormat("wwww-mmmm-yyyy")
	if converted != "Monday-January-2006" {
		t.Error("Expected Monday-January-2006, got ", converted)
	}

	// Convert format i-s to Go format.
	converted = temporal.ConvertFormat("i-s")
	if converted != "4-5" {
		t.Error("Expected 4-5, got ", converted)
	}

	// Convert format ii-ss to Go format.
	converted = temporal.ConvertFormat("ss-ii")
	if converted != "05-04" {
		t.Error("Expected 05-04, got ", converted)
	}

	// Convert format ss-u to Go format.
	converted = temporal.ConvertFormat("ss-000000")
	if converted != "05-000000" {
		t.Error("Expected 05-000000, got ", converted)
	}

	// Convert format dd-z-hh to Go format.
	converted = temporal.ConvertFormat("dd-z-hh")
	if converted != "02-±07:00-03" {
		t.Error("Expected 02-±0700-03, got ", converted)
	}

	// Convert format zh-d to Go format.
	converted = temporal.ConvertFormat("zh-d")
	if converted != "±07-2" {
		t.Error("Expected ±07-2, got ", converted)
	}

	// Convert format w-zzz to Go format.
	converted = temporal.ConvertFormat("www-zzz")
	if converted != "Mon-MST" {
		t.Error("Expected Mon-MST, got ", converted)
	}

	// Convert format yy-zz to Go format.
	converted = temporal.ConvertFormat("yy-z")
	if converted != "06-±07:00" {
		t.Error("Expected 06-±07:00, got ", converted)
	}

	// Convert format zzzz-ss to Go format.
	converted = temporal.ConvertFormat("zzzz-ss")
	if converted != "GMT-07:00-05" {
		t.Error("Expected GMT-07:00-05, got ", converted)
	}

	// Convert format M-ddd-yy to Go format.
	converted = temporal.ConvertFormat("mmm-ddd-yy")
	if converted != "Jan-002-06" {
		t.Error("Expected Jan-002-06, got ", converted)
	}

	// Convert format a-h to Go format.
	converted = temporal.ConvertFormat("a-h")
	if converted != "pm-3" {
		t.Error("Expected pm-3, got ", converted)
	}

	// Convert format w-A to Go format.
	converted = temporal.ConvertFormat("www-aa")
	if converted != "Mon-PM" {
		t.Error("Expected Mon-PM, got ", converted)
	}

	// Convert format M-yy-m to Go format.
	converted = temporal.ConvertFormat("mmm-yy")
	if converted != "Jan-06" {
		t.Error("Expected Jan-06, got ", converted)
	}

	// Convert format time.Layout format to Go format.
	converted = temporal.ConvertFormat(time.Layout)
	if converted != "01/02 03:04:05PM '06 -0700" {
		t.Error("Expected 01/02 03:04:05PM '06 -0700, got ", converted)
	}

	// Convert format time.UnixDate format to Go format.
	converted = temporal.ConvertFormat(time.UnixDate)
	if converted != "Mon Jan _2 15:04:05 MST 2006" {
		t.Error("Expected Mon Jan _2 15:04:05 MST 2006, got ", converted)
	}

	// Convert 12-01-01 from format yyyy-mm-dd to dd/mm/yyyy format.
	date, _ := temporal.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")
	if date != "01/01/2012" {
		t.Error("Expected 01/01/2012, got ", date)
	}

	// Convert 01/24/1984 from format mm/dd/yyyy to dd-mm-yyyy format.
	date, _ = temporal.Convert("01/24/1984", "mm/dd/yyyy", "dd-mm-yyyy")
	if date != "24-01-1984" {
		t.Error("Expected 24-01-1984, got ", date)
	}

	// Convert 17/09/1991 from format dd/mm/yyyy to dd/mm/yyyy format.
	date, _ = temporal.Convert("17/09/1991", "dd/mm/yyyy", "dd/mm/yyyy")
	if date != "17/09/1991" {
		t.Error("Expected 19/09/1991, got ", date)
	}

	// Returns empty string and error.
	date, err := temporal.Convert(time.RFC3339, time.RFC3339, "dd-mm-yyyy")
	if date != "" && err == nil {
		t.Errorf("Expected %v, got nil", err)
	}

	// Test Microseconds
	date, err = temporal.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss.000000")
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if date != "2012-06-03 12:00:AM:00.123000" {
		t.Error("Expected 2012-06-03 12:00:AM:00.123000, got ", date)
	}

}

func BenchmarkConvertFormat(b *testing.B) {
	// Benchmarking for temporal.ConvertFormat function
	for i := 0; i < b.N; i++ {
		temporal.ConvertFormat("dd-mm-yyyy")
	}
}

func BenchmarkConvert(b *testing.B) {
	// Benchmarking for Convert function
	for i := 0; i < b.N; i++ {
		temporal.Convert("01/24/1984", "dd/mm/yyyy", "yyyy/dd/mm")
	}
}
