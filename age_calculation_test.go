package gotime

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/v2/internal/utils"
)

func TestAge(t *testing.T) {
	tests := []struct {
		name       string
		birthDate  time.Time
		asOf       time.Time
		wantYears  int
		wantMonths int
		wantDays   int
	}{
		{
			name:       "exact years",
			birthDate:  time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 5, 15, 0, 0, 0, 0, time.UTC),
			wantYears:  35,
			wantMonths: 0,
			wantDays:   0,
		},
		{
			name:       "years and months",
			birthDate:  time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 7, 15, 0, 0, 0, 0, time.UTC),
			wantYears:  35,
			wantMonths: 2,
			wantDays:   0,
		},
		{
			name:       "years, months, and days",
			birthDate:  time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC),
			wantYears:  35,
			wantMonths: 1,
			wantDays:   22, // From May 15 to June 15 (1 month), then June 15 to July 7 (22 days)
		},
		{
			name:       "leap year birthday",
			birthDate:  time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			wantYears:  25,
			wantMonths: 0,
			wantDays:   0, // Go's AddDate correctly handles Feb 29 -> Mar 1 in non-leap years
		},
		{
			name:       "same day",
			birthDate:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			wantYears:  0,
			wantMonths: 0,
			wantDays:   0,
		},
		{
			name:       "future birth date",
			birthDate:  time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			wantYears:  0,
			wantMonths: 0,
			wantDays:   0,
		},
		{
			name:       "end of month calculations",
			birthDate:  time.Date(1990, 1, 31, 0, 0, 0, 0, time.UTC),
			asOf:       time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
			wantYears:  35,
			wantMonths: 0,
			wantDays:   28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			years, months, days := Age(tt.birthDate, tt.asOf)
			utils.AssertEqual(t, tt.wantYears, years)
			utils.AssertEqual(t, tt.wantMonths, months)
			utils.AssertEqual(t, tt.wantDays, days)
		})
	}
}

func TestAgeWithoutAsOf(t *testing.T) {
	// Test that Age works without asOf parameter (uses current time)
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	years, months, days := Age(birthDate)

	// Should return positive values for a date in the past
	if years < 0 || months < 0 || days < 0 {
		t.Errorf("Age() with past date should return non-negative values, got years=%d, months=%d, days=%d", years, months, days)
	}
}

func TestYearsBetween(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		want      float64
		tolerance float64
	}{
		{
			name:      "exact year",
			start:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      1.0,
			tolerance: 0.01,
		},
		{
			name:      "half year",
			start:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
			want:      0.5,
			tolerance: 0.05,
		},
		{
			name:      "multiple years",
			start:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      5.0,
			tolerance: 0.01,
		},
		{
			name:      "reverse order",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      5.0,
			tolerance: 0.01,
		},
		{
			name:      "same time",
			start:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      0.0,
			tolerance: 0.001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := YearsBetween(tt.start, tt.end)
			if abs := func(x float64) float64 {
				if x < 0 {
					return -x
				}
				return x
			}; abs(got-tt.want) > tt.tolerance {
				t.Errorf("YearsBetween() = %v, want %v (±%v)", got, tt.want, tt.tolerance)
			}
		})
	}
}

func TestMonthsBetween(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		want      float64
		tolerance float64
	}{
		{
			name:      "exact month",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC),
			want:      1.0,
			tolerance: 0.05,
		},
		{
			name:      "half month",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 16, 0, 0, 0, 0, time.UTC),
			want:      0.5,
			tolerance: 0.1,
		},
		{
			name:      "multiple months",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			want:      6.0,
			tolerance: 0.1,
		},
		{
			name:      "same time",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      0.0,
			tolerance: 0.001,
		},
		{
			name:      "reverse order",
			start:     time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      6.0,
			tolerance: 0.1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MonthsBetween(tt.start, tt.end)
			if abs := func(x float64) float64 {
				if x < 0 {
					return -x
				}
				return x
			}; abs(got-tt.want) > tt.tolerance {
				t.Errorf("MonthsBetween() = %v, want %v (±%v)", got, tt.want, tt.tolerance)
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		want  int
	}{
		{
			name:  "one day",
			start: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
			want:  1,
		},
		{
			name:  "one week",
			start: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC),
			want:  7,
		},
		{
			name:  "leap year",
			start: time.Date(2020, 2, 28, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			want:  2, // 2020 is a leap year
		},
		{
			name:  "reverse order",
			start: time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  7,
		},
		{
			name:  "same day",
			start: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DaysBetween(tt.start, tt.end)
			utils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestWeeksBetween(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		want      float64
		tolerance float64
	}{
		{
			name:      "exact week",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC),
			want:      1.0,
			tolerance: 0.01,
		},
		{
			name:      "half week",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 4, 12, 0, 0, 0, time.UTC),
			want:      0.5,
			tolerance: 0.01,
		},
		{
			name:      "multiple weeks",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 29, 0, 0, 0, 0, time.UTC),
			want:      4.0,
			tolerance: 0.01,
		},
		{
			name:      "same time",
			start:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      0.0,
			tolerance: 0.001,
		},
		{
			name:      "reverse order",
			start:     time.Date(2025, 1, 29, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:      4.0,
			tolerance: 0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WeeksBetween(tt.start, tt.end)
			if abs := func(x float64) float64 {
				if x < 0 {
					return -x
				}
				return x
			}; abs(got-tt.want) > tt.tolerance {
				t.Errorf("WeeksBetween() = %v, want %v (±%v)", got, tt.want, tt.tolerance)
			}
		})
	}
}

