package temporal

import (
	"errors"
	"time"
)

// Parse parses a date string and returns the time value it represents.
// It accepts a date string and a simple format string such as "yyyy-mm-dd".
func Parse(layout, value string) (time.Time, error) {
	convertedFormat := ConvertFormat(layout)
	if len(convertedFormat) > 1 {
		return time.Time{}, errors.New("Ordinals are not supported")
	}
	parsedTime, err := time.Parse(convertedFormat[0], value)
	return parsedTime, err
}
