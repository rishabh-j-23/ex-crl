package utils_test

import (
	"github.com/rishabh-j-23/ex-crl/utils"
	"testing"
)

func TestCookieStoragePath(t *testing.T) {
	path := utils.GetCookieStoragePath()
	if path == "" {
		t.Errorf("GetCookieStoragePath returned empty string")
	}
}
