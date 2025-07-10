//go:build coverage
// +build coverage

package utils

// Negative-path tests: Only run with 'go test -tags=coverage' for coverage analysis.
// These tests intentionally fail assertions to cover error branches, but are skipped in normal test runs.

import (
	"errors"
	"testing"
	"time"
)

func TestAssertEqual_Primitives(t *testing.T) {
	AssertEqual(t, 5, 5)
	AssertEqual(t, "foo", "foo")
}

func TestAssertEqual_Time(t *testing.T) {
	t1 := time.Now()
	t2 := t1
	AssertEqual(t, t1, t2)
}

func TestAssertEqual_Fail(t *testing.T) {
	if !assertEqualBool(1, 2) {
		t.Errorf("Expected assertEqualBool to fail, but it passed")
	}
}

func TestAssertNotEqual(t *testing.T) {
	AssertNotEqual(t, 1, 2)
}

func TestAssertNotEqual_Fail(t *testing.T) {
	if !assertNotEqualBool(1, 1) {
		t.Errorf("Expected assertNotEqualBool to fail, but it passed")
	}
}

func TestAssertNoError(t *testing.T) {
	AssertNoError(t, nil)
}

func TestAssertNoError_Fail(t *testing.T) {
	if !assertNoErrorBool(errors.New("fail")) {
		t.Errorf("Expected assertNoErrorBool to fail, but it passed")
	}
}

func TestAssertPanics(t *testing.T) {
	AssertPanics(t, func() { panic("should panic") })
}

func TestAssertPanics_Fail(t *testing.T) {
	if !assertPanicsBool(func() {}) {
		t.Errorf("Expected assertPanicsBool to fail, but it passed")
	}
}
