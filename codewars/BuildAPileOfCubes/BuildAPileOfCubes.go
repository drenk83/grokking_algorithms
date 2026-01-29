package buildapileofcubes

func FindNb(m int) int {
	n := 0

	for m > 0 {
		n += 1
		m -= (n * n * n)
	}

	if m == 0 {
		return n
	}

	return -1
}
