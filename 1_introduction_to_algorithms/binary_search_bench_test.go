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

func Benchmark_binarySearch_withBitShift_1(b *testing.B) {
	for b.Loop() {
		binarySearch_withBitShift(arr1, indx)
	}
}

func Benchmark_binarySearch_withBitShift_2(b *testing.B) {
	for b.Loop() {
		binarySearch_withBitShift(arr2, indx)
	}
}

func Benchmark_binarySearch_withBitShift_3(b *testing.B) {
	for b.Loop() {
		binarySearch_withBitShift(arr3, indx)
	}
}

func Benchmark_binarySearch_withBitShift_4(b *testing.B) {
	for b.Loop() {
		binarySearch_withBitShift(arr4, indx)
	}
}

func makeArr(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}
