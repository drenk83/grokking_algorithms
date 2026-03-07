package main

import "fmt"

func Search(arr []int) (int, int) {
	res1, res2 := arr[0], arr[1]
	min := res2 - res1

	for i := 1; i < len(arr) - 1; i++ {
		prev := arr[i]
		next := arr[i + 1]

		temp := next - prev
		if temp < min {
			res1 = prev
			res2 = next
			min = temp
		}
	}
	return res1, res2
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])	
	}

	x, y := Search(arr)
	fmt.Print(x, " ", y, "\n")
}
