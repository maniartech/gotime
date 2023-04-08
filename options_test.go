package temporal_test

import (
	"testing"

	"github.com/maniartech/temporal"
)

func TestDisableCache(t *testing.T) {
	// Test Case for DisableCache()
	temporal.Options.DisableCache()

	if temporal.Options.IsCacheEnabled() {
		t.Errorf("Expected %v, got, %v", false, true)
	}
}

func TestEnableCache(t *testing.T) {
	// Test Case for EnableCache()
	temporal.Options.EnableCache()
	if !temporal.Options.IsCacheEnabled() {
		t.Errorf("Expected %v, got, %v", true, false)
	}
}

func BenchmarkDisableCache(b *testing.B) {
	// Benchmarking for DisableCache function
	for i := 0; i < b.N; i++ {
		temporal.Options.DisableCache()
	}
}

func BenchmarkEnableCache(b *testing.B) {
	// Benchmarking for EnableCache function
	for i := 0; i < b.N; i++ {
		temporal.Options.EnableCache()
	}
}
