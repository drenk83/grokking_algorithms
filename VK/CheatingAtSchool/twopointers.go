package main

import "fmt"

func SortArr(arr []int) []int {
	ptr := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i != ptr {
				arr[ptr], arr[i] = arr[i], arr[ptr]
			}
			ptr++
		}
	}

	return arr
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	arr = SortArr(arr)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()
}
