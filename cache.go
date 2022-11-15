package dateutils

// This file contains constants that maps django date format to go date format

var cache = map[string]string{}

func ClearCache() {
	cache = map[string]string{}
}
