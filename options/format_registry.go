package options

import "time"

var converterRegistry map[string]FormatConverter
var defaultConverter FormatConverter

type FormatConverter interface {
	// Convert converts the datetime format to the Go's time.Time format.
	Format(dt time.Time, layout string) string

	// Parse parses the datetime format to the Go's time.Time format.
	Parser(layout, value string) (string, error)
}

func RegisterConverter(name string, converter FormatConverter) {
	converterRegistry[name] = converter
}

func DefaultConverter() FormatConverter {
	if defaultConverter == nil {
		panic("Default converter not set")
	}

	return defaultConverter
}
