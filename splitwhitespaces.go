package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	start := -1

	for i, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			if start != -1 {
				result = append(result, s[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}

	if start != -1 {
		result = append(result, s[start:])
	}
	return result
}
