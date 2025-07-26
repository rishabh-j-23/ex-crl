package assert_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"testing"
)

func TestErrIsNil_Basic(t *testing.T) {
	assert.ErrIsNil(nil, "should not panic")
}
