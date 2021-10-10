package slovakia

import (
	"strconv"

	"github.com/ltns35/go-vat/countries"
)

type slovakia struct {
	countries.Country
}

var VAT = slovakia{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"SK",
			"SVK",
			"703",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(SK)([1-9]\\d[2346-9]\\d{7})$",
			},
		},
	},
}

func (s slovakia) Calc(vat string) bool {

	expect := 0

	vatNum, _ := strconv.Atoi(vat)

	checkDigit := vatNum % 11

	return checkDigit == expect
}

func (s slovakia) GetCountry() countries.Country {
	return s.Country
}
