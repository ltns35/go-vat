package switzerland

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type switzerland struct {
	countries.Country
}

var VAT = switzerland{
	Country: countries.Country{
		Name: "Switzerland",
		Codes: []string{
			"CH",
			"CHE",
			"756",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					5,
					4,
					3,
					2,
					7,
					6,
					5,
					4,
				},
			},
			Regex: []string{
				"^(CHE)(\\d{9})(MWST|TVA|IVA)?$",
			},
		},
	},
}

func (s switzerland) Calc(vat string) bool {

	total := 0
	for i := 0; i < 8; i++ {

		num := utils.IntAt(vat, i)
		total += num * s.Rules.Multipliers["common"][i]
	}

	// Establish check digits
	total = 11 - (total % 11)
	if total == 10 {
		return false
	}

	if total == 11 {
		total = 0
	}

	// Check to see if the check digit given is correct, If not, we have an error with the VAT number
	expect := utils.IntAt(vat, 8)

	return total == expect
}

func (s switzerland) GetCountry() countries.Country {
	return s.Country
}
