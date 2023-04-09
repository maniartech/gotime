package temporal

import (
	"fmt"
	"time"
)

type Formatter func(time.Time, string) (string, error)
type Parser func(string, string) (time.Time, error)

func Format(dt time.Time, layout string) string {
	convertedLayouts := ConvertFormat(layout)
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

// // Format converts a simple format string to a time.Time format string. It
// // accepts a simple format string such as "yyyy-mm-dd" and returns a converted
// // format string such as "2006-01-02".
// func Format(dt time.Time, format string) (string, error) {
// 	convertedLayouts := ConvertFormat(format)

// 	return dt.Format(convertedLayouts), nil
// }

// func FormatStr(dt string, format string, inputFormat ...string) (string, error) {
// 	convertedLayouts := ConvertFormat(format)
// 	ip := "yyyy-mm-ddThh:mm:ssZ"
// 	if len(inputFormat) > 0 {
// 		ip = inputFormat[0]
// 	}
// 	inputFormatFormatted := ConvertFormat(ip)

// 	// Convert string to time.Time using RFC3339 format
// 	t, err := time.Parse(inputFormatFormatted, dt)
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Format(convertedLayouts), nil
// }

// func FormatInt(dt int64, format string) (string, error) {
// 	convertedLayouts := ConvertFormat(format)

// 	return time.Unix(dt, 0).Format(convertedLayouts), nil
// }

// func FormatUint(dt uint64, format string) (string, error) {
// 	convertedLayouts := ConvertFormat(format)

// 	return time.Unix(int64(dt), 0).Format(convertedLayouts), nil
// }
