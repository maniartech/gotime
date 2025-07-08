package gotime

import "github.com/maniartech/gotime/internal/nites"

// Convert parses a date-time string according to the source format and
// reformats it using the target format. Both formats use NITES
// (Natural and Intuitive Time Expression Syntax) format specifiers.
//
// Example:
//	converted, err := Convert("2022-12-31", "yyyy-mm-dd", "dt mmmm, yyyy")
//	if err != nil {
//		// handle error
//	}
//	// converted: "31st December, 2022"
func Convert(value, fromLayout, toLayout string) (string, error) {
	return nites.Convert(value, fromLayout, toLayout)
}
