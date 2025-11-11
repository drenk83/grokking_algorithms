package main

import "fmt"

func main() {
	fmt.Println()
}

// Сложность: O(log n)
func binarySearch(arr []int, x int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		// для большего ускорения можно использовать побитовый сдвиг
		// mid := start + (end-start)>>1  но так менее читаемо
		guess := arr[mid]

		switch {
		case guess == x:
			return mid
		case guess > x:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1
}
