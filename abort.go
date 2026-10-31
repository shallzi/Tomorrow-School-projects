package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{}
	arr = append(arr, a, b, c, d, e)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	middle := len(arr) / 2
	return arr[middle]
}
