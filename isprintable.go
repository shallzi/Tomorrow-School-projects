package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if i < ' ' || i > '~' {
			return false
		}
	}
	return true
}
