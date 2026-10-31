package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return []int(nil)
	}
	arr := []rune(str)
	result := make([]int, len(arr))
	for i, b := range arr {
		result[i] = int(b)
	}
	return result
}
