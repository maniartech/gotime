package cache_test

import (
	"testing"

	"github.com/maniartech/gotime/internal/cache"
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

func TestSetAndGet(t *testing.T) {
	// Ensure cache is enabled
	cache.Enable()

	// Test setting and getting values
	testKey := "test_key"
	testValue := "test_value"

	cache.Set(testKey, testValue)
	retrievedValue := cache.Get(testKey)

	if retrievedValue != testValue {
		t.Errorf("Expected %v, got %v", testValue, retrievedValue)
	}

	// Test getting non-existent key
	nonExistentValue := cache.Get("non_existent_key")
	if nonExistentValue != nil {
		t.Errorf("Expected nil for non-existent key, got %v", nonExistentValue)
	}
}

func TestSetAndGetWithDisabledCache(t *testing.T) {
	// Test Set and Get when cache is disabled
	cache.Disable()

	testKey := "disabled_test_key"
	testValue := "disabled_test_value"

	// Set should do nothing when cache is disabled
	cache.Set(testKey, testValue)

	// Get should return nil when cache is disabled
	retrievedValue := cache.Get(testKey)
	if retrievedValue != nil {
		t.Errorf("Expected nil when cache is disabled, got %v", retrievedValue)
	}

	// Re-enable cache for other tests
	cache.Enable()
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
