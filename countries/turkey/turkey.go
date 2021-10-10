package turkey

import (
	"github.com/ltns35/go-vat/countries"
)

type turkey struct {
	countries.Country
}

var VAT = turkey{
	Country: countries.Country{
		Name: "Turkey",
		Codes: []string{
			"TR",
			"TUR",
			"792",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(TR)([0-9]{10})$",
			},
		},
	},
}

func (t turkey) Calc(vat string) bool {
	return len(vat) == 10
}

func (t turkey) GetCountry() *countries.Country {
	return &t.Country
}
