package estonia

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type estonia struct {
	countries.Country
}

var VAT = estonia{
	Country: countries.Country{
		Name: "Estonia",
		Codes: []string{
			"EE",
			"EST",
			"233",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					3,
					7,
					1,
					3,
					7,
					1,
					3,
					7,
				},
			},
			Regex: []string{
				"^(EE)(10\\d{7})$",
			},
		},
	},
}

func (e estonia) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 8; i++ {
		num := utils.IntAt(vat, i)
		total += num * e.Rules.Multipliers["common"][i]
	}

	// Establish check digits using modulus 10.
	total = 10 - (total % 10)
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 8)
	return total == expect
}

func (e estonia) GetCountry() countries.Country {
	return e.Country
}
