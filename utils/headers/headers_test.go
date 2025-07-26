package headers_test

import (
	"github.com/rishabh-j-23/ex-crl/utils/headers"
	"testing"
)

func TestHeadersConstants(t *testing.T) {
	if headers.ContentType != "Content-Type" {
		t.Errorf("ContentType constant incorrect")
	}
}
