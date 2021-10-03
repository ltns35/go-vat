package countries

type hungary struct {
	Country
}

var Hungary = hungary{
	Country: Country{
		Name: "Hungary",
		Codes: []string{
			"HU",
			"HUN",
			"348",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					9,
					7,
					3,
					1,
					9,
					7,
					3,
				},
			},
			Regex: []string{
				"^(HU)(\\d{8})$",
			},
		},
	},
}

func (h hungary) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 7; i++ {
		num, _ := getIntFromStringAt(vat, i)
		total += num * h.Rules.Multipliers["common"][i]
	}

	// Establish check digit.
	total = 10 - (total % 10)
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect, _ := getIntFromStringAt(vat, 7)

	return total == expect
}

func (h hungary) GetCountry() Country {
	return h.Country
}
