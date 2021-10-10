package andorra

import (
	"github.com/ltns35/go-vat/countries"
)

type andorra struct {
	countries.Country
}

var VAT = andorra{
	Country: countries.Country{
		Name: "Andorra",
		Codes: []string{
			"AD",
			"AND",
			"020",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(AD)([fealecdgopuFEALECDGOPU]{1}\\d{6}[fealecdgopuFEALECDGOPU]{1})$",
			},
		},
	},
}

func (a andorra) Calc(vat string) bool {
	return len(vat) == 8
}

func (a andorra) GetCountry() *countries.Country {
	return &a.Country
}
