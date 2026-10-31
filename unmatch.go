package piscine

func Unmatch(a []int) int {
	mapp := make(map[int]int)
	for i := 0; i < len(a); i++ {
		mapp[a[i]]++
	}
	for _, val := range a {
		if mapp[val] == 1 {
			return val
		}
	}
	return 0
}
