package f3

import (
	"net"
	"net/http"
	"time"
)

var (
	// DefaultEndPoint is the default endpoint that f3.NewClient will use.
	DefaultEndPoint = "https://api.f3.tech/v1"

	// DefaultTimeout is the default timeout to be used for the http client.
	DefaultTimeout = time.Second * 5

	// DefaultOrganizationId is the default organization identifier to be used, when a new account is created. The variable
	// is initialized from the environment variable F3_CLIENT_ORG_ID
	DefaultOrganizationId string

	// DefaultIntegrationOrganizationId is the default organization identifier used for integration tests.
	DefaultIntegrationOrganizationId = "f8e803fe-6962-4811-94dc-e558610cbe78"

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
