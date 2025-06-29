package countries

import (
	"regexp"

	"github.com/ltns35/go-vat/utils"
)

type lithuania struct {
	Country
}

var Lithuania = lithuania{
	Country: Country{
		Name: "Lithuania",
		Codes: []string{
			"LT",
			"LTU",
			"440",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"short": {
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					1,
				},
				"med": {
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					1,
					2,
				},
				"alt": {
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					1,
					2,
					3,
					4,
				},
			},
			Check: "^\\d{10}1",
			Regex: []string{
				"^(LT)(\\d{9}|\\d{12})$",
			},
		},
	},
}

func (l lithuania) Calc(vat string) bool {
	return check9DigitVat(vat, l.Rules) || check12DigitVat(vat, l.Rules)
}

func (l lithuania) GetCountry() Country {
	return l.Country
}

func extractDigit1(vat string, multiplierList []int, key int) int {
	num := utils.IntAt(vat, key)

	return num * multiplierList[key]
}

func doubleCheckCalculation(vat string, total int, rules CountryRules) int {
	result := total
	if result%11 == 10 {
		result = 0
		for i := range 8 {
			result += extractDigit1(vat, rules.Multipliers["short"], i)
		}
	}

	return result
}

func extractDigit2(vat string, total int) int {
	result := total

	for i := range 8 {
		num := utils.IntAt(vat, i)
		result += num * (i + 1)
	}

	return result
}

func checkDigitLithuania(total int) int {
	result := total % 11
	if result == 10 {
		result = 0
	}

	return result
}

func check9DigitVat(vat string, rules CountryRules) bool {
	// 9 character VAT numbers are for legal persons
	total := 0

	if len(vat) == 9 {
		// 8th character must be one
		regex := regexp.MustCompile("^\\d{7}1")
		if !regex.MatchString(vat) {
			return false
		}

		// Extract the next digit and multiply by the counter+1.
		total = extractDigit2(vat, total)

		// Can have a double check digit calculation!
		total = doubleCheckCalculation(vat, total, rules)

		// Establish check digit.
		total = checkDigitLithuania(total)

		// Compare it with the last character of the VAT number. If it's the same, then it's valid.
		expect := utils.IntAt(vat, 8)

		return total == expect
	}

	return false
}

func extractDigit12(vat string, total int, rules CountryRules) int {
	result := total

	for k := range 11 {
		result += extractDigit1(vat, rules.Multipliers["med"], k)
	}

	return result
}

func doubleCheckCalculation12(vat string, total int, rules CountryRules) int {
	result := total

	if total%11 == 10 {
		result = 0

		for l := range 11 {
			result += extractDigit1(vat, rules.Multipliers["alt"], l)
		}
	}

	return result
}

func check12DigitVat(vat string, rules CountryRules) bool {
	total := 0

	// 12 character VAT numbers are for temporarily registered taxpayers
	if len(vat) == 12 {
		if rules.Check == "" {
			return false
		}

		// 11th character must be one
		regex := regexp.MustCompile(vat)

		if !regex.MatchString(vat) {
			return false
		}

		// Extract the next digit and multiply by the counter+1.
		total = extractDigit12(vat, total, rules)

		// Can have a double check digit calculation!
		total = doubleCheckCalculation12(vat, total, rules)

		// Establish check digit.
		total = checkDigitLithuania(total)

		// Compare it with the last character of the VAT number. If it's the same, then it's valid.
		expect := utils.IntAt(vat, 11)

		return total == expect
	}

	return false
}
