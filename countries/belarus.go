package countries

type belarus struct {
	Country
}

var Belarus = belarus{
	Country: Country{
		Name: "Belarus",
		Codes: []string{
			"BY",
			"BLR",
			"112",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(BY)(\\d{9})$",
			},
		},
	},
}

func (b belarus) Calc(vat string) bool {
	return len(vat) == 9
}

func (b belarus) GetCountry() Country {
	return b.Country
}
