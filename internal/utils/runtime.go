package utils

import (
	"runtime"
	"strconv"
	"strings"
)

var RuntimeVersion int

func init() {

	// Calculate the current version of the Go runtime. It fetches
	// the version from the runtime.Version() function. It generates
	// the version in the form of 100, 119, 120 etc. where 100 is Go 1.0,
	// 119 is Go 1.19 and 120 is Go 1.20. This is done to make it easier
	// to compare versions.
	version := runtime.Version()
	version = strings.TrimPrefix(version, "go")

	// Split the version string into two parts,
	// the major and minor version.
	parts := strings.Split(version, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	RuntimeVersion = major*100 + minor
}
