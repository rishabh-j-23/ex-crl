package models

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Name       string            `json:"name"`
	HttpMethod string            `json:"http-method"`
	Endpoint   string            `json:"endpoint"`
	Headers    map[string]string `json:"headers"`
	Body       any               `json:"body"`
}

func CreateRequest(r Request, p ProjectConfig) (*http.Request, error) {
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

	for k, v := range p.GlobalHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
