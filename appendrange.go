package piscine

func AppendRange(min, max int) []int {
	start := min
	end := max

	var slice []int

	if min > max {
		return []int(nil)
	}

	for i := start; i < end; i++ {
		slice = append(slice, i)
	}
	return slice
}
