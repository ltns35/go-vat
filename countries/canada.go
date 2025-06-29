package countries

import (
	"fmt"
	"regexp"
)

type canada struct {
	Country
}

var Canada = canada{
	Country: Country{
		Name: "Canada",
		Codes: []string{
			"CA",
			"CAN",
			"124",
		},
		Rules: CountryRules{
			Regex: []string{
				"^(CA)(\\d{7})$",
				"^(CA)(\\d{9})$",
				"^(CA)(\\d{9}R[C|M|P|T]\\d{4})$",
				"^(CA)(PST\\d{8})$",
				"^(CA)(\\d{10}TQ\\d{4})$",
			},
		},
	},
}

func (c canada) Calc(vat string) bool {
	vat = fmt.Sprintf("CA%s", vat)

	return c.isGST(vat) ||
		c.isBusinessName(vat) ||
		c.isPSTBritishColumbia(vat) ||
		c.isPSTManitoba(vat) ||
		c.isPSTSaskatchewan(vat) ||
		c.isQSTQuebec(vat)
}

func (c canada) GetCountry() Country {
	return c.Country
}

func (c canada) isBusinessName(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[1])

	return regex.MatchString(vat)
}

func (c canada) isGST(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[2])

	return regex.MatchString(vat)
}

func (c canada) isPSTBritishColumbia(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[3])

	return regex.MatchString(vat)
}

func (c canada) isPSTManitoba(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[0])

	return regex.MatchString(vat)
}

func (c canada) isPSTSaskatchewan(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[0])

	return regex.MatchString(vat)
}

func (c canada) isQSTQuebec(vat string) bool {
	regex := regexp.MustCompile(c.Rules.Regex[4])

	return regex.MatchString(vat)
}
