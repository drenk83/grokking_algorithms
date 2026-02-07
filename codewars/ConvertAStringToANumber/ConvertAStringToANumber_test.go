package convertastringtoanumber

import (
	"testing"
)

func TestStringToNumber(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1234", 1234},
		{"605", 605},
		{"1405", 1405},
		{"-7", -7},
	}

	for _, test := range tests {
		t.Run("based test", func(t *testing.T) {
			got := StringToNumber(test.input)
			want := test.output

			if got != want {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
