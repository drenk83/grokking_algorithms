package validparentheses

func isValid(s string) bool {
	hash := [3]int{}
	// 0: (), 1: [], 2: {}

	for _, v := range s {
		switch v {
		case '(':
			hash[0]++
		case ')':
			hash[0]--
		case '[':
			hash[1]++
		case ']':
			hash[1]--
		case '{':
			hash[2]++
		case '}':
			hash[2]--
		}
	}
}
