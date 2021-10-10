package albania

import (
	"github.com/ltns35/go-vat/countries"
)

type albania struct {
	countries.Country
}

var VAT = albania{
	Country: countries.Country{
		Name: "Albania",
		Codes: []string{
			"AL",
			"ALB",
			"008",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(AL)([JKL]\\d{8}[A-Z])$",
			},
		},
	},
}

func (a albania) Calc(vat string) bool {
	return len(vat) == 10
}

func (a albania) GetCountry() *countries.Country {
	return &a.Country
}
