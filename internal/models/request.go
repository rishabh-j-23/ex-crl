package models

type Request struct {
	Name       string            `json:"name"`
	HttpMethod string            `json:"http-method"`
	Endpoint   string            `json:"endpoint"`
	Headers    map[string]string `json:"headers"`
	Body       any               `json:"body"`
}
