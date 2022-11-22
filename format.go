package dateutils

import (
	"fmt"
	"time"
)

// Format converts a simple format string to a time.Time format string. It
// accepts a simple format string such as "yyyy-mm-dd" and returns a converted
// format string such as "2006-01-02".
func Format[T DateTime](dt T, format string) (string, error) {
	convertedFormat := ConvertFormat(format)
	var dtAny interface{} = dt

	switch v := dtAny.(type) {
	case time.Time:
		return v.Format(convertedFormat), nil
	case string:
		// Convert string to time.Time using RFC3339 format
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return "", err
		}
		return t.Format(convertedFormat), nil
	case int64:
		return time.Unix(v, 0).Format(convertedFormat), nil
	case uint64:
		return time.Unix(int64(v), 0).Format(convertedFormat), nil
	default:
		return "", fmt.Errorf(InvalidType, dt)
	}
}
