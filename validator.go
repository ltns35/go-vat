package vat

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/albania"
	"github.com/ltns35/go-vat/countries/andorra"
	"github.com/ltns35/go-vat/countries/argentina"
	"github.com/ltns35/go-vat/countries/austria"
	"github.com/ltns35/go-vat/countries/belarus"
	"github.com/ltns35/go-vat/countries/belgium"
	"github.com/ltns35/go-vat/countries/brazil"
	"github.com/ltns35/go-vat/countries/bulgaria"
	"github.com/ltns35/go-vat/countries/croatia"
	"github.com/ltns35/go-vat/countries/cyprus"
	"github.com/ltns35/go-vat/countries/czech_republic"
	"github.com/ltns35/go-vat/countries/denmark"
	"github.com/ltns35/go-vat/countries/estonia"
	"github.com/ltns35/go-vat/countries/finland"
	"github.com/ltns35/go-vat/countries/france"
	"github.com/ltns35/go-vat/countries/germany"
	"github.com/ltns35/go-vat/countries/greece"
	"github.com/ltns35/go-vat/countries/hungary"
	"github.com/ltns35/go-vat/countries/ireland"
	"github.com/ltns35/go-vat/countries/italy"
	"github.com/ltns35/go-vat/countries/latvia"
	"github.com/ltns35/go-vat/countries/liechtenstein"
	"github.com/ltns35/go-vat/countries/lithuania"
	"github.com/ltns35/go-vat/countries/luxembourg"
	"github.com/ltns35/go-vat/countries/malta"
	"github.com/ltns35/go-vat/countries/netherlands"
	"github.com/ltns35/go-vat/countries/north_macedonia"
	"github.com/ltns35/go-vat/countries/norway"
	"github.com/ltns35/go-vat/countries/peru"
	"github.com/ltns35/go-vat/countries/poland"
	"github.com/ltns35/go-vat/countries/portugal"
	"github.com/ltns35/go-vat/countries/romania"
	"github.com/ltns35/go-vat/countries/russia"
	"github.com/ltns35/go-vat/countries/san_marino"
	"github.com/ltns35/go-vat/countries/serbia"
	"github.com/ltns35/go-vat/countries/slovakia"
	"github.com/ltns35/go-vat/countries/slovenia"
	"github.com/ltns35/go-vat/countries/spain"
	"github.com/ltns35/go-vat/countries/sweden"
	"github.com/ltns35/go-vat/countries/switzerland"
	"github.com/ltns35/go-vat/countries/turkey"
	"github.com/ltns35/go-vat/countries/ukraine"
	"github.com/ltns35/go-vat/countries/united_kingdom"
)

var countriesVATDoesNotStartWithCountryCode = []string{
	brazil.VAT.Name,
}

var allCountries = []countries.Calculer{
	albania.VAT,
	andorra.VAT,
	argentina.VAT,
	austria.VAT,
	belarus.VAT,
	belgium.VAT,
	brazil.VAT,
	bulgaria.VAT,
	croatia.VAT,
	cyprus.VAT,
	czech_republic.VAT,
	denmark.VAT,
	estonia.VAT,
	finland.VAT,
	france.VAT,
	germany.VAT,
	greece.VAT,
	hungary.VAT,
	italy.VAT,
	ireland.VAT,
	latvia.VAT,
	liechtenstein.VAT,
	lithuania.VAT,
	luxembourg.VAT,
	malta.VAT,
	north_macedonia.VAT,
	norway.VAT,
	netherlands.VAT,
	peru.VAT,
	poland.VAT,
	portugal.VAT,
	romania.VAT,
	russia.VAT,
	san_marino.VAT,
	serbia.VAT,
	slovakia.VAT,
	slovenia.VAT,
	spain.VAT,
	sweden.VAT,
	switzerland.VAT,
	turkey.VAT,
	ukraine.VAT,
	united_kingdom.VAT,
}

func makeResult(vat string, isValid bool, country *countries.Country) *CheckResult {

	var checkResult CheckResult

	checkResult.Value = vat
	checkResult.IsValid = isValid
	checkResult.Country = country

	if country != nil {
		checkResult.IsSupportedCountry = true
	}

	return &checkResult
}

func removeExtraChars(vat string) string {

	regex := regexp.MustCompile(`(\s|-|\.|\/)+`)

	vat = strings.ToUpper(vat)
	vat = regex.ReplaceAllString(vat, "")

	return vat
}

func getCountryCodes(country *countries.Country) []string {

	codes := country.Codes

	if country.Name == greece.VAT.Name {
		codes = append(codes, "EL")
	} else if country.Name == liechtenstein.VAT.Name {
		codes = append(codes, "FL")
	}

	return codes
}

func getCountry(vat string, countriesList []countries.Calculer) (countries.Calculer, error) {

	for _, c := range countriesList {

		country := c.GetCountry()
		if startsWithCode(vat, country) || (!isVATStartWithCountryCode(country.Name) && isVATStartWithNumber(vat)) {
			return c, nil
		}
	}

	return nil, errors.New("cannot retrieve a supported country")
}

func startsWithCode(vat string, country *countries.Country) bool {
	countryCodes := getCountryCodes(country)

	for _, code := range countryCodes {
		if strings.HasPrefix(vat, code) {
			return true
		}
	}

	return false
}

func isVATStartWithCountryCode(countryName string) bool {

	for _, country := range countriesVATDoesNotStartWithCountryCode {
		if country == countryName {
			return false
		}
	}

	return true
}

func isVATStartWithNumber(vat string) bool {

	regex := regexp.MustCompile("^\\d{2}")

	return regex.MatchString(vat)
}

func isVatValidToRegexp(vat string, regexArr []string) (bool, *regexp.Regexp) {

	for _, regexStr := range regexArr {

		regex := regexp.MustCompile(regexStr)
		isValid := regex.MatchString(vat)

		if isValid {
			return true, regex
		}
	}

	return false, nil
}

func isVatValid(vat string, country countries.Calculer) bool {

	regexpValidRes, regex := isVatValidToRegexp(vat, country.GetCountry().Rules.Regex)

	if !regexpValidRes {
		return false
	}

	foundResults := regex.FindAllStringSubmatch(vat, -1)
	cleanVAT := foundResults[0][2]

	return country.Calc(cleanVAT)
}

func CheckVAT(vat string, countriesList ...countries.Calculer) (*CheckResult, error) {

	if strings.Trim(vat, " ") == "" {
		return makeResult(vat, false, nil), nil
	}

	cleanVAT := removeExtraChars(vat)

	if len(countriesList) == 0 {
		countriesList = allCountries
	}

	country, err := getCountry(cleanVAT, countriesList)
	if err != nil {
		return nil, err
	}

	isValid := isVatValid(cleanVAT, country)

	return makeResult(cleanVAT, isValid, country.GetCountry()), nil
}
