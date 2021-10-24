package countries

type liechtenstein struct {
	Country
}

var Liechtenstein = liechtenstein{
	Country: Country{
		Name: "Liechtenstein",
		Codes: []string{
			"LI",
			"LIE",
			"807",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(FL)(\\d{11})$",
			},
		},
	},
}

func (l liechtenstein) Calc(vat string) bool {
	return len(vat) == 11
}

func (l liechtenstein) GetCountry() Country {
	return l.Country
}
