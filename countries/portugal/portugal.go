package portugal

import (
	"strconv"

	"github.com/ltns35/go-vat/countries"
)

type portugal struct {
	countries.Country
}

var VAT = portugal{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"PT",
			"PRT",
			"020",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					9,
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
				"^(PT)(\\d{9})$",
			},
		},
	},
}

func (p portugal) GetCountry() countries.Country {
	return p.Country
}

func (p portugal) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 8; i++ {
		total += int(vat[i]) * p.Rules.Multipliers["common"][i]
	}

	// Establish check digits subtracting modulus 11 from 11.
	total = 11 - (total % 11)
	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	lastChar := string(vat[len(vat)-1])
	expect, err := strconv.Atoi(lastChar)
	if err != nil {
		return false
	}

	return total == expect
}