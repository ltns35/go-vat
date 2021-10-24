package countries

type ukraine struct {
	Country
}

var Ukraine = ukraine{
	Country: Country{
		Name: "Ukraine",
		Codes: []string{
			"UA",
			"UKR",
			"804",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(UA)(\\d{12})$",
			},
		},
	},
}

func (u ukraine) Calc(vat string) bool {
	return len(vat) == 12
}

func (u ukraine) GetCountry() Country {
	return u.Country
}
