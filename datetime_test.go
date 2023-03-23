package temporal_test

import (
	"testing"

	"github.com/maniartech/temporal"
)

func BenchmarkDate(b *testing.B) {
	d := temporal.Date()
	for i := 0; i < b.N; i++ {
		d.UTC()
	}
}
