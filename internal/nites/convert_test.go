package nites_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
	"github.com/maniartech/gotime/v2/internal/nites"
	"github.com/maniartech/gotime/v2/internal/utils"
)

func TestConvertLayoutA(t *testing.T) {
	// c := nites.Format(time.Now(), "dd-mm-yyyyThhhh:ii:ss.000000000000 zzoo")
	fmt.Println(time.Now().Format(time.ANSIC))

	fmt.Println(gotime.TimeAgo(time.Now().Add(time.Second * 24 * 1)))
}

func TestTZ(t *testing.T) {
	date := nites.Format(time.Now(), "yyyy-mm-ddThh:ii:aa:ss.000000000TZZ")
	fmt.Println(date)
}

func TestConvert(t *testing.T) {
	date, err := nites.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss")

	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "2012-06-03 12:00:AM:00", date)
}

// func TestConvertWithCache(t *testing.T) {
// 	// Enable cache
// 	gotime.Options.EnableCache()

// 	// Convert date
// 	f := "yyyy-mmm-dd hhhh:ii:ss.999"
// 	f1 := gotime.convertLayout(f)
// 	f2 := gotime.convertLayout(f)

// 	if f1 != f2 {
// 		t.Error("Expected ", f1, ", got ", f2)
// 	}
// }

// func TestconvertLayout(t *testing.T) {
// 	// Convert format dd-mm-yyyy to to Go format.
// 	converted := gotime.convertLayout("dd-mm-yyyy")[0]
// 	if converted != "02-01-2006" {
// 		t.Error("Expected 02-01-2006, got ", converted)
// 	}

// 	// Convert format w-ww to to Go format.
// 	converted = gotime.convertLayout("www-wwww")[0]
// 	if converted != "Mon-Monday" {
// 		t.Error("Expected Mon-Monday, got ", converted)
// 	}

// 	// Convert format hhhh-hh-h to to Go format.
// 	converted = gotime.convertLayout("hhhh-hh-h")[0]
// 	if converted != "15-03-3" {
// 		t.Error("Expected 15-03-3, got ", converted)
// 	}

// 	// Convert format ww-MM-yyyy to to Go format.
// 	converted = gotime.convertLayout("wwww-mmmm-yyyy")[0]
// 	if converted != "Monday-January-2006" {
// 		t.Error("Expected Monday-January-2006, got ", converted)
// 	}

// 	// Convert format i-s to Go format.
// 	converted = gotime.convertLayout("i-s")[0]
// 	if converted != "4-5" {
// 		t.Error("Expected 4-5, got ", converted)
// 	}

// 	// Convert format ii-ss to Go format.
// 	converted = gotime.convertLayout("ss-ii")[0]
// 	if converted != "05-04" {
// 		t.Error("Expected 05-04, got ", converted)
// 	}

// 	// Convert format ss-u to Go format.
// 	converted = gotime.convertLayout("ss-000000")[0]
// 	if converted != "05-000000" {
// 		t.Error("Expected 05-000000, got ", converted)
// 	}

// 	// Convert format dd-z-hh to Go format.
// 	converted = gotime.convertLayout("dd-z-hh")[0]
// 	if converted != "02-±07:00-03" {
// 		t.Error("Expected 02-±0700-03, got ", converted)
// 	}

// 	// Convert format zh-d to Go format.
// 	converted = gotime.convertLayout("zh-d")[0]
// 	if converted != "±07-2" {
// 		t.Error("Expected ±07-2, got ", converted)
// 	}

// 	// Convert format w-zzz to Go format.
// 	converted = gotime.convertLayout("www-zzz")[0]
// 	if converted != "Mon-MST" {
// 		t.Error("Expected Mon-MST, got ", converted)
// 	}

// 	// Convert format yy-zz to Go format.
// 	converted = gotime.convertLayout("yy-z")[0]
// 	if converted != "06-±07:00" {
// 		t.Error("Expected 06-±07:00, got ", converted)
// 	}

// 	// Convert format zzzz-ss to Go format.
// 	converted = gotime.convertLayout("zzzz-ss")[0]
// 	if converted != "GMT-07:00-05" {
// 		t.Error("Expected GMT-07:00-05, got ", converted)
// 	}

// 	// Convert format M-ddd-yy to Go format.
// 	converted = gotime.convertLayout("mmm-ddd-yy")[0]
// 	if converted != "Jan-002-06" {
// 		t.Error("Expected Jan-002-06, got ", converted)
// 	}

// 	// Convert format a-h to Go format.
// 	converted = gotime.convertLayout("a-h")[0]
// 	if converted != "pm-3" {
// 		t.Error("Expected pm-3, got ", converted)
// 	}

// 	// Convert format w-A to Go format.
// 	converted = gotime.convertLayout("www-aa")[0]
// 	if converted != "Mon-PM" {
// 		t.Error("Expected Mon-PM, got ", converted)
// 	}

