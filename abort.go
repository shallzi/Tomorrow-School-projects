package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{}
	arr = append(arr, a, b, c, d, e)
	middle := len(arr) / 2
	return arr[middle]
}
