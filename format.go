package temporal

import (
	"time"
)

type Formatter func(time.Time, string) (string, error)
type Parseer func(string, string) (time.Time, error)

// Format converts a simple format string to a time.Time format string. It
// accepts a simple format string such as "yyyy-mm-dd" and returns a converted
// format string such as "2006-01-02".
func Format(dt time.Time, format string) (string, error) {
	convertedFormat := ConvertFormat(format)

	return dt.Format(convertedFormat), nil
}

func FormatStr(dt string, format string, inputFormat ...string) (string, error) {
	convertedFormat := ConvertFormat(format)
	ip := "yyyy-mm-ddThh:mm:ssZ"
	if len(inputFormat) > 0 {
		ip = inputFormat[0]
	}
	inputFormatFormatted := ConvertFormat(ip)

	// Convert string to time.Time using RFC3339 format
	t, err := time.Parse(inputFormatFormatted, dt)
	if err != nil {
		return "", err
	}
	return t.Format(convertedFormat), nil
}

func FormatInt(dt int64, format string) (string, error) {
	convertedFormat := ConvertFormat(format)

	return time.Unix(dt, 0).Format(convertedFormat), nil
}

func FormatUint(dt uint64, format string) (string, error) {
	convertedFormat := ConvertFormat(format)

	return time.Unix(int64(dt), 0).Format(convertedFormat), nil
}
