package temporal

// FormatConverter is an interface for converting specifed string datetime
// format that that supported by the datetime package to the Go's time.Time format.
// It basically translates the datetime format to the Go's time.Time format.
type FormatConverter interface {
	// Convert converts the datetime format to the Go's time.Time format.
	Convert(format string) (string, error)
}
