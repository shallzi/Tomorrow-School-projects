package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	direction := 0
	for i := 0; i < len(a)-1; i++ {
		res := f(a[i], a[i+1])
		if res != 0 {
			direction = res
			break
		}
	}

	if direction == 0 {
		return true
	}

	for i := 0; i < len(a); i++ {
		res := f(a[i], a[i+1])
		if direction > 0 && res < 0 {
			return false
		}
		if direction < 0 && res > 0 {
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
