package core_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/core"
	"testing"
)

func TestAddRequest_InvalidArgs(t *testing.T) {
	// Simulate missing arguments
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing arguments, got none")
		}
	}()
	core.AddRequest("GET", "", "/api/test") // requestName is empty
}

func TestAddRequest_ValidArgs(t *testing.T) {
	// Should not panic
	core.AddRequest("GET", "test", "/api/test")
}
