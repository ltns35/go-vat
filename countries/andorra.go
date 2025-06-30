package countries

import (
	"strconv"
	"strings"

	"github.com/ltns35/go-vat/utils"
)

type andorra struct {
	Country
}

var Andorra = andorra{
	Country: Country{
		Name: "Andorra",
		Codes: []string{
			"AD",
			"AND",
			"020",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(AD)([fealecdgopuFEALECDGOPU]{1}\\d{6}[fealecdgopuFEALECDGOPU]{1})$",
			},
		},
	},
}

func (a andorra) Calc(vat string) bool {
	firstLetter := strings.ToUpper(utils.StringAt(vat, 0))

	number, err := strconv.Atoi(vat[1:7])
	if err != nil {
		return false
	}

	if !strings.Contains("ACDEFGLOPU", firstLetter) {
		return false
	}

	if firstLetter == "F" && number > 699999 {
		return false
	}

	if strings.Contains("AL", firstLetter) && number > 699999 && number < 800000 {
		return false
	}

	return true
}

func (a andorra) GetCountry() Country {
	return a.Country
}
