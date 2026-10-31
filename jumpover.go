package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	result := ""
	arr := []rune(str)
	for index, i := range arr {
		if index != 0 && (index+1)%3 == 0 {
			result += string(i)
		}
	}
	return result + "\n"
}