func TestDurationInWords(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{
			name: "zero duration",
			d:    0,
			want: "0 seconds",
		},
		{
			name: "seconds only",
			d:    30 * time.Second,
			want: "30 seconds",
		},
		{
			name: "one second",
			d:    1 * time.Second,
			want: "1 second",
		},
		{
			name: "minutes only",
			d:    5 * time.Minute,
			want: "5 minutes",
		},
		{
			name: "one minute",
			d:    1 * time.Minute,
			want: "1 minute",
		},
		{
			name: "hours only",
			d:    3 * time.Hour,
			want: "3 hours",
		},
		{
			name: "one hour",
			d:    1 * time.Hour,
			want: "1 hour",
		},
		{
			name: "days only",
			d:    2 * 24 * time.Hour,
			want: "2 days",
		},
		{
			name: "one day",
			d:    24 * time.Hour,
			want: "1 day",
		},
		{
			name: "hours and minutes",
			d:    2*time.Hour + 30*time.Minute,
			want: "2 hours 30 minutes",
		},
		{
			name: "days and hours",
			d:    25 * time.Hour,
			want: "1 day 1 hour",
		},
		{
			name: "complex duration",
			d:    2*24*time.Hour + 3*time.Hour + 45*time.Minute,
			want: "2 days 3 hours",
		},
		{
			name: "negative duration",
			d:    -2 * time.Hour,
			want: "-2 hours",
		},
		{
			name: "very small duration",
			d:    500 * time.Millisecond,
			want: "less than 1 second",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DurationInWords(tt.d)
			utils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestIsValidAge(t *testing.T) {
	now := time.Date(2025, 1, 7, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		birthDate time.Time
		asOf      *time.Time
		want      bool
	}{
		{
			name:      "valid age 25",
			birthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:      &now,
			want:      true,
		},
		{
			name:      "valid age 0 (newborn)",
			birthDate: now,
			asOf:      &now,
			want:      true,
		},
		{
			name:      "valid age 100",
			birthDate: time.Date(1925, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:      &now,
			want:      true,
		},
		{
			name:      "valid age 149 (near edge case)",
			birthDate: time.Date(1876, 1, 7, 0, 0, 0, 0, time.UTC), // 149 years before now
			asOf:      &now,
			want:      true,
		},
		{
			name:      "invalid future birth date",
			birthDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:      &now,
			want:      false,
		},
		{
			name:      "invalid age over 150",
			birthDate: time.Date(1874, 1, 1, 0, 0, 0, 0, time.UTC),
			asOf:      &now,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			if tt.asOf != nil {
				got = IsValidAge(tt.birthDate, *tt.asOf)
			} else {
				got = IsValidAge(tt.birthDate)
			}
			utils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestIsValidAgeWithoutAsOf(t *testing.T) {
	// Test that IsValidAge works without asOf parameter (uses current time)
	validBirthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	futureBirthDate := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)

	utils.AssertEqual(t, true, IsValidAge(validBirthDate))
	utils.AssertEqual(t, false, IsValidAge(futureBirthDate))
}

// Benchmark tests
func BenchmarkAge(b *testing.B) {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	asOf := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Age(birthDate, asOf)
	}
}

func BenchmarkYearsBetween(b *testing.B) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsBetween(start, end)
	}
}

func BenchmarkDurationInWords(b *testing.B) {
	d := 2*24*time.Hour + 3*time.Hour + 45*time.Minute

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DurationInWords(d)
	}
}

// Example tests for documentation
func ExampleAge() {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	asOf := time.Date(2025, 7, 7, 0, 0, 0, 0, time.UTC)

	years, months, days := Age(birthDate, asOf)
	fmt.Printf("Age: %d years, %d months, %d days", years, months, days)
	// Output: Age: 35 years, 1 months, 22 days
}

func ExampleYearsBetween() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	years := YearsBetween(start, end)
	fmt.Printf("%.1f years", years)
	// Output: 5.5 years
}

func ExampleDurationInWords() {
	d := 2*time.Hour + 30*time.Minute
	result := DurationInWords(d)
	fmt.Println(result)
	// Output: 2 hours 30 minutes
}

func ExampleIsValidAge() {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	valid := IsValidAge(birthDate)
	fmt.Println(valid)
	// Output: true
}
