package countries

type kazakhstan struct {
	Country
}

var Kazakhstan = kazakhstan{
	Country: Country{
		Name: "Kazakhstan",
		Codes: []string{
			"KZ",
			"KAZ",
			"398",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(KZ)(\\d{12})$",
			},
		},
	},
}

func (k kazakhstan) Calc(vat string) bool {
	return len(vat) == 12
}

func (k kazakhstan) GetCountry() Country {
	return k.Country
}
