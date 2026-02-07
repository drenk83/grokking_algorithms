package convertastringtoanumber

import "strconv"

func StringToNumber(str string) int {
	out, _ := strconv.Atoi(str)
	return out
}
