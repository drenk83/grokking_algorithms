package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}

func main() {
	fmt.Println("test out")

	x, y := 3, 7
	
	fmt.Println(Add(x, y))
	fmt.Println(Minus(x, y))
}

