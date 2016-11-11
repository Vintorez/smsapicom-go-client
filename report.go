package smsapicom

type Report struct {
	Count int              `json:"count"`
	List  []reportByNumber `json:"list"`
}

type reportByNumber struct {
	Id              string  `json:"id"`
	Idx             string  `json:"idx"`
	Points          float64 `json:"points"`
	Number          string  `json:"number"`
	SubmittedNumber string  `json:"submitted_number"`
	DateSent        int64   `json:"date_sent"`
	Status          string  `json:"status"`
	Error           string  `json:"error"`
}

type DetailedReport struct {
	Report
	Length  int    `json:"length"`
	Parts   int    `json:"parts"`
	Message string `json:"message"`
}

type ErrorReport struct {
	Message        string          `json:"message"`
	Error          int             `json:"error"`
	InvalidNumbers []invalidNumber `json:"invalid_numbers"`
}

type invalidNumber struct {
	Number          string `json:"number"`
	SubmittedNumber string `json:"submitted_number"`
	Message         string `json:"message"`
}

type rawReport struct {
	Count          int
	List           []reportByNumber
	Length         int
	Parts          int
	Message        string
	Error          int
	InvalidNumbers []invalidNumber
}

func (r rawReport) HasError() bool {
	return r.Error != 0
}

func (r rawReport) ToError() *Error {
	return &Error{code: r.Error, message: r.Message}
}

func (r rawReport) ToReport() *Report {
	return &Report{
		Count: r.Count,
		List:  r.List,
	}
}

func (r rawReport) ToDetailedReport() *DetailedReport {
	return &DetailedReport{
		Report:  Report{Count: r.Count, List: r.List},
		Length:  r.Length,
		Parts:   r.Parts,
		Message: r.Message,
	}
}
