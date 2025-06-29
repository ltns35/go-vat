package countries

import (
	"github.com/ltns35/go-vat/utils"
)

type greece struct {
	Country
}

var Greece = greece{
	Country: Country{
		Name: "Greece",
		Codes: []string{
			"GR",
			"GRC",
			"300",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					256,
					128,
					64,
					32,
					16,
					8,
					4,
					2,
				},
			},
			Regex: []string{
				"^(EL)(\\d{9})$",
			},
		},
	},
}

func (g greece) Calc(vat string) bool {
	total := 0

	// eight character numbers should be prefixed with an 0.
	newVat := vat

	if len(vat) == 8 {
		newVat = "0" + vat
	}

	// Extract the next digit and multiply by the counter.
	for i := range 8 {
		num := utils.IntAt(newVat, i)
		total += num * g.Rules.Multipliers["common"][i]
	}

	// Establish check digit.
	total %= 11

	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 8)

	return total == expect
}

func (g greece) GetCountry() Country {
	return g.Country
}
