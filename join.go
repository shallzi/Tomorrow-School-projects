package piscine

func Join(strs []string, sep string) string {
	var result string

	for i, s := range strs {
		if i < len(strs)-1 {
			result += s + sep
		} else {
			result += s
		}
	}
	return result
}
