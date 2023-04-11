package temporal

import (
	"time"

	"github.com/maniartech/temporal/internal/idf"
	"github.com/maniartech/temporal/options"
)

// Format converts a simple format string to a time.Time format string. It
// accepts a simple format string such as "yyyy-mm-dd" and returns a converted
// format string such as "2006-01-02".
func Format(dt time.Time, layout string) string {
	return options.DefaultConverter().Format(dt, layout)
}

func FormatUnix(sec int64, nsec int64, layout string) string {
	return options.DefaultConverter().Format(
		time.Unix(sec, nsec), layout,
	)
}

func FormatTimestamp(timestamp int64, layout string) string {
	return options.DefaultConverter().Format(time.Unix(timestamp, 0), layout)
}

// Converts the date-time from one string format to another.
func Convert(dt, from, to string) (string, error) {
	return idf.Convert(dt, from, to)
}
