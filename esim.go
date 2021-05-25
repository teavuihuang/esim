package esim

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

/*
Internal function to check if a string is numeric
*/
func isNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

/*
The struct containing the output from DecodeAndVerifyEid()
*/
type EidData struct {
	EidIndustryIdentifier             string
	EidCountryCode                    string
	EidIssuerIdentifier               string
	EidPlatformAndOsVersions          string
	EidAdditionalIssuerInfo           string
	EidIndividualIdentificationNumber string
	EidCheckDigits                    string
	EidVerificationSuccessful         bool
}

/*
Decode and verify the eUICC IDentifier (EID) used in the context of Remote Provisioning
and Management of the eUICC (eSIM) in according to GSM Association Official Document
SGP.02 (Remote Provisioning of Embedded UICC Technical Specification) and SGP.22
(RSP Technical Specification) for EID using the ITU-T E.118 (ITU-T Recommendation E.118,
the international telecommunication charge card) based scheme.
*/
func DecodeAndVerifyEid(eid string) (EidData, error) {

	if !isNumeric(eid) {
		return EidData{}, errors.New("EID is not numeric")
	}
	if len(eid) != 32 {
		return EidData{}, errors.New("EID is not 32 characters")
	}
	if eid[:2] != "89" {
		return EidData{}, errors.New("EID is not using the ITU-T E.118 based scheme")
	}

	/*
		The 2 check digits are calculated as follows:
		1. Replace the 2 check digits by 2 digits of 0
		2. Using the resulting 32 digits as a decimal integer
		3. Compute the remainder of that number on division by 97
		4. Subtract the remainder from 98, and use the decimal result for the 2 check digits
	*/
	eid_cd1, _ := strconv.ParseInt(eid[30:32], 10, 64)
	big00, _ := new(big.Int).SetString(eid[:30]+"00", 0)
	big97, _ := new(big.Int).SetString("97", 0)
	bigMm, _ := new(big.Int).SetString("0", 0)
	big00.DivMod(big00, big97, bigMm)
	eid_cd2 := 98 - bigMm.Uint64()

	eid_decoded := EidData{
		EidIndustryIdentifier:             eid[:2],
		EidCountryCode:                    eid[2:5],
		EidIssuerIdentifier:               eid[5:8],
		EidPlatformAndOsVersions:          eid[8:13],
		EidAdditionalIssuerInfo:           eid[13:18],
		EidIndividualIdentificationNumber: eid[18:30],
		EidCheckDigits:                    eid[30:32],
		EidVerificationSuccessful:         uint64(eid_cd1) == eid_cd2,
	}
	return eid_decoded, nil
}

/*
Display the EidData on the console
*/
func ShowEidData(eidData EidData) {
	fmt.Println("EidIndustryIdentifier              ", eidData.EidIndustryIdentifier)
	fmt.Println("EidCountryCode                     ", eidData.EidCountryCode)
	fmt.Println("EidIssuerIdentifier                ", eidData.EidIssuerIdentifier)
	fmt.Println("EidPlatformAndOsVersions           ", eidData.EidPlatformAndOsVersions)
	fmt.Println("EidAdditionalIssuerInfo            ", eidData.EidAdditionalIssuerInfo)
	fmt.Println("EidIndividualIdentificationNumber  ", eidData.EidIndividualIdentificationNumber)
	fmt.Println("EidCheckDigits                     ", eidData.EidCheckDigits)
	fmt.Println("EidVerificationSuccessful          ", eidData.EidVerificationSuccessful)
}
