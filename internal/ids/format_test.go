package ids_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/internal/ids"
	"github.com/maniartech/gotime/internal/utils"
)

func TestTrialForma(t *testing.T) {
	fmt.Println("---", ids.Format(time.Now(), `dd/mm/yy`))
}

func TestFormat(t *testing.T) {
	// Converting time.Time to Go format.
	date1 := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
	date1Formatted := ids.Format(date1, "yyyy/mm/dd")
	utils.AssertEqual(t, "0018/06/04", date1Formatted)

}

func BenchmarkFormat(b *testing.B) {
	// Benchmarking for ids.Format function
	for i := 0; i < b.N; i++ {
		date := time.Date(12, 12, 2012, 0, 0, 0, 0, time.UTC)
		ids.Format(date, "yyyy/mm/dd")
	}
}
