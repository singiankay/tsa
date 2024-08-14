package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type RegexValidator struct {
	phoneRegex *regexp.Regexp
}

func NewRegexValidator() *RegexValidator {
	phoneRegex := regexp.MustCompile(`^(?:\+?(61))?(?:\d{1,4})\d{6,10}$`)
	return &RegexValidator{phoneRegex: phoneRegex}
	// phoneRegex := `^(?:\+?(61))? ?(?:\((?=.*\)))?(0?[2-57-8])\)? ?(\d\d(?:[- ](?=\d{3})|(?!\d\d[- ]?\d[- ]))\d\d[- ]?\d[- ]?\d{3})$`
	// phoneRegex := regexp.MustCompile(phoneRegex)
	// phoneRegex := regexp.MustCompile(`^(\+61|0)?(4\d{8}|[2378]\d{8})$`)
	// phoneRegex := regexp.MustCompile(`^(?:\+?61\s?)?(?:\(\d{2}\)\s?|\d{2}\s?)?\d{4}\s?\d{4}$|^(?:\+?61\s?)?04\d{2}\s?\d{3}\s?\d{3}$|^(?:\+?61\s?)?(13|1300|1800)\s?\d{3}\s?\d{3}$`)
	// phoneRegex := regexp.MustCompile(`^(\+61|0) ?(2|3|7|8|4[0-9]|5[0-9]|6[0-9]|9[0-9]) ?\d{4} ?\d{4}$`)
	// phoneRegex := regexp.MustCompile(`^(?:\(\d{2}\)\s?|\d{2}\s?)?\d{4}\s?\d{4}$|^04\d{2}\s?\d{3}\s?\d{3}$`)
	// phoneRegex := regexp.MustCompile(`^(?:\+?61\s?)?(?:\(\d{2}\)\s?|\d{2}\s?)?\d{4}\s?\d{4}$|^(?:\+?61\s?)?04\d{2}\s?\d{3}\s?\d{3}$`)
}

func (v *RegexValidator) Validate(phoneNumber string) (string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")

	if v.phoneRegex.MatchString(phoneNumber) {
		if !strings.HasPrefix(phoneNumber, "+61") {
			phoneNumber = fmt.Sprintf("+61%s", phoneNumber[1:])
		}
		return phoneNumber, nil
	}
	return "", errors.New("invalid phone number")
}