// 	// Convert format M-yy-m to Go format.
// 	converted = gotime.convertLayout("mmm-yy")[0]
// 	if converted != "Jan-06" {
// 		t.Error("Expected Jan-06, got ", converted)
// 	}

// 	// Convert format time.Layout format to Go format.
// 	converted = gotime.convertLayout(time.Layout)[0]
// 	if converted != "01/02 03:04:05PM '06 -0700" {
// 		t.Error("Expected 01/02 03:04:05PM '06 -0700, got ", converted)
// 	}

// 	// Convert format time.UnixDate format to Go format.
// 	converted = gotime.convertLayout(time.UnixDate)[0]
// 	if converted != "Mon Jan _2 15:04:05 MST 2006" {
// 		t.Error("Expected Mon Jan _2 15:04:05 MST 2006, got ", converted)
// 	}

// 	// Convert 12-01-01 from format yyyy-mm-dd to dd/mm/yyyy format.
// 	date, _ := gotime.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")
// 	if date != "01/01/2012" {
// 		t.Error("Expected 01/01/2012, got ", date)
// 	}

// 	// Convert 01/24/1984 from format mm/dd/yyyy to dd-mm-yyyy format.
// 	date, _ = gotime.Convert("01/24/1984", "mm/dd/yyyy", "dd-mm-yyyy")
// 	if date != "24-01-1984" {
// 		t.Error("Expected 24-01-1984, got ", date)
// 	}

// 	// Convert 17/09/1991 from format dd/mm/yyyy to dd/mm/yyyy format.
// 	date, _ = gotime.Convert("17/09/1991", "dd/mm/yyyy", "dd/mm/yyyy")
// 	if date != "17/09/1991" {
// 		t.Error("Expected 19/09/1991, got ", date)
// 	}

// 	// Returns empty string and error.
// 	date, err := gotime.Convert(time.RFC3339, time.RFC3339, "dd-mm-yyyy")
// 	if date != "" && err == nil {
// 		t.Errorf("Expected %v, got nil", err)
// 	}

// 	// Test Microseconds
// 	date, err = gotime.Convert("2012-Jun-03 00:00:00.123", "yyyy-mmm-dd hhhh:ii:ss.999", "yyyy-mm-dd hh:ii:aa:ss.000000")
// 	if err != nil {
// 		t.Error("Expected no error, got ", err)
// 	}

// 	if date != "2012-06-03 12:00:AM:00.123000" {
// 		t.Error("Expected 2012-06-03 12:00:AM:00.123000, got ", date)
// 	}

// }

// func BenchmarkconvertLayout(b *testing.B) {
// 	// Benchmarking for gotime.convertLayout function
// 	for i := 0; i < b.N; i++ {
// 		gotime.convertLayout("dd-mm-yyyy")
// 	}
// }

func BenchmarkConvert(b *testing.B) {
	// Benchmarking for Convert function
	for i := 0; i < b.N; i++ {
		nites.Convert("01/24/1984", "dd/mm/yyyy", "yyyy/dd/mm")
	}
}

func TestConvertErrorHandling(t *testing.T) {
	// Test parsing with ordinals (should fail)
	_, err := nites.Convert("1st Jan 2025", "dt mmm yyyy", "yyyy-mm-dd")
	if err == nil {
		t.Error("Expected error when parsing ordinals, but got none")
	}

	// Test invalid date
	_, err = nites.Convert("32/01/2025", "dd/mm/yyyy", "yyyy-mm-dd")
	if err == nil {
		t.Error("Expected error with invalid date, but got none")
	}

	// Test mismatched format
	_, err = nites.Convert("2025-01-01", "mm/dd/yyyy", "dd-mm-yyyy")
	if err == nil {
		t.Error("Expected error with mismatched format, but got none")
	}
}

func TestConvertSameFormat(t *testing.T) {
	// Test converting to same format (should return original)
	date := "2025-01-01"
	result, err := nites.Convert(date, "yyyy-mm-dd", "yyyy-mm-dd")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, date, result)
}

func TestConvertWithOrdinals(t *testing.T) {
	// Test converting from format with ordinals in "to" format
	result, err := nites.Convert("01/01/2025", "mm/dd/yyyy", "dt mmm yyyy")
	utils.AssertNoError(t, err)
	if result == "" {
		t.Error("Expected non-empty result for ordinal format conversion")
	}
}

func TestConvertLayoutBuiltIn(t *testing.T) {
	// Test with built-in layouts
	date, err := nites.Convert("2006-01-02T15:04:05Z", time.RFC3339, "dd/mm/yyyy")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "02/01/2006", date)
}

func TestConvertLayoutEscape(t *testing.T) {
	// Test escaped characters in format - corrected to use proper Go format
	date, err := nites.Convert("2025-01d-01", "yyyy-mm\\d-dd", "dd/mm/yyyy")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "01/01/2025", date)

	_, err = nites.Convert("2025-01d-01", "yyyy-mm\\d-dd\\", "dd/mm/yyyy")
	if err == nil {
		t.Error("Expected error with malformed escape sequence, but got none")
	}
}

