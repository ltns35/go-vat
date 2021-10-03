package countries

import (
	"strconv"
)

type germany struct {
	Country
}

var Germany = germany{
	Country: Country{
		Name: "Germany",
		Codes: []string{
			"DE",
			"DEU",
			"276",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(DE)([1-9]\\d{8})$",
			},
		},
	},
}

func (g germany) Calc(vat string) bool {

	// Checks the check digits of a German VAT number.
	var product = 10
	var sum = 0
	var checkDigit = 0

	for i := 0; i < 8; i++ {
		// Extract the next digit and implement peculiar algorithm!.
		num, _ := getIntFromStringAt(vat, i)
		sum = (num + product) % 10

		if sum == 0 {
			sum = 10
		}

		product = (2 * sum) % 11
	}

	// Establish check digit.
	if 11-product == 10 {
		checkDigit = 0
	} else {
		checkDigit = 11 - product
	}

	// Compare it with the last character of the VAT number.
	// If the same, then it is a valid
	expectStr := vat[8:]
	expect, _ := strconv.Atoi(expectStr)

	return checkDigit == expect
}

func (g germany) GetCountry() Country {
	return g.Country
}
