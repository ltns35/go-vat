package countries

import (
	"math"
)

type sweden struct {
	Country
}

var Sweden = sweden{
	Country: Country{
		Name: "Sweden",
		Codes: []string{
			"SE",
			"SWE",
			"752",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(SE)(\\d{10}01)$",
			},
		},
	},
}

func (s sweden) Calc(vat string) bool {

	// Calculate R where R = R1 + R3 + R5 + R7 + R9, and Ri = INT(Ci/5) + (Ci*2) modulo 10
	var R float64 = 0
	for i := 0; i < 9; i = i + 2 {
		num, _ := getIntFromStringAt(vat, i)
		R += math.Floor(float64(num/5)) + float64((num*2)%10)
	}

	// Calculate S where S = C2 + C4 + C6 + C8
	S := 0
	for j := 1; j < 9; j = j + 2 {
		num, _ := getIntFromStringAt(vat, j)
		S += num
	}

	checkDigit := (10 - ((int(R) + S) % 10)) % 10

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect, _ := getIntFromStringAt(vat, 9)

	return checkDigit == expect
}

func (s sweden) GetCountry() Country {
	return s.Country
}