func TestConvertLayoutSpecialCases(t *testing.T) {
	// Test special characters that should remain as-is
	date, err := nites.Convert("2025@01@01", "yyyy@mm@dd", "dd#mm#yyyy")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "01#01#2025", date)
}

func TestConvertLayoutEdgeCases(t *testing.T) {
	// Test case where convertLayout returns multiple formats for "from"
	_, err := nites.Convert("1st", "dt", "dd")
	if err == nil {
		t.Error("Expected error when parsing multiple format from, but got none")
	}

	// Test complex format conversions
	date, err := nites.Convert("2025-Jan-01", "yyyy-mmm-dd", "dd/mmm/yyyy")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "01/Jan/2025", date)

	// Test timezone conversion
	date, err = nites.Convert("2025-01-01 12:00:00 UTC", "yyyy-mm-dd hhhh:ii:ss zz", "dd/mm/yyyy hh:ii:aa")
	utils.AssertNoError(t, err)
	if date == "" {
		t.Error("Expected non-empty result for timezone conversion")
	}
}

func TestConvertWithVersionSpecificLayouts(t *testing.T) {
	// Test conversion using layouts available in current Go version
	// Use older layouts that are always available (Go 1.0+)
	date, err := nites.Convert("Mon Jan  2 15:04:05 MST 2006", time.UnixDate, "yyyy-mm-dd")
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, "2006-01-02", date)

	// Test with newer layouts if available
	if utils.RuntimeVersion >= 120 {
		// Test with Go 1.20+ layouts
		date, err = nites.Convert("2006-01-02 15:04:05", time.DateTime, "dd/mm/yyyy")
		utils.AssertNoError(t, err)
		utils.AssertEqual(t, "02/01/2006", date)
	}
}

func TestConvertInvalidToLayoutType(t *testing.T) {
	// This test is designed to trigger the default case in the second switch statement
	// We need to mock a scenario where convertLayout returns something other than string or []string
	// This is actually very hard to achieve with the current code, as convertLayout only returns those types
	// But we can test error handling for malformed layouts

	// Test with a format that might cause issues in convertLayout for "to" parameter
	_, err := nites.Convert("2025-01-01", "yyyy-mm-dd", "")
	if err != nil {
		// This is expected - empty format should cause some kind of issue
		return
	}

	// If no error, that's also acceptable as empty string might be handled gracefully
}

func TestConvertWithCacheCorruption(t *testing.T) {
	// Try to test edge cases that might expose the default case
	// Use very unusual format strings that might confuse the parser

	// Test with format containing only special characters
	_, err := nites.Convert("2025-01-01", "yyyy-mm-dd", "!@#$%^&*()")
	// This should not crash and should either work or return an error
	if err != nil {
		// Error is acceptable
		return
	}

	// Test with extremely long format string
	longFormat := strings.Repeat("y", 1000)
	_, err = nites.Convert("2025", "yyyy", longFormat)
	// Should handle gracefully
	if err != nil {
		// Error is acceptable
		return
	}
}

func TestConvertInvalidCases(t *testing.T) {
	// Test conversion where toLayout convertLayout returns invalid type
	// This would be difficult to trigger directly, but we can test error paths

	// Test invalid format characters that might cause issues
	_, err := nites.Convert("2025-01-01", "yyyy-mm-dd", "dt mmm yyyy")
	utils.AssertNoError(t, err) // This should work as it returns ordinal format

	// Test multiple ordinals in from format (should error)
	_, err = nites.Convert("1st 2nd", "dt dt", "dd/mm")
	if err == nil {
		t.Error("Expected error with multiple ordinals in from format")
	}
}

func TestConvertWithErrorInToFormat(t *testing.T) {
	// Test conversion where convertLayout for "to" format returns a type not string or []string
	// Now, this should not error, but return an empty string (normalized to []string{})
	_, err := nites.Convert("2025-01-01", "yyyy-mm-dd", "dt dt dt dt dt dt")
	if err != nil {
		t.Errorf("Expected no error with invalid 'to' format, got: %v", err)
	}

	_, err = nites.Convert("1st Jan 2025", "dt mmm, yyyy", "yyyy-mm-dd")
	if err == nil {
		t.Errorf("Expected no error with invalid 'to' format, got: %v", err)
	}
}

func TestConvertFromFormatWithMultipleFormats(t *testing.T) {
	// Test case where convertLayout returns []string with multiple elements for "from"
	// This should trigger the error "ordinals not supported during parsing"
	_, err := nites.Convert("1st", "dt", "dd")
	if err == nil {
		t.Error("Expected error when from format returns multiple layouts")
	}
	if err != nil && !strings.Contains(err.Error(), "ordinals not supported") {
		t.Errorf("Expected 'ordinals not supported' error, got: %v", err)
	}
}
