package dateutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFormat(t *testing.T) {
	converted := ConvertFormat("dd-mm-yyyy")

	assert.Equal(t, "02-01-2006", converted)
}

func BenchmarkConvertFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertFormat("dd-mm-yyyy")
	}
}
