package countries

import (
	"math"
	"strconv"

	"github.com/ltns35/go-vat/countries/utils"
)

type unitedKingdom struct {
	Country
}

var UnitedKingdom = unitedKingdom{
	Country: Country{
		Name: "United Kingdom",
		Codes: []string{
			"GB",
			"GBR",
			"826",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					8,
					7,
					6,
					5,
					4,
					3,
					2,
				},
			},
			Regex: []string{
				"^(GB)?(\\d{9})$",
				"^(GB)?(\\d{12})$",
				"^(GB)?(GD\\d{3})$",
				"^(GB)?(HA\\d{3})$",
			},
		},
	},
}

func (u unitedKingdom) Calc(vat string) bool {

	// Government departments
	beginStr := vat[:2]
	if beginStr == "GD" {
		return isGovernmentDepartment(vat)
	}

	// Health authorities
	if beginStr == "HA" {
		return isHealthAuthorities(vat)
	}

	// Standard and commercial numbers
	return isStandardOrCommercialNumber(vat, u.Rules.Multipliers["common"])
}

func (u unitedKingdom) GetCountry() Country {
	return u.Country
}

func isGovernmentDepartment(vat string) bool {
	const expect = 500
	num := utils.IntAt(vat, 2)

	return num < expect
}

func isHealthAuthorities(vat string) bool {
	const expect = 499
	num := utils.IntAt(vat, 2)

	return num > expect
}

func isStandardOrCommercialNumber(vat string, multipliers []int) bool {

	var total float64 = 0

	// 0 VAT numbers disallowed!
	zeroVAT, _ := strconv.Atoi(vat)
	if zeroVAT == 0 {
		return false
	}

	// Check range is OK for modulus 97 calculation
	noStr := vat[:7]
	no, _ := strconv.Atoi(noStr)

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 7; i++ {
		num := utils.IntAt(vat, i)
		total += float64(num * multipliers[i])
	}

	// Old numbers use a simple 97 modulus, but new numbers use an adaptation of that (less 55). Our
	// VAT number could use either system, so we check it against both.

	// Establish check digits by subtracting 97 from total until negative.
	checkDigit := total

	for checkDigit > 0 {
		checkDigit = checkDigit - 97
	}

	// Get the absolute value and compare it with the last two characters of the VAT number. If the
	// same, then it is a valid traditional check digit. However, even then the number must fit within
	// certain specified ranges.
	checkDigit = math.Abs(checkDigit)

	lastDigitsStr := vat[7:9]
	lastDigits, _ := strconv.Atoi(lastDigitsStr)

	if checkDigit == float64(lastDigits) && no < 9990001 && (no < 100000 || no > 999999) && (no < 9490001 || no > 9700000) {
		return true
	}

	// Now try the new method by subtracting 55 from the check digit if we can - else add 42

	if checkDigit >= 55 {
		checkDigit = checkDigit - 55
	} else {
		checkDigit = checkDigit + 42
	}

	return checkDigit == float64(lastDigits) && no > 1000000
}
