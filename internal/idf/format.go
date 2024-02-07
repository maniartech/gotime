package idf

import (
	"fmt"
	"strconv"
	"time"
)

// This file contains the IDF (Intuitive Date Format) conversion functions
// for the gotime package.

func Format(dt time.Time, layout string) string {
	convertedLayouts, _ := convertLayout(layout, false)
	if str, ok := convertedLayouts.(string); ok {
		return dt.Format(str)
	}

	return formatStrs(dt, convertedLayouts.([]string))
}

func formatStrs(dt time.Time, convertedLayouts []string) string {
	converted := make([]any, 0, len(convertedLayouts))
	for _, f := range convertedLayouts {

		ordinalItem := ""
		if f == "dt" {
			ordinalItem = strconv.Itoa(dt.Day())
		} else if f == "mt" {
			ordinalItem = strconv.Itoa(int(dt.Month()))
		}

		if ordinalItem != "" {
			switch s := dt.Day(); {
			case s == 1 || s == 21 || s == 31:
				converted = append(converted, ordinalItem+"st")
			case s == 2 || s == 22:
				converted = append(converted, ordinalItem+"nd")
			case s == 3 || s == 23:
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
