package countries

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ltns35/go-vat/countries/utils"
)

type netherlands struct {
	Country
}

var Netherlands = netherlands{
	Country: Country{
		Name: "Netherlands",
		Codes: []string{
			"NL",
			"NLD",
			"528",
		},
		Rules: CountryRules{
			Multipliers: map[string][]int{
				"common": {
					9,
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
				"^(NL)(\\d{9}B\\d{2})$",
			},
			Additional: []string{
				"^(\\d{9})B\\d{2}$",
			},
		},
	},
}

func (n netherlands) Calc(vat string) bool {

	regex := regexp.MustCompile(`[\ \-\_]`)

	vat = regex.ReplaceAllString(vat, "")
	vat = strings.ToUpper(vat)

	if len(n.Rules.Additional) == 0 {
		return false
	}

	regex = regexp.MustCompile(n.Rules.Additional[0])

	if !regex.MatchString(vat) {
		return false
	}

	numb := vat

	vatFull := fmt.Sprintf("NL%s", vat)
	chars := strings.Split(vatFull, "")

	characterValues := make([]int, len(chars))
	for i, str := range chars {
		char := []rune(str)
		characterValues[i] = getCharValue(char[0])
	}

	total := 0

	// Extract the next digit and multiply by the counter.
	for i := 0; i < 8; i++ {
		num := utils.IntAt(numb, i)
		total += num * n.Rules.Multipliers["common"][i]
	}

	// Establish check digits by getting modulus 11.
	total = total % 11
	if total > 9 {
		total = 0
	}

	// Compare it with the last character of the VAT number. If it's the same, then it's valid.
	expect := utils.IntAt(numb, 8)

	concat := ""
	for _, value := range characterValues {
		concat += fmt.Sprintf("%d", value)
	}

	// is either 11 proof or 97 mod proof.
	return total == expect || isNinetySevenMod(concat)
}

func (n netherlands) GetCountry() Country {
	return n.Country
}

func getCharValue(char rune) int {

	// if one of these set values
	if char == '+' {
		return 36
	}

	if char == '*' {
		return 37
	}

	// if A...Z return code VAL -55
	code := char - 55
	if code > 9 && code < 91 {
		return int(code)
	}

	num := utils.IntAt(string(char), 0)

	return num
}

func isNinetySevenMod(value string) bool {

	remainder := mod(value, 97)

	return remainder == 1
}

// custom module function, to check module on values above Number limit
func mod(value string, divisor int) int {

	// Initialize result
	res := 0

	chars := strings.Split(value, "")

	for _, char := range chars {
		num := utils.IntAt(char, 0)
		res = (res*10 + num) % divisor
	}

	return res
}
