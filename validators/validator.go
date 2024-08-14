package validator

type Validator interface {
	Validate(phonenumber string) (string, error)
}

type PhoneNumberValidator struct {
	validator Validator
}

func NewPhoneNumberValidator (useLibPhoneNumber bool) *PhoneNumberValidator {
	var validator Validator
	if useLibPhoneNumber {
		validator = NewLibPhoneNumberValidator()
	} else {
		validator = NewRegexValidator()
	}
	return &PhoneNumberValidator{validator: validator,}
}

func (p *PhoneNumberValidator) Validate(phoneNumber string) (string, error) {
	if libValidator, ok := p.validator.(*LibPhoneNumberValidator); ok {
		return libValidator.Validate(phoneNumber)
	}
	return p.validator.Validate(phoneNumber)
}