package countries

import (
	"math"
	"strconv"
)

type belgium struct {
	Country
}

var Belgium = belgium{
	Country: Country{
		Name: "Belgium",
		Codes: []string{
			"BE",
			"BEL",
			"056",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(BE)([01]?\\d{9})$",
			},
		},
	},
}

func (b belgium) Calc(vat string) bool {
	newVat := vat
	if len(vat) == 9 {
		newVat = "0" + vat
	}

	strNum := newVat[:8]
	num, _ := strconv.Atoi(strNum)

	check := 97 - math.Mod(float64(num), 97)

	lastDigitsStr := newVat[8:10]
	lastDigits, _ := strconv.Atoi(lastDigitsStr)

	return check == float64(lastDigits)
}

func (b belgium) GetCountry() Country {
	return b.Country
}
