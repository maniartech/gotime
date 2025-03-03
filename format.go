package gotime

import (
	"time"

	"github.com/maniartech/gotime/internal/idfs"
)

// Format is a utility function that takes a time.Time value and a layout string
// as input, then converts the time value to a formatted string based on the
// layout. If the layout is empty, the function uses the RFC3339 layout by default.
//
// Example usage:
//
//	formattedDate := Format(time.Now(), "yyyy-mm-dd")
//	fmt.Println(formattedDate) // 2022-12-31
func Format(dt time.Time, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return idfs.Format(dt, layout)
}

// FormatUnix takes the Unix time in seconds and nanoseconds,
// as well as a layout string, and returns a formatted string
// based on the layout. If the layout is empty, the function uses
// the RFC3339 layout by default.
//
// Example usage:
//
//	formattedDate, err := FormatUnix(1609459200, 0, "yyyy-mm-dd")
//	if err != nil {
//		log.Printf("Failed to format date: %v", err)
//		return
//	}
//	fmt.Println(formattedDate) // 2021-01-01
func FormatUnix(sec int64, nsec int64, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return idfs.Format(time.Unix(sec, nsec), layout)
}

// FormatTimestamp takes a Unix timestamp in seconds and a layout string, then
// returns a formatted string based on the layout. If the layout is empty, the
// function uses the RFC3339 layout by default.
//
// Example usage:
//
//	formattedDate, err := FormatTimestamp(1609459200, "yyyy-mm-dd")
//	if err != nil {
//		log.Printf("Failed to format date: %v", err)
//		return
//	}
//	fmt.Println(formattedDate) // 2021-01-01
func FormatTimestamp(timestamp int64, layout string) string {
	if layout == "" {
		// Layout is RFC3339 by default
		layout = time.RFC3339
	}
	return idfs.Format(time.Unix(timestamp, 0), layout)
}
