package kazakhstan

import (
	"github.com/ltns35/go-vat/countries"
)

type kazakhstan struct {
	countries.Country
}

var VAT = kazakhstan{
	Country: countries.Country{
		Name: "Kazakhstan",
		Codes: []string{
			"KZ",
			"KAZ",
			"398",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(KZ)(\\d{12})$",
			},
		},
	},
}

func (k kazakhstan) Calc(vat string) bool {
	return len(vat) == 12
}

func (k kazakhstan) GetCountry() *countries.Country {
	return &k.Country
}
