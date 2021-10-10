package slovenia

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type slovenia struct {
	countries.Country
}

var VAT = slovenia{
	Country: countries.Country{
		Name: "Slovenia",
		Codes: []string{
			"SI",
			"SVN",
			"705",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					8,
					7,
					6,
					5,
					4,
					3,
					2,
				},
			},
			Regex: []string{
				"^(SI)([1-9]\\d{7})$",
			},
		},
	},
}

func (s slovenia) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 7; i++ {
		num := utils.IntAt(vat, i)
		total += num * s.Rules.Multipliers["common"][i]
	}

	// Establish check digits using modulus 11
	total = 11 - (total % 11)
	if total == 10 {
		total = 0
	}

	// Compare the number with the last character of the VAT number. If it is the
	// same, then it's a valid check digit.
	expect := utils.IntAt(vat, 7)

	return !!(total != 11 && total == expect)
}

func (s slovenia) GetCountry() *countries.Country {
	return &s.Country
}
