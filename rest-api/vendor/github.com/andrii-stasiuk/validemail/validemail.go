package validemail

import "regexp"

// Validator struct
type Validator struct {
	regex *regexp.Regexp
}

// New function creates new object of Validator
func New() *Validator {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return &Validator{regex: re}
}

// EMailValidator metod checks if email-address is correct
func (v Validator) EMailValidator(email string) bool {
	return v.regex.MatchString(email)
}
