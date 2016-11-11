package smsapicom

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// A client manages communication with smsapi.com API
type client struct {
	// HTTP client
	httpClient *http.Client
	// User agent
	userAgent string
	// API base URL
	baseUrl *url.URL
	// Credentials which is used for authentication during API request
	username string
	password string
}

// NewClient returns a new smsapi.com API client.
// This will load default http.Client if httpClient is nil.
func newClient(username, password string, httpClient *http.Client) *client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(ApiUrl)

	return &client{
		httpClient: httpClient,
		userAgent:  UserAgent,
		baseUrl:    baseUrl,
		username:   username,
		password:   password,
	}
}

// NewRequest returns a new Request given a method, URL, and optional body.
func (c *client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	urlPath, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	fullUrl := c.baseUrl.ResolveReference(urlPath)

	r, err := http.NewRequest(method, fullUrl.String(), body)
	if err != nil {
		return nil, err
	}

	r.SetBasicAuth(c.username, c.password)

	r.Header.Add("User-Agent", c.userAgent)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Accept-Charset", "utf-8")
	if method == "POST" {
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	return r, nil
}

// Do sends an HTTP request and returns an HTTP response, following
// policy as configured on the client.
func (c *client) Do(r *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := NewResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	if v != nil {
		err := json.NewDecoder(response.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}
