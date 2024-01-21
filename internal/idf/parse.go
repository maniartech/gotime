package idf

import (
	"time"
)

// Parse parses a date string and returns the time value it represents.
// It accepts a date string and a simple format string such as "yyyy-mm-dd".
func Parse(layout, value string, loc ...*time.Location) (time.Time, error) {
	convertedFormat, err := convertLayout(layout, true)
	if err != nil {
		return time.Time{}, err
	}

	if str, ok := convertedFormat.(string); ok {

		if len(loc) > 0 {
			return time.ParseInLocation(str, value, loc[0])
		}

		return time.Parse(str, value)
	}

	return time.Time{}, nil
}
