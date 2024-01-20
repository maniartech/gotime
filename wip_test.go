package temporal_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/temporal"
)

func TestWIP(t *testing.T) {
	fmt.Println("---", temporal.Format(time.Now(), "dt mmm 'yy"))
}
