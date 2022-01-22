package f3

// err is internally used to return errors.
type err struct {
	errCode int32
	errMsg  string
}

// Err is an interface to return errors, which is compatible to the standard error interface, but additionally supports
// an error code. I do not get why errors are returned as strings, comparing a string is much more expensive than
// comparing an integer (it needs to be done byte by byte). Apart from that the error message may contain volatile
// information, for example like "The file '/tmp/test.foo' was not found". So we use this to have the option to compare
// the error code to find out the error class.
type Err interface {
	Error() string
	ErrorCode() int32
}

// Error returns the human-readable error message.
func (e err) Error() string {
	return e.errMsg
}

// ErrorCode returns the machine-readable error reason, basically the error class. Sometimes this is helpful to
// recover from certain errors.
func (e err) ErrorCode() int32 {
	return e.errCode
}

const (
	ErrUnknown             int32 = iota // An error that has no error code happened.
	ErrEmptyHostName       int32 = iota // Returned when a new client is created with an empty host name.
	ErrIllegalPort         int32 = iota // Returned when a new client is given with an invalid port.
	ErrInvalidUuid         int32 = iota // When an identifier must be a UUID and is not valid UUID string.
	ErrInvalidStatus       int32 = iota // When an invalid status is given
	ErrMissingStatusReason int32 = iota // When a status with mandatory reason is given without reason.
	ErrCountryUnknown      int32 = iota // When a new account could not be initialized for the provided country.
	ErrBankIdMissing       int32 = iota // When a bank-id is required, but none was given.
	ErrBankIdTooShort      int32 = iota // When a bank-id is too short.
	ErrBankIdTooLong       int32 = iota // When a bank-id is too long.
	ErrAccountNil          int32 = iota // When the account structure is nil.
	ErrJsonStringify       int32 = iota // When serialization to JSON fails.
	ErrJsonParse           int32 = iota // When parsing JSON fails.
	ErrRequest             int32 = iota // When the request itself failed.
	ErrResponse            int32 = iota // When the response is not of the expected form.
)

// Internally used as default message.
const (
	msgEmptyHostName       = "Empty hostname"
	msgIllegalPort         = "Illegal port, must be a value between 1 and 65535"
	msgInvalidUuid         = "Invalid UUID format"
	msgMissingStatusReason = "The status 'failed' requires that a reason is given"
	msgCountryUnknown      = "The given country is not well defined in the client, manual setup is required"
	msgBankIdMissing       = "Empty (missing) bank-id given, but a valid bank-id is required in this country"
	msgAccountNil          = "The account object is nil"
	msgMissingBody         = "Missing body"
	msgInvalidResponse     = "Invalid response"
)
