package countries

import (
	"github.com/ltns35/go-vat/utils"
)

type finland struct {
	Country
}

var Finland = finland{
	Country: Country{
		Name: "Finland",
		Codes: []string{
			"FI",
			"FIN",
			"246",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					7,
					9,
					10,
					5,
					8,
					4,
					2,
				},
			},
			Regex: []string{
				"^(FI)(\\d{8})$",
			},
		},
	},
}

func (f finland) Calc(vat string) bool {
	total := 0

	// Extract the next digit and multiply by the counter.
	for i := range 7 {
		num := utils.IntAt(vat, i)
		total += num * f.Rules.Multipliers["common"][i]
	}

	// Establish check digit.
	total = 11 - (total % 11)
	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 7)

	return total == expect
}

func (f finland) GetCountry() Country {
	return f.Country
}
