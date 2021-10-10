package france

import (
	"regexp"
	"strconv"

	"github.com/ltns35/go-vat/countries"
)

type france struct {
	countries.Country
}

var VAT = france{
	Country: countries.Country{
		Name: "France",
		Codes: []string{
			"FR",
			"FRA",
			"250",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(FR)(\\d{11})$",
				"^(FR)([A-HJ-NP-Z]\\d{10})$",
				"^(FR)(\\d[A-HJ-NP-Z]\\d{9})$",
				"^(FR)([A-HJ-NP-Z]{2}\\d{9})$",
			},
		},
	},
}

func (f france) Calc(vat string) bool {

	// Checks the check digits of a French VAT number.
	regex := regexp.MustCompile("^\\d{11}$")

	if !regex.MatchString(vat) {
		return true
	}

	// Extract the last nine digits as an integer.
	totalStr := vat[2:]
	total, _ := strconv.Atoi(totalStr)

	// Establish check digit.
	total = (total*100 + 12) % 97

	// Compare it with the first 2 characters of the VAT number.
	// If it's the same, then it's valid.
	expectStr := vat[:2]
	expect, _ := strconv.Atoi(expectStr)

	return total == expect
}

func (f france) GetCountry() countries.Country {
	return f.Country
}
