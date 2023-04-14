package idf

import (
	"time"
)

// Parse parses a date string and returns the time value it represents.
// It accepts a date string and a simple format string such as "yyyy-mm-dd".
func Parse(layout, value string) (time.Time, error) {
	convertedFormat, err := convertLayout(layout, true)
	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.Parse(convertedFormat[0], value)
	return parsedTime, err
}
