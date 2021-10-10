package sweden

import (
	"math"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/utils"
)

type sweden struct {
	countries.Country
}

var VAT = sweden{
	Country: countries.Country{
		Name: "VAT",
		Codes: []string{
			"SE",
			"SWE",
			"752",
		},
		Rules: countries.CountryRules{
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
		num := utils.IntAt(vat, i)
		R += math.Floor(float64(num/5)) + float64((num*2)%10)
	}

	// Calculate S where S = C2 + C4 + C6 + C8
	S := 0
	for j := 1; j < 9; j = j + 2 {
		num := utils.IntAt(vat, j)
		S += num
	}

	checkDigit := (10 - ((int(R) + S) % 10)) % 10

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(vat, 9)

	return checkDigit == expect
}

func (s sweden) GetCountry() countries.Country {
	return s.Country
}
