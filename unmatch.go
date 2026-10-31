package piscine

func Unmatch(a []int) int {
	mapp := make(map[int]int)
	for i := 0; i < len(a); i++ {
		mapp[i]++
	}
	for i := 0; i < len(mapp); i++ {
		if mapp[i] == 1 {
			return i
		}
	}
	return 0
}
