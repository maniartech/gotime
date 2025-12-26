package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
)

func TestTri(t *testing.T) {
	// tm is time at 2PM
	tm := time.Date(2022, time.January, 1, 14, 9, 0, 0, time.UTC)
	println(gotime.Format(tm, "hhhh:ii"))  // Expected: `14:09`
	println(gotime.Format(tm, "hh:ii aa")) // Expected: `02:09 PM`
}
