package quicksort

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := make([]int, 0, len(arr)/2)
	greater := make([]int, 0, len(arr)/2)

	for _, v := range arr[1:] {
		if pivot > v {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(quicksort(less), pivot), quicksort(greater)...)
}
