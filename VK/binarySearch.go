package main

import "errors"

var ErrNotFound = errors.New("not found")
var ErrArrayEmpty = errors.New("array is empty")

func BinarySearch(arr []int, target int) (int, error) {
	n := len(arr)
	if n == 0 {
		return -1, ErrArrayEmpty
	}

	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid, nil
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, ErrNotFound
}

func LeftmostBinarySearch(arr []int, target int) (int, error) {
	n := len(arr)
	if n == 0 {
		return -1, ErrArrayEmpty
	}

	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if arr[left] == target {
		return left, nil
	}

	return -1, ErrNotFound
}

func RightmostBinarySearch(arr []int, target int) (int, error) {
	n := len(arr)
	if n == 0 {
		return -1, ErrArrayEmpty
	}

	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left+1)/2

		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	if arr[right] == target {
		return right, nil
	}

	return -1, ErrNotFound
}

func TernarySearch(arr []int, target int) (int, error) {
	n := len(arr)
	if n == 0 {
		return -1, ErrArrayEmpty
	}

	left := 0
	right := n - 1

	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1, nil
		}
		if arr[mid2] == target {
			return mid2, nil
		}

		if arr[mid1] > target {
			right = mid1 - 1
		} else if arr[mid2] < target {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}

	return -1, ErrNotFound
}

func ExponentialSearch(arr []int, target int) (int, int, error) {
	if len(arr) == 0 {
		return -1, -1, ErrArrayEmpty
	}

	if arr[0] == target {
		return 0, 0, nil
	}

	border := 1
	n := len(arr)

	for border < n && arr[border] < target {
		border *= 2
	}

	left := border / 2
	right := min(border, n-1)

	return left, right, nil
}
