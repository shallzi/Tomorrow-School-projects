package piscine

func Index(s string, toFind string) int {
	slen := len(s)
	flen := len(toFind)
	if slen < 0 || flen < 0 {
		return 0
	}
	if slen < flen {
		return -1
	}
	for i := 0; i <= slen-flen; i++ {
		if s[i:i+flen] == toFind {
			return i
		}
	}
	return -1
}
