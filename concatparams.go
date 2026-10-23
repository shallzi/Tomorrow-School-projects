package piscine

func ConcatParams(args []string) string {
	srez := args

	if len(srez) == 0 {
		return ""
	}

	result := ""

	for i := 0; i < len(srez); i++ {
		result += srez[i]
		if i+1 < len(srez) {
			result += "\n"
		}
	}
	return result
}
