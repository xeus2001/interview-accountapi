package f3

import "fmt"

// Err is an interface to return errors, which is compatible to the standard error interface.
type Err interface {
	Unwrap() error
	Error() string
	ErrorCode() int
}

// err is the internal implementation of the Err interface.
type err struct {
	code    int
	message string
	cause   error
}

func pErr(e err) *err {
	return &e
}

// Error returns the human-readable err message.
func (e err) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.message)
}

// ErrorCode returns the machine-readable error code.
func (e err) ErrorCode() int {
	return e.code
}

// Unwrap returns the cause.
func (e err) Unwrap() error {
	return e.cause
}

const (
	// ErrUnknown is an unknown error, for example if the err structure has not been initialized correctly.
	ErrUnknown int = iota

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
	ErrAccountNil          int = iota // When the account structure is nil.
	ErrJsonStringify       int = iota // When serialization to JSON fails.
	ErrJsonParse           int = iota // When parsing JSON fails.
	ErrRequest             int = iota // When the request itself failed.
	ErrResponse            int = iota // When the response is not of the expected form.
)

// Internally used as default message.
const (
	msgUuidCreationFailed  = "UUID creation failed"
	msgEmptyHostName       = "Empty hostname"
	msgIllegalPort         = "Illegal port, must be a value between 1 and 65535"
	msgInvalidUuid         = "Invalid UUID format"
	msgMissingStatusReason = "The status 'failed' requires that a cause is given"
	msgCountryUnknown      = "The given country is not well defined in the client, manual setup is required"
	msgBankIdMissing       = "Empty (missing) bank-id given, but a valid bank-id is required in this country"
	msgAccountNil          = "The account object is nil"
	msgMissingBody         = "Missing body"
	msgInvalidResponse     = "Invalid response"
)
