package validator

import (
	"errors"
	"fmt"

	"github.com/ttacon/libphonenumber"
)

type LibPhoneNumberValidator struct {}

func NewLibPhoneNumberValidator() *LibPhoneNumberValidator {
	return &LibPhoneNumberValidator{}
}

func (v *LibPhoneNumberValidator) Validate(phoneNumber string) (string, error) {
	parsedNumber, err := libphonenumber.Parse(phoneNumber, "AU")
	if err != nil {
		return "", fmt.Errorf("error parsing phone number: %v", err)
	}
	isValid := libphonenumber.IsValidNumber(parsedNumber)
	if !isValid {
		return "", errors.New("the phone number is not valid")
	}

	isAustralianNumber := libphonenumber.GetRegionCodeForNumber(parsedNumber) == "AU"
	if !isAustralianNumber {
		return "", fmt.Errorf("the phone number is not from Australia: %v", err)
	}

	e164Format := libphonenumber.Format(parsedNumber, libphonenumber.E164)
	return e164Format, nil
}
