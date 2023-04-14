package idf

import (
	"fmt"
	"time"
)

// This file contains the IDF (Intuitive Date Format) conversion functions
// for the temporal package.

func Format(dt time.Time, layout string) string {
	convertedLayouts, _ := convertLayout(layout, false)
	return format(dt, convertedLayouts)
}

func format(dt time.Time, convertedLayouts []string) string {
	converted := make([]any, 0, len(convertedLayouts))
	for _, f := range convertedLayouts {
		if f == "dt" {
			switch s := dt.Day(); {
			case s == 1 || s == 21 || s == 31:
				converted = append(converted, fmt.Sprintf("%dst", s))
			case s == 2 || s == 22:
				converted = append(converted, fmt.Sprintf("%dnd", s))
			case s == 3 || s == 23:
				converted = append(converted, fmt.Sprintf("%drd", s))
			default:
				converted = append(converted, fmt.Sprintf("%dth", s))
			}
		} else {
			converted = append(converted, dt.Format(f))
		}
	}

	return fmt.Sprint(converted...)
}
