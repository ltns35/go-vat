package countries

import "github.com/ltns35/go-vat/utils"

type australia struct {
	Country
}

var Australia = australia{
	Country: Country{
		Name: "Australia",
		Codes: []string{
			"AU",
			"AUS",
			"036",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					10, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19,
				},
			},
			Regex: []string{
				`^(AU)?(\d{11})$`,
			},
		},
	},
}

func (a australia) Calc(vat string) bool {
	total := 0

	for i := range len(vat) {
		num := utils.IntAt(vat, i)

		// If the first digit, subtract 1
		if i == 0 {
			num -= 1
		}

		total += num * a.Rules.Multipliers["common"][i]
	}

	return total > 0 && total%89 == 0
}

func (a australia) GetCountry() Country {
	return a.Country
}
