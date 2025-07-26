package models_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"testing"
)

func TestProjectConfigStruct(t *testing.T) {
	c := models.ProjectConfig{Name: "foo"}
	if c.Name != "foo" {
		t.Errorf("ProjectConfig struct basic usage failed")
	}
}
