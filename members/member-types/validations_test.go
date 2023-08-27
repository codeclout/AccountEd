package membertypes

import (
	"errors"
	"strings"
	"sync"
	"testing"
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
	var wg sync.WaitGroup

	addresses := []string{"something.null", "x.y@example.com", "Barry Gibbs <bg@example.com>"}
	prime := []*PrimaryMemberStartRegisterIn{
		&PrimaryMemberStartRegisterIn{Username: &addresses[0]},
		&PrimaryMemberStartRegisterIn{Username: &addresses[1]},
		&PrimaryMemberStartRegisterIn{Username: &addresses[2]},
	}

	cases := []struct {
		a string
		b error
	}{
		{a: "", b: ErrorInvalidEmail},
		{a: "x.y@example.com", b: nil},
		{a: "bg@example.com", b: nil},
	}

	for i, address := range cases {
		wg.Add(1)
		e := ValidatePrimaryMember(prime[i], &wg)
		wg.Wait()

		if *prime[i].Username != address.a || !errors.Is(e, address.b) {
			t.Errorf("expected address: %s e: %t \n received address: %s e: %t", *prime[i].Username, e, address.a, address.b)
		}
	}
}

func TestValidatePin(t *testing.T) {
	pins := []string{"7472", "382023475", "301%132$", "472651", "4276291", "32824746169"}

	cases := []struct {
		a string
		b error
	}{
		{a: "", b: ErrorPinInvalid},
		{a: "ok", b: nil},
		{a: "", b: ErrorPinInvalid},
		{a: "", b: ErrorPinInvalid},
		{a: "ok", b: nil},
		{a: "", b: ErrorPinInvalid},
	}

	for i, pin := range cases {
		c, ok := ValidatePin(&pins[i])
		id := *c

		if id != pin.a || !errors.Is(ok, pin.b) {
			t.Errorf("expected pin: %s ok: %t \n received pin: %s ok: %t", id, ok, pin.a, pin.b)
		}
	}
}
