package selectionsort

func SelectionSort(arr []int) {
	for i := range arr {
		smallest := findSmallest(arr, i)
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
}

func findSmallest(arr []int, index int) int {
	smallest := arr[index]
	out := index
	for ; index < len(arr); index++ {
		if arr[index] < smallest {
			smallest = arr[index]
			out = index
		}
	}
	return out
}
