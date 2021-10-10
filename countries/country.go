package countries

type Country struct {
	Name  string       `json:"name"`
	Codes []string     `json:"codes"`
	Rules CountryRules `json:"rules"`
}

func (c Country) Calc(vat string) bool {
	return false
}

func (c Country) GetCountry() *Country {
	return &c
}

type Calculer interface {
	Calc(vat string) bool
	GetCountry() *Country
}

type CountryRules struct {
	Multipliers map[string][]int  `json:"multipliers"`
	TypeFormats map[string]string `json:"typeFormats"`
	Lookup      []int             `json:"lookup"`
	Check       string            `json:"check"`
	Regex       []string          `json:"regex"`
	Additional  []string          `json:"additional"`
}
