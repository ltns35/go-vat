package romania

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type romania struct {
	countries.Country
}

var VAT = romania{
	Country: countries.Country{
		Name: "Romania",
		Codes: []string{
			"RO",
			"ROU",
			"642",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					7,
					5,
					3,
					2,
					1,
					7,
					5,
					3,
					2,
				},
			},
			Regex: []string{
				"^(RO)([1-9]\\d{1,9})$",
			},
		},
	},
}

func (r romania) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	vatLength := len(vat)

	multipliers := r.Rules.Multipliers["common"][10-vatLength:]

	for i := 0; i < vatLength-1; i++ {
		num := utils.IntAt(vat, i)
		total += num * multipliers[i]
	}

	// Establish check digits by getting modulus 11.
	total = (10 * total) % 11
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, vatLength-1)

	return total == expect
}

func (r romania) GetCountry() countries.Country {
	return r.Country
}
