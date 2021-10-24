package countries

type hongKong struct {
	Country
}

var HongKong = hongKong{
	Country: Country{
		Name: "Hong Kong",
		Codes: []string{
			"HK",
			"HKG",
			"344",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(HK)(\\d{8})$",
			},
		},
	},
}

func (h hongKong) Calc(vat string) bool {
	return len(vat) == 8
}

func (h hongKong) GetCountry() Country {
	return h.Country
}
