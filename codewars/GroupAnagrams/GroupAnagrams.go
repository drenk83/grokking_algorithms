package groupanagrams

// hash: key: [26]int, value: []string
// time: o(n * m), memory: o(n * m)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	hash := make(map[[26]int][]string)

	for _, v := range strs {
		arr := [26]int{}
		for i := range v {
			arr[v[i]-'a']++
		}

		hash[arr] = append(hash[arr], v)
	}

	result := make([][]string, 0, len(hash))
	for key, _ := range hash {
		result = append(result, hash[key])
	}

	return result
}
