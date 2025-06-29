package countries

import (
	"strconv"
)

type slovakia struct {
	Country
}

var Slovakia = slovakia{
	Country: Country{
		Name: "Slovakia",
		Codes: []string{
			"SK",
			"SVK",
			"703",
		},
		Rules: CountryRules{
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

func (s slovakia) GetCountry() Country {
	return s.Country
}
