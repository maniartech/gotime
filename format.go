package gotime

import (
	"time"

	"github.com/maniartech/gotime/internal/nites"
)

// Format converts a time.Time value to a formatted string using the specified layout.
// If layout is empty, RFC3339 format is used by default.
//
// The layout uses NITES (Natural and Intuitive Time Expression Syntax) format
// specifiers like "yyyy-mm-dd" instead of Go's reference time format.
//
// Example:
//
//	formatted := Format(time.Now(), "yyyy-mm-dd")
//	// formatted: "2025-07-08"
//
//	formatted = Format(time.Now(), "mmmm dd, yyyy")
//	// formatted: "July 08, 2025"
func Format(dt time.Time, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return nites.Format(dt, layout)
}

// FormatUnix converts Unix time (seconds and nanoseconds) to a formatted string
// using the specified layout. If layout is empty, RFC3339 format is used by default.
//
// Example:
//
//	formatted := FormatUnix(1609459200, 0, "yyyy-mm-dd")
//	// formatted: "2021-01-01"
func FormatUnix(sec int64, nsec int64, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return nites.Format(time.Unix(sec, nsec), layout)
}

// FormatTimestamp converts a Unix timestamp (seconds) to a formatted string
// using the specified layout. If layout is empty, RFC3339 format is used by default.
//
// Example:
//
//	formatted := FormatTimestamp(1609459200, "yyyy-mm-dd")
//	// formatted: "2021-01-01"
func FormatTimestamp(timestamp int64, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return nites.Format(time.Unix(timestamp, 0), layout)
}
