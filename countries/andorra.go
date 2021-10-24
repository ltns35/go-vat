package countries

type andorra struct {
	Country
}

var Andorra = andorra{
	Country: Country{
		Name: "Andorra",
		Codes: []string{
			"AD",
			"AND",
			"020",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(AD)([fealecdgopuFEALECDGOPU]{1}\\d{6}[fealecdgopuFEALECDGOPU]{1})$",
			},
		},
	},
}

func (a andorra) Calc(vat string) bool {
	return len(vat) == 8
}

func (a andorra) GetCountry() Country {
	return a.Country
}
