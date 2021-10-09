package vat

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ltns35/go-vat/countries"
)

var countriesVATDoesNotStartWithCountryCode = []string{
	countries.Brazil.Name,
}

var allCountries = []countries.Calculer{
	countries.Andorra,
	countries.Austria,
	countries.Belgium,
	countries.Brazil,
	countries.Bulgaria,
	countries.Croatia,
	countries.Cyprus,
	countries.CzechRepublic,
	countries.Denmark,
	countries.Estonia,
	countries.Finland,
	countries.France,
	countries.Germany,
	countries.Greece,
	countries.Hungary,
	countries.Italy,
	countries.Ireland,
	countries.Latvia,
	countries.Lithuania,
	countries.Luxembourg,
	countries.Malta,
	countries.Norway,
	countries.Portugal,
	countries.Serbia,
	countries.Slovakia,
	countries.Slovenia,
	countries.Spain,
	countries.Sweden,
}

func removeExtraChars(vat string) string {

	regex := regexp.MustCompile(`(\s|-|\.|\/)+`)

	vat = strings.ToUpper(vat)
	vat = regex.ReplaceAllString(vat, "")

	return vat
}

func getCountryCodes(country *countries.Country) []string {

	codes := country.Codes

	if country.Name == countries.Greece.Name {
		codes = append(codes, "EL")
	}

	return codes
}

func getCountry(vat string, countriesList []countries.Calculer) (countries.Calculer, error) {

	for _, c := range countriesList {

		country := c.GetCountry()
		if startsWithCode(vat, &country) || (!isVATStartWithCountryCode(country.Name) && isVATStartWithNumber(vat)) {
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

func CheckVAT(vat string, countriesList ...countries.Calculer) (bool, error) {

	if strings.Trim(vat, " ") == "" {
		return false, nil
	}

	cleanVAT := removeExtraChars(vat)

	if len(countriesList) == 0 {
		countriesList = allCountries
	}

	country, err := getCountry(cleanVAT, countriesList)
	if err != nil {
		return false, err
	}

	return isVatValid(cleanVAT, country), nil
}
