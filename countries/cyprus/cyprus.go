package cyprus

import (
	"strconv"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type cyprus struct {
	countries.Country
}

var VAT = cyprus{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"CY",
			"CYP",
			"196",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(CY)([0-59]\\d{7}[A-Z])$",
			},
		},
	},
}

func (c cyprus) Calc(vat string) bool {

	// Not allowed to start with '12'
	numStr := vat[:2]
	num, _ := strconv.Atoi(numStr)

	if num == 12 {
		return false
	}

	// Extract the next digit and multiply by the counter.
	total := extractAndMultiplyByCounter(vat, 0)

	// Establish check digit using modulus 26, and translate to char. equivalent.
	total = total % 26
	totalChar := string(rune(total + 65))

	// Check to see if the check digit given is correct
	expect := utils.StringAt(vat, 8)

	return totalChar == expect
}

func (c cyprus) GetCountry() countries.Country {
	return c.Country
}

func extractAndMultiplyByCounter(vat string, total int) int {

	result := total
	for i := 0; i < 8; i++ {

		num := utils.IntAt(vat, i)

		if i%2 == 0 {
			switch num {
			case 0:
				num = 1
				break
			case 1:
				num = 0
				break
			case 2:
				num = 5
				break
			case 3:
				num = 7
				break
			case 4:
				num = 9
				break
			default:
				num = num*2 + 3
			}
		}

		result += num
	}

	return result
}
