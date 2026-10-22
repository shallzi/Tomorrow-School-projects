package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return []int(nil)
	}

	size := max - min
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i + min
	}

	return slice
}
