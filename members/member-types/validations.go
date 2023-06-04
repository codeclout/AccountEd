package membertypes

import (
	"net/mail"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	ErrorInvalidEmail = errors.New("invalid email")
	ErrorMemberName   = errors.New("invalid name")
	ErrorPayloadSize  = errors.New("invalid payload size")
	ErrorPinInvalid   = errors.New("invalid pin")
)

func ValidateUsernamePayloadSize(in []byte) error {
	if len(in) > 360 {
		return ErrorPayloadSize
	}

	return nil
}

func ValidateName(name *string) (atomic.Value, error) {
	var atom atomic.Value

	c := cases.Title(language.English, cases.NoLower)
	t := strings.TrimSpace(*name)
	p := `(?m)[\p{Sm}\p{Nd}\p{Sc}\p{Sk}\p{Sm}\p{So}\p{Pe}\p{Ps}&#@*\\\/\.]`

	if ok, e := regexp.MatchString(p, c.String(t)); e == nil && ok {
		atom.Store("")
		return atom, ErrorMemberName
	}

	atom.Store(c.String(*name))
	return atom, nil
}

func ValidateEmail(email *string) (atomic.Value, error) {
	var atom atomic.Value

	a, e := mail.ParseAddress(*email)
	if e != nil {
		atom.Store("")
		return atom, ErrorInvalidEmail
	}

	if ok, _ := regexp.MatchString(`(?m)\.[a-z]{2,24}`, *email); !ok {
		atom.Store("")
		return atom, ErrorInvalidEmail
	}

	atom.Store(a.Address)
	return atom, nil
}

func ValidatePin(pin *string) (*string, error) {
	if ok, e := regexp.MatchString(`(?m)^[0-9]{7,10}$`, *pin); ok && e == nil {
		x := "ok"
		return &x, nil
	}

	y := ""
	return &y, ErrorPinInvalid
}

func ValidatePrimaryMember(in *PrimaryMemberStartRegisterIn, wg *sync.WaitGroup) error {
	out := make(chan error, 1)

	wg.Add(1)

	go func() {
		defer wg.Done()

		email, e := ValidateEmail(in.Username)
		if e != nil {
			out <- ErrorInvalidEmail
			return
		}

		v := email.Load()
		x := v.(string)
		in.Username = &x

		out <- nil
	}()

	wg.Wait()
	return <-out
}

func ValidateParentGuardian(in *ParentGuardian, wg *sync.WaitGroup) error {
	defer wg.Done()
	out := make(chan error)

	go func() {
		names := []*string{in.Member.LegalFirstName, in.Member.LegalLastName}
		for i, n := range names {
			name, e := ValidateName(n)
			if e != nil {
				out <- errors.Wrap(e, "name validation failed")
			}

			v := name.Load()
			x := v.(string)
			names[i] = &x
		}
	}()

	return <-out
}

func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
		return true
	}

	return false
}
