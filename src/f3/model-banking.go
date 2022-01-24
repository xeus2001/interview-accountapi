package f3

import (
	"fmt"
	"github.com/xeus2001/interview-accountapi/src/iso"
)

// BankCodeString is a special_u16 string that identifies a bank, for example in Germany the "Bankleitzahl" (BLZ).
type BankCodeString string

// BankCountryVerifier is an interface that can be used to verify attributes against restrictions for certain
// countries.
type BankCountryVerifier interface {
	// CountryCode returns the ISO-3166 country code for which these restrictions are to be applied.
	CountryCode() iso.CountryCodeString

	// BankCode returns the code for the banks of the country.
	BankCode() *BankCodeString

	// ValidateBankId verifies if the given value is a valid bank identifier in this country, for example in Germany
	// the "Bankleitzahl" (BLZ).
	ValidateBankId(id string) Err
}

// special_u16 is just an alias for a 16 bit special value.
type special_u16 uint16

var allBankCountryVerifier = map[iso.CountryCodeString]BankCountryVerifier{}

const (
	special_none         special_u16 = iota
	special_blz          special_u16 = iota
	special_aba          special_u16 = iota
	special_uk_sort_code special_u16 = iota
)

type info struct {
	value    string
	min      int
	max      int
	special  special_u16
	required bool
}

type bankInfo struct {
	countryCode   iso.CountryCodeString
	bankId        info
	bic           info
	bankIdCode    info
	accountNumber info
}

func (i bankInfo) CountryCode() iso.CountryCodeString {
	return i.countryCode
}

func (i bankInfo) BankCode() *BankCodeString {
	if len(i.bankIdCode.value) > 0 {
		code := BankCodeString(i.bankIdCode.value)
		return &code
	}
	return nil
}

func (i bankInfo) ValidateBankId(id string) Err {
	length := len(id)
	if length < i.bankId.min {
		return err{ErrBankIdTooShort, fmt.Sprintf("The bank-id must be less than %d letters, given: %d", i.bankId.min, length), nil}
	}
	if i.bankId.max > 0 && length > i.bankId.max {
		return err{ErrBankIdTooLong, fmt.Sprintf("The bank-id must not be longer than %d letters, given: %d", i.bankId.max, length), nil}
	}
	// TODO: handle special_u16
	if length == 0 && i.bankId.required {
		return err{ErrBankIdMissing, msgBankIdMissing, nil}
	}
	return nil
}

// GetBankVerifier returns the bank verifier for the given country or nil, if no verifier is available for this country.
func GetBankVerifier(country iso.CountryCodeString) BankCountryVerifier {
	verifier, ok := allBankCountryVerifier[country]
	if !ok {
		return nil
	}
	return verifier
}
