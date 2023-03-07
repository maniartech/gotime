package temporal_test

import (
	"testing"

	"github.com/maniartech/temporal"
)

func TestDisableCache(t *testing.T) {
	// Test Case for DisableCache()
	temporal.DisableCache()
	if cache != nil {
		t.Error("Expected nil, got, ", cache)
	}
}

func TestEnableCache(t *testing.T) {
	// Test Case for EnableCache()
	temporal.EnableCache()
	if temporal.cache == nil {
		t.Errorf("Expected %v, got, nil", cache)
	}
}

func BenchmarkDisableCache(b *testing.B) {
	// Benchmarking for DisableCache function
	for i := 0; i < b.N; i++ {
		temporal.DisableCache()
	}
}

func BenchmarkEnableCache(b *testing.B) {
	// Benchmarking for EnableCache function
	for i := 0; i < b.N; i++ {
		temporal.EnableCache()
	}
}
