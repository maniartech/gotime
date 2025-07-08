package idfs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/internal/cache"
	"github.com/maniartech/gotime/internal/idfs"
	"github.com/maniartech/gotime/internal/utils"
)

func TestParseWithOrdinalsShouldError(t *testing.T) {
	// Clear cache to avoid pollution from other tests
	cache.Disable()
	cache.Enable()

	// Test with day ordinal (dt)
	_, err := idfs.Parse("dt mmm yyyy", "31st 12 2025")
	if err == nil {
		t.Error("Expected error when parsing with day ordinal (dt), but got none")
	}

	// Test with month ordinal (mt)
	_, err = idfs.Parse("dd mt yyyy", "31 12th 2025")
	if err == nil {
		t.Error("Expected error when parsing with month ordinal (mt), but got none")
	}

	// Test with both day and month ordinals
	_, err = idfs.Parse("dt mt yyyy", "31st 12th 2025")
	if err == nil {
		t.Error("Expected error when parsing with both day and month ordinals, but got none")
	}
}

func TestParse(t *testing.T) {
	// Test case for parsing the date 24-01-1984
	format, _ := idfs.Parse("dd-mm-yyyy", "24-01-1984")
	correctTime := time.Date(1984, 1, 24, 0, 0, 0, 0, time.UTC)
	if !format.Equal(correctTime) {
		t.Errorf("Expected %v, got, %v", correctTime, format)
	}
}

func TestParseWithLocation(t *testing.T) {
	// Test case for parsing the date 24-01-1984
	format, _ := idfs.ParseInLocation("dd-mm-yyyy", "24-01-1984", time.FixedZone("IST", 5.5*60*60))
	correctTime := time.Date(1984, 1, 24, 0, 0, 0, 0, time.FixedZone("IST", 5.5*60*60))
	if !format.Equal(correctTime) {
		t.Errorf("Expected %v, got, %v", correctTime, format)
	}
}

func BenchmarkParse(b *testing.B) {

	// Benchmarking for Parse function
	for i := 0; i < b.N; i++ {
		idfs.Parse("24-01-1984", "dd-mm-yyyy")
	}
}

func TestTrial(t *testing.T) {
	tm := time.Now().Add(8 * time.Hour)
	tmu := tm.UTC()

	// truncate time from tm and tmu and store them in tru and truu
	// tru := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	truu := time.Date(tmu.Year(), tmu.Month(), tmu.Day(), 0, 0, 0, 0, tmu.Location())

	fmt.Println(truu, truu.Local())
}

func TestParseErrorHandling(t *testing.T) {
	// Test invalid date format
	_, err := idfs.Parse("dd/mm/yyyy", "invalid-date")
	if err == nil {
		t.Error("Expected error with invalid date format, but got none")
	}

	// Test parsing with proper layout but invalid date values
	_, err = idfs.Parse("dd/mm/yyyy", "32/13/2025")
	if err == nil {
		t.Error("Expected error with invalid date values, but got none")
	}
}

func TestParseInLocationErrorHandling(t *testing.T) {
	// Test invalid date format
	_, err := idfs.ParseInLocation("dd/mm/yyyy", "invalid-date", time.UTC)
	if err == nil {
		t.Error("Expected error with invalid date format, but got none")
	}

	// Test parsing with proper layout but invalid date values
	_, err = idfs.ParseInLocation("dd/mm/yyyy", "32/13/2025", time.UTC)
	if err == nil {
		t.Error("Expected error with invalid date values, but got none")
	}
}

func TestParseEdgeCases(t *testing.T) {
	// Test when convertLayout returns error
	_, err := idfs.Parse("invalid-format-{{{", "01/01/2025")
	if err == nil {
		t.Error("Expected error with invalid format, but got none")
	}

	// Test with built-in layouts that return []string from convertLayout
	// This should trigger the "return time.Time{}, nil" path
	result, err := idfs.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Error("Expected no error for built-in layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when convertLayout returns []string for built-in layout")
	}
}

func TestParseInLocationEdgeCases(t *testing.T) {
	// Test when convertLayout returns error
	_, err := idfs.ParseInLocation("invalid-format-{{{", "01/01/2025", time.UTC)
	if err == nil {
		t.Error("Expected error with invalid format, but got none")
	}

	// Test with built-in layouts that return []string from convertLayout
	// This should trigger the "return time.Time{}, nil" path
	result, err := idfs.ParseInLocation(time.Kitchen, "3:04PM", time.UTC)
	if err != nil {
		t.Error("Expected no error for built-in layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when convertLayout returns []string for built-in layout")
	}
}

