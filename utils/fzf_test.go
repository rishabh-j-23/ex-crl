package utils_test

import (
	"github.com/rishabh-j-23/ex-crl/utils"
	"testing"
)

func TestFzfFromList_Basic(t *testing.T) {
	items := []string{"a", "b", "c"}
	out := utils.FzfFromList(items, false)
	if out == nil {
		t.Errorf("FzfFromList returned nil")
	}
}
