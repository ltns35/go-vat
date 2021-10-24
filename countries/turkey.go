package countries

type turkey struct {
	Country
}

var Turkey = turkey{
	Country: Country{
		Name: "Turkey",
		Codes: []string{
			"TR",
			"TUR",
			"792",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(TR)(\\d{10})$",
			},
		},
	},
}

func (t turkey) Calc(vat string) bool {
	return len(vat) == 10
}

func (t turkey) GetCountry() Country {
	return t.Country
}
