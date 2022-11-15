package dateutils

import (
	"testing"
)

func TestConvertFormat(t *testing.T) {
	converted := ConvertFormat("dd-mm-yyyy")

	if converted != "02-01-2006" {
		t.Error("Expected 02-01-2006, got ", converted)
	}
}

func BenchmarkConvertFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertFormat("dd-mm-yyyy")
	}
}
