package main

import (
	"fmt"
	"time"
)

// DaysInMonth returns the number of days in a given month of a given year
func DaysInMonth(year, month int) int {
	// Go to the first day of the next month, then subtract one day
	nextMonth := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	return int(nextMonth.Sub(firstOfMonth).Hours() / 24)
}

func Age(birthDate time.Time, asOf time.Time) (years, months, days int) {
	// Use a more reliable approach: count years first, then months, then days
	birth := time.Date(birthDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, birthDate.Location())
	reference := time.Date(asOf.Year(), asOf.Month(), asOf.Day(), 0, 0, 0, 0, asOf.Location())

	fmt.Printf("Birth: %s\n", birth.Format("2006-01-02"))
	fmt.Printf("AsOf:  %s\n", reference.Format("2006-01-02"))

	// Start with the birth year and increment until we can't add more years
	currentDate := birth
	for {
		nextYear := currentDate.AddDate(1, 0, 0)
		if nextYear.After(reference) {
			break
		}
		years++
		currentDate = nextYear
		fmt.Printf("Added year %d, current date: %s\n", years, currentDate.Format("2006-01-02"))
	}

	// Now add months
	for {
		nextMonth := currentDate.AddDate(0, 1, 0)
		if nextMonth.After(reference) {
			break
		}
		months++
		currentDate = nextMonth
		fmt.Printf("Added month %d, current date: %s\n", months, currentDate.Format("2006-01-02"))
	}

	// Finally, count remaining days
	for currentDate.Before(reference) {
		days++
		currentDate = currentDate.AddDate(0, 0, 1)
		if days <= 5 { // Limit output for debugging
			fmt.Printf("Added day %d, current date: %s\n", days, currentDate.Format("2006-01-02"))
		}
	}

	return years, months, days
}

func main() {
	// Test case 1: years, months, and days
	fmt.Println("=== Test 1: years, months, and days ===")
	birthDate1 := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	asOf1 := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)
	years1, months1, days1 := Age(birthDate1, asOf1)
	fmt.Printf("Expected: 35 years, 1 months, 23 days\n")
	fmt.Printf("Got:      %d years, %d months, %d days\n\n", years1, months1, days1)

	// Test case 2: leap year birthday
	fmt.Println("=== Test 2: leap year birthday ===")
	birthDate2 := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
	asOf2 := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
	years2, months2, days2 := Age(birthDate2, asOf2)
	fmt.Printf("Expected: 25 years, 0 months, 1 days\n")
	fmt.Printf("Got:      %d years, %d months, %d days\n\n", years2, months2, days2)
}
