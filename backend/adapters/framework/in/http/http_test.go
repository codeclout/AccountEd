package http

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	appPorts := []string{"", "3000", "240", "75535", "65535"}

	tcases := []struct {
		result string
	}{
		{result: ":8088"},
		{result: ":3000"},
		{result: ":8088"},
		{result: ":8088"},
		{result: ":65535"},
	}

	for i, tc := range tcases {
		e := os.Setenv("PORT", appPorts[i])

		if e != nil {
			t.Errorf("%s", e)
		}

		if c := getPort(); c != tc.result {
			t.Errorf("expected %s \n received %s", c, tc.result)
		}

		var _ = os.Unsetenv("PORT")
	}
}