func TestParseBuiltInLayouts(t *testing.T) {
	// Test Parse with built-in layout (convertLayout returns []string)
	// This should return zero time because Parse expects a single string, not []string
	result, err := idfs.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Error("Expected no error for built-in layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when using built-in layout in Parse")
	}

	// Test ParseInLocation with built-in layout
	result, err = idfs.ParseInLocation(time.Kitchen, "3:04PM", time.UTC)
	if err != nil {
		t.Error("Expected no error for built-in layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when using built-in layout in ParseInLocation")
	}

	// Test with more built-in layouts to ensure we cover the []string path
	result, err = idfs.Parse(time.ANSIC, "Mon Jan _2 15:04:05 2006")
	if err != nil {
		t.Error("Expected no error for ANSIC layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when using ANSIC layout in Parse")
	}

	result, err = idfs.ParseInLocation(time.UnixDate, "Mon Jan _2 15:04:05 MST 2006", time.UTC)
	if err != nil {
		t.Error("Expected no error for UnixDate layout, but got:", err)
	}
	if !result.IsZero() {
		t.Error("Expected zero time when using UnixDate layout in ParseInLocation")
	}
}

func TestParseWithBuiltInLayout(t *testing.T) {
	// Test Parse with built-in layout
	// Built-in layouts return []string from convertLayout, so Parse should return zero time
	result, err := idfs.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Errorf("Expected no error with built-in layout, but got: %v", err)
	}
	// When convertLayout returns []string (not string), Parse returns zero time
	if !result.IsZero() {
		t.Errorf("Expected zero time when convertLayout returns []string, got %v", result)
	}
}

func TestParseInLocationWithBuiltInLayout(t *testing.T) {
	// Test ParseInLocation with built-in layout
	// Built-in layouts return []string from convertLayout, so ParseInLocation should return zero time
	result, err := idfs.ParseInLocation(time.RFC3339, "2006-01-02T15:04:05Z", time.UTC)
	if err != nil {
		t.Errorf("Expected no error with built-in layout, but got: %v", err)
	}
	// When convertLayout returns []string (not string), ParseInLocation returns zero time
	if !result.IsZero() {
		t.Errorf("Expected zero time when convertLayout returns []string, got %v", result)
	}
}

func TestParseWithGo120PlusLayouts(t *testing.T) {
	// Test with newer Go layouts (1.20+) to ensure []string path is covered
	if utils.RuntimeVersion >= 120 {
		// Test Parse with DateTime layout
		result, err := idfs.Parse(time.DateTime, "2006-01-02 15:04:05")
		if err != nil {
			t.Errorf("Expected no error with DateTime layout, but got: %v", err)
		}
		if !result.IsZero() {
			t.Errorf("Expected zero time with DateTime layout, got %v", result)
		}

		// Test ParseInLocation with DateOnly layout
		result, err = idfs.ParseInLocation(time.DateOnly, "2006-01-02", time.UTC)
		if err != nil {
			t.Errorf("Expected no error with DateOnly layout, but got: %v", err)
		}
		if !result.IsZero() {
			t.Errorf("Expected zero time with DateOnly layout, got %v", result)
		}

		// Test ParseInLocation with TimeOnly layout
		result, err = idfs.ParseInLocation(time.TimeOnly, "15:04:05", time.UTC)
		if err != nil {
			t.Errorf("Expected no error with TimeOnly layout, but got: %v", err)
		}
		if !result.IsZero() {
			t.Errorf("Expected zero time with TimeOnly layout, got %v", result)
		}
	}
}

func TestParseInLocationWithOrdinalsShouldError(t *testing.T) {
	// Clear cache to avoid pollution from other tests
	cache.Disable()
	cache.Enable()

	loc := time.UTC

	// Test with day ordinal (dt)
	_, err := idfs.ParseInLocation("dt mmm yyyy", "31st 12 2025", loc)
	if err == nil {
		t.Error("Expected error when parsing with day ordinal (dt) in location, but got none")
	}

	// Test with month ordinal (mt)
	_, err = idfs.ParseInLocation("dd mt yyyy", "31 12th 2025", loc)
	if err == nil {
		t.Error("Expected error when parsing with month ordinal (mt) in location, but got none")
	}

	// Test with both day and month ordinals
	_, err = idfs.ParseInLocation("dt mt yyyy", "31st 12th 2025", loc)
	if err == nil {
		t.Error("Expected error when parsing with both day and month ordinals in location, but got none")
	}
}
