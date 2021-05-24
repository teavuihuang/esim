package esim

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

func isNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

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

/*
This eSIM module processes eUICC IDentifier (EID) used in the context of Remote Provisioning
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
		eidIndustryIdentifier:             eid[:2],
		eidCountryCode:                    eid[2:5],
		eidIssuerIdentifier:               eid[5:8],
		eidPlatformAndOsVersions:          eid[8:13],
		eidAdditionalIssuerInfo:           eid[13:18],
		eidIndividualIdentificationNumber: eid[18:30],
		eidCheckDigits:                    eid[30:32],
		eidVerificationSuccessful:         uint64(eid_cd1) == eid_cd2,
	}
	return eid_decoded, nil
}

func ShowEidData(eidData EidData) {
	fmt.Println("eidIndustryIdentifier              ", eidData.eidIndustryIdentifier)
	fmt.Println("eidCountryCode                     ", eidData.eidCountryCode)
	fmt.Println("eidIssuerIdentifier                ", eidData.eidIssuerIdentifier)
	fmt.Println("eidPlatformAndOsVersions           ", eidData.eidPlatformAndOsVersions)
	fmt.Println("eidAdditionalIssuerInfo            ", eidData.eidAdditionalIssuerInfo)
	fmt.Println("eidIndividualIdentificationNumber  ", eidData.eidIndividualIdentificationNumber)
	fmt.Println("eidCheckDigits                     ", eidData.eidCheckDigits)
	fmt.Println("eidVerificationSuccessful          ", eidData.eidVerificationSuccessful)
}
