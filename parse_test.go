package dateutils

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	// Test case for parsing the date 24-01-1984
	format, _ := Parse("24-01-1984", "dd-mm-yyyy")
	correctTime := time.Date(1984, 1, 24, 0, 0, 0, 0, time.UTC)
	if !format.Equal(correctTime) {
		t.Errorf("Expected %v, got, %v", correctTime, format)
	}
}
