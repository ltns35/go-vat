package countries

type peru struct {
	Country
}

var Peru = peru{
	Country: Country{
		Name: "Peru",
		Codes: []string{
			"PE",
			"PER",
			"604",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(PE)(\\d{11})$",
			},
		},
	},
}

func (p peru) Calc(vat string) bool {
	return len(vat) == 11
}

func (p peru) GetCountry() Country {
	return p.Country
}
