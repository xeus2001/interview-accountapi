package f3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// ErrEmptyHostName is returned when a new client is created with an empty host name.
	ErrEmptyHostName int32 = iota
	// ErrIllegalPort is returned when a new client is given with an invalid port.
	ErrIllegalPort int32 = iota
)

const (
	msgEmptyHostName = "Empty hostname"
	msgIllegalPort   = "Illegal port, must be a value between 1 and 65535"
)

// HTTP headers.
const (
	headerUserAgent = "User-Agent"
	headerAccept    = "Accept"
)

// HTTP header values.
const (
	userAgentName       = "f3.Client"
	mimeApplicationJson = "application/json"
)

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

// Client is a structure to store options needed to access a specific Form3 service. This object is stateless.
type Client struct {
	host          string
	port          int32
	endpoint      string
	epHealthCheck *string
	httpClient    http.Client
}

// healthyResponse is the response return by the health-check API.
type healthyResponse struct {
	// status is the status returned by the service, which may be nil.
	Status *string `json:"status,omitempty"`
}

// NewClient creates a new Form3 client for the given host and port using the given request timeout.
func NewClient(host string, port int32, timeout time.Duration) (*Client, Err) {
	if len(host) == 0 {
		return nil, err{ErrEmptyHostName, msgEmptyHostName}
	}
	if port < 1 || port > 65535 {
		return nil, err{ErrIllegalPort, msgIllegalPort}
	}
	client := Client{host: host, port: port, endpoint: fmt.Sprintf("http://%s:%d/v1/", host, port), httpClient: http.Client{Timeout: timeout}}
	return &client, nil
}

// IsHealthy tests if the service is alive and responsive within the set request timeout.
func (c *Client) IsHealthy() bool {
	// This avoids too many string concatenations and doing it dynamically saves memory,
	// when a client does not need certain APIs.
	if c.epHealthCheck == nil {
		ep := fmt.Sprintf("%s/health", c.endpoint)
		c.epHealthCheck = &ep
	}
	req, err := http.NewRequest(http.MethodGet, *c.epHealthCheck, nil)
	if err != nil {
		return false
	}
	req.Header.Set(headerUserAgent, userAgentName)
	req.Header.Set(headerAccept, mimeApplicationJson)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false
	}
	if resp.Body == nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	response := healthyResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false
	}
	if response.Status == nil {
		return false
	}
	return *response.Status == "up"
}
