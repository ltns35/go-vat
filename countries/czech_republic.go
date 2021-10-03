package countries

import (
	"math"
	"regexp"
	"strconv"
)

type czechRepublic struct {
	Country
}

var CzechRepublic = czechRepublic{
	Country: Country{
		Name: "Czech Republic",
		Codes: []string{
			"CZ",
			"CZE",
			"203",
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
			Lookup: []int{
				8,
				7,
				6,
				5,
				4,
				3,
				2,
				1,
				0,
				9,
				8,
			},
			Regex: []string{
				"^(CZ)(\\d{8,10})(\\d{3})?$",
			},
			Additional: []string{
				"^\\d{8}$",
				"^[0-5][0-9][0|1|5|6]\\d[0-3]\\d\\d{3}$",
				"^6\\d{8}$",
				"^\\d{2}[0-3|5-8]\\d[0-3]\\d\\d{4}$",
			},
		},
	},
}

func (c czechRepublic) Calc(vat string) bool {

	if len(c.Rules.Additional) == 0 {
		return false
	}

	return isLegalEntities(vat, c.Rules.Multipliers["common"], c.Rules.Additional[0]) ||
		isIndividualType2(vat, c.Rules.Multipliers["common"], c.Rules.Additional[2], c.Rules.Lookup...) ||
		isIndividualType3(vat, c.Rules.Additional[3]) ||
		isIndividualType1(vat, c.Rules.Additional[1])

}

func (c czechRepublic) GetCountry() Country {
	return c.Country
}

func isLegalEntities(vat string, multipliers []int, additional string) bool {

	regex := regexp.MustCompile(additional)

	if regex.MatchString(vat) {

		total := 0

		// Extract the next digit and multiply by the counter.
		for i := 0; i < len(multipliers); i++ {
			num, _ := getIntFromStringAt(vat, i)
			total += num * multipliers[i]
		}

		// Establish check digit.
		total = 11 - (total % 11)
		if total == 10 {
			total = 0
		} else if total == 11 {
			total = 1
		}

		// Compare it with the last character of the VAT number. If it's the same, then it's valid.
		expect, _ := getIntFromStringAt(vat, 7)

		return total == expect
	}

	return false
}

func isIndividualType1(vat string, additional string) bool {

	regex := regexp.MustCompile(additional)
	if regex.MatchString(vat) {

		numStr := vat[:2]
		num, _ := strconv.Atoi(numStr)

		return num <= 62
	}

	return false
}

func isIndividualType2(vat string, multipliers []int, additional string, lookup ...int) bool {

	regex := regexp.MustCompile(additional)

	if regex.MatchString(vat) {

		var total float64 = 0

		// Extract the next digit and multiply by the counter.
		for i := 0; i < len(multipliers); i++ {
			num, _ := getIntFromStringAt(vat, i+1)
			total += float64(num) * float64(multipliers[i])
		}

		// Establish check digit.
		var a float64

		if math.Mod(total, 11) == 0 {
			a = total + 11
		} else {
			a = math.Ceil(float64(total)/11) * 11
		}

		pointer := a - total - 1

		// Convert calculated check digit according to a lookup table
		if len(lookup) == 0 {
			return false
		}

		expect, _ := getIntFromStringAt(vat, 8)

		return lookup[int(pointer)] == expect
	}

	return false
}

func isIndividualType3(vat string, additional string) bool {

	regex := regexp.MustCompile(additional)

	if regex.MatchString(vat) {

		num1, _ := strconv.Atoi(vat[:2])
		num2, _ := strconv.Atoi(vat[2:4])
		num3, _ := strconv.Atoi(vat[4:6])
		num4, _ := strconv.Atoi(vat[6:8])
		num5, _ := strconv.Atoi(vat[8:])

		temp := num1 + num2 + num3 + num4 + num5

		numFull, _ := strconv.Atoi(vat)
		expect := numFull%11 == 0

		return !!(temp%11 == 0 && expect)
	}

	return false
}
