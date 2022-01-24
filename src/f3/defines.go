package f3

import (
	uuid "github.com/nu7hatch/gouuid"
	"os"
	"time"
)

// DefaultTimeout is the default timeout to be used for the http client.
var DefaultTimeout time.Duration = time.Second * 5

// DefaultOrganizationId is the default organization identifier to be used, when a new account is created. The variable
// is initialized from the environment variable F3_ORG_ID
var DefaultOrganizationId *uuid.UUID

// LocalOrganizationId is a special_u16 identifier used for integration tests, feel free to use it as well for own tests.
var LocalOrganizationId *uuid.UUID

// ProductionEndPoint the production endpoint.
const ProductionEndPoint = "https://api.form3.tech/v1"

// LocalEndPoint the default docker endpoint for integration tests and other tests.
const LocalEndPoint = "http://localhost:8080/v1"

func init() {
	text, set := os.LookupEnv("F3_ORG_ID")
	if set {
		uuid, e := uuid.ParseHex(text)
		if e == nil {
			DefaultOrganizationId = uuid
		}
	}
	uuid, e := uuid.ParseHex("f8e803fe-6962-4811-94dc-e558610cbe78")
	if e == nil {
		LocalOrganizationId = uuid
	}
}
