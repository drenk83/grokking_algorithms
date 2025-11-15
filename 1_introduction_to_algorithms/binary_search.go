package binarysearch

import "fmt"

// Сложность: O(log n)
func binarySearch(arr []int, x int) (int, error) {
	if len(arr) == 0 {
		return -1, fmt.Errorf("len(arr) == 0")
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		// для большего ускорения можно использовать побитовый сдвиг
		// mid := left + (right-left)>>1   но так менее читаемо
		guess := arr[mid]

		switch {
		case guess == x:
			return mid, nil
		case guess > x:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1, fmt.Errorf("no in arr")
}

func binarySearch_withBitShift(arr []int, x int) (int, error) {
	if len(arr) == 0 {
		return -1, fmt.Errorf("len(arr) == 0")
	}

	left, right := 0, len(arr)-1
	for left <= right {
		//mid := left + (right-left)/2
		mid := left + (right-left)>>1
		guess := arr[mid]

		switch {
		case guess == x:
			return mid, nil
		case guess > x:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1, fmt.Errorf("no in arr")
}
