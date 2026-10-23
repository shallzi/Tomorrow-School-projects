package piscine

func Split(s, sep string) []string {
	result := []string{}
	size := len(sep)
	if size == 0 {
		return result
	}
	start := 0
	for i := 0; i < len(s)-size; {
		if s[i:i+size] == sep {
			result = append(result, s[start:i])
			i += size
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result
}
