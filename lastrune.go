package piscine

func LastRune(s string) rune {
	srez := []rune(s)
	if len(srez) == 0 {
		return 0
	}
	return srez[len(srez)-1]
}
