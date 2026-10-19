package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	for {
		isPrime := true

		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return nb
		}
		nb++
	}
}
