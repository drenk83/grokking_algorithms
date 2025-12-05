// quicksort_test.go
package quicksort

import (
	"slices"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "unsorted with duplicates",
			input:    []int{3, 6, 1, 2, 3, 5, 1, 4},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 6},
		},
		{
			name:     "two elements",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := quicksort(test.input)
			if !slices.Equal(result, test.expected) {
				t.Errorf("quicksort(%v) = %v, want %v", test.input, result, test.expected)
			}
		})
	}
}
