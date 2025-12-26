package gotime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
)

func TestWIP(t *testing.T) {
	dt, _ := gotime.ParseInLocation("yyyy-mm-dd", "2022-12-31", time.FixedZone("ABC", 5.5*60*60))
	fmt.Println(dt)
}
