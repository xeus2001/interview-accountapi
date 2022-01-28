package f3

import (
	"github.com/google/uuid"
)

// AccountStatusString is an alias for a string that represents the account status.
type AccountStatusString string

const (
	// StatusPending represents a pending (not yet confirmed) account.
	StatusPending = AccountStatusString("pending")

	// StatusConfirmed represents a confirmed account.
	StatusConfirmed = AccountStatusString("confirmed")

	// StatusClosed is an FPS only account status string and requires AccountAttr.StatusReason to be set.
	StatusClosed = AccountStatusString("closed")

	// StatusFailed is an SEPA and FPS Indirect (LHV) only account status string.
	StatusFailed = AccountStatusString("failed")

	// TypeAccount is the type for accounts.
	TypeAccount = "accounts"
)

// AccountsEnvelope is an envelope for a list of accounts.
type AccountsEnvelope struct {
	Data []*Account `json:"data"`
}

// AccountEnvelope is an envelope for a single account.
type AccountEnvelope struct {
	Data *Account `json:"data"`
}

// Account represents the account resource with attributes.
type Account struct {
	Resource
	// Attr are the attributes of the account.
	Attr *AccountAttr `json:"attributes,omitempty"`
}

// AccountAttr are the account specific attributes.
type AccountAttr struct {
	AlternativeNames        []string            `json:"alternative_names,omitempty"`
	Name                    []string            `json:"name,omitempty"`
	AccountClassification   string              `json:"account_classification,omitempty"`
	AccountNumber           string              `json:"account_number,omitempty"` // A unique account number will automatically be generated if not provided. If provided, the account number is not validated.
	BankId                  string              `json:"bank_id,omitempty"`
	BankIdCode              string              `json:"bank_id_code,omitempty"`
	BaseCurrency            string              `json:"base_currency,omitempty"`
	CustomerId              *string             `json:"customer_id,omitempty"`
	Bic                     string              `json:"bic,omitempty"`
	Country                 string              `json:"country,omitempty"`
	Iban                    string              `json:"iban,omitempty"` // Will be calculated from other fields if not supplied.
	SecondaryIdentification string              `json:"secondary_identification,omitempty"`
	Status                  AccountStatusString `json:"status,omitempty"`
	StatusReason            *string             `json:"status_reason,omitempty"`
	Switched                bool                `json:"switched,omitempty"`
	AccountMatchingOptOut   bool                `json:"account_matching_opt_out,omitempty"`
	JointAccount            bool                `json:"joint_account,omitempty"`
}

// WithStatusPending set the status to StatusPending.
func (aa *AccountAttr) WithStatusPending() *AccountAttr {
	aa.Status = StatusPending
	aa.StatusReason = nil
	return aa
}

// WithStatusConfirmed set the status to StatusConfirmed.
func (aa *AccountAttr) WithStatusConfirmed() *AccountAttr {
	aa.Status = StatusConfirmed
	aa.StatusReason = nil
	return aa
}

// WithStatusClosed set the status to StatusClosed, requires reason.
func (aa *AccountAttr) WithStatusClosed(reason string) *AccountAttr {
	aa.Status = StatusClosed
	aa.StatusReason = &reason
	return aa
}

// WithStatusFailed set the status to StatusFailed.
func (aa *AccountAttr) WithStatusFailed() *AccountAttr {
	aa.Status = StatusFailed
	aa.StatusReason = nil
	return aa
}

// NewAccount is a small helper method to create a basic structure for a new account. The created structure can be
// modified after creation or directly used to create a new account. If the organization-id is nil, then the
// DefaultOrganizationId is used. The customerId is optional and if given, it is set.
func NewAccount(
	organizationId *string,
	country string,
	bankId string,
	bankIdCode string,
	accountHolder string,
	accountNumber string,
	accountCurrency string,
	customerId string,
) *Account {
	pAccount := new(Account)
	pAccount.Type = TypeAccount
	pAccount.Id = uuid.New().String()
	if organizationId == nil {
		organizationId = &DefaultOrganizationId
	}
	pAccount.OrganisationId = *organizationId
	pAccount.Attr = new(AccountAttr)
	attr := pAccount.Attr
	attr.WithStatusConfirmed()
	attr.Country = country
	attr.BankId = bankId
	attr.BankIdCode = bankIdCode
	attr.Name = []string{accountHolder}
	attr.AccountNumber = accountNumber
	attr.BaseCurrency = accountCurrency
	if len(customerId) > 0 {
		attr.CustomerId = &customerId
	}
	return pAccount
}
