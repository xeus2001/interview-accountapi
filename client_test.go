//go:build int
// +build int

package f3_test

import (
	"flag"
	"fmt"
	f3 "github.com/xeus2001/interview-accountapi"
	"testing"
	"time"
)

var wait = flag.Int("f3.wait", 60, "The time to wait for the account API in seconds")

func TestClient(t *testing.T) {
	fmt.Printf("Execute integration tests against endpoint: '%s'\n", *f3.DefaultEndPoint)
	client := f3.NewClient()
	client.WithEndPoint(*f3.DefaultEndPoint)

	// Wait for endpoint
	WAIT_MAX := time.Second * time.Duration(*wait)
	START := time.Now()
	fmt.Printf("Wait for Account API to become available ")
	for !client.IsHealthy() && time.Since(START) < WAIT_MAX {
		print(".")
		time.Sleep(time.Millisecond * 1000)
	}
	println()
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
