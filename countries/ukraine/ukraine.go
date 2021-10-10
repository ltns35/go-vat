package ukraine

import (
	"github.com/ltns35/go-vat/countries"
)

type ukraine struct {
	countries.Country
}

var VAT = ukraine{
	Country: countries.Country{
		Name: "Ukraine",
		Codes: []string{
			"UA",
			"UKR",
			"804",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(UA)([0-9]{12})$",
			},
		},
	},
}

func (u ukraine) Calc(vat string) bool {
	return len(vat) == 12
}

func (u ukraine) GetCountry() *countries.Country {
	return &u.Country
}
