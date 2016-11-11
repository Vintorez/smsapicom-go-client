package smsapicom

import "fmt"

const (
	// Internal errors
	SendErr       = 1001
	RequestErr    = 1002
	DecodeJsonErr = 1003

	// API errors
	RequestApiErr               = 8
	InvalidMessageApiErr        = 11
	IncorrectPartsCountApiErr   = 12
	InvalidNumberApiErr         = 13
	WrongSenderNameApiErr       = 14
	InvalidFlashMsgApiErr       = 17
	InvalidNumberOfParamsApiErr = 18
)

type Error struct {
	code    int
	message string
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, message: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

func (e Error) GetCode() int {
	return e.code
}

func (e *Error) SetCode(code int) {
	e.code = code
}

func (e Error) GetMessage() string {
	return e.message
}

func (e *Error) SetMessage(msg string) {
	e.message = msg
}
