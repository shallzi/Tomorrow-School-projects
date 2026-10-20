package piscine

func FirstRune(s string) rune {
	srez := []rune(s)
	if len(srez) == 0 {
		return 0
	}
	return srez[0]
}
