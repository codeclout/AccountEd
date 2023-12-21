package validations

import (
	"errors"
	"strings"
	"sync"
	"testing"

	member_types "github.com/codeclout/AccountEd/members/member-types"
)

func TestValidateUsernamePayloadSize(t *testing.T) {
	payloads := []string{
		"caexdmsrhedwybldobabtviszgntitqccbdojmccptzyyvrp",
		strings.Repeat("abcd", 80) + "b",
		"scwrkecuvbnyieytohbjchrrxkxfqvqcjzxbpidbjxuiityjdbbwseiqpdedkywfoujexznzpqdlaqzikeryfwgmqxbnknf",
	}

	cases := []struct {
		a []byte
		b error
	}{
		{a: []byte(payloads[0]), b: nil},
		{a: []byte(payloads[1]), b: ErrorPayloadSize},
		{a: []byte(payloads[2]), b: nil},
	}

	for _, v := range cases {
		e := ValidateUsernamePayloadSize(v.a)

		if !errors.Is(v.b, e) {
			t.Errorf("expected error: %s\n received error: %s", ErrorPayloadSize.Error(), e.Error())
		}
	}
}

func TestValidateName(t *testing.T) {
	names := []string{"brian", "<script>", "O'Brien", "lópez", "Gutiérrez", "de jong", "Adébáyọ̀", "()", "Trey$"}

	cases := []struct {
		a string
		b error
	}{
		{a: "Brian", b: nil},
		{a: "", b: ErrorMemberName},
		{a: "O'Brien", b: nil},
		{a: "López", b: nil},
		{a: "Gutiérrez", b: nil},
		{a: "De Jong", b: nil},
		{a: "Adébáyọ̀", b: nil},
		{a: "", b: ErrorMemberName},
		{a: "", b: ErrorMemberName},
	}

	for i, tc := range cases {
		c, ok := ValidateName(&names[i])
		name := c.Load()

		if name != tc.a || !errors.Is(tc.b, ok) {
			t.Errorf("expected name: %s %t \n received name: %s %t", name, ok, tc.a, tc.b)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	addresses := []string{"x@example.com", "x@example", "x.y@example.com", "x-y@example.com", "x_y@example.com", "Barry Gibbs <bg@example.com>"}

	cases := []struct {
		a string
		b error
	}{
		{a: "x@example.com", b: nil},
		{a: "", b: ErrorInvalidEmail},
		{a: "x.y@example.com", b: nil},
		{a: "x-y@example.com", b: nil},
		{a: "x_y@example.com", b: nil},
		{a: "bg@example.com", b: nil},
	}

	for i, address := range cases {
		c, ok := ValidateEmail(&addresses[i])
		email := c.Load()

		if email != address.a || !errors.Is(ok, address.b) {
			t.Errorf("expected address: %s ok: %t \n received address: %s ok: %t", email, ok, address.a, address.b)
		}
	}
}

func TestValidatePrimaryMember(t *testing.T) {
	wg := &sync.WaitGroup{}

	addresses := []string{"something.null", "x.y@example.com", "Barry Gibbs <bg@example.com>"}
	prime := []*member_types.PrimaryMemberStartRegisterIn{
		&member_types.PrimaryMemberStartRegisterIn{MemberID: &addresses[0]},
		&member_types.PrimaryMemberStartRegisterIn{MemberID: &addresses[1]},
		&member_types.PrimaryMemberStartRegisterIn{MemberID: &addresses[2]},
	}

	cases := []struct {
		a error
	}{
		{a: ErrorInvalidEmail},
		{a: nil},
		{a: nil},
	}

	for i, cas := range cases {
		e := ValidatePrimaryMember(prime[i], wg)

		if !errors.Is(e, cas.a) {
			t.Errorf("expected e: %t \n received e: %t", e, cas.a)
		}
	}
}

func TestValidatePin(t *testing.T) {
	pins := []string{"7472", "382023475", "301%132$", "472651", "4276291", "328247461693282474616932824746169"}

	cases := []struct {
		a error
	}{
		{a: ErrorPinInvalid},
		{a: nil},
		{a: ErrorPinInvalid},
		{a: nil},
		{a: nil},
		{a: ErrorPinInvalid},
	}

	for i, cas := range cases {
		e := ValidatePin(&pins[i])

		if !errors.Is(e, cas.a) {
			t.Errorf("expected e: %t \n received e: %t", e, cas.a)
		}
	}
}
