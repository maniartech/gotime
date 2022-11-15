package dateutils

import (
	"strings"
	"time"
)

var builtInFormats map[string]bool = map[string]bool{

	time.Layout:      true,
	time.ANSIC:       true,
	time.UnixDate:    true,
	time.RubyDate:    true,
	time.RFC822:      true,
	time.RFC822Z:     true,
	time.RFC850:      true,
	time.RFC1123:     true,
	time.RFC1123Z:    true,
	time.RFC3339:     true,
	time.RFC3339Nano: true,

	// Handy time stamps.
	time.Kitchen:    true,
	time.Stamp:      true,
	time.StampMilli: true,
	time.StampMicro: true,
	time.StampNano:  true,
}

// ConvertFormat converts this library datetime format to a go format.
// It loops through each character of the supplied format string
// and checks if it is a valid format character. If it is, it
// converts it to the go format.
// The function convert following format:
// yy     -> 06         Two digit year with leading zero
// yyyy   -> 2006       Four digit year
// m      -> 1          Month without leading zero
// mm     -> 01         Month in two digits with leading zero
// mmm    -> Jan        Month in short name
// mmmm   -> January    Month in full name
// d      -> 2          Day without leading zero
// dd     -> 02         Day in two digits with leading zero
// ddd    -> 002        Zero padded day of year
// w      -> 1          Three letter weekday name
// ww     -> Monday     Full weekday name
// h      -> 3          Hour in 12 hour format without leading zero
// hh     -> 03         Hour in 12 hour format with leading zero
// H      -> 3          Hour in 24 hour format without leading zero
// HH     -> 15         Hour in 24 hour format with leading zero
// a      -> pm         am/pm
// A      -> PM         AM/PM
// ii     -> 04         Minute with leading zero
// i      -> 4          Minute without leading zero
// ss     -> 05         Second with leading zero
// s      -> 5          Second without leading zero
// u      -> 000000     Microsecond
// z      -> ±0700      UTC offset
// zh     -> ±07        Numeric timezone hour with hours only
// zz     -> ±07:00     UTC offset with colon
// zzz    -> MST        Timezone abbreviation
// zzzz   -> GMT-07:00  Timezone in long format
func ConvertFormat(f string) string {

	// If the format is a built-in format, return as is.
	if ok := builtInFormats[f]; ok {
		return f
	}

	// Otherwise, convert the format.
	to := strings.Builder{}

	// getNext returns the next characters until the
	// next count characters are reached.
	getNext := func(i int, c int) string {
		start := i + 1
		end := i + c + 1
		if end <= len(f) {
			return f[start:end]
		}
		return ""
	}

	i := 0
	l := len(f)
	for i < l {
		c := rune(f[i])

		if c == 'd' {
			if getNext(i, 2) == "dd" {
				to.WriteString("002")
				i += 2
			} else if getNext(i, 1) == "d" {
				to.WriteString("02")
				i++
			} else {
				to.WriteString("2")
			}
		} else if c == 'w' {
			if getNext(i, 1) == "w" {
				to.WriteString("Monday")
				i++
			} else {
				to.WriteString("Mon")
			}
		} else if c == 'm' {
			if getNext(i, 1) == "m" {
				to.WriteString("01")
				i++
			} else {
				to.WriteString("1")
			}
		} else if c == 'M' {
			if getNext(i, 1) == "M" {
				to.WriteString("January")
				i++
			} else {
				to.WriteString("Jan")
			}
		} else if c == 'y' {
			if getNext(i, 3) == "yyy" {
				to.WriteString("2006")
				i += 3
			} else if getNext(i, 1) == "y" {
				to.WriteString("06")
				i++
			}
		} else if c == 'h' {
			if getNext(i, 2) == "hh" {
				to.WriteString("15")
				i += 2
			} else if getNext(i, 1) == "h" {
				to.WriteString("03")
				i++
			} else {
				to.WriteString("3")
			}
		} else if c == 'a' {
			to.WriteString("pm")
		} else if c == 'A' {
			to.WriteString("PM")
		} else if c == 'H' {
			if getNext(i, 1) == "H" {
				to.WriteString("15")
				i++
			} else {
				to.WriteString("3")
			}
		} else if c == 'i' {
			if getNext(i, 1) == "i" {
				to.WriteString("04")
				i++
			} else {
				to.WriteString("4")
			}
		} else if c == 's' {
			if getNext(i, 1) == "s" {
				to.WriteString("05")
				i++
			} else {
				to.WriteString("5")
			}
		} else if c == 'u' {
			to.WriteString("000000")
		} else if c == 'z' {
			if getNext(i, 3) == "zzz" {
				to.WriteString("GMT-07:00")
				i += 3
			} else if getNext(i, 2) == "zz" {
				to.WriteString("MST")
				i += 2
			} else if getNext(i, 1) == "z" {
				to.WriteString("±07:00")
				i++
			} else if getNext(i, 1) == "h" {
				to.WriteString("±07")
				i++
			} else {
				to.WriteString("±0700")
			}
		} else {
			to.WriteRune(c)
		}

		i += 1
	}

	return to.String()
}
