package countries

import (
	"regexp"

	"github.com/ltns35/go-vat/countries/utils"
)

type latvia struct {
	Country
}

var Latvia = latvia{
	Country: Country{
		Name: "Latvia",
		Codes: []string{
			"LV",
			"LVA",
			"428",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					9,
					1,
					4,
					8,
					3,
					10,
					2,
					5,
					7,
					6,
				},
			},
			Regex: []string{
				"^(LV)(\\d{11})$",
			},
		},
	},
}

func (l latvia) Calc(vat string) bool {

	total := 0

	// Differentiate between legal entities and natural bodies. For the latter we simply check that
	// the first six digits correspond to valid DDMMYY dates.
	regex := regexp.MustCompile("^[0-3]")

	if regex.MatchString(vat) {
		regex = regexp.MustCompile("^[0-3][0-9][0-1][0-9]")

		return regex.MatchString(vat)
	} else {

		// Extract the next digit and multiply by the counter.
		for i := 0; i < 10; i++ {
			num := utils.IntAt(vat, i)
			total += num * l.Rules.Multipliers["common"][i]
		}

		// Establish check digits by getting modulus 11.
		firstDigit := utils.IntAt(vat, 0)
		if total%11 == 4 && firstDigit == 9 {
			total = total - 45
		}

		switch val := total % 11; {
		case val == 4:
			total = 4 - val
			break
		case val > 4:
			total = 14 - val
			break
		case val < 4:
			total = 3 - val
			break
		}

		// Compare it with the last character of the VAT number. If it's the same, then it's valid.
		expect := utils.IntAt(vat, 10)

		return total == expect
	}
}

func (l latvia) GetCountry() Country {
	return l.Country
}
