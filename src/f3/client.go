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
)

// NewClient creates a new Form3 client for the given endpoint using default settings.
func NewClient(endpoint string) *Client {
	client := Client{endpoint: endpoint, httpClient: http.Client{Timeout: DefaultTimeout, Transport: DefaultTransport}}
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

// HttpClient returns the underlying http client being used. If the default created by NewClient is not sufficient,
// modify this before using.
func (c *Client) HttpClient() http.Client {
	return c.httpClient
}

// IsHealthy test if the service is alive and responsive within the set request timeout.
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

// CreateAccount creates the given account and returns the server response or an error, when the account creation
// failed.
func (c *Client) CreateAccount(account *Account) (*Account, Err) {
	if account == nil {
		return nil, err{ErrAccountNil, msgAccountNil, nil}
	}
	envelope := AccountEnvelope{account}
	jsonBytes, e := json.Marshal(envelope)
	if e != nil {
		return nil, err{ErrJsonStringify, e.Error(), e}
	}

	req, e := http.NewRequest(http.MethodPost, c.accountUri, bytes.NewBuffer(jsonBytes))
	if e != nil {
		return nil, err{ErrUnknown, e.Error(), e}
	}
	req.Header.Set(headerUserAgent, userAgentName)
	req.Header.Set(headerContentType, mimeApplicationJson)
	req.Header.Set(headerAccept, mimeApplicationJson)
	resp, e := c.httpClient.Do(req)
	if e != nil {
		// How do we know why this request failed?
		return nil, err{ErrRequest, e.Error(), e}
	}
	if resp.Body == nil {
		return nil, err{ErrRequest, msgMissingBody, nil}
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, err{ErrRequest, e.Error(), e}
	}
	e = json.Unmarshal(body, &envelope)
	if e != nil {
		var errResponse *ErrorResponse
		e = json.Unmarshal(body, &errResponse)
		if e != nil {
			return nil, err{ErrJsonParse, e.Error(), e}
		}
		return nil, err{ErrResponse, errResponse.ErrorMessage, e}
	}
	if envelope.Data == nil {
		return nil, err{ErrResponse, msgInvalidResponse, nil}
	}
	return envelope.Data, nil
}

func (c *Client) FetchAccount(accountId string) (*Account, Err) {
	uri := fmt.Sprintf("%s/%s", c.accountUri, url.QueryEscape(accountId))
	req, e := http.NewRequest(http.MethodGet, uri, nil)
	if e != nil {
		return nil, err{ErrUnknown, e.Error(), e}
	}
	req.Header.Set(headerUserAgent, userAgentName)
	req.Header.Set(headerAccept, mimeApplicationJson)
	resp, e := c.httpClient.Do(req)
	if e != nil {
		// How do we know why this request failed?
		return nil, err{ErrRequest, e.Error(), e}
	}
	if resp.Body == nil {
		return nil, err{ErrRequest, msgMissingBody, nil}
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, err{ErrRequest, e.Error(), e}
	}
	var envelope AccountEnvelope
	e = json.Unmarshal(body, &envelope)
	if e != nil {
		var errResponse *ErrorResponse
		e = json.Unmarshal(body, &errResponse)
		if e != nil {
			return nil, err{ErrJsonParse, e.Error(), e}
		}
		return nil, err{ErrResponse, errResponse.ErrorMessage, e}
	}
	if envelope.Data == nil {
		return nil, err{ErrResponse, msgInvalidResponse, nil}
	}
	return envelope.Data, nil
}

func (c *Client) DeleteAccount(accountId string, version uint64) Err {
	uri := fmt.Sprintf("%s/%s?version=%d", c.accountUri, url.QueryEscape(accountId), version)
	req, e := http.NewRequest(http.MethodDelete, uri, nil)
	if e != nil {
		return err{ErrUnknown, e.Error(), e}
	}
	req.Header.Set(headerUserAgent, userAgentName)
	req.Header.Set(headerAccept, mimeApplicationJson)
	resp, e := c.httpClient.Do(req)
	if e != nil {
		// How do we know why this request failed?
		return err{ErrRequest, e.Error(), e}
	}
	if resp.Body == nil {
		return err{ErrRequest, msgMissingBody, nil}
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return err{ErrRequest, e.Error(), e}
	}
	var envelope AccountEnvelope
	e = json.Unmarshal(body, &envelope)
	if e != nil {
		var errResponse *ErrorResponse
		e = json.Unmarshal(body, &errResponse)
		if e != nil {
			return err{ErrJsonParse, e.Error(), e}
		}
		return err{ErrResponse, errResponse.ErrorMessage, e}
	}
	if envelope.Data == nil {
		return err{ErrResponse, msgInvalidResponse, nil}
	}
	return nil
}

/*
Host: api.form3.tech (note this is different when using the Staging Environment)
Date: [The date and time the request is made]
Accept: application/vnd.api+json

Requests that contain body also require the following headers:

Content-Type: application/vnd.api+json
Content-Length: [Length of the submitted content]

time.Parse ?
*/
