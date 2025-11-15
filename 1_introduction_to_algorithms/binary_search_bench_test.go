package binarysearch

import (
	"testing"
)

var (
	arr1 = makeArr(1_000_000)
	arr2 = makeArr(10_000_000)
	arr3 = makeArr(100_000_000)
	arr4 = makeArr(1_000_000_000)
	indx = 17
)

var tests = []struct {
	name string
	arr  []int
	id   int
}{
	{"arr1", arr1, indx},
	{"arr2", arr2, indx},
	{"arr3", arr3, indx},
	{"arr4", arr4, indx},
}

func Benchmark_binarySearch(b *testing.B) {
	for _, t := range tests {
		b.Run(t.name, func(b *testing.B) {
			for b.Loop() {
				binarySearch(t.arr, t.id)
			}
		})
	}
}

func Benchmark_binarySearch_withBitShift(b *testing.B) {
	for _, t := range tests {
		b.Run(t.name, func(b *testing.B) {
			for b.Loop() {
				binarySearch_withBitShift(t.arr, t.id)
			}
		})
	}
}

func makeArr(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}
