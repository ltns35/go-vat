package croatia

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type croatia struct {
	countries.Country
}

var VAT = croatia{
	Country: countries.Country{
		Name: "Croatia",
		Codes: []string{
			"HR",
			"HRV",
			"191",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(HR)(\\d{11})$",
			},
		},
	},
}

func (c croatia) Calc(vat string) bool {

	// Checks the check digits of a Croatian VAT number using ISO 7064, MOD 11-10 for check digit.
	product := 10
	sum := 0

	for i := 0; i < 10; i++ {
		// Extract the next digit and implement the algorithm
		num := utils.IntAt(vat, i)

		sum = (num + product) % 10
		if sum == 0 {
			sum = 10
		}

		product = (2 * sum) % 11
	}

	// Now check that we have the right check digit
	expect := utils.IntAt(vat, 10)

	return (product+expect)%10 == 1
}

func (c croatia) GetCountry() countries.Country {
	return c.Country
}
