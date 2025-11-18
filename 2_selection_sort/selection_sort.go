package selectionsort

func SelectionSort(arr []int) {
	for i := range arr {
		smallest := findSmallest(arr, i)
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
}

func findSmallest(arr []int, index int) int {
	out := index
	for i := index; i < len(arr); i++ {
		if arr[i] < arr[out] {
			out = i
		}
	}
	return out
}
