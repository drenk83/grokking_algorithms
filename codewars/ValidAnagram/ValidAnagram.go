package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashS := [26]int{}
	hashT := [26]int{}

	for i := range s {
		hashS[s[i]-'a']++
		hashT[t[i]-'a']++
	}

	if hashS == hashT {
		return true
	}
	return false
}
