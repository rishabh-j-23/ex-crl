package utils_test

import (
	"github.com/rishabh-j-23/ex-crl/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCurrentProjectName(t *testing.T) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	tmp := os.TempDir()
	dir := filepath.Join(tmp, "ex-crl-test-proj")
	os.MkdirAll(dir, 0755)
	os.Chdir(dir)
	if got := utils.GetCurrentProjectName(); got != "ex-crl-test-proj" {
		t.Errorf("expected ex-crl-test-proj, got %s", got)
	}
}

func TestGetRequestsDir_EnvOverride(t *testing.T) {
	os.Setenv("EX_CRL_REQUESTS_DIR", "/tmp/reqs")
	if got := utils.GetRequestsDir(); got != "/tmp/reqs" {
		t.Errorf("expected /tmp/reqs, got %s", got)
	}
	os.Unsetenv("EX_CRL_REQUESTS_DIR")
}

func TestGetRequestsDir_Default(t *testing.T) {
	os.Unsetenv("EX_CRL_REQUESTS_DIR")
	projectDir := utils.GetProjectDir()
	expected := filepath.Join(projectDir, "requests")
	if got := utils.GetRequestsDir(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
