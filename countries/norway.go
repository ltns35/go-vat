package countries

type norway struct {
	Country
}

var Norway = norway{
	Country: Country{
		Name: "Norway",
		Codes: []string{
			"NO",
			"NOR",
			"578",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					3,
					2,
					7,
					6,
					5,
					4,
					3,
					2,
				},
			},
			Regex: []string{
				"^(NO)(\\d{9})(MVA)?$",
			},
		},
	},
}

func (n norway) Calc(vat string) bool {

	total := 0
	// See http://www.brreg.no/english/coordination/number.html

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 8; i++ {
		num, _ := getIntFromStringAt(vat, i)
		total += num * n.Rules.Multipliers["common"][i]
	}

	// Establish check digits by getting modulus 11. Check digits > 9 are invalid
	total = 11 - (total % 11)

	if total == 11 {
		total = 0
	}

	if total < 10 {

		// Compare it with the last character of the VAT number. If it's the same, then it's valid.
		expect, _ := getIntFromStringAt(vat, 8)

		return total == expect
	}

	return false
}

func (n norway) GetCountry() Country {
	return n.Country
}
