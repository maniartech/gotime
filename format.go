package temporal

import (
	"time"

	"github.com/maniartech/temporal/internal/idf"
)

// Format is a utility function that takes a time.Time value and a layout string
// as input, then converts the time value to a formatted string based on the
// layout.
//
// Example usage:
//
//	formattedDate := Format(time.Now(), "yyyy-mm-dd")
func Format(dt time.Time, layout string) string {
	return idf.Format(dt, layout)
}

// FormatUnix takes the Unix time in seconds and nanoseconds,
// as well as a layout string, and returns a formatted string
// based on the layout.
//
// Example usage:
//
//	formattedDate := FormatUnix(1609459200, 0, "yyyy-mm-dd")
func FormatUnix(sec int64, nsec int64, layout string) string {
	return idf.Format(
		time.Unix(sec, nsec), layout,
	)
}

// FormatTimestamp takes a Unix timestamp in seconds and a layout string, then
// returns a formatted string based on the layout.
//
// Example usage:
//
//	formattedDate := FormatTimestamp(1609459200, "yyyy-mm-dd")
func FormatTimestamp(timestamp int64, layout string) string {
	return idf.Format(time.Unix(timestamp, 0), layout)
}
