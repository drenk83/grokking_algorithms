package findkclosestelements

// Бинарный поиск
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := left + (right-left)/2

		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}

/*
Дан отсортированный по неубыванию массив целых чисел a, индекс элемента
index и целое число k.
Необходимо вернуть в любом порядке k чисел из массива, которые являются
ближайшими по значению к элементу a[index].
*/

// Два указателя
func findClosestElements_easy(arr []int, k int, x int) []int {
	left := x - 1
	right := x

	for range k {
		if left < 0 {
			right++
		} else if right >= len(arr) {
			left--
		} else {
			if (arr[x] - arr[left]) <= (arr[right] - arr[x]) {
				left--
			} else {
				right++
			}
		}
	}

	return arr[left+1 : right]
}
