package piscine

func IterativePower(nb int, power int) int {
	digit := nb

	result := 1

	if power < 0 {
		return 0
	}

	for i := 1; i <= power; i++ {
		result *= digit
	}
	return result
}
