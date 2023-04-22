package utils

import "time"

const v120 int = 120
const v100 int = 100

// builtInLayouts is a map of built-in formats. It is used to
// check if the specified format is a built-in format.
// https://pkg.go.dev/time#pkg-constants
var BuiltInLayouts map[string]int = map[string]int{

	// Built-in formats.

	// Layout - 2006-01-02 15:04:05.999999999 -0700 MST
	time.Layout: v100,

	// ANSIC - Mon Jan _2 15:04:05 2006
	time.ANSIC: v100,

	// UnixDate - Mon Jan _2 15:04:05 MST 2006
	time.UnixDate: v100,

	// RubyDate - Mon Jan 02 15:04:05 -0700 2006
	time.RubyDate: v100,

	// RFC822 - 02 Jan 06 15:04 MST
	time.RFC822: v100,

	// RFC822Z - 02 Jan 06 15:04 -0700
	time.RFC822Z: v100,

	// RFC850 - Monday, 02-Jan-06 15:04:05 MST
	time.RFC850: v100,

	// RFC1123 - Mon, 02 Jan 2006 15:04:05 MST
	time.RFC1123: v100,

	// RFC1123Z - Mon, 02 Jan 2006 15:04:05 -0700
	time.RFC1123Z: v100,

	// RFC3339 - 2006-01-02T15:04:05Z07:00
	time.RFC3339: v100,

	// RFC3339Nano - 2006-01-02T15:04:05.999999999Z07:00
	time.RFC3339Nano: v100,

	// Kitchen - 3:04PM
	time.Kitchen: v100,

	// Handy time stamps.

	// Stamp - Jan _2 15:04:05
	time.Stamp: v100,

	// StampMilli - Jan _2 15:04:05.000
	time.StampMilli: v100,

	// StampMicro - Jan _2 15:04:05.000000
	time.StampMicro: v100,

	// StampNano - Jan _2 15:04:05.000000000
	time.StampNano: v100,

	// // DateTime - 2006-01-02 15:04:05
	// time.DateTime: v120,

	// // Date - 2006-01-02
	// time.DateOnly: v120,

	// // Time - 15:04:05
	// time.TimeOnly: v120,
}
