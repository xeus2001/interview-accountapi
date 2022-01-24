package f3

import "github.com/xeus2001/interview-accountapi/src/iso"

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
	AlternativeNames        []string               `json:"alternative_names,omitempty"`
	Name                    []string               `json:"name,omitempty"`
	AccountClassification   string                 `json:"account_classification,omitempty"`
	AccountNumber           string                 `json:"account_number,omitempty"` // A unique account number will automatically be generated if not provided. If provided, the account number is not validated.
	BankId                  string                 `json:"bank_id,omitempty"`
	BankIdCode              BankCodeString         `json:"bank_id_code,omitempty"`
	BaseCurrency            iso.CurrencyCodeString `json:"base_currency,omitempty"`
	CustomerId              string                 `json:"customer_id,omitempty"`
	Bic                     string                 `json:"bic,omitempty"`
	Country                 iso.CountryCodeString  `json:"country,omitempty"`
	Iban                    string                 `json:"iban,omitempty"` // Will be calculated from other fields if not supplied.
	SecondaryIdentification string                 `json:"secondary_identification,omitempty"`
	Status                  AccountStatusString    `json:"status,omitempty"`
	StatusReason            *string                `json:"status_reason,omitempty"`
	Switched                bool                   `json:"switched,omitempty"`
	AccountMatchingOptOut   bool                   `json:"account_matching_opt_out,omitempty"`
	JointAccount            bool                   `json:"joint_account,omitempty"`
}

// SetStatusPending set the status to StatusPending.
func (aa *AccountAttr) SetStatusPending() *AccountAttr {
	aa.Status = StatusPending
	aa.StatusReason = nil
	return aa
}

// SetStatusConfirmed set the status to StatusConfirmed.
func (aa *AccountAttr) SetStatusConfirmed() *AccountAttr {
	aa.Status = StatusConfirmed
	aa.StatusReason = nil
	return aa
}

// SetStatusClosed set the status to StatusClosed, requires reason.
func (aa *AccountAttr) SetStatusClosed(reason string) *AccountAttr {
	aa.Status = StatusClosed
	aa.StatusReason = &reason
	return aa
}

// SetStatusFailed set the status to StatusFailed.
func (aa *AccountAttr) SetStatusFailed() *AccountAttr {
	aa.Status = StatusFailed
	aa.StatusReason = nil
	return aa
}
