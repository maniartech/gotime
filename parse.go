package temporal

import (
	"time"

	"github.com/maniartech/temporal/internal/idf"
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
func Parse(layout, value string, loc ...*time.Location) (time.Time, error) {
	return idf.Parse(layout, value, loc...)
}
