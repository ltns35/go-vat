package hungary

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type hungary struct {
	countries.Country
}

var VAT = hungary{
	Country: countries.Country{
		Name: "Hungary",
		Codes: []string{
			"HU",
			"HUN",
			"348",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					9,
					7,
					3,
					1,
					9,
					7,
					3,
				},
			},
			Regex: []string{
				"^(HU)(\\d{8})$",
			},
		},
	},
}

func (h hungary) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 7; i++ {
		num := utils.IntAt(vat, i)
		total += num * h.Rules.Multipliers["common"][i]
	}

	// Establish check digit.
	total = 10 - (total % 10)
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 7)

	return total == expect
}

func (h hungary) GetCountry() *countries.Country {
	return &h.Country
}
