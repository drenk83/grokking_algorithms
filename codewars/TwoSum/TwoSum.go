package twosum

func twoSum(nums []int, target int) []int {
	// key = num, value = indx
	hash := make(map[int]int)

	for idx, num := range nums {
		diff := target - num

		if found, ok := hash[diff]; ok {
			return []int{found, idx}
		}
		hash[num] = idx
	}
	return []int{}
}
