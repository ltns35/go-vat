package belarus

import (
	"github.com/ltns35/go-vat/countries"
)

type belarus struct {
	countries.Country
}

var VAT = belarus{
	Country: countries.Country{
		Name: "Belarus",
		Codes: []string{
			"BY",
			"BLR",
			"112",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(BY)(\\d{9})$",
			},
		},
	},
}

func (b belarus) Calc(vat string) bool {
	return len(vat) == 9
}

func (b belarus) GetCountry() *countries.Country {
	return &b.Country
}
