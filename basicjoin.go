package piscine

func BasicJoin(elems []string) string {
	var result string

	for _, s := range elems {
		result += s
	}
	return result
}
