package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
	"github.com/stretchr/testify/assert"
)

func TestMonth(t *testing.T) {
	// MonthStart - write a test that loops through all the months and checks that the first day of the month is correct
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	totalDaysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for _, month := range months {
		expectedDate := temporal.Date(2023, month, 1)
		functionDate := temporal.MonthStart(time.Date(2023, time.Month(month), 1, 11, 2, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}

	// MonthEnd
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), totalDaysInMonth[month-1], 23, 59, 59, 999999999, time.UTC)
		functionDate := temporal.MonthEnd(time.Date(2023, time.Month(month), 1, 11, 22, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}

	// LastMonth
	expectedDate := time.Now().AddDate(0, -1, 0)
	functionDate := temporal.LastMonth()

	assert.Equal(t, expectedDate, functionDate)

	// NextMonth
	expectedDate = time.Now().AddDate(0, 1, 0)
	functionDate = temporal.NextMonth()

	assert.Equal(t, expectedDate, functionDate)
	
}
