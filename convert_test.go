package dateutils

import (
	"time"
	"testing"
)

func TestConvertFormat(t *testing.T) {
	converted := ConvertFormat("dd-mm-yyyy")

	if converted != "02-01-2006" {
		t.Error("Expected 02-01-2006, got ", converted)
	}

	converted = ConvertFormat("w-ww")
	if converted != "Mon-Monday" {
		t.Error("Expected Mon-Monday, got ", converted)
	}

	converted = ConvertFormat("hhh-hh-h")
	if converted != "15-03-3" {
		t.Error("Expected 15-03-3, got ", converted)
	}

	converted = ConvertFormat("ww-MM-yyyy")
	if converted != "Monday-January-2006" {
		t.Error("Expected Monday-January-2006, got ", converted)
	}

	converted = ConvertFormat("i-s")
	if converted != "4-5" {
		t.Error("Expected 4-5, got ", converted)
	}

	converted = ConvertFormat("ss-ii")
	if converted != "05-04" {
		t.Error("Expected 05-04, got ", converted)
	}

	converted = ConvertFormat("ss-u")
	if converted != "05-000000" {
		t.Error("Expected 05-000000, got ", converted)
	}

	converted = ConvertFormat("dd-z-hh")
	if converted != "02-±0700-03" {
		t.Error("Expected 02-±0700-03, got ", converted)
	}

	converted = ConvertFormat("zh-d")
	if converted != "±07-2" {
		t.Error("Expected ±07-2, got ", converted)
	}

	converted = ConvertFormat("w-zzz")
	if converted != "Mon-MST" {
		t.Error("Expected Mon-MST, got ", converted)
	}

	converted = ConvertFormat("yy-zz")
	if converted != "06-±07:00" {
		t.Error("Expected 06-±07:00, got ", converted)
	}

	converted = ConvertFormat("zzzz-ss")
	if converted != "GMT-07:00-05" {
		t.Error("Expected GMT-07:00-05, got ", converted)
	}

	converted = ConvertFormat("M-ddd-yy")
	if converted != "Jan-002-06" {
		t.Error("Expected Jan-002-06, got ", converted)
	}

	converted = ConvertFormat("a-h")
	if converted != "pm-3" {
		t.Error("Expected pm-3, got ", converted)
	}

	converted = ConvertFormat("w-A")
	if converted != "Mon-PM" {
		t.Error("Expected Mon-PM, got ", converted)
	}

	converted = ConvertFormat("M-yy")
	if converted != "Jan-06" {
		t.Error("Expected Jan-06, got ", converted)
	}

	converted = ConvertFormat(time.Layout)
	if converted != "01/02 03:04:05PM '06 -0700" {
		t.Error("Expected 01/02 03:04:05PM '06 -0700, got ", converted)
	}

	converted = ConvertFormat(time.UnixDate)
	if converted != "Mon Jan _2 15:04:05 MST 2006" {
		t.Error("Expected Mon Jan _2 15:04:05 MST 2006, got ", converted)
	}
}

func BenchmarkConvertFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertFormat("dd-mm-yyyy")
	}
}
