package options_test

import (
	"testing"

	"github.com/maniartech/temporal/options"
)

func TestDisableCache(t *testing.T) {
	// Test Case for DisableCache()
	options.DisableCache()

	if options.IsCacheEnabled() {
		t.Errorf("Expected %v, got, %v", false, true)
	}
}

func TestEnableCache(t *testing.T) {
	// Test Case for EnableCache()
	options.EnableCache()
	if !options.IsCacheEnabled() {
		t.Errorf("Expected %v, got, %v", true, false)
	}
}

func BenchmarkDisableCache(b *testing.B) {
	// Benchmarking for DisableCache function
	for i := 0; i < b.N; i++ {
		options.DisableCache()
	}
}

func BenchmarkEnableCache(b *testing.B) {
	// Benchmarking for EnableCache function
	for i := 0; i < b.N; i++ {
		options.EnableCache()
	}
}
