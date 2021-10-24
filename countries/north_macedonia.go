package countries

type northMacedonia struct {
	Country
}

var NorthMacedonia = northMacedonia{
	Country: Country{
		Name: "North Macedonia",
		Codes: []string{
			"MK",
			"MKD",
			"807",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(MK)(\\d{13})$",
			},
		},
	},
}

func (n northMacedonia) Calc(vat string) bool {
	return len(vat) == 13
}

func (n northMacedonia) GetCountry() Country {
	return n.Country
}
