package countries

type sanMarino struct {
	Country
}

var SanMarino = sanMarino{
	Country: Country{
		Name: "San Marino",
		Codes: []string{
			"SM",
			"SMR",
			"674",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(SM)(\\d{5})$",
			},
		},
	},
}

func (s sanMarino) Calc(vat string) bool {
	return len(vat) == 5
}

func (s sanMarino) GetCountry() Country {
	return s.Country
}
