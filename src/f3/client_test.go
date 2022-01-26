package f3_test

import (
	"github.com/xeus2001/interview-accountapi/src/f3"
	"testing"
)

func createTestClient() *f3.Client {
	return f3.NewClient(f3.IntegrationEndPoint)
}

func TestClient_IsHealthy(t *testing.T) {
	client := createTestClient()
	if !client.IsHealthy() {
		t.Fatalf("Health check for service failed")
	}
}

func TestClient_CreateAccount(t *testing.T) {
	client := createTestClient()
	account := createTestAccount()
	account.Id = f3.IntegrationTestAccountId
	created, e := client.CreateAccount(account)
	if e != nil {
		t.Fatalf("Failed to create test account: %s", e.Error())
	}
	if created == nil {
		t.Fatal("Received nil as created account")
	}
}

func TestClient_FetchAccount(t *testing.T) {
	client := createTestClient()
	account := createTestAccount()
	account.Id = f3.IntegrationTestAccountId
	created, e := client.CreateAccount(account)
	if e != nil {
		t.Errorf("Failed to create test account: %s", e.Error())
	}
	if created == nil {
		t.Errorf("Received nil as created account")
	}
	fetched, e := client.FetchAccount(f3.IntegrationTestAccountId)
	if e != nil {
		t.Fatalf("Failed to fetch test account: %s", e.Error())
	}
	if fetched == nil {
		t.Fatalf("Failed to fetch test account, returned account is nil")
	}
}

func TestClient_DeleteAccount(t *testing.T) {
	client := createTestClient()
	e := client.DeleteAccount(f3.IntegrationTestAccountId, 0)
	if e != nil {
		t.Errorf("Deleting the test account failed, reason: %s", e.Error())
	}
}
