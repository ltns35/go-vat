package countries

import (
	"regexp"

	"github.com/ltns35/go-vat/utils"
)

type ireland struct {
	Country
}

var Ireland = ireland{
	Country: Country{
		Name: "Ireland",
		Codes: []string{
			"IE",
			"IRL",
			"372",
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
			TypeFormats: map[string]string{
				"first": "^\\d[A-Z*+]",
				"third": "^\\d{7}[A-Z][AH]$",
			},
			Regex: []string{
				"^(IE)(\\d{7}[A-W])$",
				"^(IE)([7-9][A-Z*+)]\\d{5}[A-W])$",
				"^(IE)(\\d{7}[A-W][AH])$",
			},
		},
	},
}

func (i ireland) Calc(vat string) bool {
	if _, ok := i.Rules.TypeFormats["first"]; !ok {
		return false
	} else if _, ok = i.Rules.TypeFormats["third"]; !ok {
		return false
	}

	newVat := vat
	// If the code is type 1 format, we need to convert it to the new before performing the validation.
	regex := regexp.MustCompile(i.Rules.TypeFormats["first"])

	if regex.MatchString(vat) {
		firstStr := vat[:1]
		middleStr := vat[2:7]
		lastStr := vat[7:]

		newVat = "0" + middleStr + firstStr + lastStr
	}

	total := 0

	for j := range 7 {
		// Extract the next digit and multiply by the counter.
		num := utils.IntAt(newVat, j)
		total += num * i.Rules.Multipliers["common"][j]
	}

	// If the number is type 3 then we need to include the trailing A or H in the calculation
	regex = regexp.MustCompile(i.Rules.TypeFormats["third"])

	if regex.MatchString(newVat) {
		// Add in a multiplier for the character A (1*9=9) or H (8*9=72)
		isH := utils.StringAt(newVat, 8)

		if isH == "H" {
			total += 72
		} else {
			total += 9
		}
	}

	// Establish check digit using modulus 23, and translate to char. equivalent.
	total %= 23

	totalStr := string(rune(total + 64))

	if total == 0 {
		totalStr = "W"
	}

	// Compare it with the eighth character of the VAT number. If it's the same, then it's valid.
	expect := utils.StringAt(newVat, 7)

	return totalStr == expect
}

func (i ireland) GetCountry() Country {
	return i.Country
}
