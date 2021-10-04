package countries

type estonia struct {
	Country
}

var Estonia = estonia{
	Country: Country{
		Name: "Estonia",
		Codes: []string{
			"EE",
			"EST",
			"233",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					3,
					7,
					1,
					3,
					7,
					1,
					3,
					7,
				},
			},
			Regex: []string{
				"^(EE)(10\\d{7})$",
			},
		},
	},
}

func (e estonia) Calc(vat string) bool {

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 8; i++ {
		num, _ := getIntFromStringAt(vat, i)
		total += num * e.Rules.Multipliers["common"][i]
	}

	// Establish check digits using modulus 10.
	total = 10 - (total % 10)
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect, _ := getIntFromStringAt(vat, 8)
	return total == expect
}

func (e estonia) GetCountry() Country {
	return e.Country
}
