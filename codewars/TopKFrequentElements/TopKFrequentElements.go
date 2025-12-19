package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	// key: num, value: frequent
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	buck := make([][]int, len(nums)+1)
	for num, value := range hash {
		buck[value] = append(buck[value], num)
	}

	result := make([]int, 0, k)
	for i := len(buck) - 1; i >= 0 && len(result) < k; i-- {
		if buck[i] != nil {
			result = append(result, buck[i]...)
		}
	}

	return result[:k]
}
