package validations

import (
	"encoding/base64"
	"net/mail"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	membertypes "github.com/codeclout/AccountEd/members/member-types"
)

var (
	ErrorInvalidEmail = errors.New("invalid email")
	ErrorMemberName   = errors.New("invalid name")
	ErrorPayloadSize  = errors.New("invalid payload size")
	ErrorPinInvalid   = errors.New("invalid pin")
)

func ValidateBase64(in string) error {
	_, e := base64.StdEncoding.DecodeString(in)

	return e
}

// ValidateUsernamePayloadSize checks if the size of the input byte slice is within the acceptable limit.
// It returns an error if the input byte slice is larger than 360 bytes.
func ValidateUsernamePayloadSize(in []byte) error {
	if len(in) > 320 {
		return ErrorPayloadSize
	}

	return nil
}

// ValidateName checks if the given name is valid by applying a transformation to Title case, removing trailing white spaces,
// and verifying that the name does not contain prohibited characters. It returns an atomic.Value storing the transformed name and any
// related errors, such as ErrorMemberName for prohibited characters. It accepts a string pointer as input.
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

// ValidateEmail checks if the provided email is valid. It accepts a string pointer to the email address as input.
// The function returns an atomic.Value containing the parsed email address and an error which can be ErrorInvalidEmail.
// The email address validation is performed using the net/mail.ParseAddress function and a regex check for the email domain's TLD.
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

// ValidatePin checks if the provided PIN is valid by matching it against a regular expression pattern.
// It accepts a string pointer to the PIN as input and returns a pointer to the result and any related error.
// If the validation fails, an error is returned as ErrorPinInvalid. Valid PINs should contain 7 to 10 digits.
func ValidatePin(pin *string) (*string, error) {
	if ok, e := regexp.MatchString(`(?m)^[0-9]{7,10}$`, *pin); ok && e == nil {
		x := "ok"
		return &x, nil
	}

	y := ""
	return &y, ErrorPinInvalid
}

// ValidatePrimaryMember
/*
 * ValidatePrimaryMember checks a primary member's email.
 * It only prepares the member for registration if the email is valid.

 * @param in pointer to PrimaryMemberStartRegisterIn, which contains the email to be validated
 * @param wg pointer to sync.WaitGroup used to wait until the email validation is done
 * @return error which can be either nil or ErrorInvalidEmail
 */
func ValidatePrimaryMember(in *membertypes.PrimaryMemberStartRegisterIn, wg *sync.WaitGroup) error {
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

//
// func ValidateParentGuardian(in *member_types.ParentGuardian, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	out := make(chan error)
//
// 	go func() {
// 		names := []*string{in.Member.LegalFirstName, in.Member.LegalLastName}
// 		for i, n := range names {
// 			name, e := ValidateName(n)
// 			if e != nil {
// 				out <- errors.Wrap(e, "name validation failed")
// 			}
//
// 			v := name.Load()
// 			x := v.(string)
// 			names[i] = &x
// 		}
// 	}()
//
// 	return <-out
// }

func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
		return true
	}

	return false
}
