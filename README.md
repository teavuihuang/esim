# Esim


This Go eSIM module processes eUICC IDentifier (EID) used in the context of Remote Provisioning
and Management of the eUICC (eSIM) in according to GSM Association Official Document
SGP.02 (Remote Provisioning of Embedded UICC Technical Specification) and SGP.22
(RSP Technical Specification) for EID using the ITU-T E.118 (ITU-T Recommendation E.118,
the international telecommunication charge card) based scheme. See [GSMA eSIM specifications](https://www.gsma.com/esim/esim-specification/) for more info. Mobile devices that support eSIM include Apple iPhone 12 Pro Max, Samsung Galaxy S20 Ultra, Huawei P40 Pro, Google Pixel 4 XL etc.


## Decoded EID Data
~~~ go
type EidData struct {
	eidIndustryIdentifier             string
	eidCountryCode                    string
	eidIssuerIdentifier               string
	eidPlatformAndOsVersions          string
	eidAdditionalIssuerInfo           string
	eidIndividualIdentificationNumber string
	eidCheckDigits                    string
	eidVerificationSuccessful         bool
}
~~~


## Usage
~~~ go
// main.go
package main

import (
	"fmt"

	"github.com/teavuihuang/esim"
)

func main() {
	// A slice of eSIM EIDs
	eids := []string{
		"89001012012341234012345678901224", "89001567010203040506070809101152",
		"89044011112233441122334411223321", "89044011112233441122334411223373",
		"A9033023426100000000000859956802", "9033023426100000000000859956802",
		"789033023426100000000000859956802", "72001012012341234012345678901224",
	}

	// Decode and verify the EIDs
	// Show the decoded data if succesful, else show the error message
	for _, eid := range eids {
		eidData, err := esim.DecodeAndVerifyEid(eid)
		fmt.Println("\n--- eid", eid)
		if err == nil {
			esim.ShowEidData(eidData)
		} else {
			fmt.Println("err", err)
		}
	}

}
~~~


## Sample Output
~~~
--- eid 89001012012341234012345678901224
eidIndustryIdentifier               89
eidCountryCode                      001
eidIssuerIdentifier                 012
eidPlatformAndOsVersions            01234       
eidAdditionalIssuerInfo             12340       
eidIndividualIdentificationNumber   123456789012
eidCheckDigits                      24
eidVerificationSuccessful           true

--- eid 89001567010203040506070809101152
eidIndustryIdentifier               89
eidCountryCode                      001
eidIssuerIdentifier                 567
eidPlatformAndOsVersions            01020
eidAdditionalIssuerInfo             30405
eidIndividualIdentificationNumber   060708091011
eidCheckDigits                      52
eidVerificationSuccessful           true

--- eid 89044011112233441122334411223321
eidIndustryIdentifier               89
eidCountryCode                      044
eidIssuerIdentifier                 011
eidPlatformAndOsVersions            11223
eidAdditionalIssuerInfo             34411
eidIndividualIdentificationNumber   223344112233
eidCheckDigits                      21
eidVerificationSuccessful           true

--- eid 89044011112233441122334411223373
eidIndustryIdentifier               89
eidCountryCode                      044
eidIssuerIdentifier                 011
eidPlatformAndOsVersions            11223
eidAdditionalIssuerInfo             34411
eidIndividualIdentificationNumber   223344112233
eidCheckDigits                      73
eidVerificationSuccessful           false

--- eid A9033023426100000000000859956802
err EID is not numeric

--- eid 9033023426100000000000859956802
err EID is not 32 characters

--- eid 789033023426100000000000859956802
err EID is not 32 characters

--- eid 72001012012341234012345678901224
err EID is not using the ITU-T E.118 based scheme
~~~
