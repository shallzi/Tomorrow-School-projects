package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, number := range s {
		digit := int(number - '0')
		result = result*10 + digit
	}
	return result
}
