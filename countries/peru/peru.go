package peru

import (
	"github.com/ltns35/go-vat/countries"
)

type peru struct {
	countries.Country
}

var VAT = peru{
	Country: countries.Country{
		Name: "Peru",
		Codes: []string{
			"PE",
			"PER",
			"604",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(PE)([0-9]{11})$",
			},
		},
	},
}

func (p peru) Calc(vat string) bool {
	return len(vat) == 11
}

func (p peru) GetCountry() *countries.Country {
	return &p.Country
}
