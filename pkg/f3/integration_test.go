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
	t.Run("IsHealthy", func(t *testing.T) {
		if !client.IsHealthy() {
			t.Fatalf("Health check for service failed")
		}
	})
	t.Run("CreateAccount", func(t *testing.T) {
		account := createTestAccount()
		account.Id = f3.IntegrationTestAccountId
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
	t.Run("DeleteAccount", func(t *testing.T) {
		e := client.DeleteAccount(f3.IntegrationTestAccountId, 0)
		if e != nil {
			t.Errorf("Deleting the test account failed, reason: %s", e.Error())
		}
	})
}
