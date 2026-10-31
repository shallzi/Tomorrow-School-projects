package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	check := make(map[string]int)

	// Если строка абсолютно пустая, возвращаем пустую мапу
	if len(str) == 0 {
		return check
	}

	word := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			// Добавляем слово в мапу (даже если оно пустое "")
			check[word]++
			word = "" // Сбрасываем для следующего слова
		} else {
			word += string(str[i])
		}
	}

	// Добавляем последнее оставшееся слово
	check[word]++

	return check
}
