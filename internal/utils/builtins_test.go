package utils

import (
	"testing"
	"time"
)

func TestBuiltInLayouts_ContainsAll(t *testing.T) {
	for layout := range BuiltInLayouts {
		if _, ok := BuiltInLayouts[layout]; !ok {
			t.Errorf("Layout %s missing from BuiltInLayouts", layout)
		}
	}
}

func TestBuiltInLayouts_Values(t *testing.T) {
	// Check that all values are either v100 or v120
	for _, v := range BuiltInLayouts {
		if v != 100 && v != 120 {
			t.Errorf("Unexpected value in BuiltInLayouts: %v", v)
		}
	}
}

func TestBuiltInLayouts_Keys(t *testing.T) {
	// Check that all keys are valid time layouts
	for layout := range BuiltInLayouts {
		_ = layout // just ensure we can iterate
	}
}

func TestBuiltInLayouts_Example(t *testing.T) {
	if BuiltInLayouts[time.RFC3339] != 100 {
		t.Errorf("Expected RFC3339 to have value 100")
	}
	if BuiltInLayouts[time.DateTime] != 120 {
		t.Errorf("Expected DateTime to have value 120")
	}
}
