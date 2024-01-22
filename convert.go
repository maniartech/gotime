package temporal

import "github.com/maniartech/temporal/internal/idf"

// ParseAndFormat is a utility function that takes a date-time string, a source
// format string (from), and a target format string (to). The function parses
// the input date-time string according to the source format, and then formats
// the parsed date-time value using the target format. The function returns the
// formatted date-time string and an error if the parsing or formatting fails.
//
// Example usage:
//
//	formattedDate, err := ParseAndFormat("2022-12-31", "yyyy-mm-dd", "dt mmmm, yyyy")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(formattedDate) // 31st December, 2022
func Convert(value, fromLayout, toLayout string) (string, error) {
	return idf.Convert(value, fromLayout, toLayout)
}
