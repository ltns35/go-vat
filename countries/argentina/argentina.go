package argentina

import (
	"github.com/ltns35/go-vat/countries"
)

type argentina struct {
	countries.Country
}

var VAT = argentina{
	Country: countries.Country{
		Name: "Argentina",
		Codes: []string{
			"AR",
			"ARG",
			"032",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(AR)(\\d{11})$",
			},
		},
	},
}

func (a argentina) Calc(vat string) bool {
	return len(vat) == 11
}

func (a argentina) GetCountry() *countries.Country {
	return &a.Country
}
