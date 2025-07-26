package app_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/app"
	"os"
	"path/filepath"
	"testing"
)

func TestAddRequest_InvalidArgs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing arguments, got none")
		}
	}()
	app.AddRequest("GET", "", "/api/test") // requestName is empty
}

func TestAddRequest_ValidArgs(t *testing.T) {
	requestsDir := filepath.Join(os.TempDir(), "ex-crl-test-requests")
	os.RemoveAll(requestsDir)
	os.MkdirAll(requestsDir, 0755)
	defer os.RemoveAll(requestsDir)
	os.Setenv("EX_CRL_REQUESTS_DIR", requestsDir)

	os.Setenv("EX_CRL_SKIP_EDITOR", "1")
	os.Setenv("EX_CRL_TEST_MODE", "1")
	app.AddRequest("GET", "dup", "/api/dup")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected os.Exit or panic for duplicate request, got none")
		}
	}()
	os.Setenv("EX_CRL_SKIP_EDITOR", "1")
	app.AddRequest("GET", "dup", "/api/dup")
}
