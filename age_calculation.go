package gotime

import (
	"fmt"
	"time"
)

// Age calculates the precise age in years, months, and days between a birth date and a reference date.
// If no reference date is provided, it uses the current time.
//
// The calculation accounts for leap years and varying month lengths.
//
// Example:
//
//	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
//	asOf := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
//	years, months, days := gotime.Age(birthDate, asOf)
//	// Returns: 35, 1, 23 (35 years, 1 month, 23 days)
func Age(birthDate time.Time, asOf ...time.Time) (years, months, days int) {
	var ref time.Time
	if len(asOf) > 0 {
		ref = asOf[0]
	} else {
		ref = time.Now()
	}

	// Ensure birth date is before reference date
	if birthDate.After(ref) {
		return 0, 0, 0
	}

	// Use a more reliable approach: count years first, then months, then days
	birth := time.Date(birthDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, birthDate.Location())
	reference := time.Date(ref.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, ref.Location())

	// Start with the birth year and increment until we can't add more years
	currentDate := birth
	for {
		nextYear := currentDate.AddDate(1, 0, 0)
		if nextYear.After(reference) {
			break
		}
		years++
		currentDate = nextYear
	}

	// Now add months
	for {
		nextMonth := currentDate.AddDate(0, 1, 0)
		if nextMonth.After(reference) {
			break
		}
		months++
		currentDate = nextMonth
	}

	// Finally, count remaining days
	for currentDate.Before(reference) {
		days++
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return years, months, days
}

// YearsBetween calculates the precise number of years between two dates as a float64.
// The result includes fractional years based on the exact time difference.
//
// Example:
//
//	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
//	years := gotime.YearsBetween(start, end)
//	// Returns: approximately 5.5 years
func YearsBetween(start, end time.Time) float64 {
	// Ensure start is before end
	if start.After(end) {
		start, end = end, start
	}

	duration := end.Sub(start)

	// Average number of hours in a year (accounting for leap years)
	// 365.2425 days per year * 24 hours per day
	const hoursPerYear = 365.2425 * 24

	return duration.Hours() / hoursPerYear
}

// MonthsBetween calculates the precise number of months between two dates as a float64.
// The result includes fractional months based on the exact time difference.
//
// Example:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC)
//	months := gotime.MonthsBetween(start, end)
//	// Returns: approximately 6.5 months
func MonthsBetween(start, end time.Time) float64 {
	// Ensure start is before end
	if start.After(end) {
		start, end = end, start
	}

	duration := end.Sub(start)

	// Average number of hours in a month
	// 365.2425 days per year / 12 months * 24 hours per day
	const hoursPerMonth = (365.2425 / 12) * 24

	return duration.Hours() / hoursPerMonth
}

// DaysBetween calculates the number of days between two dates.
// This is a convenience function that returns the integer number of days.
//
// Example:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC)
//	days := gotime.DaysBetween(start, end)
//	// Returns: 7
func DaysBetween(start, end time.Time) int {
	// Ensure start is before end
	if start.After(end) {
		start, end = end, start
	}

	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

// WeeksBetween calculates the precise number of weeks between two dates as a float64.
// The result includes fractional weeks based on the exact time difference.
//
// Example:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 11, 0, 0, 0, 0, time.UTC)
//	weeks := gotime.WeeksBetween(start, end)
//	// Returns: approximately 1.43 weeks
func WeeksBetween(start, end time.Time) float64 {
	// Ensure start is before end
	if start.After(end) {
		start, end = end, start
	}

	duration := end.Sub(start)
	const hoursPerWeek = 7 * 24

	return duration.Hours() / hoursPerWeek
}

// DurationInWords returns a human-readable representation of a duration.
// It formats the duration in the most appropriate unit(s).
//
// Example:
//
//	d := 2*time.Hour + 30*time.Minute
//	result := gotime.DurationInWords(d)
//	// Returns: "2 hours 30 minutes"
func DurationInWords(d time.Duration) string {
	if d == 0 {
		return "0 seconds"
	}

	// Handle negative durations
	negative := d < 0
	if negative {
		d = -d
	}

	var parts []string

	// Extract time components
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	// Build the string with appropriate units
	if days > 0 {
		if days == 1 {
			parts = append(parts, "1 day")
		} else {
			parts = append(parts, fmt.Sprintf("%d days", days))
		}
	}

	if hours > 0 {
		if hours == 1 {
			parts = append(parts, "1 hour")
		} else {
			parts = append(parts, fmt.Sprintf("%d hours", hours))
		}
	}

	if minutes > 0 {
		if minutes == 1 {
			parts = append(parts, "1 minute")
		} else {
			parts = append(parts, fmt.Sprintf("%d minutes", minutes))
		}
	}

	if seconds > 0 && len(parts) < 2 { // Only show seconds if we don't have 2+ larger units
		if seconds == 1 {
			parts = append(parts, "1 second")
		} else {
			parts = append(parts, fmt.Sprintf("%d seconds", seconds))
		}
	}

	// Handle very small durations
	if len(parts) == 0 {
		return "less than 1 second"
	}

	// Join parts appropriately
	var result string
	if len(parts) == 1 {
		result = parts[0]
	} else if len(parts) == 2 {
		result = parts[0] + " " + parts[1]
	} else {
		// For 3+ parts, use only the first 2 most significant units
		result = parts[0] + " " + parts[1]
	}

	if negative {
		result = "-" + result
	}

	return result
}

// IsValidAge checks if the given birth date results in a valid age (not negative, not unreasonably old).
// This is useful for validating user input for birth dates.
//
// Example:
//
//	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
//	valid := gotime.IsValidAge(birthDate)
//	// Returns: true
func IsValidAge(birthDate time.Time, asOf ...time.Time) bool {
	var ref time.Time
	if len(asOf) > 0 {
		ref = asOf[0]
	} else {
		ref = time.Now()
	}

	// Birth date cannot be in the future
	if birthDate.After(ref) {
		return false
	}

	// Calculate age in years
	years := YearsBetween(birthDate, ref)

	// Reasonable age limits (0 to 150 years) - use <= for inclusive range
	return years >= 0 && years <= 150
}
