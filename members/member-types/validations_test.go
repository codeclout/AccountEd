package memberTypes

import "testing"

func TestValidateName(t *testing.T) {
  names := []string{"brian", "<script>", "O'Brien", "lópez", "Gutiérrez", "de jong", "Adébáyọ̀", "()", "Trey$"}

  tcases := []struct {
    a string
    b bool
  }{
    {a: "Brian", b: true},
    {a: "invalid", b: false},
    {a: "O'Brien", b: true},
    {a: "López", b: true},
    {a: "Gutiérrez", b: true},
    {a: "De Jong", b: true},
    {a: "Adébáyọ̀", b: true},
    {a: "invalid", b: false},
    {a: "invalid", b: false},
  }

  for i, tc := range tcases {
    c, ok := ValidateName(&names[i])
    name := *c

    if name != tc.a || tc.b != ok {
      t.Errorf("expected name: %s %t \n received name: %s %t", name, ok, tc.a, tc.b)
    }
  }
}

func TestValidateEmail(t *testing.T) {
  addresses := []string{"x@example.com", "x@example", "x.y@example.com", "x-y@example.com", "x_y@example.com", "Barry Gibbs <bg@example.com>"}

  tcases := []struct {
    a string
    b bool
  }{
    {a: "x@example.com", b: true},
    {a: "invalid", b: false},
    {a: "x.y@example.com", b: true},
    {a: "x-y@example.com", b: true},
    {a: "x_y@example.com", b: true},
    {a: "bg@example.com", b: true},
  }

  for i, address := range tcases {
    c, ok := ValidateEmail(&addresses[i])
    email := *c

    if email != address.a || ok != address.b {
      t.Errorf("expected address: %s ok: %t \n received address: %s ok: %t", email, ok, address.a, address.b)
    }
  }
}

func TestValidatePin(t *testing.T) {
  pins := []string{"7472", "382023475", "301%132$", "472651", "4276291", "32824746169"}

  tcases := []struct {
    a string
    b bool
  }{
    {a: "invalid", b: false},
    {a: "ok", b: true},
    {a: "invalid", b: false},
    {a: "invalid", b: false},
    {a: "ok", b: true},
    {a: "invalid", b: false},
  }

  for i, pin := range tcases {
    c, ok := ValidatePin(&pins[i])
    id := *c

    if id != pin.a || ok != pin.b {
      t.Errorf("expected pin: %s ok: %t \n received pin: %s ok: %t", id, ok, pin.a, pin.b)
    }
  }
}