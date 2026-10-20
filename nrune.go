package piscine

func NRune(s string, n int) rune {
	srez := []rune(s)
	letter := n - 1
	if len(srez) == 0 || n < 1 || n > len(srez) {
		return 0
	}
	return srez[letter]
}
