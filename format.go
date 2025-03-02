package gotime

import (
	"errors"
	"time"

	"github.com/maniartech/gotime/internal/idfs"
)

// Format is a utility function that takes a time.Time value and a layout string
// as input, then converts the time value to a formatted string based on the
// layout. Returns an error if the layout is invalid.
//
// Example usage:
//
//	formattedDate, err := Format(time.Now(), "yyyy-mm-dd")
//	if err != nil {
//		log.Printf("Failed to format date: %v", err)
//		return
//	}
//	fmt.Println(formattedDate) // 2022-12-31
func Format(dt time.Time, layout string) (string, error) {
	if layout == "" {
		return "", errors.New("layout cannot be empty")
	}
	return idfs.Format(dt, layout), nil
}

// FormatUnix takes the Unix time in seconds and nanoseconds,
// as well as a layout string, and returns a formatted string
// based on the layout. Returns an error if the layout is invalid.
//
// Example usage:
//
//	formattedDate, err := FormatUnix(1609459200, 0, "yyyy-mm-dd")
//	if err != nil {
//		log.Printf("Failed to format date: %v", err)
//		return
//	}
//	fmt.Println(formattedDate) // 2021-01-01
func FormatUnix(sec int64, nsec int64, layout string) (string, error) {
	if layout == "" {
		return "", errors.New("layout cannot be empty")
	}
	return idfs.Format(time.Unix(sec, nsec), layout), nil
}

// FormatTimestamp takes a Unix timestamp in seconds and a layout string, then
// returns a formatted string based on the layout. Returns an error if the layout is invalid.
//
// Example usage:
//
//	formattedDate, err := FormatTimestamp(1609459200, "yyyy-mm-dd")
//	if err != nil {
//		log.Printf("Failed to format date: %v", err)
//		return
//	}
//	fmt.Println(formattedDate) // 2021-01-01
func FormatTimestamp(timestamp int64, layout string) (string, error) {
	if layout == "" {
		return "", errors.New("layout cannot be empty")
	}
	return idfs.Format(time.Unix(timestamp, 0), layout), nil
}
