package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, number := range s {
		if number < '0' || number > '9' {
			return 0
		}
		digit := int(number - '0')
		result = result*10 + digit
	}
	return result
}
