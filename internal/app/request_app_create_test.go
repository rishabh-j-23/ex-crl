package app_test

import (
	app "github.com/rishabh-j-23/ex-crl/internal/app"
	"github.com/rishabh-j-23/ex-crl/internal/models"

	"testing"
)

func TestCreateRequest_Basic(t *testing.T) {
	reqModel := models.Request{
		Name:        "test",
		HttpMethod:  "POST",
		Endpoint:    "/api/test",
		Headers:     map[string]string{"X-Test": "yes"},
		Body:        map[string]any{"foo": "bar"},
		QueryParams: map[string]string{"q": "1"},
		PathParams:  map[string]string{"id": "123"},
	}
	proj := models.ProjectConfig{
		Name:          "proj",
		ActiveEnv:     models.Environment{Name: "dev", BaseUrl: "https://api.example.com"},
		GlobalHeaders: map[string]string{"X-Global": "g"},
	}

	httpReq, err := app.CreateRequest(reqModel, proj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpReq.Method != "POST" {
		t.Errorf("expected POST, got %s", httpReq.Method)
	}
	if httpReq.URL.Scheme != "https" || httpReq.URL.Host != "api.example.com" {
		t.Errorf("unexpected URL: %s", httpReq.URL.String())
	}
	if httpReq.Header.Get("X-Global") != "g" {
		t.Errorf("global header missing")
	}
	if httpReq.Header.Get("X-Test") != "yes" {
		t.Errorf("request header missing")
	}
	if httpReq.URL.Query().Get("q") != "1" {
		t.Errorf("query param missing")
	}
}

func TestCreateRequest_PathParams(t *testing.T) {
	reqModel := models.Request{
		Name:       "test",
		HttpMethod: "GET",
		Endpoint:   "/api/:id/:type",
		PathParams: map[string]string{"id": "42", "type": "user"},
	}
	proj := models.ProjectConfig{
		Name:      "proj",
		ActiveEnv: models.Environment{Name: "dev", BaseUrl: "http://localhost:8080"},
	}

	httpReq, err := app.CreateRequest(reqModel, proj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpReq.URL.Path != "/api/42/user" {
		t.Errorf("expected /api/42/user, got %s", httpReq.URL.Path)
	}
}
