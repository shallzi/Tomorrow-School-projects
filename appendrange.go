package piscine

func AppendRange(min, max int) []int {
	start := min
	end := max

	var slice []int

	if min < 0 || min > max || max <= 0 {
		return slice
	}

	for i := start; i < end; i++ {
		slice = append(slice, i)
	}
	return slice
}
