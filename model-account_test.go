package interview_accountapi_test

import (
	"github.com/google/uuid"
	"github.com/xeus2001/interview-accountapi"
	"testing"
)

func createTestAccount() *interview_accountapi.Account {
	return interview_accountapi.NewAccount(
		&interview_accountapi.DefaultIntegrationOrganizationId,
		"GB",
		"400300",
		"GBDSC",
		"Alexander Lowey-Weber",
		"41426819",
		"GBP",
		"tests")
}

func TestNewAccount(t *testing.T) {
	account := createTestAccount()
	if account == nil {
		t.Fatal("Failed to create account, NewAccount returned nil")
	}
	if _, e := uuid.Parse(account.Id); e != nil {
		t.Errorf("The account.Id that was generated is no valid UUID: %s", account.Id)
	}
	if _, e := uuid.Parse(account.OrganisationId); e != nil {
		t.Errorf("The account.OrganisationId that is no valid UUID: %s", account.OrganisationId)
	}
	if interview_accountapi.TypeAccount != account.Type {
		t.Errorf("The account.Type is not valid, expected '%s', but found '%s'", interview_accountapi.TypeAccount, account.Type)
	}
	if account.Version != nil {
		t.Errorf("The accout has a version %d, this must not be, the version for new accounts should be nil", account.Version)
	}
	if account.Attr == nil {
		t.Fatal("The account.Attr is nil")
	}
	attr := account.Attr
	if attr.Status != interview_accountapi.StatusConfirmed {
		t.Errorf("The account.Attr.Status should have been '%s', but was: %s", interview_accountapi.StatusConfirmed, attr.Status)
	}
	if "GB" != attr.Country {
		t.Errorf("The account.Attr.Country should have been GB, but was: %s", attr.Country)
	}
	if "400300" != attr.BankId {
		t.Errorf("The account.Attr.BankId should have been 400300, but was: %s", attr.BankId)
	}
	if "GBDSC" != attr.BankIdCode {
		t.Errorf("The account.Attr.BankIdCode should have been GBDSC, but was: %s", attr.BankIdCode)
	}
	if attr.Name == nil {
		t.Error("The account.Attr.Name is nil")
	} else {
		if len(attr.Name) != 1 {
			t.Errorf("The account.Attr.Name should have been an array of the length 1, but has the length %d", len(attr.Name))
		}
		if "Alexander Lowey-Weber" != attr.Name[0] {
			t.Errorf("The account.Attr.Name[] should have been 'Alexander Lowey-Weber', but is: '%s'", attr.Name[0])
		}
	}
	if "41426819" != attr.AccountNumber {
		t.Errorf("The account.Attr.AccountNumber should have been '41426819', but was: %s", attr.AccountNumber)
	}
	if "GBP" != attr.BaseCurrency {
		t.Errorf("The account.Attr.BaseCurrency should have been 'GBP', but was: %s", attr.BaseCurrency)
	}
	if attr.CustomerId == nil {
		t.Error("The account.Attr.CustomerId is nil")
	} else if "tests" != *attr.CustomerId {
		t.Errorf("The account.Attr.CustomerId should have been 'tests', but was: %s", *attr.CustomerId)
	}
}

func TestAccountAttr_WithStatusClosed(t *testing.T) {
	account := createTestAccount()
	attr := account.Attr
	attr.WithStatusClosed("Just for fun")
	if attr.Status != interview_accountapi.StatusClosed {
		t.Errorf("The account.Attr.Status should have been '%s', but was: %s", interview_accountapi.StatusClosed, attr.Status)
	}
	if attr.StatusReason == nil {
		t.Errorf("The account.Attr.StatusReason should have been 'Just for fun', but was nil")
	} else if "Just for fun" != *attr.StatusReason {
		t.Errorf("The account.Attr.StatusReason should have been 'Just for fun', but was: %s", *attr.StatusReason)
	}
}

func TestAccountAttr_WithStatusFailed(t *testing.T) {
	account := createTestAccount()
	attr := account.Attr
	attr.WithStatusFailed()
	if attr.Status != interview_accountapi.StatusFailed {
		t.Errorf("The account.Attr.Status should have been '%s', but was: %s", interview_accountapi.StatusFailed, attr.Status)
	}
	if attr.StatusReason != nil {
		t.Errorf("The account.Attr.StatusReason should have been nil, but was: %s", *attr.StatusReason)
	}
}

func TestAccountAttr_WithStatusConfirmed(t *testing.T) {
	account := createTestAccount()
	attr := account.Attr
	attr.WithStatusConfirmed()
	if attr.Status != interview_accountapi.StatusConfirmed {
		t.Errorf("The account.Attr.Status should have been '%s', but was: %s", interview_accountapi.StatusConfirmed, attr.Status)
	}
	if attr.StatusReason != nil {
		t.Errorf("The account.Attr.StatusReason should have been nil, but was: %s", *attr.StatusReason)
	}
}

func TestAccountAttr_WithStatusPending(t *testing.T) {
	account := createTestAccount()
	attr := account.Attr
	attr.WithStatusPending()
	if attr.Status != interview_accountapi.StatusPending {
		t.Errorf("The account.Attr.Status should have been '%s', but was: %s", interview_accountapi.StatusPending, attr.Status)
	}
	if attr.StatusReason != nil {
		t.Errorf("The account.Attr.StatusReason should have been nil, but was: %s", *attr.StatusReason)
	}
}
