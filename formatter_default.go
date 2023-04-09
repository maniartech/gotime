package temporal

import (
	"fmt"
	"strings"
	"time"
)

func TryFormat(t time.Time, f string) string {
	// Built-in format, return as is
	if _, ok := builtInFormats[f]; ok {
		return f
	}

	// If the format is cached, return the cached value
	if Options.cache != nil {
		if v, ok := Options.cache[f]; ok {
			return v[0]
		}
	}

	// Convert format to lower case for case insensitive matching
	f = strings.ToLower(f)

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
						// Special case for "dt" which is the day of the month with a suffix
						// (e.g. 1st, 2nd, 3rd, 4th, etc.)
						switch s := t.Day(); {
						case s == 1 || s == 21 || s == 31:
							to.WriteString(fmt.Sprintf("%dst", s))
						case s == 2 || s == 22:
							to.WriteString(fmt.Sprintf("%dnd", s))
						case s == 3 || s == 23:
							to.WriteString(fmt.Sprintf("%drd", s))
						default:
							to.WriteString(fmt.Sprintf("%dth", s))
						}

					} else {
						to.WriteString(val)
					}
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
	c := to.String()
	fmt.Println(">>>", c)

	if Options.cache != nil {
		Options.cache[f] = []string{c}
	}

	return t.Format(c)
}
