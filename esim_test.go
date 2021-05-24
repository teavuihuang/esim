package esim

import (
	"testing"
)

/*
Examples of valid EIDs are:
8900 1012 0123 4123 4012 3456 7890 1224
8900 1567 01020304 0506 0708 0910 1152
8904 4011 1122 3344 1122 3344 1122 3321
*/

func TestEidOk1(t *testing.T) {
	_, err := DecodeAndVerifyEid("89001012012341234012345678901224")
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidOk2(t *testing.T) {
	_, err := DecodeAndVerifyEid("89001567010203040506070809101152")
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidOk3(t *testing.T) {
	_, err := DecodeAndVerifyEid("89044011112233441122334411223321")
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidNonNumeric(t *testing.T) {
	_, err := DecodeAndVerifyEid("A9033023426100000000000859956802")
	if err == nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidTooShort(t *testing.T) {
	_, err := DecodeAndVerifyEid("9033023426100000000000859956802")
	if err == nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidTooLong(t *testing.T) {
	_, err := DecodeAndVerifyEid("789033023426100000000000859956802")
	if err == nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestEidNotItuScheme(t *testing.T) {
	_, err := DecodeAndVerifyEid("72001012012341234012345678901224")
	if err == nil {
		t.Fatalf(`Error: %v`, err)
	}
}

func TestShowEidData(t *testing.T) {
	dummy_data := EidData{}
	ShowEidData(dummy_data)
}
