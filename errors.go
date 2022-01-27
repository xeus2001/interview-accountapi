package interview_accountapi

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

	// ErrUuidCreationFailed is returned when creating a UUID fails.
	ErrUuidCreationFailed int = iota

	// ErrNoVerifierAvailableForCountry is returned when there is no verifier available for the requested country.
	ErrNoVerifierAvailableForCountry int = iota

	// ErrBankIdMissing is returned when a bank-id is required, but none was given.
	ErrBankIdMissing int = iota

	// ErrBankIdTooShort is returned when a bank-id is too short.
	ErrBankIdTooShort int = iota

	// ErrBankIdTooLong is returned when a bank-id is too long.
	ErrBankIdTooLong int = iota

	ErrEmptyHostName       int = iota // Returned when a new client is created with an empty host name.
	ErrIllegalPort         int = iota // Returned when a new client is given with an invalid port.
	ErrInvalidUuid         int = iota // When an identifier must be a UUID and is not valid UUID string.
	ErrInvalidStatus       int = iota // When an invalid status is given
	ErrMissingStatusReason int = iota // When a status with mandatory cause is given without cause.
)

// Internally used as default message.
const (
	msgUnknownErrorWhileCreatingMessage = "Unknown error while creating the request"
	msgAccountNil                       = "The account object is nil"
	msgUuidCreationFailed               = "UUID creation failed"
	msgEmptyHostName                    = "Empty hostname"
	msgIllegalPort                      = "Illegal port, must be a value between 1 and 65535"
	msgInvalidUuid                      = "Invalid UUID format"
	msgMissingStatusReason              = "The status 'failed' requires that a cause is given"
	msgCountryUnknown                   = "The given country is not well defined in the client, manual setup is required"
	msgBankIdMissing                    = "Empty (missing) bank-id given, but a valid bank-id is required in this country"
	msgMissingBody                      = "Missing body"
	msgInvalidResponse                  = "Invalid response"
)
