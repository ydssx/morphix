package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient interface {
	Get(url string, result interface{}) error
	Post(url string, payload interface{}, result interface{}) error
	}

// Request represents a HTTP request
type Request struct {
	client      *resty.Client
	headers     map[string]string
	query       map[string]string
	formData    map[string]string
	timeout     time.Duration
	contentType string
}

// NewRequest creates a new Request instance with default settings
func NewRequest() *Request {
	return &Request{
		client:   resty.New(),
		headers:  make(map[string]string),
		query:    make(map[string]string),
		formData: make(map[string]string),
		timeout:  time.Second * 10,
	}
}

// SetHeader sets the request header
func (r *Request) SetHeader(key, value string) *Request {
	r.headers[key] = value
	return r
}

// SetHeaders sets the request headers from the given map.
// The key-value pairs in the headers map will be set as the request headers.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for key, value := range headers {
		r.headers[key] = value
	}
	return r
}

// SetQuery sets the query parameters for the request
func (r *Request) SetQuery(key, value string) *Request {
	r.query[key] = value
	return r
}

// SetQuerys sets multiple query parameters from a map.
// The key-value pairs in the querys map will be set as the query parameters.
func (r *Request) SetQuerys(querys map[string]string) *Request {
	for k, v := range querys {
		r.SetQuery(k, v)
	}
	return r
}

// SetFormData sets the form data for the request
func (r *Request) SetFormData(key, value string) *Request {
	r.formData[key] = value
	return r
}

// SetTimeout sets the timeout duration for the request
func (r *Request) SetTimeout(timeout time.Duration) *Request {
	r.timeout = timeout
	return r
}

// WithContentType sets the content type of the request.
func (r *Request) WithContentType(contentType string) *Request {
	r.contentType = contentType
	return r
}

// Get performs a GET request to the specified URL and stores the response in the provided result variable.
//
//  - url: the URL to send the GET request to.
//  - result: a pointer to a variable that will store the response data.
// error: an error that occurred during the request, if any.
// Return type: error, returns any error that occurred during the request.
func (r *Request) Get(url string, result interface{}) error {
	req := r.client.R()
	r.addHeaders(req)
	r.addQueryParams(req)
	r.setContentType(req)
	resp, err := req.Get(url)
	if err != nil {
		return err
	}
	return r.handleResponse(resp, result)
}

// Post sends a POST request to the specified URL with the given payload and stores the response in the provided result.
//
// Parameters:
//  - url: The URL to send the request to.
//  - payload: The data to send in the request body.
//  - result: A pointer to the variable where the response will be stored.
//
// Returns:
//  - error: An error if there was a problem sending the request or handling the response.
func (r *Request) Post(url string, payload interface{}, result interface{}) error {
	req := r.client.R()
	r.addHeaders(req)
	r.addQueryParams(req)
	r.setContentType(req)
	resp, err := req.SetBody(payload).Post(url)
	if err != nil {
		return err
	}
	return r.handleResponse(resp, result)
}

// Put sends a PUT request.
func (r *Request) Put(url string, payload interface{}, result interface{}) error {
	req := r.client.R()
	r.addHeaders(req)
	r.addQueryParams(req)
	r.setContentType(req)
	resp, err := req.SetBody(payload).Put(url)
	if err != nil {
		return err
	}
	return r.handleResponse(resp, result)
}

// Delete sends a DELETE request.
func (r *Request) Delete(url string, result interface{}) error {
	req := r.client.R()
	r.addHeaders(req)
	r.addQueryParams(req)
	r.setContentType(req)
	resp, err := req.Delete(url)
	if err != nil {
		return err
	}
	return r.handleResponse(resp, result)
}

func (r *Request) addHeaders(req *resty.Request) {
	for key, value := range r.headers {
		req.SetHeader(key, value)
	}
}

func (r *Request) addQueryParams(req *resty.Request) {
	for key, value := range r.query {
		req.SetQueryParam(key, value)
	}
}

func (r *Request) setContentType(req *resty.Request) {
	if r.contentType != "" {
		req.SetHeader("Content-Type", r.contentType)
	}
}

func (r *Request) handleResponse(resp *resty.Response, result interface{}) (err error) {
	if resp.StatusCode() >= http.StatusBadRequest {
		return fmt.Errorf("request failed with status code %d", resp.StatusCode())
	}
	if resp != nil {
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return err
		}
	}

	return nil
}
