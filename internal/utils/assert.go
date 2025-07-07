package utils

import (
	"reflect"
	"testing"
	"time"
)

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
