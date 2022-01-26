package f3

import (
	"net"
	"net/http"
	"os"
	"time"
)

const (
	F3_CLIENT_ORG_ID       = "F3_CLIENT_ORG_ID"
	F3_CLIENT_INT_ORG_ID   = "F3_CLIENT_INT_ORG_ID"
	F3_CLIENT_PRD_ENDPOINT = "F3_CLIENT_PRD_ENDPOINT"
	F3_CLIENT_INT_ENDPOINT = "F3_CLIENT_INT_ENDPOINT"
	F3_TEST_ACCOUNT_ID     = "F3_TEST_ACCOUNT_ID"
)

var (
	// DefaultTimeout is the default timeout to be used for the http client.
	DefaultTimeout time.Duration = time.Second * 5

	// DefaultOrganizationId is the default organization identifier to be used, when a new account is created. The variable
	// is initialized from the environment variable F3_CLIENT_ORG_ID
	DefaultOrganizationId string

	// DefaultIntegrationOrganizationId is the default organization identifier used for integration tests. Can be overridden
	// using the F3_CLIENT_INT_ORG_ID environment variable.
	DefaultIntegrationOrganizationId string = "f8e803fe-6962-4811-94dc-e558610cbe78"

	// ProductionEndPoint the production endpoint. Can be overridden using the F3_CLIENT_PRD_ENDPOINT environment variable.
	ProductionEndPoint = "https://api.form3.tech/v1"

	// IntegrationEndPoint the endpoint for the integration tests. Can be overridden by declaring the
	// F3_CLIENT_INT_ENDPOINT environment variable.
	IntegrationEndPoint = "http://localhost:8080/v1"

	// IntegrationTestAccountId is the id to be used for the integration test account, which is created and deleted.
	// Can be overridden using the environment variable F3_TEST_ACCOUNT_ID
	IntegrationTestAccountId = "f8e803fe-6962-4811-94dc-000000000000"

	// DefaultDialer is the default dialer that is being used when new Form 3 clients are created.
	DefaultDialer = &net.Dialer{}

	// DefaultTransport is the default transport configuration to be used when new Form 3 clients are created.
	DefaultTransport = &http.Transport{
		DialContext:     DefaultDialer.DialContext,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
)

func init() {
	if text, ok := os.LookupEnv(F3_CLIENT_ORG_ID); ok {
		DefaultOrganizationId = text
	}
	if text, ok := os.LookupEnv(F3_CLIENT_INT_ORG_ID); ok {
		DefaultOrganizationId = text
	}

	if text, ok := os.LookupEnv(F3_TEST_ACCOUNT_ID); ok {
		IntegrationTestAccountId = text
	}

	// TODO: Should we trim a trailing slash?
	if text, ok := os.LookupEnv(F3_CLIENT_PRD_ENDPOINT); ok {
		ProductionEndPoint = text
	}
	if text, ok := os.LookupEnv(F3_CLIENT_INT_ENDPOINT); ok {
		IntegrationEndPoint = text
	}
}
