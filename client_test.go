//go:build int
// +build int

package interview_accountapi_test

import (
	"github.com/xeus2001/interview-accountapi"
	"testing"
)

func createTestClient() *interview_accountapi.Client {
	return interview_accountapi.NewClient(interview_accountapi.IntegrationEndPoint)
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
	account.Id = interview_accountapi.IntegrationTestAccountId
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
	fetched, e := client.FetchAccount(interview_accountapi.IntegrationTestAccountId)
	if e != nil {
		t.Fatalf("Failed to fetch test account: %s", e.Error())
	}
	if fetched == nil {
		t.Fatalf("Failed to fetch test account, returned account is nil")
	}
}

func TestClient_DeleteAccount(t *testing.T) {
	client := createTestClient()
	e := client.DeleteAccount(interview_accountapi.IntegrationTestAccountId, 0)
	if e != nil {
		t.Errorf("Deleting the test account failed, reason: %s", e.Error())
	}
}
