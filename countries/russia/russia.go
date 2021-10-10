package russia

import (
	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type russia struct {
	countries.Country
}

var VAT = russia{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"RU",
			"RUS",
			"643",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"m_1": {
					2,
					4,
					10,
					3,
					5,
					9,
					4,
					6,
					8,
					0,
				},
				"m_2": {
					7,
					2,
					4,
					10,
					3,
					5,
					9,
					4,
					6,
					8,
					0,
				},
				"m_3": {
					3,
					7,
					2,
					4,
					10,
					3,
					5,
					9,
					4,
					6,
					8,
					0,
				},
			},
			Regex: []string{
				"^(RU)(\\d{10}|\\d{12})$",
			},
		},
	},
}

func (r russia) Calc(vat string) bool {
	return check10DigitINN(vat, r.Rules) || check12DigitINN(vat, r.Rules)
}

func (r russia) GetCountry() countries.Country {
	return r.Country
}

func check10DigitINN(vat string, rules countries.CountryRules) bool {

	total := 0

	if len(vat) == 10 {
		for i := 0; i < 10; i++ {
			num := utils.IntAt(vat, i)
			total += num * rules.Multipliers["m_1"][i]
		}

		total = total % 11
		if total > 9 {
			total = total % 10
		}

		// Compare it with the last character of the VAT number. If it is the same, then it's valid
		expect := utils.IntAt(vat, 9)

		return total == expect
	}

	return false
}

func check12DigitINN(vat string, rules countries.CountryRules) bool {

	total1 := 0
	total2 := 0

	if len(vat) == 12 {
		for j := 0; j < 11; j++ {
			num := utils.IntAt(vat, j)
			total1 += num * rules.Multipliers["m_2"][j]
		}

		total1 = total1 % 11

		if total1 > 9 {
			total1 = total1 % 10
		}

		for k := 0; k < 11; k++ {
			num := utils.IntAt(vat, k)
			total2 += num * rules.Multipliers["m_3"][k]
		}

		total2 = total2 % 11
		if total2 > 9 {
			total2 = total2 % 10
		}

		// Compare the first check with the 11th character and the second check with the 12th and last
		// character of the VAT number. If they're both the same, then it's valid
		eleventh := utils.IntAt(vat, 10)
		expect := total1 == eleventh

		twelfth := utils.IntAt(vat, 11)
		expect2 := total2 == twelfth

		return expect && expect2
	}

	return false
}
