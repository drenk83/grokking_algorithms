package findallanagramsinastring

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	hashP := [26]int{}
	window := [26]int{}
	result := make([]int, 0, len(s)/len(p)+1)

	for i, _ := range p {
		hashP[p[i]-'a']++
		window[s[i]-'a']++
	}

	if hashP == window {
		result = append(result, 0)
	}

	for i := len(p); i < len(s); i++ {
		window[s[i]-'a']++
		window[s[i-len(p)]-'a']--

		if window == hashP {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
