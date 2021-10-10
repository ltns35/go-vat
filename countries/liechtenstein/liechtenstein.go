package liechtenstein

import (
	"github.com/ltns35/go-vat/countries"
)

type liechtenstein struct {
	countries.Country
}

var VAT = liechtenstein{
	Country: countries.Country{
		Name: "Liechtenstein",
		Codes: []string{
			"LI",
			"LIE",
			"807",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(FL)(\\d{11})$",
			},
		},
	},
}

func (l liechtenstein) Calc(vat string) bool {
	return len(vat) == 11
}

func (l liechtenstein) GetCountry() *countries.Country {
	return &l.Country
}
