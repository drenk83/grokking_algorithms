package binarysearch

import (
	"testing"
)

var (
	arr1 = makeArr(100)
	arr2 = makeArr(1000)
	arr3 = makeArr(10_000)
	arr4 = makeArr(100_000)
	arr5 = makeArr(1_000_000)
	arr6 = makeArr(10_000_000)
	arr7 = makeArr(100_000_000)
	arr8 = makeArr(1_000_000_000)

	indx = 17
)

func Benchmark_binarySearch_1(b *testing.B) {
	for b.Loop() {
		binarySearch(arr1, indx)
	}
}

func Benchmark_binarySearch_2(b *testing.B) {
	for b.Loop() {
		binarySearch(arr2, indx)
	}
}

func Benchmark_binarySearch_3(b *testing.B) {
	for b.Loop() {
		binarySearch(arr3, indx)
	}
}

func Benchmark_binarySearch_4(b *testing.B) {
	for b.Loop() {
		binarySearch(arr4, indx)
	}
}

func Benchmark_binarySearch_5(b *testing.B) {
	for b.Loop() {
		binarySearch(arr5, indx)
	}
}

func Benchmark_binarySearch_6(b *testing.B) {
	for b.Loop() {
		binarySearch(arr6, indx)
	}
}

func Benchmark_binarySearch_7(b *testing.B) {
	for b.Loop() {
		binarySearch(arr7, indx)
	}
}

func Benchmark_binarySearch_8(b *testing.B) {
	for b.Loop() {
		binarySearch(arr8, indx)
	}
}

func makeArr(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}
