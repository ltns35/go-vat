package greece

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type greece struct {
	countries.Country
}

var VAT = greece{
	Country: countries.Country{
		Name: "Greece",
		Codes: []string{
			"GR",
			"GRC",
			"300",
		},
		Rules: countries.CountryRules{
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
	for i := 0; i < 8; i++ {
		num := utils.IntAt(newVat, i)
		total += num * g.Rules.Multipliers["common"][i]
	}

	// Establish check digit.
	total = total % 11

	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 8)

	return total == expect
}

func (g greece) GetCountry() countries.Country {
	return g.Country
}
