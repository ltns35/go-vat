package brazil

import (
	"strconv"
	"strings"

	"github.com/ltns35/go-vat/countries"
)

type brazil struct {
	countries.Country
}

var VAT = brazil{
	Country: countries.Country{
		Name: "Brazil",
		Codes: []string{
			"BR",
			"BRA",
			"076",
		},
		Rules: countries.CountryRules{
			Regex: []string{
				"^(BR)?(\\d{14}|\\d{2}\\.\\d{3}\\.\\d{3}\\/\\d{4}-\\d{2})$",
			},
		},
	},
}

func (b brazil) Calc(vat string) bool {

	fields := strings.Split(vat, "")

	numbers := make([]int, len(fields))

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return false
		}

		numbers[i] = num
	}

	if isRepeatedArray(numbers) {
		return false
	}

	var validators = []int{
		6,
		5,
		4,
		3,
		2,
		9,
		8,
		7,
		6,
		5,
		4,
		3,
		2,
	}

	checkers := generateCheckSums(numbers, validators)

	return numbers[12] == getRemaining(checkers[0]) && numbers[13] == getRemaining(checkers[1])
}

func (b brazil) GetCountry() countries.Country {
	return b.Country
}

// Generate check sums. Multiply numbers to validators and sum them to generate
// check sums, they're used to checking if numbers are valid.
// @param numbers - Numbers used to generate checkers.
// @param validators - Validators used to generate checkers.
func generateCheckSums(numbers []int, validators []int) []int {

	var checker = []int{
		0,
		0,
	}

	for i, validator := range validators {

		if i == 0 {
			checker[0] = 0
		} else {
			checker[0] = checker[0] + numbers[i-1]*validator
		}

		checker[1] = checker[1] + numbers[i]*validator
	}

	return checker
}

func isRepeatedArray(numbers []int) bool {

	for _, num := range numbers {
		if numbers[0] != num {
			return false
		}
	}

	return true
}

func getRemaining(num int) int {

	if num%11 < 2 {
		return 0
	} else {
		return 11 - (num % 11)
	}

}
