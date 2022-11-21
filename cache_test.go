package dateutils

import (
	"testing"
)

func TestDisableCache(t *testing.T) {
	// Test Case for DisableCache()
	DisableCache()
	if cache != nil {
		t.Error("Expected nil, got, ", cache)
	}
}

func TestEnableCache(t *testing.T) {
	// Test Case for EnableCache()
	EnableCache()
	if cache == nil {
		t.Errorf("Expected %v, got, nil", cache)
	}
}

func BenchmarkDisableCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DisableCache()
	}
}

func BenchmarkEnableCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnableCache()
	}
}
