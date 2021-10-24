package countries

import (
	"math"

	"github.com/ltns35/go-vat/countries/utils"
)

type austria struct {
	Country
}

var Austria = austria{
	Country: Country{
		Name: "Austria",
		Codes: []string{
			"AT",
			"AUT",
			"040",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					1,
					2,
					1,
					2,
					1,
					2,
					1,
				},
			},
			Regex: []string{
				"^(AT)U(\\d{8})$",
			},
		},
	},
}

func (a austria) Calc(vat string) bool {

	var total float64 = 0

	for i := 0; i < 7; i++ {

		num := utils.IntAt(vat, i)

		var temp = float64(num * a.Rules.Multipliers["common"][i])

		if temp > 9 {
			total += math.Floor(temp/10) + math.Mod(temp, 10)
		} else {
			total += temp
		}
	}

	total = 10 - math.Mod(total+4, 10)
	if total == 10 {
		total = 0
	}

	lastNum := utils.IntAt(vat, len(vat)-1)

	return int(total) == lastNum
}

func (a austria) GetCountry() Country {
	return a.Country
}
