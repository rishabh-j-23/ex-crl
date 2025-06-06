package models

// Request defines the structure of an HTTP request with support for dynamic
// path and query parameters, custom headers, and a JSON body.
type Request struct {
	Name        string            `json:"name"`
	HttpMethod  string            `json:"http-method"`
	Endpoint    string            `json:"endpoint"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"query-params"`
	PathParams  map[string]string `json:"path-params"`
	Body        any               `json:"body"`
}
