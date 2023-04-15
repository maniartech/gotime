package cache_test

import (
	"testing"

	"github.com/maniartech/temporal/cache"
)

func TestDisableCache(t *testing.T) {
	// Test Case for DisableCache()
	cache.Disable()

	if cache.IsEnabled() {
		t.Errorf("Expected %v, got, %v", false, true)
	}
}

func TestEnableCache(t *testing.T) {
	// Test Case for EnableCache()
	cache.Enable()
	if !cache.IsEnabled() {
		t.Errorf("Expected %v, got, %v", true, false)
	}
}

func BenchmarkDisableCache(b *testing.B) {
	// Benchmarking for DisableCache function
	for i := 0; i < b.N; i++ {
		cache.Disable()
	}
}

func BenchmarkEnableCache(b *testing.B) {
	// Benchmarking for EnableCache function
	for i := 0; i < b.N; i++ {
		cache.Enable()
	}
}
