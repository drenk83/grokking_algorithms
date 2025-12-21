package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		output [][]string
	}{
		{
			name:   "Example 2: empty string",
			input:  []string{""},
			output: [][]string{{""}},
		},
		{
			name:   "Example 3: single character",
			input:  []string{"a"},
			output: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := groupAnagrams(tc.input)

			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, result)
			}
		})
	}
}
