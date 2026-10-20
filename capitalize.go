package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if !inWord {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				inWord = true
			} else {
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			inWord = false
		}
	}

	return string(runes)
}
