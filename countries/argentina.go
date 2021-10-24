package countries

type argentina struct {
	Country
}

var Argentina = argentina{
	Country: Country{
		Name: "Argentina",
		Codes: []string{
			"AR",
			"ARG",
			"032",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(AR)(\\d{11})$",
			},
		},
	},
}

func (a argentina) Calc(vat string) bool {
	return len(vat) == 11
}

func (a argentina) GetCountry() Country {
	return a.Country
}
