package vat

import (
	"github.com/ltns35/go-vat/countries"
)

type CheckResult struct {
	Value              string             `json:"value"`
	IsValid            bool               `json:"isValid"`
	IsSupportedCountry bool               `json:"isSupportedCountry"`
	Country            *countries.Country `json:"country"`
}
