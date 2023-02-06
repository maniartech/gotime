package temporal

import "fmt"

const errInvalidFormat = "invalid-format"
const errInvalidFormatMsg = "The specified format is invalid: %s"

const errInvalidArgument = "invalid-argument"
const errInvalidArgumentMsg = "The argument '%s' is invalid"
const errInvalidArgumentExpectingMsg = "The argument '%s' is invalid, expecting '%s'"

const errInvalidType = "invalid-type"
const errInvalidTypeMsg = "Invalid type: %s"
const errInvalidTypeExpectingMsg = "Invalid type: %s, expecting %s"

type TemporalError struct {

	// Code is the error code, e.g. "invalid-argument". It is used to identify
	// the error type.
	Code string `json:"code"`

	// Message is the error message, e.g. "Invalid argument: %s". It is used to
	// provide more information about the error.
	Message string `json:"message"`
}

func (e *TemporalError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func newEror(code string, messageFmt string, args ...interface{}) error {
	return &TemporalError{
		Code:    code,
		Message: fmt.Sprintf(messageFmt, args...),
	}
}
