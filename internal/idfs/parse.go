package idfs

import (
	"time"
)

// Parse parses a date string and returns the time value it represents.
// It accepts a date string and a simple format string such as "yyyy-mm-dd".
// The layout uses the intuitive date format (IDF) syntax, which is more
// human-readable than Go's reference time layout.
//
// Example:
//
//	time, err := Parse("dd-mm-yyyy", "24-01-1984")
func Parse(layout, value string) (time.Time, error) {
	convertedFormat, err := convertLayout(layout, true)
	if err != nil {
		return time.Time{}, err
	}

	if str, ok := convertedFormat.(string); ok {
		return time.Parse(str, value)
	}

	return time.Time{}, nil
}

// ParseInLocation parses a date string in the given location and returns the time value.
// It's like Parse but allows specifying a timezone location for the parsed time.
// The layout uses the intuitive date format (IDF) syntax.
//
// Example:
//
//	loc := time.FixedZone("IST", 5.5*60*60)
//	time, err := ParseInLocation("dd-mm-yyyy", "24-01-1984", loc)
func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	convertedFormat, err := convertLayout(layout, true)
	if err != nil {
		return time.Time{}, err
	}

	if str, ok := convertedFormat.(string); ok {
		return time.ParseInLocation(str, value, loc)
	}

	return time.Time{}, nil
}
