package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func Sorted(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
