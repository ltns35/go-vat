# Go VAT

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https:/github.com/ltns35/go-vat)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https:/github.com/ltns35/go-vat/graphs/commit-activity)

Check the validity of a VAT number without any HTTP request.

This go library is based on the original [jsVAT](https://github.com/se-panfilov/jsvat) for JS/TS.

## Features

- [x] Check the validity of VAT numbers.
- [ ] Retrieve the list of available taxes for each country.
- [x] Extendable with custom countries/rules
- [x] Offline

## How to use

```go

vatResult, err := vat.CheckVAT("ADE000000E")
if err != nil {
// Handle error
}

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
- [x] Austria
- [x] Belgium
- [x] Brazil
- [x] Bulgaria
- [x] Croatia
- [x] Cyprus
- [x] Czech Republic
- [x] Denmark
- [x] Estonia
- [ ] Europe
- [x] Finland
- [x] France
- [x] Germany
- [x] Greece
- [x] Hungary
- [x] Ireland
- [x] Italy
- [x] Latvia
- [x] Lithuania
- [x] Luxembourg
- [x] Malta
- [x] Netherlands
- [x] North Macedonia
- [x] Norway
- [x] Poland
- [x] Portugal
- [x] Romania
- [x] Russia
- [x] Serbia
- [x] Slovakia
- [x] Slovenia
- [x] Spain
- [x] Sweden
- [x] Switzerland
- [x] United Kingdom

## LICENSE

The library is under the MIT license, so you can use it for free for commercial products.