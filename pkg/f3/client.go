package f3

import (
	"bytes"
	"encoding/json"
	"errors"
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
	client := Client{endpoint: DefaultEndPoint, httpClient: http.Client{Timeout: DefaultTimeout, Transport: DefaultTransport}}
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
		if e == nil {
			req, e = http.NewRequest(method, uri, bytes.NewBuffer(jsonBytes))
		}
	} else {
		req, e = http.NewRequest(method, uri, nil)
	}
	if e == nil && req != nil {
		req.Header.Set(headerUserAgent, userAgentName)
		req.Header.Set(headerContentType, mimeForm3Json)
		req.Header.Set(headerAccept, mimeForm3Json)
		return req, nil
	}
	return nil, err{code: ErrGeneric, msg: "Unknown error while creating the request", cause: e, req: req}
}

// parseResponse parses the JSON of the given response into the given object. If no object is given, no response is expected.
func parseResponse[T any](req *http.Request, resp *http.Response, object *T) Err {
	var (
		e    error
		body []byte
	)
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	body, e = ioutil.ReadAll(resp.Body)
	if e == nil {
		e = json.Unmarshal(body, object)
	}
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return nil
	}
	if body != nil {
		var errResponse *ErrorResponse
		e = json.Unmarshal(body, &errResponse)
		if errResponse != nil {
			e = errors.New(errResponse.ErrorMessage)
		}
	}
	if resp.StatusCode == 404 {
		return err{code: ErrNotFound, msg: "Not found", cause: e, req: req, resp: resp}
	}
	return err{code: ErrBadRequest, msg: "Bad Request: The given payload was invalid", cause: e, req: req, resp: resp}
}

// CreateAccount creates the given account and returns the new account as returned from the server or an error, when
// the account creation failed.
func (c *Client) CreateAccount(account *Account) (*Account, Err) {
	var (
		req  *http.Request
		resp *http.Response
		er   Err
		e    error
	)
	if account != nil {
		envelope := AccountEnvelope{account}
		req, er = createRequest(http.MethodPost, c.accountUri, &envelope)
		if er == nil && req != nil {
			resp, e = c.httpClient.Do(req)
			if e == nil && resp != nil {
				er = parseResponse(req, resp, &envelope)
				if er == nil && envelope.Data != nil {
					return envelope.Data, nil
				}
			}
		}
	}
	if er != nil {
		return nil, er
	}
	return nil, err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
}

// FetchAccount returns the account with the given id or ErrNotFound if the account does not exist.
func (c *Client) FetchAccount(accountId string) (*Account, Err) {
	var (
		req  *http.Request
		resp *http.Response
		er   Err
		e    error
	)
	uri := fmt.Sprintf("%s/%s", c.accountUri, url.QueryEscape(accountId))
	req, er = createRequest(http.MethodGet, uri, (*any)(nil))
	if er == nil && req != nil {
		resp, e = c.httpClient.Do(req)
		if e == nil && resp != nil {
			var envelope AccountEnvelope
			er = parseResponse(req, resp, &envelope)
			if er == nil && envelope.Data != nil {
				return envelope.Data, nil
			}
		}
	}
	if er != nil {
		return nil, er
	}
	return nil, err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
}

// DeleteAccount deletes the account with the given id and return nil. If the account does not exist, ErrNotFound is
// returned.
func (c *Client) DeleteAccount(accountId string, version uint64) Err {
	var (
		req  *http.Request
		resp *http.Response
		er   Err
		e    error
	)
	uri := fmt.Sprintf("%s/%s?version=%d", c.accountUri, url.QueryEscape(accountId), version)
	req, er = createRequest(http.MethodDelete, uri, (*any)(nil))
	if er == nil && req != nil {
		resp, e = c.httpClient.Do(req)
		if e == nil && resp != nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
				return nil
			}
			if resp.StatusCode == 409 {
				return err{code: ErrConflict, msg: "Conflict, version does not match", req: req, resp: resp}
			}
			if resp.StatusCode == 404 {
				return err{code: ErrNotFound, msg: "Account does not exist", req: req, resp: resp}
			}
		}
	}
	if er != nil {
		return er
	}
	return err{code: ErrRequest, msg: "Request failed", cause: e, req: req, resp: resp}
}
