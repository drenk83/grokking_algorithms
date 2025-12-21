package findkclosestelements

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			arr:      []int{1, 1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findClosestElements(tt.arr, tt.k, tt.x)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestFindClosestElements_easy(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{2, 3, 5, 7, 11}, 2, 3, []int{5, 7}},
		{[]int{4, 12, 15, 15, 24}, 3, 1, []int{12, 15, 15}},
		{[]int{2, 3, 5, 7, 11}, 2, 2, []int{3, 5}}, // или {5, 7}
	}

	for _, tt := range tests {
		got := findClosestElements_easy(tt.arr, tt.k, tt.x)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findClosestElements(%v, %d, %d) = %v, want %v",
				tt.arr, tt.k, tt.x, got, tt.want)
		}
	}
}
