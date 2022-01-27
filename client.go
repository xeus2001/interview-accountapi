package f3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	headerUserAgent   = "User-Agent"
	headerAccept      = "Accept"
	headerContentType = "Content-Type"

	userAgentName       = "f3.Client;v=0.1"
	mimeApplicationJson = "application/json"
	mimeForm3Json       = "application/vnd.api+json"
)

// NewClient creates a new Form3 client bound to the production endpoint and setup with defaults.
func NewClient() *Client {
	client := Client{endpoint: *DefaultEndPoint, httpClient: http.Client{Timeout: DefaultTimeout, Transport: DefaultTransport}}
	client.healthCheckUri = fmt.Sprintf("%s/health", client.endpoint)
	client.accountUri = fmt.Sprintf("%s/organisation/accounts", client.endpoint)
	return &client
}

// Client is an abstraction above a http.Client bound to a specific Form3 endpoint.
type Client struct {
	endpoint       string
	healthCheckUri string
	accountUri     string
	httpClient     http.Client
}

// WithEndPoint rebinds the endpoint of the client.
func (c *Client) WithEndPoint(endpoint string) *Client {
	c.endpoint = endpoint
	c.healthCheckUri = fmt.Sprintf("%s/health", endpoint)
	c.accountUri = fmt.Sprintf("%s/organisation/accounts", endpoint)
	return c
}

// HttpClient returns the underlying http client being used. If the default created by NewClient is not sufficient,
// modify this before using.
func (c *Client) HttpClient() http.Client {
	return c.httpClient
}

// IsHealthy tests if the service is alive and responsive within the set request timeout.
func (c *Client) IsHealthy() bool {
	var (
		req  *http.Request
		resp *http.Response
		e    error
		raw  []byte
	)
	if req, e = http.NewRequest(http.MethodGet, c.healthCheckUri, nil); e == nil {
		req.Header.Set(headerUserAgent, userAgentName)
		req.Header.Set(headerAccept, mimeApplicationJson)
		if resp, e = c.httpClient.Do(req); e == nil && resp.Body != nil {
			//goland:noinspection GoUnhandledErrorResult
			defer resp.Body.Close()
			if raw, e = ioutil.ReadAll(resp.Body); e == nil {
				response := HealthyResponse{}
				if e = json.Unmarshal(raw, &response); e == nil {
					return response.Status != nil && *response.Status == "up"
				}
			}
		}
	}
	return false
}

// createRequest creates a new request and returns it. If an object is given, this is JSON serialized and attached as body.
func createRequest[T any](method string, uri string, object *T) (*http.Request, Err) {
	var (
		req *http.Request
		e   error
	)
	if object != nil {
		var jsonBytes []byte
		jsonBytes, e = json.Marshal(object)
		if e != nil {
			return nil, err{code: ErrJsonStringify, msg: "Failed to stringify", cause: e}
		}
		req, e = http.NewRequest(method, uri, bytes.NewBuffer(jsonBytes))
	} else {
		req, e = http.NewRequest(method, uri, nil)
	}
	if e != nil {
		return nil, err{code: ErrGeneric, msg: "Unknown error while creating the request", cause: e, req: req}
	}
	req.Header.Set(headerUserAgent, userAgentName)
	req.Header.Set(headerContentType, mimeForm3Json)
	req.Header.Set(headerAccept, mimeForm3Json)
	return req, nil
}

// parseResponse parses the JSON of the given response into the given object. If no object is given, no response is expected.
func parseResponse[T any](req *http.Request, resp *http.Response, object *T) Err {
	if resp.Body == nil {
		return err{code: ErrRequest, msg: "Body is nil", req: req, resp: resp}
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return err{code: ErrRequest, msg: "Failed to read bytes of returned body", cause: e, req: req, resp: resp}
	}
	if object == nil {
		return nil
	}
	e = json.Unmarshal(body, object)
	if e != nil {
		var errResponse *ErrorResponse
		e2 := json.Unmarshal(body, &errResponse)
		if e2 != nil {
			return err{code: ErrJsonParse, msg: "Failed to parse response", cause: e, req: req, resp: resp}
		}
		return err{code: ErrResponse, msg: errResponse.ErrorMessage, cause: e, req: req, resp: resp}
	}
	return nil
}

// CreateAccount creates the given account and returns the server response or an error, when the account creation
// failed.
func (c *Client) CreateAccount(account *Account) (*Account, Err) {
	if account == nil {
		return nil, err{code: ErrGeneric, msg: "Account is nil"}
	}
	envelope := AccountEnvelope{account}
	req, err0 := createRequest(http.MethodPost, c.accountUri, &envelope)
	if err0 != nil {
		return nil, err0
	}
	resp, e := c.httpClient.Do(req)
	if e != nil {
		return nil, err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
	}
	err0 = parseResponse(req, resp, &envelope)
	if err0 != nil {
		return nil, err0
	}
	if envelope.Data == nil {
		return nil, err{code: ErrResponse, msg: "Received an empty envelope", req: req, resp: resp}
	}
	return envelope.Data, nil
}

func (c *Client) FetchAccount(accountId string) (*Account, Err) {
	uri := fmt.Sprintf("%s/%s", c.accountUri, url.QueryEscape(accountId))
	req, err0 := createRequest(http.MethodGet, uri, (*any)(nil))
	if err0 != nil {
		return nil, err0
	}
	resp, e := c.httpClient.Do(req)
	if e != nil {
		return nil, err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
	}
	var envelope AccountEnvelope
	err0 = parseResponse(req, resp, &envelope)
	if err0 != nil {
		return nil, err0
	}
	if envelope.Data == nil {
		return nil, err{code: ErrResponse, msg: "Received an empty envelope", req: req, resp: resp}
	}
	return envelope.Data, nil
}

func (c *Client) DeleteAccount(accountId string, version uint64) Err {
	uri := fmt.Sprintf("%s/%s?version=%d", c.accountUri, url.QueryEscape(accountId), version)
	req, err0 := createRequest(http.MethodDelete, uri, (*any)(nil))
	if err0 != nil {
		return err0
	}
	resp, e := c.httpClient.Do(req)
	if e != nil {
		return err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
	}
	return parseResponse(req, resp, (*any)(nil))
}
