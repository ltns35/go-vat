package countries

type albania struct {
	Country
}

var Albania = albania{
	Country: Country{
		Name: "Albania",
		Codes: []string{
			"AL",
			"ALB",
			"008",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(AL)([JKL]\\d{8}[A-Z])$",
			},
		},
	},
}

func (a albania) Calc(vat string) bool {
	return len(vat) == 10
}

func (a albania) GetCountry() Country {
	return a.Country
}
