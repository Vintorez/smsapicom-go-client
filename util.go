package smsapicom

import (
	"encoding/json"
	"net/http"
)

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	e := struct{ Message string }{}
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return err
	}

	return &Error{code: r.StatusCode, message: e.Message}
}
