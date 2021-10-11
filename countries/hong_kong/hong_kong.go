package hong_kong

import (
	"github.com/ltns35/go-vat/countries"
)

type hongKong struct {
	countries.Country
}

var VAT = hongKong{
	Country: countries.Country{
		Name: "Hong Kong",
		Codes: []string{
			"HK",
			"HKG",
			"344",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(HK)(\\d{8})$",
			},
		},
	},
}

func (h hongKong) Calc(vat string) bool {
	return len(vat) == 8
}

func (h hongKong) GetCountry() *countries.Country {
	return &h.Country
}
