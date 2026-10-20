package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	hasDigits := false
	isSignFound := false

	for _, i := range s {
		if i >= '0' && i <= '9' {
			hasDigits = true
			result = result*10 + int(r-'0')
		} else if !hasDigits && !isSignFound {
			if r == '-' {
				sign = -1
				isSignFound = true
			} else if r == '+' {
				isSignFound = true
			}
		}
	}

	if !hasDigits {
		return 0
	}

	return result * sign
}
