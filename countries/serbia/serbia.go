package serbia

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type serbia struct {
	countries.Country
}

var VAT = serbia{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"RS",
			"SRB",
			"688",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(RS)(\\d{9})$",
			},
		},
	},
}

func (s serbia) Calc(vat string) bool {

	// Checks the check digits of a Serbian VAT number using ISO 7064, MOD 11-10 for check digit.
	product := 10
	sum := 0

	for i := 0; i < 8; i++ {
		num := utils.IntAt(vat, i)
		sum = (num + product) % 10
		if sum == 0 {
			sum = 10
		}
		product = (2 * sum) % 11
	}

	// Now check that we have the right check digit
	lastDigit := utils.IntAt(vat, 8)
	checkDigit := (product + lastDigit) % 10

	const expect = 1

	return checkDigit == expect
}

func (s serbia) GetCountry() countries.Country {
	return s.Country
}
