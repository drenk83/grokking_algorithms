package codewars

import (
	"math/rand"
	"testing"
)

var benchData []int

func init() {
	size := 10_000_000
	benchData = make([]int, size)

	for i := range benchData {
		benchData[i] = rand.Int()
	}
}

func BenchmarkSmallest_V20(b *testing.B) {
	for b.Loop() {
		_ = Smallest_V20(benchData)
	}
}

func BenchmarkSmallest_V21(b *testing.B) {
	for b.Loop() {
		_ = Smallest_V21(benchData)
	}
}

func BenchmarkSmallest_Slices(b *testing.B) {
	for b.Loop() {
		_ = Smallest_Slices(benchData)
	}
}

func BenchmarkSmallest_Sort(b *testing.B) {
	temp := make([]int, len(benchData))
	copy(temp, benchData)

	for b.Loop() {
		_ = Smallest_Sort(temp)
	}
}
