package countries

type bolivia struct {
	Country
}

var Bolivia = bolivia{
	Country: Country{
		Name: "Bolivia",
		Codes: []string{
			"BO",
			"BOL",
			"068",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(BO)(\\d{7})$",
			},
		},
	},
}

func (b bolivia) Calc(vat string) bool {
	return len(vat) == 7
}

func (b bolivia) GetCountry() Country {
	return b.Country
}
