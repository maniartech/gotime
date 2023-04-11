package temporal

import (
	"errors"
	"strings"
	"time"
)

// convertLayout converts this library datetime format to a go format.
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
func convertLayout(f string, forParsing bool) ([]string, error) {
	// Built-in format, return as is
	if version, ok := builtInLayouts[f]; ok {
		if runtimeVersion >= version {
			return []string{f}, nil
		}
	}

	// If the format is cached, return the cached value
	if Options.cache != nil {
		if v, ok := Options.cache[f]; ok {
			return v, nil
		}
	}

	// Convert format to lower case for case insensitive matching
	f = strings.ToLower(f)
	converted := []string{}

	// Initialize a map of format conversions
	conversions := map[string][][]string{
		"y": {{"yyyy", "2006"}, {"yy", "06"}},
		"m": {{"mmmm", "January"}, {"mmm", "Jan"}, {"mm", "01"}, {"m", "1"}},
		"d": {{"ddd", "002"}, {"dd", "02"}, {"dt", ""}, {"d", "2"}},
		"w": {{"wwww", "Monday"}, {"www", "Mon"}},
		"h": {{"hhh", "15"}, {"hh", "03"}, {"h", "3"}},
		"a": {{"aa", "PM"}, {"a", "pm"}},
		"i": {{"ii", "04"}, {"i", "4"}},
		"s": {{"ss", "05"}, {"s", "5"}},

		// Timezone
		"z": {
			{"zzzz", "GMT-07:00"},
			{"zzz", "MST"},
			{"zhh", "±0700"},

			{"zz", "MST"},
			{"zh", "±07"},

			{"z", "±07:00"},
		},
	}

	// Initialize a new string builder
	to := strings.Builder{}

	// Loop through the input format
	i := 0
	for i < len(f) {
		c := f[i]

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

						converted = append(converted, to.String()) // Append the converted format
						converted = append(converted, key)         // Append the value to the converted format
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
	if finalConvert != "" {
		converted = append(converted, finalConvert)
	}
	// converted = append(converted, to.String())
	if Options.cache != nil {
		Options.cache[f] = converted
	}

	return converted, nil
}

// Convert function converts a datetime from one string format to another.
// It takes the datetime string in the single format and converts it to the expected output.
// It returns an error when the format is not supported.
func Convert(dt string, from string, to string) (string, error) {
	if from == to {
		return dt, nil
	}

	// // Convert the format to go format.
	fromConverted, err := convertLayout(from, true)
	if err != nil {
		return "", err
	}

	toConverted, _ := convertLayout(to, false)

	if len(fromConverted) > 1 {
		return "", errors.New(errOrdinalsNotSupported)
	}

	parsed, err := time.Parse(fromConverted[0], dt)
	if err != nil {
		return "", err
	}

	return format(parsed, toConverted), nil
}
