package piscine

func IsPrime(nb int) bool {
	Or := false

	if nb <= 0 || nb == 1 {
		return Or
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return Or
		}
	}
	return true
}
