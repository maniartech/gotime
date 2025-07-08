package nites_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/internal/nites"
	"github.com/maniartech/gotime/internal/utils"
)

func TestTrialForma(t *testing.T) {
	fmt.Println("---", nites.Format(time.Now(), `dd/mm/yy`))
}

func TestFormat(t *testing.T) {
	// Converting time.Time to Go format.
	date1 := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
	date1Formatted := nites.Format(date1, "yyyy/mm/dd")
	utils.AssertEqual(t, "0018/06/04", date1Formatted)

}

func TestOrdinalFormatting(t *testing.T) {
	// Test day ordinals
	date1st := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)    // 1st January
	date2nd := time.Date(2025, 2, 2, 0, 0, 0, 0, time.UTC)    // 2nd February
	date3rd := time.Date(2025, 3, 3, 0, 0, 0, 0, time.UTC)    // 3rd March
	date21st := time.Date(2025, 11, 21, 0, 0, 0, 0, time.UTC) // 21st November

	// Test day ordinals
	utils.AssertEqual(t, "1st", nites.Format(date1st, "dt"))
	utils.AssertEqual(t, "2nd", nites.Format(date2nd, "dt"))
	utils.AssertEqual(t, "3rd", nites.Format(date3rd, "dt"))
	utils.AssertEqual(t, "21st", nites.Format(date21st, "dt"))

	// Test month ordinals (this was the bug - was using day value for month ordinals)
	utils.AssertEqual(t, "1st", nites.Format(date1st, "mt"))   // January = 1st
	utils.AssertEqual(t, "2nd", nites.Format(date2nd, "mt"))   // February = 2nd
	utils.AssertEqual(t, "3rd", nites.Format(date3rd, "mt"))   // March = 3rd
	utils.AssertEqual(t, "11th", nites.Format(date21st, "mt")) // November = 11th (not 21st)
}

func TestOrdinalEdgeCases(t *testing.T) {
	// Test edge cases for ordinals: 11th, 12th, 13th (should all be "th", not "st", "nd", "rd")
	date11th := time.Date(2025, 11, 11, 0, 0, 0, 0, time.UTC)
	date12th := time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC)
	date13th := time.Date(2025, 1, 13, 0, 0, 0, 0, time.UTC)
	date22nd := time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)
	date23rd := time.Date(2025, 3, 23, 0, 0, 0, 0, time.UTC)

	// Day ordinals
	utils.AssertEqual(t, "11th", nites.Format(date11th, "dt"))
	utils.AssertEqual(t, "12th", nites.Format(date12th, "dt"))
	utils.AssertEqual(t, "13th", nites.Format(date13th, "dt"))
	utils.AssertEqual(t, "22nd", nites.Format(date22nd, "dt"))
	utils.AssertEqual(t, "23rd", nites.Format(date23rd, "dt"))

	// Month ordinals
	utils.AssertEqual(t, "11th", nites.Format(date11th, "mt")) // November
	utils.AssertEqual(t, "12th", nites.Format(date12th, "mt")) // December
}

func BenchmarkFormat(b *testing.B) {
	// Benchmarking for nites.Format function
	for i := 0; i < b.N; i++ {
		date := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
		nites.Format(date, "yyyy/mm/dd")
	}
}
