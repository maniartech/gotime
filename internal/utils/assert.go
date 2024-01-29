package utils

import (
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func AssertNotEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v to not equal %v", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}

func AssertPanics(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but got none")
		}
	}()
	f()
}
