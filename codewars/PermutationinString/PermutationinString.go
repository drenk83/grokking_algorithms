package permutationinstring

// O(n)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	hashS1 := [26]int{}
	window := [26]int{}

	for i, _ := range s1 {
		hashS1[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	if hashS1 == window {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		window[s2[i]-'a']++
		window[s2[i-len(s1)]-'a']--

		if hashS1 == window {
			return true
		}
	}
	return false
}
