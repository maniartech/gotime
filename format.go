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

// Convert is a utility function that takes a date-time string, a source
// format string (from), and a target format string (to). The function parses
// the input date-time string according to the source format, and then formats
// the parsed date-time value using the target format. The function returns the
// formatted date-time string and an error if the parsing or formatting fails.
//
// Example usage:
//
//	formattedDate, err := Convert("2022-12-31", "yyyy-mm-dd", "dt mmmm, yyyy")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(formattedDate) // 31st December, 2022
func Convert(value, fromLayout, toLayout string) (string, error) {
	return idf.Convert(value, fromLayout, toLayout)
}

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
