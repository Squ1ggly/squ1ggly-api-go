package functions

import (
	"net/http"
	"squ1ggly/squ1ggly-api-go/types"
	"strings"
)

func Fetch(url string, opts types.RequestOptions) (*http.Response, error) {
	client := &http.Client{}

	// If the method is not provided, default to GET
	method := opts.Method
	if method == "" {
		method = "GET"
	}

	// If the body is provided, create a request with the body
	var req *http.Request
	var err error
	if opts.Body != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(opts.Body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	// Set optional headers if provided
	if opts.Headers != nil {
		for key, value := range opts.Headers {
			req.Header.Set(key, value)
		}
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
