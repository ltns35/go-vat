package countries

import (
	"strconv"
)

type portugal struct {
	Country
}

var Portugal = portugal{
	Country: Country{
		Name: "Portugal",
		Codes: []string{
			"PT",
			"PRT",
			"020",
		},
		Rules: CountryRules{
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

func (p portugal) GetCountry() Country {
	return p.Country
}

func (p portugal) Calc(vat string) bool {
	total := 0

	// Extract the next digit and multiply by the counter.
	for i := range 8 {
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
