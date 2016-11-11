package smsapicom

import (
	"net/http"
	"net/url"
	"strings"
)

type MessageParams struct {
	details bool `json:"details"`
	Test    bool `json:"test"`
	Fast    bool `json:"fast"`
	Flash   bool `json:"flash"`
}

type SmsService struct {
	client  *client
	urlPath string
}

func NewSmsService(username, password string, httpClient *http.Client) *SmsService {
	return &SmsService{
		client:  newClient(username, password, httpClient),
		urlPath: "sms.do",
	}
}

// Send sends message with options.
func (s *SmsService) Send(from, to, message string, params *MessageParams) (*Report, error) {
	data := s.createRequestData(from, to, message, params, false)

	raw, err := s.send(data)
	if err != nil {
		return nil, err
	}

	return s.getReport(raw)
}

// SendWithDetailedReport sends message with options and returns detailed report.
func (s *SmsService) SendWithDetailedReport(from, to, message string, params *MessageParams) (*DetailedReport, error) {
	data := s.createRequestData(from, to, message, params, true)

	raw, err := s.send(data)
	if err != nil {
		return nil, err
	}

	return s.getDetailedReport(raw)
}

func (s *SmsService) createRequestData(from, to, message string, params *MessageParams, details bool) url.Values {
	data := url.Values{}
	data.Set("format", ApiFormat)
	data.Set("from", from)
	data.Add("to", to)
	data.Set("message", message)
	if details {
		data.Set("details", "1")
	}
	if params != nil {
		if params.Test {
			data.Set("test", "1")
		}
		if params.Fast {
			data.Set("fast", "1")
		}
		if params.Flash {
			data.Set("flash", "1")
		}
	}

	return data
}

func (s *SmsService) send(data url.Values) (*rawReport, *Error) {
	r, err := s.client.NewRequest("POST", s.urlPath, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, NewError(RequestErr, err.Error())
	}

	raw := &rawReport{}
	_, err = s.client.Do(r, raw)
	if err != nil {
		return nil, NewError(SendErr, err.Error())
	}

	return raw, nil
}

func (s *SmsService) getReport(raw *rawReport) (*Report, *Error) {
	if raw.HasError() {
		return nil, raw.ToError()
	}

	return raw.ToReport(), nil
}

func (s *SmsService) getDetailedReport(raw *rawReport) (*DetailedReport, *Error) {
	if raw.HasError() {
		return nil, raw.ToError()
	}

	return raw.ToDetailedReport(), nil
}
