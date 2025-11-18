package recursion

import (
	"testing"
)

func TestFact(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"Zero", 0, 1},
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 6},
		{"Four", 4, 24},
		{"Five", 5, 120},
		{"Six", 6, 720},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := fact(test.in)
			if got != test.want {
				t.Errorf("Test name: %s\n Input: %d\n, Want: %d\n Got: %d\n",
					test.name, test.in, test.want, got)
			}
		})
	}
}

func TestExtendedEuclideanAlgorithm(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		gcd     int
		x       int
		y       int
		errBool bool
	}{
		{"12 5", 12, 5, 1, -2, 5, false},
		{"30 12", 30, 12, 6, 1, -2, false},
		{"55 34", 55, 34, 1, 13, -21, false},
		{"48 16", 48, 16, 16, 0, 1, false},
		{"42 42", 42, 42, 42, 0, 1, false},
		{"1071 462", 1071, 462, 21, -3, 7, false},
		{"240 46", 240, 46, 2, -9, 47, false},
		{"1000 35", 1000, 35, 5, 2, -57, false},
		{"2 1", 2, 1, 1, 0, 1, false},
		{"1 13", 1, 13, 0, 0, 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotGCD, gotX, gotY, err := ExtendedEuclideanAlgorithm(test.a, test.b)

			if (err != nil) != test.errBool {
				t.Errorf("Test name: %s\n Input: %d, %d\n Want: %d, %d, %d\n Got: %d, %d, %d",
					test.name, test.a, test.b, test.gcd, test.x, test.y, gotGCD, gotX, gotY)
			}

			if test.gcd != gotGCD || test.x != gotX || test.y != gotY {
				t.Errorf("Test name: %s\n Input: %d, %d\n Want: %d, %d, %d\n Got: %d, %d, %d",
					test.name, test.a, test.b, test.gcd, test.x, test.y, gotGCD, gotX, gotY)
			}
		})
	}
}

func TestDivideWithRemainder(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		q       int
		r       int
		errBool bool
	}{
		{
			name: "positive numbers",
			a:    17, b: 5, q: 3, r: 2, errBool: false,
		},
		{
			name: "exact division",
			a:    20, b: 4, q: 5, r: 0, errBool: false,
		},
		{
			name: "divide by 1",
			a:    42, b: 1, q: 42, r: 0, errBool: false,
		},
		{
			name: "divide by -1",
			a:    42, b: -1, q: -42, r: 0, errBool: false,
		},
		{
			name: "negative dividend",
			a:    -17, b: 5, q: -3, r: -2, errBool: false,
		},
		{
			name: "negative divisor",
			a:    17, b: -5, q: -3, r: 2, errBool: false,
		},
		{
			name: "both negative",
			a:    -17, b: -5, q: 3, r: -2, errBool: false,
		},
		{
			name: "zero dividend",
			a:    0, b: 7, q: 0, r: 0, errBool: false,
		},
		{
			name: "division by zero",
			a:    10, b: 0, q: 0, r: 0, errBool: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotQ, gotR, gotErr := DivideWithRemainder(test.a, test.b)

			if (gotErr != nil) != test.errBool {
				t.Errorf("Test name: %s\n Input: %d, %d\n Want: %d, %d\n Got: %d, %d",
					test.name, test.a, test.b, test.q, test.r, gotQ, gotR)
			}
			if test.r != gotR || test.q != gotQ {
				t.Errorf("Test name: %s\n Input: %d, %d\n Want: %d, %d\n Got: %d, %d",
					test.name, test.a, test.b, test.q, test.r, gotQ, gotR)
			}
		})
	}
}
