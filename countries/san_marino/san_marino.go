package san_marino

import (
	"github.com/ltns35/go-vat/countries"
)

type sanMarino struct {
	countries.Country
}

var VAT = sanMarino{
	Country: countries.Country{
		Name: "San Marino",
		Codes: []string{
			"SM",
			"SMR",
			"674",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(SM)([0-9]{5})$",
			},
		},
	},
}

func (s sanMarino) Calc(vat string) bool {
	return len(vat) == 5
}

func (s sanMarino) GetCountry() *countries.Country {
	return &s.Country
}
