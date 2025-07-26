package app_test

import (
	"github.com/rishabh-j-23/ex-crl/internal/app"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"os"
	"testing"
)

func TestGetEnvByName_Found(t *testing.T) {
	cfg := models.ProjectConfig{
		ConfiguredEnv: []models.Environment{
			{Name: "dev", BaseUrl: "http://dev"},
			{Name: "prod", BaseUrl: "http://prod"},
		},
	}
	env := app.GetEnvByName("prod", cfg)
	if env == nil || env.Name != "prod" || env.BaseUrl != "http://prod" {
		t.Errorf("expected prod env, got %+v", env)
	}
}

func TestGetEnvByName_NotFound(t *testing.T) {
	cfg := models.ProjectConfig{
		ConfiguredEnv: []models.Environment{
			{Name: "dev", BaseUrl: "http://dev"},
		},
	}
	env := app.GetEnvByName("staging", cfg)
	if env != nil {
		t.Errorf("expected nil, got %+v", env)
	}
}

func TestSwitchEnv_SameEnv(t *testing.T) {
	// Setup a temp project config file
	dir := t.TempDir()
	os.Setenv("EX_CRL_REQUESTS_DIR", dir)
	os.Setenv("EX_CRL_TEST_MODE", "1")
	defer os.Unsetenv("EX_CRL_REQUESTS_DIR")
	defer os.Unsetenv("EX_CRL_TEST_MODE")

	// Save config to file
	f := dir + "/projectconfig.json"
	file, _ := os.Create(f)
	file.Close()
	os.WriteFile(f, []byte(`{"project":"proj","active-env":{"name":"dev","base-url":"http://dev"},"configured-env":[{"name":"dev","base-url":"http://dev"},{"name":"prod","base-url":"http://prod"}]}`), 0644)
	os.Setenv("EX_CRL_PROJECT_CONFIG", f)
	defer os.Unsetenv("EX_CRL_PROJECT_CONFIG")
	// Should not panic or error
	app.SwitchEnv("dev")
}
