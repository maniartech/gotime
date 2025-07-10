package utils

import (
	"reflect"
	"testing"
	"time"
)

// AssertEqual checks equality and fails the test if not equal (for use in all packages)
func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	switch e := expected.(type) {
	case time.Time:
		a, ok := actual.(time.Time)
		if !ok || !e.Equal(a) {
			t.Errorf("Expected %v but got %v", e, actual)
		}
	default:
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}

// AssertNotEqual checks inequality and fails the test if equal (for use in all packages)
func AssertNotEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v to not equal %v", expected, actual)
	}
}

// AssertNoError fails the test if err is not nil (for use in all packages)
func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}

// AssertPanics fails the test if f does not panic (for use in all packages)
func AssertPanics(t *testing.T, f func()) {
	didPanic := false
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}
	}()
	f()
	if !didPanic {
		t.Errorf("Expected panic but got none")
	}
}

// Internal helpers for negative-path coverage tests (return bool, do not use outside assert_test.go)
func assertEqualBool(expected interface{}, actual interface{}) bool {
	switch e := expected.(type) {
	case time.Time:
		a, ok := actual.(time.Time)
		if !ok || !e.Equal(a) {
			return true
		}
	default:
		if !reflect.DeepEqual(expected, actual) {
			return true
		}
	}
	return false
}

func assertNotEqualBool(expected interface{}, actual interface{}) bool {
	if expected == actual {
		return true
	}
	return false
}

func assertNoErrorBool(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func assertPanicsBool(f func()) bool {
	didPanic := false
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}
	}()
	f()
	if !didPanic {
		return true
	}
	return false
}
