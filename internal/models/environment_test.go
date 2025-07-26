package models_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"testing"
)

func TestEnvironmentStruct(t *testing.T) {
	e := models.Environment{Name: "foo"}
	if e.Name != "foo" {
		t.Errorf("Environment struct basic usage failed")
	}
}
