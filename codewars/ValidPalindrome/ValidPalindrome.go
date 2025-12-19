package validpalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	normalized := strings.ToLower(s)

	left := 0
	right := len(normalized) - 1

	for left < right {
		if !unicode.IsLetter(rune(normalized[left])) && !unicode.IsNumber(rune(normalized[left])) {
			left++
			continue
		}

		if !unicode.IsLetter(rune(normalized[right])) && !unicode.IsNumber(rune(normalized[right])) {
			right--
			continue
		}

		if normalized[left] != normalized[right] {
			return false
		}

		right--
		left++
	}

	return true
}
