package countries

import (
	"math"
	"strconv"

	"github.com/ltns35/go-vat/countries/utils"
)

type italy struct {
	Country
}

var Italy = italy{
	Country: Country{
		Name: "Italy",
		Codes: []string{
			"IT",
			"ITA",
			"380",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					1,
					2,
					1,
					2,
					1,
					2,
					1,
					2,
					1,
					2,
				},
			},
			Regex: []string{
				"^(IT)(\\d{11})$",
			},
		},
	},
}

func (i italy) Calc(vat string) bool {

	beginNumStr := vat[:7]
	beginNum, _ := strconv.Atoi(beginNumStr)

	if beginNum == 0 {
		return false
	}

	// The last three digits are the issuing office, and cannot exceed more 201, unless 999 or 888
	tempStr := vat[7:10]
	temp, _ := strconv.Atoi(tempStr)

	if temp < 1 || (temp > 201 && temp != 999 && temp != 888) {
		return false
	}

	var total float64 = 0

	// Extract the next digit and multiply by the appropriate
	for j := 0; j < 10; j++ {

		num := utils.IntAt(vat, j)
		temp = num * i.Rules.Multipliers["common"][j]

		if temp > 9 {
			total += math.Floor(float64(temp/10)) + math.Mod(float64(temp), 10)
		} else {
			total += float64(temp)
		}

	}

	// Establish check digit.
	total = 10 - math.Mod(total, 10)
	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 10)

	return total == float64(expect)
}

func (i italy) GetCountry() Country {
	return i.Country
}
