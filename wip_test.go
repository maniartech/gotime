package temporal_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/temporal"
)

func TestWIP(t *testing.T) {
	dt, _ := temporal.Parse("yyyy-mm-dd", "2022-12-31", time.FixedZone("ABC", 5.5*60*60))
	fmt.Println(dt)
}
