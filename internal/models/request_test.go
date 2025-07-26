package models_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"testing"
)

func TestRequestStruct(t *testing.T) {
	r := models.Request{Name: "foo"}
	if r.Name != "foo" {
		t.Errorf("Request struct basic usage failed")
	}
}
