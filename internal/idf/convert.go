package idf

import (
	"errors"
	"strings"

	"github.com/maniartech/temporal/cache"
	"github.com/maniartech/temporal/internal/utils"
)

// Convert function converts a datetime from one string format to another.
// It takes the datetime string in the single format and converts it to the expected output.
// It returns an error when the format is not supported.
func Convert(dt string, from string, to string) (string, error) {
	t, err := Parse(from, dt)
	if err != nil {
		return "", err
	}

	return Format(t, to), nil
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
// www    -> 1          Three letter weekday name
// wwww   -> Monday     Full weekday name
// h      -> 3          Hour in 12 hour format without leading zero
// hh     -> 03         Hour in 12 hour format with leading zero
// hhh    -> 15         Hour in 24 hour format without leading zero
// a      -> pm         am/pm
// A      -> PM         AM/PM
// ii     -> 04         Minute with leading zero
// i      -> 4          Minute without leading zero
// ss     -> 05         Second with leading zero
// s      -> 5          Second without leading zero
// 0      -> 0          Microsecond with leading zero
// 9      -> 9          Microsecond without leading zero

// z      -> ±0700      UTC offset
// zh     -> ±07        Numeric timezone hour with hours only
// zz     -> ±07:00     UTC offset with colon
// zzz    -> MST        Timezone abbreviation
// zzzz   -> GMT-07:00  Timezone in long format
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
	f = strings.ToUpper(f)
	var converted interface{}

	// Initialize a map of format conversions
	conversions := map[string][][]string{
		"Y": {{"YYYY", "2006"}, {"YY", "06"}},
		"M": {{"MMMM", "January"}, {"MMM", "Jan"}, {"MM", "01"}, {"MT", ""}, {"M", "1"}},
		"D": {{"DDD", "002"}, {"DD", "02"}, {"DT", ""}, {"D", "2"}},
		"W": {{"WWWW", "Monday"}, {"WWW", "Mon"}},
		"H": {{"HHH", "15"}, {"HH", "03"}, {"H", "3"}},
		"A": {{"AA", "PM"}, {"A", "pm"}},
		"I": {{"II", "04"}, {"I", "4"}},
		"S": {{"SS", "05"}, {"S", "5"}},

		// Timezone
		"Z": {
			{"ZZZZ", "GMT-07:00"},
			{"ZZZ", "MST"},
			{"ZHH", "±0700"},

			{"ZZ", "MST"},
			{"ZH", "±07"},

			{"Z", "±07:00"},
		},
	}

	// Initialize a new string builder
	to := strings.Builder{}

	// Loop through the input format
	i := 0
	for i < len(f) {
		c := f[i]

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
			to.WriteString(string(c))
			i++
			continue
		}

		// Valid format character, check for the longest possible match
		for _, keyVal := range conv {
			key, val := keyVal[0], keyVal[1]
			iEnd := i + len(key)

			// Check if the len of key + i is less than the len of the format
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
					break
				}
			}
		}

		// If we get here, we didn't find a match, add the character as is
		if i < len(f) {
			to.WriteString(string(f[i]))
			i++
		}
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
