package nites

import (
	"errors"
	"strings"
	"time"

	"github.com/maniartech/gotime/internal/cache"
	"github.com/maniartech/gotime/internal/utils"
)

// Convert function converts a datetime from one string format to another.
// It takes the datetime string in the single format and converts it to the expected output.
// It returns an error when the format is not supported.
func Convert(dt string, from string, to string) (string, error) {
	if from == to {
		return dt, nil
	}

	// Convert the format to go format. While parsing
	// the from layout, it may return an error if the format
	// contains ordinals (mt, dt).
	fromConverted, err := convertLayout(from, true)
	if err != nil {
		return "", err
	}

	var fromLayout string
	switch v := fromConverted.(type) {
	case []string:
		fromLayout = v[0]
	default:
		fromLayout, _ = v.(string)
	}

	toLayout, _ := convertLayout(to, false) // ConvertLayout never returns an error when forParsing is false

	t, err := time.Parse(fromLayout, dt)
	if err != nil {
		return "", err
	}

	switch v := toLayout.(type) {
	case []string:
		return formatStrs(t, v), nil
	default:
		vstr, _ := v.(string)
		return t.Format(vstr), nil
	}
}

// convertLayout converts this library datetime format to a go format.
// It loops through each character of the supplied format string
// and checks if it is a valid format character. If it is, it
// converts it to the go format.
//
// The function convert following format:
// yy     -> 06         Two digit year with leading zero
// yyyy   -> 2006       Four digit year
// m      -> 1          Month without leading zero
// mm     -> 01         Month in two digits with leading zero
// mt    ->  1st        Month in ordinal format with leading zero (not supported during parsing)
// mmm    -> Jan        Month in short name
// mmmm   -> January    Month in full name
// d      -> 2          Day without leading zero
// dd     -> 02         Day in two digits with leading zero
// dt     -> 2nd        Day in ordinal format with leading zero (not supported during parsing)

// ddd    -> 002        Zero padded day of year
// www    -> Mon        Three letter weekday name
// wwww   -> Monday     Full weekday name

// h      -> 3          Hour in 12 hour format without leading zero
// hh     -> 03         Hour in 12 hour format with leading zero
// hhh    -> 15         Hour in 24 hour format without leading zero
// a      -> pm         am/pm (lowercase)
// aa     -> PM         AM/PM (uppercase)
// ii     -> 04         Minute with leading zero
// i      -> 4          Minute without leading zero
// ss     -> 05         Second with leading zero
// s      -> 5          Second without leading zero
// 0      -> 0          Microsecond with leading zero
// 9      -> 9          Microsecond without leading zero

// z      -> Z      		The Z literal represents UTC
// zz     -> MST        Timezone abbreviation
// o     -> ±07     		Timezone offset with leading zero (only hours)
// oo    -> ±0700       Timezone offset with leading zero without colon
// ooo   -> ±07:00      Timezone offset with leading zero with colon
func convertLayout(f string, forParsing bool) (interface{}, error) {
	// Built-in format, return as is
	if version, ok := utils.BuiltInLayouts[f]; ok {
		if utils.RuntimeVersion >= version {
			return []string{f}, nil
		}
	}

	// If the format is cached, return the cached value
	if v := cache.Get(f); v != nil {
		return v, nil
	}

	// Convert format to lower case for case insensitive matching
	var converted interface{}

	// Initialize a map of format conversions
	conversions := map[string][][]string{
		"y": {{"yyyy", "2006"}, {"yy", "06"}},
		"m": {{"mmmm", "January"}, {"mmm", "Jan"}, {"mm", "01"}, {"mt", ""}, {"m", "1"}},
		"d": {{"ddd", "002"}, {"dd", "02"}, {"db", "_2"}, {"dt", ""}, {"d", "2"}}, // dt for ordinals
		"w": {{"wwww", "Monday"}, {"www", "Mon"}},
		"h": {{"hhh", "15"}, {"hh", "03"}, {"h", "3"}},
		"a": {{"aa", "PM"}, {"a", "pm"}},
		"i": {{"ii", "04"}, {"i", "4"}},
		"s": {{"ss", "05"}, {"s", "5"}},

		// Timezone
		"z": {
			{"zz", "MST"},
			{"z", "Z"},
		},
		"o": {{"ooo", "-07:00"}, {"oo", "-0700"}, {"o", "-07"}},
	}

	// Initialize a new string builder
	to := strings.Builder{}

	// Loop through the input format
	i := 0
	for i < len(f) {
		c := f[i]

		// If c is uppercase, convert it to lowercase
		if c >= 'A' && c <= 'Z' {
			c += 32
		}

		// Check if the current character is an escape character
		if f[i] == '\\' {
			// Check if we're not at the end of the format string
			if i+1 < len(f) {
				// Append the next character as a literal, ignoring its format meaning
				to.WriteString(string(f[i+1]))
				i += 2
			} else {
				// We're at the end of the format string; append the escape character
				to.WriteString(string(f[i]))
				i++
			}
			continue
		}

		// Check if the current character is a valid format character
		conv, ok := conversions[string(c)]
		if !ok {
			// Not a valid format character, add as is
			to.WriteString(string(f[i]))
			i++
			continue
		}

		// Valid format character, check for the longest possible match
		for _, keyVal := range conv {
			key, val := keyVal[0], keyVal[1]
			iEnd := i + len(key)

			// Check if the len of key + i is less than the len of the format.
			if iEnd <= len(f) {
				if f[i:iEnd] == key {
					if val == "" {
						if forParsing {
							return nil, errors.New(errOrdinalsNotSupported)
						}
						if converted == nil {
							converted = []string{}
						}
						converted = append(converted.([]string), to.String()) // Append the converted format
						converted = append(converted.([]string), key)         // Append the value to the converted format
						to.Reset()
					}
					to.WriteString(val)
					i += len(key)
					goto MatchFound
				}
			}
		}

		// If we get here, we didn't find a match, add the character as is
		if i < len(f) {
			to.WriteString(string(f[i]))
			i++
		}
	MatchFound:
	}

	// Cache the converted format
	finalConvert := to.String()
	if converted == nil {
		cache.Set(f, finalConvert)
		return finalConvert, nil
	}

	converted = append(converted.([]string), finalConvert)

	cache.Set(f, converted)
	return converted, nil
}
