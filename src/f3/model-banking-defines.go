package f3

import "github.com/xeus2001/interview-accountapi/src/iso/countryCode"

func init() {
	// TODO: Add all bank information.
	allBankCountryVerifier[countryCode.UnitedKingdom] = bankInfo{
		countryCode:   countryCode.UnitedKingdom,
		bankId:        info{required: true, min: 6, max: 6, special: special_uk_sort_code},
		bic:           info{required: true},
		bankIdCode:    info{required: true, value: "GBDSC"},
		accountNumber: info{min: 8, max: 8},
	}
	allBankCountryVerifier[countryCode.Germany] = bankInfo{
		countryCode:   countryCode.Germany,
		bankId:        info{required: true, min: 8, max: 8},
		bic:           info{},
		bankIdCode:    info{required: true, value: "DEBLZ"},
		accountNumber: info{min: 7, max: 7},
	}
	allBankCountryVerifier[countryCode.Usa] = bankInfo{
		countryCode:   countryCode.Usa,
		bankId:        info{required: true, min: 9, max: 9, special: special_aba},
		bic:           info{},
		bankIdCode:    info{required: true, value: "CHBCC"},
		accountNumber: info{min: 6, max: 17},
	}
}
