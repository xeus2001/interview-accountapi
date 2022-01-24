package f3

import (
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/xeus2001/interview-accountapi/src/iso"
	"io/ioutil"
	"net/http"
)

const (
	headerUserAgent   = "User-Agent"
	headerAccept      = "Accept"
	headerContentType = "Content-Type"

	userAgentName       = "f3.Client"
	mimeApplicationJson = "application/json"
)

// The response return by the health-check API.
type healthyResponse struct {
	// status is the status returned by the service, which may be nil.
	Status *string `json:"status,omitempty"`
}

// Create a new account and fills it with default values. The created structure can be modified or directly used
// to create a new account. The country is required and dependent on the country the bank-id may be required too.
// If the provided country is not eligible for default values and err is returned. In this case the account structure
// must be created manually.
func NewAccount(
	organizationId *uuid.UUID,
	countryCode iso.CountryCodeString,
	bankId string,
	accountHolderName []string,
	accountNumber string,
	customerId string,
) (*Account, Err) {
	restrictions := GetBankVerifier(countryCode)
	if restrictions == nil {
		return nil, err{ErrNoVerifierAvailableForCountry, msgCountryUnknown, nil}
	}
	country, ok := iso.CountryByCode[countryCode]
	if !ok {
		return nil, err{ErrNoVerifierAvailableForCountry, msgCountryUnknown, nil}
	}
	pAccount := new(Account)
	id, e := NewUuid()
	if e != nil {
		return nil, e
	}
	pAccount.Id = id.String()
	if organizationId == nil {
		if DefaultOrganizationId == nil {
			return nil, err{ErrInvalidUuid, msgInvalidUuid, nil}
		}
		organizationId = DefaultOrganizationId
	}
	pAccount.OrganisationId = organizationId.String()
	pAccount.Attr = new(AccountAttr)
	attr := pAccount.Attr
	attr.SetStatusConfirmed()
	attr.Country = countryCode
	currencies := country.Currencies
	if len(currencies) > 0 {
		attr.BaseCurrency = currencies[0]
	}
	e = restrictions.ValidateBankId(bankId)
	if e != nil {
		return nil, e
	}
	attr.BankId = bankId
	bankIdCode := restrictions.BankCode()
	if bankIdCode != nil {
		attr.BankIdCode = *bankIdCode
	}
	// TODO: Name of the account holder, up to four lines possible.
	//
	// TODO: CoP: Primary account name. For concatenated personal names, joint account names and organisation names,
	//       use the first line. If first and last names of a personal name are separated, use the first line for
	//       first names, the second line for last names. Titles are ignored and should not be entered.
	//       required !
	//
	// TODO: SEPA Indirect: Can be a person or organisation. Only the first line is used, mininum 5 characters.
	//       required !
	attr.Name = accountHolderName
	attr.AccountNumber = accountNumber
	attr.CustomerId = customerId
	return pAccount, nil
}

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
	req, err := http.NewRequest(http.MethodGet, c.healthCheckUri, nil)
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
	//goland:noinspection GoUnhandledErrorResult
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

func (c *Client) CreateAccount(account *Account) (*Account, Err) {
	if account == nil {
		return nil, err{ErrAccountNil, msgAccountNil, nil}
	}
	accounts := AccountsEnvelope{Data: make([]*Account, 1)}
	accounts.Data[0] = account
	jsonBytes, e := json.Marshal(accounts)
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
	e = json.Unmarshal(body, &accounts)
	if e != nil {
		var errResponse *ErrorResponse
		e = json.Unmarshal(body, &errResponse)
		if e != nil {
			return nil, err{ErrJsonParse, e.Error(), e}
		}
		return nil, err{ErrResponse, errResponse.ErrorMessage, e}
	}
	if len(accounts.Data) != 1 {
		return nil, err{ErrResponse, msgInvalidResponse, nil}
	}
	return accounts.Data[0], nil
}

func (c *Client) FetchAccount(accountId string) {

}

func (c *Client) DeleteAccount(accountId string) {

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
