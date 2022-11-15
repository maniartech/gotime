package dateutils

import "time"

// Parse parses a date string and returns the time value it represents.
// It accepts a date string and a simple format string such as "yyyy-mm-dd".
func Parse(date string, format string) (time.Time, error) {
	convertedFormat := ConvertFormat(format)
	return time.Parse(convertedFormat, date)
}
