package recursion

import "errors"

func Fact(x int) int {
	if x == 1 || x == 0 {
		return 1
	} else {
		return x * Fact(x-1)
	}
}

//func EuclideanAlgorithm(a, b int) (int, error) {
// not now...
//}

// Input: 0 < b <= a
// Output: GCD(a, b), x, y   (ax + by = gcd(a, b))
func ExtendedEuclideanAlgorithm(a, b int) (int, int, int, error) {
	if a < b {
		a, b = b, a
	}
	if b <= 0 {
		return 0, 0, 0, errors.New("Division by zero or negative number")
	}

	r0, r1 := a, b
	x0, x1 := 1, 0
	y0, y1 := 0, 1

	for {
		q, r, err := DivideWithRemainder(r0, r1)
		if err != nil {
			return 0, 0, 0, err
		}

		if r == 0 {
			return r1, x1, y1, nil
		}

		x0, x1 = x1, x0-(q*x1)
		y0, y1 = y1, y0-(q*y1)

		r0, r1 = r1, r
	}
}

// a = q * b + r
// Input: a, b
// Output: q, r
func DivideWithRemainder(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Division by zero")
	}
	q := a / b
	r := a % b
	return q, r, nil
}
