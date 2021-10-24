package countries

import (
	"github.com/ltns35/go-vat/countries/utils"
)

type denmark struct {
	Country
}

var Denmark = denmark{
	Country: Country{
		Name: "Denmark",
		Codes: []string{
			"DK",
			"DNK",
			"208",
		},
		Rules: CountryRules{
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

func (d denmark) GetCountry() Country {
	return d.Country
}
