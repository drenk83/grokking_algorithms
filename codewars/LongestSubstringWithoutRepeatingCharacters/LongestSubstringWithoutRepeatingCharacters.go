package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0

	var hash [26]int

	for i := range s {
		hash[s[i]-'a']++

		if hash[s[i]-'a'] > 1 {

		}
	}
}
