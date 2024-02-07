package gotime

import (
	"time"

	"github.com/maniartech/gotime/internal/idf"
)

// Parse is a utility function that takes a date-time string and a format string
// as input, then parses the date-time string according to the format. The
// function returns a time.Time value and an error if the parsing fails.
//
// Example usage:
//
//	parsedDate, err := Parse("yyyy-mm-dd", "2022-12-31")
//	if err != nil {
//	    panic(err)
//	}
//	fmt.Println(parsedDate) // 2022-12-31 00:00:00 +0000 UTC
func Parse(layout, value string) (time.Time, error) {
	return idf.Parse(layout, value)
}

// ParseInLocation is a utility function that takes a date-time string, a format string
// as input, then parses the date-time string according to the format. The function returns a time.Time value and an error if the parsing fails. It also takes a location as input and returns the time in that location.
//
// Example usage:
//
//	parsedDate, err := ParseInLocation("yyyy-mm-dd", "2022-12-31", time.FixedZone("IST", 5.5*60*60))
//	if err != nil {
//	    panic(err)
//	}
//	fmt.Println(parsedDate) // 2022-12-31 00:00:00 +0000 UTC
func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return idf.ParseInLocation(layout, value, loc)
}
