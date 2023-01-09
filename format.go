package dateutils

import (
	"fmt"
	"time"
)

// Format converts a simple format string to a time.Time format string. It
// accepts a simple format string such as "yyyy-mm-dd" and returns a converted
// format string such as "2006-01-02".
func Format(dt any, format string) (string, error) {
	convertedFormat := ConvertFormat(format)

	switch dtType := dt.(type) {
	case time.Time:
		return dtType.Format(convertedFormat), nil
	case string:
		// Convert string to time.Time using RFC3339 format
		t, err := time.Parse(time.RFC3339, dtType)
		if err != nil {
			return "", err
		}
		return t.Format(convertedFormat), nil
	case int64:
		return time.Unix(dtType, 0).Format(convertedFormat), nil
	case uint64:
		return time.Unix(int64(dtType), 0).Format(convertedFormat), nil
	}

	return "", fmt.Errorf("Invalid type %T, expecting time.Time, string, int64 or uint64", dt)
}
