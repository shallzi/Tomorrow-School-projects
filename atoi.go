package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	startIndex := 0

	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}
	for i := startIndex; i < len(s); i++ {
		number := s[i]
		if number < '0' || number > '9' {
			return 0
		}
		digit := int(number - '0')
		result = result*10 + digit
	}
	return result * sign
}
