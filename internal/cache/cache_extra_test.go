package cache_test

import (
	"testing"

	"github.com/maniartech/gotime/v2/internal/cache"
)

func TestGetStrs(t *testing.T) {
	cache.Enable()
	cache.Set("strs", []string{"a", "b"})
	cache.Set("str", "c")
	cache.Set("other", 123)

	v := cache.GetStrs("strs")
	if len(v) != 2 || v[0] != "a" || v[1] != "b" {
		t.Errorf("Expected [a b], got %v", v)
	}

	v2 := cache.GetStrs("str")
	if len(v2) != 1 || v2[0] != "c" {
		t.Errorf("Expected [c], got %v", v2)
	}

	v3 := cache.GetStrs("other")
	if v3 != nil {
		t.Errorf("Expected nil for non-string value, got %v", v3)
	}

	cache.Disable()
	v4 := cache.GetStrs("strs")
	if v4 != nil {
		t.Errorf("Expected nil when cache is disabled, got %v", v4)
	}
}

// TestInitAndRepeatedEnableDisable covers the init logic and repeated enable/disable for 100% coverage
func TestInitAndRepeatedEnableDisable(t *testing.T) {
	// Simulate package init
	cache.Disable()
	cache.Enable()
	cache.Enable() // call twice
	cache.Disable()
	cache.Disable() // call twice
	cache.Enable()
	// Set and Get after re-enabling
	cache.Set("foo", "bar")
	if got := cache.Get("foo"); got != "bar" {
		t.Errorf("Expected bar after re-enable, got %v", got)
	}
	// Set/Get when cache is nil
	cache.Disable()
	cache.Set("foo", "baz") // should do nothing
	if got := cache.Get("foo"); got != nil {
		t.Errorf("Expected nil after disable, got %v", got)
	}
}

// TestGetStrs_KeyExistsButNil covers the case where the key exists but value is nil
func TestGetStrs_KeyExistsButNil(t *testing.T) {
	cache.Enable()
	cache.Set("nil_value", nil)
	v := cache.GetStrs("nil_value")
	if v != nil {
		t.Errorf("Expected nil for key with nil value, got %v", v)
	}
}
