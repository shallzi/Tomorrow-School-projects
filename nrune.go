package piscine

func NRune(s string, n int) rune {
	srez := []rune(s)
	if len(srez) == 0 || n < 0 {
		return 0
	}
	return srez[n]
}
