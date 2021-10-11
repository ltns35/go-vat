# Go VAT

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https:/github.com/ltns35/go-vat)
[![Maintenance](https://img.shields.io/badge/Maintained-yes-green.svg)](https:/github.com/ltns35/go-vat/graphs/commit-activity)

Check the validity of a VAT number without any HTTP request.

This go library is based on the original [jsVAT](https://github.com/se-panfilov/jsvat) for JS/TS.

The intention of this library is to offer to all gophers the ability to validate all the VAT numbers before they are
stored in a database.

## Features

- [x] Check the validity of VAT numbers.
- [ ] Display with format.
- [ ] Retrieve the list of available taxes for each country.
- [x] Extendable with custom countries/rules
- [x] Offline

## How to use

```go

// Check against all supported countries validators.
vatResult, err := vat.CheckVAT("ADE000000E")
if err != nil {
// Handle error
}

// Check ONLY against Andorra validator.
countries := []countries.Calculer{
andorra.VAT,
}
vatResult, err = vat.CheckVAT("ADE000000E", countries...)

// output: vatResult
//
//	{
//	   "value":"ADF000000F",
//	   "isValid":true,
//	   "isSupportedCountry":true,
//	   "country":{
//	      "name":"Andorra",
//	      "codes":[
//	         "AD",
//	         "AND",
//	         "020"
//	      ],
//	      "rules":{
//	         "multipliers":null,
//	         "typeFormats":null,
//	         "lookup":null,
//	         "check":"",
//	         "regex":[
//	            "^(AD)([fealecdgopuFEALECDGOPU]{1}\\d{6}[fealecdgopuFEALECDGOPU]{1})$"
//	         ],
//	         "additional":null
//	      }
//	   }
//	}
//

```

## Supported countries

- [x] Albania
- [x] Andorra
- [x] Argentina
- [x] Austria
- [x] Belarus
- [x] Belgium
- [ ] Bolivia
- [x] Brazil
- [x] Bulgaria
- [ ] Canada
- [x] Croatia
- [x] Cyprus
- [x] Czech Republic
- [x] Denmark
- [ ] Ecuador
- [x] Estonia
- [ ] Europe
- [x] Finland
- [x] France
- [x] Germany
- [x] Greece
- [ ] Guatemala
- [x] Hungary
- [ ] Indonesia
- [x] Ireland
- [x] Italy
- [ ] Kazakhstan
- [x] Latvia
- [x] Liechtenstein
- [x] Lithuania
- [x] Luxembourg
- [x] Malta
- [x] Netherlands
- [ ] New Zealand
- [ ] Nigeria
- [x] North Macedonia
- [x] Norway
- [x] Peru
- [x] Poland
- [x] Portugal
- [x] Romania
- [x] Russia
- [x] San Marino
- [x] Serbia
- [x] Slovakia
- [x] Slovenia
- [x] Spain
- [x] Sweden
- [x] Switzerland
- [x] Turkey
- [x] Ukraine
- [x] United Kingdom
- [ ] Uruguay

## Need an unsupported country

If you need a country is not yet supported by the library **open a new issue** or **create a _pull request_** to be
merged.

## Testing

The library is tested against the 2 latest stable major versions of Go.

- 1.17.x
- 1.16.x

All the validators has been tested individually to ensure the correct working, if you need more tests don't hesitate to
open a new issue with the values you want to be tested.

## LICENSE

The library is under the MIT license, so you can use it for free for commercial products.