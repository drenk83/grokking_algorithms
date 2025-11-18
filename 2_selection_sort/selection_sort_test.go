package selectionsort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"ClassicExample", []int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{"AlreadySorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"ReverseSorted", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"ManyDuplicates", []int{5, 2, 8, 5, 1, 5, 9, 3, 5, 7}, []int{1, 2, 3, 5, 5, 5, 5, 7, 8, 9}},
		{"NegativeNumbers", []int{-3, -1, -4, -2, 0, -5}, []int{-5, -4, -3, -2, -1, 0}},
		{"AllEqual", []int{7, 7, 7, 7, 7, 7}, []int{7, 7, 7, 7, 7, 7}},
		{"SingleElement", []int{42}, []int{42}},
		{"EmptyArray", []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := make([]int, len(test.arr))
			copy(got, test.arr)
			SelectionSort(got)
			if !slices.Equal(got, test.want) {
				t.Errorf("Test name: %s \n Input: %v \n Want: %v \n Got: %v ", test.name, test.arr, test.want, got)
			}
		})

	}
}
