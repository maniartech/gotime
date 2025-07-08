package nites

import (
	"fmt"
	"strconv"
	"time"
)

// This file contains the IDF (Intuitive Date Format) conversion functions
// for the gotime package.

// Format formats a time.Time value according to the given layout string.
// The layout uses intuitive date format (IDF) syntax instead of Go's reference time.
// Supports ordinal formatting for days (dt) and months (mt).
//
// Example formats:
//
//	"yyyy-mm-dd"     -> "2025-07-07"
//	"dt of mmm"      -> "7th of Jul"
//	"mt month"       -> "7th month"
//
// See convertLayout documentation for complete format specification.
func Format(dt time.Time, layout string) string {
	convertedLayouts, _ := convertLayout(layout, false)
	if str, ok := convertedLayouts.(string); ok {
		return dt.Format(str)
	}

	return formatStrs(dt, convertedLayouts.([]string))
}

// formatStrs formats a time using multiple layout strings and concatenates the results.
// This function is used when the layout contains ordinal formats (dt, mt) that require
// special processing beyond standard Go time formatting.
func formatStrs(dt time.Time, convertedLayouts []string) string {
	converted := make([]any, 0, len(convertedLayouts))
	for _, f := range convertedLayouts {

		ordinalItem := ""
		ordinalValue := 0
		switch f {
		case "dt":
			ordinalValue = dt.Day()
			ordinalItem = strconv.Itoa(ordinalValue)
		case "mt":
			ordinalValue = int(dt.Month())
			ordinalItem = strconv.Itoa(ordinalValue)
		}

		if ordinalItem != "" {
			switch ordinalValue {
			case 1, 21, 31:
				converted = append(converted, ordinalItem+"st")
			case 2, 22:
				converted = append(converted, ordinalItem+"nd")
			case 3, 23:
				converted = append(converted, ordinalItem+"rd")
			default:
				converted = append(converted, ordinalItem+"th")
			}
		} else {
			converted = append(converted, dt.Format(f))
		}
	}

	return fmt.Sprint(converted...)
}
