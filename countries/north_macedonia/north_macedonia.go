package north_macedonia

import (
	"github.com/ltns35/go-vat/countries"
)

type northMacedonia struct {
	countries.Country
}

var VAT = northMacedonia{
	Country: countries.Country{
		Name: "North Macedonia",
		Codes: []string{
			"MK",
			"MKD",
			"807",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(MK)(\\d{13})$",
			},
		},
	},
}

func (n northMacedonia) Calc(vat string) bool {
	return len(vat) == 13
}

func (n northMacedonia) GetCountry() *countries.Country {
	return &n.Country
}
