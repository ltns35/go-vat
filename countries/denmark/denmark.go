package denmark

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type denmark struct {
	countries.Country
}

var VAT = denmark{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"DK",
			"DNK",
			"208",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"common": {
					2,
					7,
					6,
					5,
					4,
					3,
					2,
					1,
				},
			},
			Regex: []string{
				"^(DK)(\\d{8})$",
			},
		},
	},
}

func (d denmark) Calc(vat string) bool {

	total := 0

	for i := 0; i < 8; i++ {
		num := utils.IntAt(vat, i)
		total += num * d.Rules.Multipliers["common"][i]
	}

	return total%11 == 0
}

func (d denmark) GetCountry() countries.Country {
	return d.Country
}
