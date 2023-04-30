package internal

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func ValidateAccountTypeId(id AccountTypeIn) (bool, error) {

	if ok, e := regexp.MatchString(`(?m)^[a-zA-Z0-9]{24}$`, id.Id); ok && e == nil {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("validation failed for %T", id))
}

func ValidateHomeschoolRegistration(p *ParentGuardian) bool {
	v = validator.New()
	if e := v.Struct(p); e != nil {
		return false
	}

	if s, ok := ValidateEmail(&p.Username); !ok {
		return false
	} else {
		p.Username = s
	}

	if ok := ValidatePin(&p.Pin); !ok {
		return false
	}

	if ok := ValidateName(&p.FirstName); !ok {
		return false
	}

	if ok := ValidateName(&p.LastName); !ok {
		return false
	}

	return true
}

func ValidateEmail(email *string) (string, bool) {
	a, e := mail.ParseAddress(*email)
	if e != nil {
		return "", false
	}

	if ok, _ := regexp.MatchString(`(?m)\.[a-z]{2,3}`, *email); !ok {
		return "", false
	}

	return a.Address, true
}

func ValidatePin(pin *string) bool {
	if ok, e := regexp.MatchString(`(?m)\d+`, *pin); ok && e == nil {
		return true
	}

	return false
}

func ValidateName(name *string) bool {
	if ok, e := regexp.MatchString(`(?m)[a-zA-Z]+`, *name); ok && e == nil {
		return true
	}

	return false
}

func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
		return true
	}

	return false
}
