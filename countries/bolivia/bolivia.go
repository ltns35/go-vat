package bolivia

import (
	"github.com/ltns35/go-vat/countries"
)

type bolivia struct {
	countries.Country
}

var VAT = bolivia{
	Country: countries.Country{
		Name: "Bolivia",
		Codes: []string{
			"BO",
			"BOL",
			"068",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(BO)(\\d{7})$",
			},
		},
	},
}

func (b bolivia) Calc(vat string) bool {
	return len(vat) == 7
}

func (b bolivia) GetCountry() *countries.Country {
	return &b.Country
}
