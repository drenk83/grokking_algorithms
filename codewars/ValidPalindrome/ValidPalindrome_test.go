package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Example 1: Palindrome with punctuation",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Example 2: Non-palindrome with spaces",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "Example 3: Empty string",
			input:    " ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
