package codewars

import (
	"slices"
	"sort"
)

func Smallest_V20(numbers []int) int {
	min := numbers[0]
	for _, v := range numbers[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Smallest_V21(numbers []int) int {
	smallest := numbers[0]
	for _, v := range numbers[1:] {
		smallest = min(smallest, v)
	}
	return smallest
}

func Smallest_Slices(numbers []int) int {
	return slices.Min(numbers)
}

func Smallest_Sort(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
