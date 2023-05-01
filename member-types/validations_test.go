package memberTypes

import "testing"

func TestValidateName(t *testing.T) {
  names := []string{"Brian", "<script>", "O'Brien", "López", "Gutiérrez", "De Jong", "Adébáyọ̀", "()", "Trey$"}

  tcases := []struct {
    result bool
  }{
    {result: true},
    {result: false},
    {result: true},
    {result: true},
    {result: true},
    {result: true},
    {result: true},
    {result: false},
    {result: false},
  }

  for i, tc := range tcases {
    if c := ValidateName(&names[i]); c != tc.result {
      t.Errorf("expected %t \n received %t", c, tc.result)
    }
  }
}
