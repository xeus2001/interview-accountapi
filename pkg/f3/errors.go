package f3

import (
	"fmt"
	"net/http"
)

// Err is an interface to return errors, which is compatible to the standard error interface.
type Err interface {
	// Unwrap returns the cause.
	Unwrap() error
	// Error returns the human-readable err message.
	Error() string
	// ErrorCode returns the machine-readable error code.
	ErrorCode() int
	// Request returns the HTTP request that is the reason for this error; if any.
	Request() *http.Request
	// Response returns the HTTP response that is the reason for this error; if any.
	Response() *http.Response
}

// err is the internal implementation of the Err interface.
type err struct {
	code  int
	msg   string
	cause error
	req   *http.Request
	resp  *http.Response
}

func errAddr(e err) *err {
	return &e
}

func (e err) Request() *http.Request {
	return e.req
}

func (e err) Response() *http.Response {
	return e.resp
}

func (e err) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.msg)
}

func (e err) ErrorCode() int {
	return e.code
}

func (e err) Unwrap() error {
	return e.cause
}

const (
	// ErrGeneric signals a generic error.
	ErrGeneric int = iota

	// ErrJsonStringify signals that serialization to JSON failed.
	ErrJsonStringify int = iota

	// ErrJsonParse signals that the JSON parsing failed.
	ErrJsonParse int = iota

	// ErrRequest signals that the request failed.
	ErrRequest int = iota

	// ErrResponse signals that the response is erroneous.
	ErrResponse int = iota
)
