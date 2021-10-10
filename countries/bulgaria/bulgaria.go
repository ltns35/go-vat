package bulgaria

import (
	"math"
	"regexp"
	"strconv"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type bulgaria struct {
	countries.Country
}

var VAT = bulgaria{
	Country: countries.Country{
		Name: "Bulgaria",
		Codes: []string{
			"BG",
			"BGR",
			"100",
		},
		Rules: countries.CountryRules{
			Multipliers: map[string][]int{
				"physical": {
					2,
					4,
					8,
					5,
					10,
					9,
					7,
					3,
					6,
				},
				"foreigner": {
					21,
					19,
					17,
					13,
					11,
					9,
					7,
					3,
					1,
				},
				"miscellaneous": {
					4,
					3,
					2,
					7,
					6,
					5,
					4,
					3,
					2,
				},
			},
			Regex: []string{
				"^(BG)(\\d{9,10})$",
			},
		},
	},
}

func (b bulgaria) Calc(vat string) bool {

	if len(vat) == 9 {
		return checkNineLengthVat(vat)
	}

	return isPhysicalPerson(vat, b.Rules.Multipliers) ||
		isForeigner(vat, b.Rules.Multipliers) ||
		miscellaneousVAT(vat, b.Rules.Multipliers)
}

func (b bulgaria) GetCountry() *countries.Country {
	return &b.Country
}

func increase(value int, vat string, from int, to int, incr int) int {

	result := value

	for i := from; i < to; i++ {

		num := utils.IntAt(vat, i)
		result += num * (i + incr)
	}

	return result
}

func increase2(value int, vat string, from int, to int, multipliers []int) int {

	result := value

	for i := from; i < to; i++ {
		num := utils.IntAt(vat, i)
		result += num * multipliers[i]
	}

	return result
}

func checkNineLengthVat(vat string) bool {

	var total float64 = 0
	temp := increase(0, vat, 0, 8, 1)

	expect := utils.IntAt(vat, 8)

	total = math.Mod(float64(temp), 11)
	if total != 10 {
		return total == float64(expect)
	}

	temp = increase(0, vat, 0, 8, 3)

	total = math.Mod(float64(temp), 11)
	if total == 10 {
		total = 0
	}

	return total == float64(expect)
}

func isPhysicalPerson(vat string, multipliers map[string][]int) bool {
	// 10 digit VAT code - see if it relates to a standard physical person
	regex := regexp.MustCompile("^\\d\\d[0-5]\\d[0-3]\\d\\d{4}$")

	if regex.MatchString(vat) {
		// Check month
		monthStr := vat[2:4]
		month, _ := strconv.Atoi(monthStr)

		if (month > 0 && month < 13) || (month > 20 && month < 33) || (month > 40 && month < 53) {

			total := increase2(0, vat, 0, 9, multipliers["physical"])

			// Establish check digit.
			total = total % 11
			if total == 10 {
				total = 0
			}

			// Check to see if the check digit given is correct, If not, try next type of person
			ninthDigit := utils.IntAt(vat, 9)
			if total == ninthDigit {
				return true
			}

		}
	}

	return false
}

func isForeigner(vat string, multipliers map[string][]int) bool {

	// Extract the next digit and multiply by the counter.
	total := increase2(0, vat, 0, 9, multipliers["foreigner"])

	// Check to see if the check digit given is correct, If not, try next type of person
	ninthDigit := utils.IntAt(vat, 9)

	return total%10 == ninthDigit
}

func miscellaneousVAT(vat string, multipliers map[string][]int) bool {

	// Finally, if not yet identified, see if it conforms to a miscellaneous VAT number
	total := increase2(0, vat, 0, 9, multipliers["miscellaneous"])

	// Establish check digit.
	total = 11 - (total % 11)
	if total == 10 {
		return false
	}

	if total == 11 {
		total = 0
	}

	// Check to see if the check digit given is correct, If not, we have an error with the VAT number
	ninthDigit := utils.IntAt(vat, 9)

	return total == ninthDigit
}
