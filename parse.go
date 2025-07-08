package gotime

import (
	"time"

	"github.com/maniartech/gotime/internal/nites"
)

// Parse parses a date-time string according to the specified layout format.
// The layout uses NITES (Natural and Intuitive Time Expression Syntax) format
// specifiers like "yyyy-mm-dd" instead of Go's reference time format.
//
// Example:
//
//	parsed, err := Parse("yyyy-mm-dd", "2022-12-31")
//	if err != nil {
//		// handle error
//	}
//	// parsed: 2022-12-31 00:00:00 +0000 UTC
func Parse(layout, value string) (time.Time, error) {
	return nites.Parse(layout, value)
}

// ParseInLocation parses a date-time string according to the specified layout format
// and returns the time in the specified location/timezone.
//
// Example:
//
//	ist := time.FixedZone("IST", 5*60*60+30*60)
//	parsed, err := ParseInLocation("yyyy-mm-dd", "2022-12-31", ist)
//	if err != nil {
//		// handle error
//	}
//	// parsed: 2022-12-31 00:00:00 +0530 IST
func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return nites.ParseInLocation(layout, value, loc)
}
