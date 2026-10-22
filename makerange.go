package piscine

func MakeRange(min, max int) []int {
	if min > max || min == max || min == 0 && max == 0 {
		return []int(nil)
	}

	size := max - min

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i + min
	}

	return slice
}
