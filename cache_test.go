package dateutils

import (
	"testing"
)

func TestDisableCache(t *testing.T) {
	DisableCache()
	if cache != nil {
		t.Error("Expected nil, got, ", cache)
	}
}

func TestEnableCache(t *testing.T) {
	EnableCache()
	if cache == nil {
		t.Errorf("Expected %v, got, nil", cache)
	}
}
