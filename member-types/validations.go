package memberTypes

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var v *validator.Validate

func TrimString(s *string) *string {
	b := strings.Trim(*s, " ")

	return &b
}

func ValidateName(name *string) bool {
	c := cases.Title(language.English, cases.NoLower)
	t := TrimString(name)

	if ok, e := regexp.MatchString(`(?m)[\p{Sm}\p{Nd}\p{Sc}\p{Sk}\p{Sm}\p{So}\p{Pe}\p{Ps}&#@*\\\/\.]`, c.String(*t)); e == nil && ok {
		return false
	}

	return true
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

	//if ok := ValidateName(&p.FirstName); !ok {
	//	return false
	//}
	//
	//if ok := ValidateName(&p.LastName); !ok {
	//	return false
	//}

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

func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
		return true
	}

	return false
}
