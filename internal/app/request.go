package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func AddRequest(httpMethod, requestName, endpoint string) {
	if httpMethod == "" || requestName == "" || endpoint == "" {
		panic("httpMethod, requestName, and endpoint are required")
	}
	requestsDir := utils.GetRequestsDir()
	assert.EnsureDirExists(requestsDir)

	filePath := filepath.Join(requestsDir, requestName+".json")

	// Prevent overwrite if file exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("Request '%s' already exists\n", requestName)
		if os.Getenv("EX_CRL_TEST_MODE") != "" {
			panic("duplicate request")
		}
		os.Exit(1)
	}

	request := models.Request{
		Name:        requestName,
		HttpMethod:  strings.ToUpper(httpMethod),
		Headers:     map[string]string{},
		Body:        map[string]any{},
		Endpoint:    endpoint,
		QueryParams: map[string]string{},
		PathParams:  map[string]string{},
	}

	data, err := json.MarshalIndent(request, "", "  ")
	assert.ErrIsNil(err, "Failed to marshal request")

	err = os.WriteFile(filePath, data, 0644)
	assert.ErrIsNil(err, "Failed to write request file")

	if os.Getenv("EX_CRL_SKIP_EDITOR") == "" {
		editor.LaunchEditor(filePath)
	}

	fmt.Printf("Request '%s' added at %s\n", requestName, filePath)
}

// CreateRequest constructs a complete *http.Request based on the provided Request
// and ProjectConfig. It builds the full URL from the base URL and endpoint,
// serializes the body (if provided), and applies path parameters, query parameters,
// and headers (with Request headers taking precedence over global ones).
func CreateRequest(r models.Request, p models.ProjectConfig) (*http.Request, error) {
	url := p.ActiveEnv.BaseUrl + r.Endpoint

	var bodyReader *bytes.Reader
	if r.Body != nil {
		bodyBytes, err := json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(r.HttpMethod, url, bodyReader)
	if err != nil {
		return nil, err
	}

	resolveRequestHeaders(req, &p, &r)
	resolveRequestQueryParams(req, &r)
	resolvePathParams(req, &r)

	return req, nil
}

// resolveRequestHeaders sets the headers for the HTTP request.
// It first applies global headers from the ProjectConfig, and then applies
// request-specific headers from the Request struct, overriding global ones if needed.
func resolveRequestHeaders(req *http.Request, p *models.ProjectConfig, r *models.Request) {
	for k, v := range p.GlobalHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
}

// resolveRequestParams updates the request's URL with query parameters defined
// in the Request.QueryParams map.
func resolveRequestQueryParams(req *http.Request, r *models.Request) {
	query := req.URL.Query()
	for k, v := range r.QueryParams {
		query.Set(k, v)
	}
	req.URL.RawQuery = query.Encode()
}

// resolvePathParams replaces path parameters in the HTTP request URL using values
// provided in the Request.PathParams map.
//
// It searches for placeholders in the format `:paramName` and replaces them with
// their corresponding values from the PathParams map. For example, if the original
// path is "/api/:id/:type" and PathParams contains {"id": "123", "type": "user"},
// the path becomes "/api/123/user".
//
// Note: This function updates the request's URL.Path in-place and should be called
// after constructing the full URL, including any base path from environment configs.
//
// To prevent partial replacements (e.g., "id" matching "idLong"), path parameter
// keys are sorted by descending length before replacement.
func resolvePathParams(req *http.Request, r *models.Request) {
	path := req.URL.Path

	keys := make([]string, 0, len(r.PathParams))
	for k := range r.PathParams {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	for _, key := range keys {
		val := r.PathParams[key]
		placeholder := ":" + key
		path = strings.ReplaceAll(path, placeholder, val)
	}

	req.URL.Path = path
}
