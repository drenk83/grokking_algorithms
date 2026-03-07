package main

import "fmt"

func main() {
	var n int
	var elem int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&elem)

	first := true
	for i := 0; i < n; i++ {
		if arr[i] != elem {
			if first {
				fmt.Print(arr[i])
				first = false
			} else {
				fmt.Print(" ")
				fmt.Print(arr[i])
			}
		}
	}
	fmt.Println()
}
