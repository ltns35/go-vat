package countries

import (
	"strconv"
)

type luxembourg struct {
	Country
}

var Luxembourg = luxembourg{
	Country: Country{
		Name: "Luxembourg",
		Codes: []string{
			"LU",
			"LUX",
			"442",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(LU)(\\d{8})$",
			},
		},
	},
}

func (l luxembourg) Calc(vat string) bool {

	expectStr := vat[6:8]
	expect, _ := strconv.Atoi(expectStr)

	// Checks the check digits of a VAT number.
	checkDigitStr := vat[:6]
	checkDigit, _ := strconv.Atoi(checkDigitStr)
	checkDigit = checkDigit % 89

	return checkDigit == expect
}

func (l luxembourg) GetCountry() Country {
	return l.Country
}
