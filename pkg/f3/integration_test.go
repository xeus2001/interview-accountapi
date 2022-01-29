//go:build int
// +build int

package f3_test

import (
	"flag"
	"fmt"
	"github.com/xeus2001/interview-accountapi/pkg/f3"
	"testing"
	"time"
)

var (
	wait     = flag.Int("f3.wait", 15, "The time to wait for the account API in seconds")
	endpoint = flag.String("f3.endpoint", f3.DefaultEndPoint, "Override the default endpoint for clients")
)

func TestClient(t *testing.T) {
	fmt.Printf("Execute integration tests against endpoint: '%s'\n", *endpoint)
	client := f3.NewClient().WithEndPoint(*endpoint)

	// Wait for endpoint
	WAIT_MAX := time.Second * time.Duration(*wait)
	START := time.Now()
	fmt.Printf("Wait %d seconds for the Account-API to become available ...\n", *wait)
	for !client.IsHealthy() && time.Since(START) < WAIT_MAX {
		time.Sleep(time.Millisecond * 1000)
	}
	// We need the account API now!
	if !client.IsHealthy() {
		t.Fatalf("Health check for service failed")
	}

	// Test the happy path, create account, fetch it and then delete it.
	t.Run("CreateAccount", func(t *testing.T) {
		account := createTestAccount(true)
		created, e := client.CreateAccount(account)
		if e != nil {
			t.Fatalf("Failed to create test account: %s", e.Error())
		}
		if created == nil {
			t.Fatal("Received nil as created account")
		}
	})
	t.Run("FetchAccount", func(t *testing.T) {
		fetched, e := client.FetchAccount(f3.IntegrationTestAccountId)
		if e != nil {
			t.Fatalf("Failed to fetch test account: %s", e.Error())
		}
		if fetched == nil {
			t.Fatalf("Failed to fetch test account, returned account is nil")
		}
	})
	t.Run("DeleteAccountWithIdWrongVersion", func(t *testing.T) {
		e := client.DeleteAccount(f3.IntegrationTestAccountId, 999)
		if e == nil {
			t.Fatalf("Deleting the test account with wrong version should have failed")
		}
		if e.ErrorCode() != f3.ErrConflict {
			t.Fatalf("Deleting the test account should have failed with %d, but reiceived error code %d", f3.ErrConflict, e.ErrorCode())
		}
	})
	t.Run("DeleteAccount", func(t *testing.T) {
		e := client.DeleteAccount(f3.IntegrationTestAccountId, 0)
		if e != nil {
			t.Errorf("Deleting the test account failed, reason: %s", e.Error())
		}
	})

	brokenClient := f3.NewClient()
	brokenClient.WithEndPoint("tj48\031fnf234f234f342")
	t.Run("CreateAccountWithInvalidEndPointURL@BROKEN", func(t *testing.T) {
		account := f3.NewAccount(
			nil,
			"GB",
			"400300",
			"GBDSC",
			"Alexander Lowey-Weber",
			"41426819",
			"GBP",
			"tests")
		created, e := brokenClient.CreateAccount(account)
		if created != nil {
			t.Fatalf("Create an account with ID '%s', this should not be possible", account.Id)
		}
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrGeneric {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrGeneric, e.ErrorCode())
		}
	})
	t.Run("DeleteAccountWithInvalidEndPointURL@BROKEN", func(t *testing.T) {
		e := brokenClient.DeleteAccount("4711", 12)
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrGeneric {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrGeneric, e.ErrorCode())
		}
	})

	brokenClient.WithEndPoint("http://localhost:0/v1")
	if brokenClient.IsHealthy() {
		t.Fatalf("The borken client must not be healthy")
	}

	t.Run("CreateAccount@BROKEN", func(t *testing.T) {
		account := createTestAccount(true)
		created, e := brokenClient.CreateAccount(account)
		if created != nil {
			t.Fatalf("Create an account with ID '%s', this should not be possible", account.Id)
		}
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrRequest {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrRequest, e.ErrorCode())
		}
	})
	t.Run("CreateInvalidAccount", func(t *testing.T) {
		account := createTestAccount(false)
		account.Id = "THIS_MUST_NOT_WORK"
		created, e := client.CreateAccount(account)
		if created != nil {
			t.Fatalf("Create an account with ID '%s', this should not be possible", account.Id)
		}
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrBadRequest {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrBadRequest, e.ErrorCode())
		}
		if e.Request() == nil {
			t.Errorf("The request is missing")
		}
		if e.Error() == "" {
			t.Errorf("Error message missing")
		}
		cause := e.Unwrap()
		if cause == nil {
			t.Errorf("Unexpected cause")
		}
	})
	t.Run("CreateAccountWithNil", func(t *testing.T) {
		created, e := client.CreateAccount(nil)
		if created != nil {
			t.Fatalf("Create an account with nil???")
		}
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrRequest {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrRequest, e.ErrorCode())
		}
	})
	t.Run("CreateAccountWithoutAttributes", func(t *testing.T) {
		account := createTestAccount(false)
		account.Attr = nil
		created, e := client.CreateAccount(account)
		if created != nil {
			t.Fatalf("Create an invalid account, expect to receive nil")
		}
		if e == nil {
			t.Fatalf("Missing error")
		}
		if e.ErrorCode() != f3.ErrBadRequest {
			t.Errorf("Invalid error, expected %d, got %d", f3.ErrBadRequest, e.ErrorCode())
		}
	})
	t.Run("DeleteNotExistingAccount", func(t *testing.T) {
		e := client.DeleteAccount(f3.IntegrationTestAccountId, 0)
		if e == nil {
			t.Fatalf("Deleting the test account should have failed")
		}
		if e.ErrorCode() != f3.ErrNotFound {
			t.Fatalf("Deleting the test account should have failed with %d, but reiceived error code %d", f3.ErrNotFound, e.ErrorCode())
		}
	})
	t.Run("DeleteNotExistingAccount@BROKEN", func(t *testing.T) {
		e := brokenClient.DeleteAccount(f3.IntegrationTestAccountId, 0)
		if e == nil {
			t.Fatalf("Deleting the test account should have failed")
		}
		if e.ErrorCode() != f3.ErrRequest {
			t.Fatalf("We expect that the request failed and we do not get any response!")
		}
	})
	t.Run("FetchNotExistingAccount", func(t *testing.T) {
		fetched, e := client.FetchAccount(f3.IntegrationTestAccountId)
		if e == nil {
			t.Fatalf("Fetched an account, we expected that it does not exist!")
		}
		if e.Response().StatusCode != 404 {
			t.Errorf("Expected status code 404, but got: %d", e.Response().StatusCode)
		}
		if e.ErrorCode() != f3.ErrNotFound {
			t.Errorf("Expected error code %d, but got %d", f3.ErrNotFound, e.ErrorCode())
		}
		if fetched != nil {
			t.Errorf("Fetch not existing account")
		}
	})
	t.Run("FetchNotExistingAccount@BROKEN", func(t *testing.T) {
		fetched, e := brokenClient.FetchAccount(f3.IntegrationTestAccountId)
		if fetched != nil {
			t.Errorf("Fetch not existing account")
		}
		if e == nil {
			t.Fatalf("Fetched an account, we expected that the request does not work!")
		}
		if e.ErrorCode() != f3.ErrRequest {
			t.Errorf("Expected error code %d, but got %d", f3.ErrRequest, e.ErrorCode())
		}
		if e.Response() != nil {
			t.Errorf("We received a response, but did not expect any!")
		}
	})
}
