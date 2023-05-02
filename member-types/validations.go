package memberTypes

import (
	"net/mail"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TrimString(s *string) *string {
	b := strings.Trim(*s, " ")

	return &b
}

func ValidateName(name *string) (*string, bool) {
	c := cases.Title(language.English, cases.NoLower)
	t := TrimString(name)
	p := `(?m)[\p{Sm}\p{Nd}\p{Sc}\p{Sk}\p{Sm}\p{So}\p{Pe}\p{Ps}&#@*\\\/\.]`

	if ok, e := regexp.MatchString(p, c.String(*t)); e == nil && ok {
		x := "invalid"
		return &x, false
	}

	n := c.String(*name)
	return &n, true
}

func ValidateEmail(email *string) (*string, bool) {
	x := "invalid"

	a, e := mail.ParseAddress(*email)
	if e != nil {
		return &x, false
	}

	if ok, _ := regexp.MatchString(`(?m)\.[a-z]{2,24}`, *email); !ok {
		return &x, false
	}

	return &a.Address, true
}

func ValidatePin(pin *string) (*string, bool) {
	if ok, e := regexp.MatchString(`(?m)^[0-9]{7,10}$`, *pin); ok && e == nil {
		x := "ok"
		return &x, true
	}

	y := "invalid"
	return &y, false
}

func ValidateHomeschoolRegistration(p *ParentGuardian) bool {

}

func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
		return true
	}

	return false
}
