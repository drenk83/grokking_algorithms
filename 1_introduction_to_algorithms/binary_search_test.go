package binarysearch

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		srch int
		want int
	}{
		{"one", []int{1}, 1, 0},
		{"two", []int{1, 2}, 2, 1},
		{"tree", []int{1, 2, 3}, 3, 2},
		{"no in arr", []int{1, 2, 3, 4, 5}, 10, -1},
		{"no in arr", []int{1, 10, 100, 1000}, 5, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := binarySearch(test.arr, test.srch)
			if got != test.want {
				t.Errorf("%s: got: %d search: %d want: %d", test.name,
					test.arr, test.srch, test.want)
			}
		})
	}
}
