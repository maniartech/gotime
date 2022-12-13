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
	dtType := fmt.Sprintf("%T", dtAny)

	if dtType == "time.Time" {
		return dtAny.(time.Time).Format(convertedFormat), nil
	} else if dtType == "string" {
		// Convert string to time.Time using RFC3339 format
		t, err := time.Parse(time.RFC3339, dtAny.(string))
		if err != nil {
			return "", err
		}
		return t.Format(convertedFormat), nil
	} else if dtType == "int64" {
		return time.Unix(dtAny.(int64), 0).Format(convertedFormat), nil
	}
	return time.Unix(int64(dtAny.(uint64)), 0).Format(convertedFormat), nil
}
