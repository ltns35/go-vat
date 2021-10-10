package luxembourg

import (
	"strconv"

	"github.com/ltns35/go-vat/countries"
)

type luxembourg struct {
	countries.Country
}

var VAT = luxembourg{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"LU",
			"LUX",
			"442",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(LU)(\\d{8})$",
			},
		},
	},
}

func (l luxembourg) Calc(vat string) bool {

	expectStr := vat[6:8]
	expect, _ := strconv.Atoi(expectStr)

	// Checks the check digits of a VAT VAT number.
	checkDigitStr := vat[:6]
	checkDigit, _ := strconv.Atoi(checkDigitStr)
	checkDigit = checkDigit % 89

	return checkDigit == expect
}

func (l luxembourg) GetCountry() countries.Country {
	return l.Country
}
