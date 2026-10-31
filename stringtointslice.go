package piscine

func StringToIntSlice(str string) []int {
	arr := []byte(str)
	result := make([]int, len(arr))
	for i, b := range arr {
		result[i] = int(b)
	}
	return result
}
