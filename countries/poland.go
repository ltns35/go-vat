package countries

import (
	"github.com/ltns35/go-vat/countries/utils"
)

type poland struct {
	Country
}

var Poland = poland{
	Country: Country{
		Name: "Poland",
		Codes: []string{
			"PL",
			"POL",
			"616",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					6,
					5,
					7,
					2,
					3,
					4,
					5,
					6,
					7,
				},
			},
			Regex: []string{
				"^(PL)(\\d{10})$",
			},
		},
	},
}

func (p poland) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 9; i++ {
		num := utils.IntAt(vat, i)
		total += num * p.Rules.Multipliers["common"][i]
	}

	// Establish check digits subtracting modulus 11 from 11.
	total = total % 11
	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 9)

	return total == expect
}

func (p poland) GetCountry() Country {
	return p.Country
}
