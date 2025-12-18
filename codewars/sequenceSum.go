package codewars

func SequenceSum(start, end, step int) int {
	if start > end {
		return 0
	}

	out := 0

	for ; start <= end; start += step {
		out += start
	}

	return out
}

// реализовать с помощью формулы
