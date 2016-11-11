package smsapicom

import "net/http"

// Response wraps http.Response. So we can add more functionality later.
type Response struct {
	*http.Response
}

func NewResponse(r *http.Response) *Response {
	return &Response{Response: r}
}
