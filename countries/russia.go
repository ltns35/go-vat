package countries

type russia struct {
	Country
}

var Russia = russia{
	Country: Country{
		Name: "Russia",
		Codes: []string{
			"RU",
			"RUS",
			"643",
		},
		Rules: CountryRules{
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

func (r russia) GetCountry() Country {
	return r.Country
}

func check10DigitINN(vat string, rules CountryRules) bool {

	total := 0

	if len(vat) == 10 {
		for i := 0; i < 10; i++ {
			num, _ := getIntFromStringAt(vat, i)
			total += num * rules.Multipliers["m_1"][i]
		}

		total = total % 11
		if total > 9 {
			total = total % 10
		}

		// Compare it with the last character of the VAT number. If it is the same, then it's valid
		expect, _ := getIntFromStringAt(vat, 9)

		return total == expect
	}

	return false
}

func check12DigitINN(vat string, rules CountryRules) bool {

	total1 := 0
	total2 := 0

	if len(vat) == 12 {
		for j := 0; j < 11; j++ {
			num, _ := getIntFromStringAt(vat, j)
			total1 += num * rules.Multipliers["m_2"][j]
		}

		total1 = total1 % 11

		if total1 > 9 {
			total1 = total1 % 10
		}

		for k := 0; k < 11; k++ {
			num, _ := getIntFromStringAt(vat, k)
			total2 += num * rules.Multipliers["m_3"][k]
		}

		total2 = total2 % 11
		if total2 > 9 {
			total2 = total2 % 10
		}

		// Compare the first check with the 11th character and the second check with the 12th and last
		// character of the VAT number. If they're both the same, then it's valid
		eleventh, _ := getIntFromStringAt(vat, 10)
		expect := total1 == eleventh

		twelfth, _ := getIntFromStringAt(vat, 11)
		expect2 := total2 == twelfth

		return expect && expect2
	}

	return false
}
