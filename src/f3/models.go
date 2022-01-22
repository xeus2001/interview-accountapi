package f3

import (
	"fmt"
	"github.com/xeus2001/interview-accountapi/src/iso"
	"github.com/xeus2001/interview-accountapi/src/iso/countryCode"
	"time"
)

// The bank ID code.
type BankIdCode string

// The error message returned by the service.
type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

// The shared attributes of all resources.
// The ID is the unique identifier of the resource; must be a UUID.
// The OrganisationID is the unique identifier of the organization that own the record; must be a UUID.
// The Type is the type of the record and set by the account API.
// The Version is a counter indicating how many times this resource has been modified. When you create a resource, it
// is automatically set to 0. Whenever the content of the resource changes, the value of version is increased.
// Used for concurrency control in Patch and Delete methods to avoid modifying an older version of the record that has
// already been changed (e.g. by an internal process at Form3).
// See: https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)
type Resource struct {
	ID             string     `json:"id,omitempty"`
	OrganisationID string     `json:"organisation_id,omitempty"`
	Type           string     `json:"type,omitempty"`
	Version        *int64     `json:"version,omitempty"`
	CreatedOn      *time.Time `json:"created_on,omitempty"`
	ModifiedOn     *time.Time `json:"modified_on,omitempty"`
}

// The account status.
type AccountStatus string

const (
	StatusPending   = AccountStatus("pending")   // All services
	StatusConfirmed = AccountStatus("confirmed") // All services
	StatusClosed    = AccountStatus("closed")    // FPS only, requires status_reason
	StatusFailed    = AccountStatus("failed")    // SEPA & FPS Indirect (LHV) only
)

// The account list either being sent or received.
type Accounts struct {
	Data []*Account `json:"data"`
}

// The account record.
type Account struct {
	Resource
	Attr *AccountAttr `json:"attributes,omitempty"`
}

// The account attributes.
type AccountAttr struct {
	AccountClassification   *string           `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool             `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string            `json:"account_number,omitempty"` // A unique account number will automatically be generated if not provided. If provided, the account number is not validated.
	AlternativeNames        []string          `json:"alternative_names,omitempty"`
	BankID                  string            `json:"bank_id,omitempty"`
	BankIDCode              BankIdCode        `json:"bank_id_code,omitempty"`
	BaseCurrency            *iso.CurrencyCode `json:"base_currency,omitempty"`
	CustomerId              string            `json:"customer_id,omitempty"`
	Bic                     string            `json:"bic,omitempty"`
	Country                 iso.CountryCode   `json:"country,omitempty"`
	Iban                    string            `json:"iban,omitempty"` // Will be calculated from other fields if not supplied.
	JointAccount            *bool             `json:"joint_account,omitempty"`
	Name                    []string          `json:"name,omitempty"`
	SecondaryIdentification string            `json:"secondary_identification,omitempty"`
	Status                  *AccountStatus    `json:"status,omitempty"`
	StatusReason            *string           `json:"status_reason,omitempty"`
	Switched                *bool             `json:"switched,omitempty"`
}

// Set the status and optionally the status reason. The reason is only taking into account for the status failed and
// in that case it is mandatory.
func (aa *AccountAttr) SetStatus(status AccountStatus, reason string) Err {
	if status == StatusFailed {
		if len(reason) == 0 {
			return err{ErrMissingStatusReason, msgMissingStatusReason}
		}
		aa.StatusReason = &reason
	} else {
		aa.StatusReason = nil
	}
	aa.Status = &status
	return nil
}

// Information about restrictions for banks.
type BankRestrictions interface {
	CountryCode() iso.CountryCode
	BankIdCode() *BankIdCode
	ValidateBankId(id string) Err
}

type special int16

const (
	special_none         special = iota
	special_blz          special = iota
	special_aba          special = iota
	special_uk_sort_code special = iota
)

type info struct {
	required bool
	value    string
	min      int
	max      int
	special  special
}

type bankInfo struct {
	countryCode   iso.CountryCode
	bankId        info
	bic           info
	bankIdCode    info
	accountNumber info
}

var (
	gb = bankInfo{
		countryCode:   countryCode.UnitedKingdom,
		bankId:        info{required: true, min: 6, max: 6, special: special_uk_sort_code},
		bic:           info{required: true},
		bankIdCode:    info{required: true, value: "GBDSC"},
		accountNumber: info{min: 8, max: 8},
	}
	de = bankInfo{
		countryCode:   countryCode.Germany,
		bankId:        info{required: true, min: 8, max: 8},
		bic:           info{},
		bankIdCode:    info{required: true, value: "DEBLZ"},
		accountNumber: info{min: 7, max: 7},
	}
	us = bankInfo{
		countryCode:   countryCode.Usa,
		bankId:        info{required: true, min: 9, max: 9, special: special_aba},
		bic:           info{},
		bankIdCode:    info{required: true, value: "CHBCC"},
		accountNumber: info{min: 6, max: 17},
	}
)

func (i bankInfo) CountryCode() iso.CountryCode {
	return i.countryCode
}

func (i bankInfo) BankIdCode() *BankIdCode {
	if len(i.bankIdCode.value) > 0 {
		code := BankIdCode(i.bankIdCode.value)
		return &code
	}
	return nil
}

// Validates the given bank-id and returns an error, when it does not fulfill the requirements for this bank information.
func (i bankInfo) ValidateBankId(id string) Err {
	len := len(id)
	if len < i.bankId.min {
		return err{ErrBankIdTooShort, fmt.Sprintf("The bank-id must be less than %d letters, given: %d", i.bankId.min, len)}
	}
	if i.bankId.max > 0 && len > i.bankId.max {
		return err{ErrBankIdTooLong, fmt.Sprintf("The bank-id must not be longer than %d letters, given: %d", i.bankId.max, len)}
	}
	// TODO: handle special
	if len == 0 && i.bankId.required {
		return err{ErrBankIdMissing, msgBankIdMissing}
	}
	return nil
}

// Returns bank restrictions for the given country or nil, if no restrictions are known. Be aware that this does not
// mean that there are no restrictions.
func BankRestrictionsByCountry(country iso.CountryCode) BankRestrictions {
	switch country {
	case countryCode.UnitedKingdom:
		return gb
	case countryCode.Germany:
		return de
	case countryCode.Usa:
		return us
	}
	return nil
}
