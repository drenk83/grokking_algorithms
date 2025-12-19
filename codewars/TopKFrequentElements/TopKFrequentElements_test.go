package topkfrequentelements

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
	}

	for i, tt := range tests {
		result := topKFrequent(tt.nums, tt.k)

		// Проверяем длину
		if len(result) != len(tt.expected) {
			t.Errorf("Test %d: expected length %d, got %d", i, len(tt.expected), len(result))
		}

		// Проверяем содержимое (порядок не важен)
		if !containsSameElements(result, tt.expected) {
			t.Errorf("Test %d: expected %v, got %v", i, tt.expected, result)
		}
	}
}

func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Создаем карту для элементов первого среза
	elementMap := make(map[int]int)
	for _, v := range a {
		elementMap[v]++
	}

	// Проверяем элементы второго среза
	for _, v := range b {
		if elementMap[v] == 0 {
			return false
		}
		elementMap[v]--
	}

	return true
}
