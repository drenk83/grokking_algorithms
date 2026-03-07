package main

import "fmt"

func lastEven(n int) int {
	result := -1
	var x int

	for range n {
		fmt.Scan(&x)
		if x%2 == 0 {
			result = x
		}
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(lastEven(n))
}
