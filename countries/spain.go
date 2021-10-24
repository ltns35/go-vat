package countries

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ltns35/go-vat/countries/utils"
)

// Do not sort, it's the algorithm.
const algorithm = "TRWAGMYFPDXBNJZSQVHLCKE"

type spain struct {
	Country
}

var Spain = spain{
	Country: Country{
		Name: "Spain",
		Codes: []string{
			"ES",
			"ESP",
			"724",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
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
				"^(ES)([A-Z]\\d{8})$",
				"^(ES)([A-HN-SW]\\d{7}[A-J])$",
				"^(ES)([0-9YZ]\\d{7}[A-Z])$",
				"^(ES)([KLMX]\\d{7}[A-Z])$",
			},
			Additional: []string{
				"^[A-H|J|U|V]\\d{8}$",
				"^[A-H|N-S|W]\\d{7}[A-J]$",
				"^[0-9|Y|Z]\\d{7}[A-Z]$",
				"^[K|L|M|X]\\d{7}[A-Z]$",
			},
		},
	},
}

func (s spain) Calc(vat string) bool {

	if len(s.Rules.Additional) == 0 {
		return false
	}

	regex := regexp.MustCompile(s.Rules.Additional[0])

	// National juridical entities
	if regex.MatchString(vat) {
		return isNationalJuridicalEntities(vat, s.Rules.Multipliers["common"])
	}

	regex = regexp.MustCompile(s.Rules.Additional[1])

	// Juridical entities other than national ones
	if regex.MatchString(vat) {
		return isNonNationalJuridical(vat, s.Rules.Multipliers["common"])
	}

	regex = regexp.MustCompile(s.Rules.Additional[2])

	// Personal number (NIF) (starting with numeric of Y or Z)
	if regex.MatchString(vat) {
		return isPersonalYtoZ(vat)
	}

	regex = regexp.MustCompile(s.Rules.Additional[3])

	// Personal number (NIF) (starting with K, L, M, or X)
	if regex.MatchString(vat) {
		return isPersonalKtoX(vat)
	}

	return false
}

func (s spain) GetCountry() Country {
	return s.Country
}

func extractDigitAndMultiplyByCounter(vat string, multipliers []int, total float64) int {

	var temp float64 = 0
	result := total
	for i := 0; i < 7; i++ {

		digit := utils.IntAt(vat, i+1)
		temp = float64(digit * multipliers[i])

		if temp > 9 {
			result += math.Floor(temp/10) + math.Mod(temp, 10)
		} else {
			result += temp
		}
	}

	return int(result)
}

func isNationalJuridicalEntities(vat string, multipliers []int) bool {

	total := extractDigitAndMultiplyByCounter(vat, multipliers, 0)

	// Now calculate the check digit itself.
	total = 10 - (total % 10)
	if total == 10 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, len(vat)-1)

	return total == expect
}

func isNonNationalJuridical(vat string, multipliers []int) bool {

	total := extractDigitAndMultiplyByCounter(vat, multipliers, 0)

	// Now calculate the check digit itself.
	total = 10 - (total % 10)
	totalStr := string(rune(total + 64))

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.StringAt(vat, len(vat)-1)

	return totalStr == expect
}

func isPersonalYtoZ(vat string) bool {

	tempNumber := vat

	firstChar := utils.StringAt(vat, 0)
	if firstChar == "Y" {
		tempNumber = strings.ReplaceAll(vat, "Y", "1")
	}
	if firstChar == "Z" {
		tempNumber = strings.ReplaceAll(vat, "Z", "2")
	}

	// Y2676382R
	// 12676382R

	lastIndex := len(tempNumber) - 1

	strNum := tempNumber[:lastIndex]
	num, _ := strconv.Atoi(strNum)

	expect := utils.StringAt(algorithm, num%23)
	lastChar := utils.StringAt(tempNumber, lastIndex)

	return lastChar == expect
}

func isPersonalKtoX(vat string) bool {
	strNum := vat[1 : len(vat)-2]
	num, _ := strconv.Atoi(strNum)

	expect := utils.StringAt(algorithm, num%23)
	lastChar := utils.StringAt(vat, len(vat)-1)

	return lastChar == expect
}
