package countries

import (
	"math"
	"strconv"

	"github.com/ltns35/go-vat/countries/utils"
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
				"^(BE)(0?\\d{9})$",
			},
		},
	},
}

func (b belgium) Calc(vat string) bool {

	newVat := vat
	if len(vat) == 9 {
		newVat = "0" + vat
	}

	numAtFirstIndex := utils.IntAt(newVat, 1)
	if numAtFirstIndex == 0 {
		return false
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
