package countries

import (
	"strconv"

	"github.com/ltns35/go-vat/countries/utils"
)

type malta struct {
	Country
}

var Malta = malta{
	Country: Country{
		Name: "Malta",
		Codes: []string{
			"MT",
			"MLT",
			"470",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					3,
					4,
					6,
					7,
					8,
					9,
				},
			},
			Regex: []string{
				"^(MT)([1-9]\\d{7})$",
			},
		},
	},
}

func (m malta) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 6; i++ {
		num := utils.IntAt(vat, i)
		total += num * m.Rules.Multipliers["common"][i]
	}

	// Establish check digits by getting modulus 37.
	total = 37 - (total % 37)

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expectStr := vat[6:]
	expect, _ := strconv.Atoi(expectStr)

	return total == expect
}

func (m malta) GetCountry() Country {
	return m.Country
}
