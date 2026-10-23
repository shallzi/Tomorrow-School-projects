package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	dec := 0
	baseLenFrom := len(baseFrom)
	for i := 0; i < len(nbr); i++ {
		digit := 0
		for j := 0; j < len(baseFrom); j++ {
			if baseFrom[j] == nbr[i] {
				digit = j
				break
			}
		}
		dec = dec*baseLenFrom + digit
	}

	if dec == 0 {
		return string(baseTo[0])
	}

	baseLenTo := len(baseTo)
	var result []byte
	for dec > 0 {
		rem := dec % baseLenTo
		result = append(result, baseTo[rem])
		dec /= baseLenTo
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
