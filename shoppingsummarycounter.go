package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	check := make(map[string]int)

	// Если строка пустая, сразу возвращаем пустую мапу
	if len(str) == 0 {
		return check
	}

	word := ""

	// Идем по каждому символу в строке
	for i := 0; i < len(str); i++ {
		// Если это пробел, значит текущее слово закончилось
		if str[i] == ' ' {
			if word != "" { // Проверяем, что слово не пустое (на случай двойных пробелов)
				check[word]++
				word = "" // Сбрасываем для следующего слова
			}
		} else {
			// Если не пробел, дописываем букву к текущему слову
			word += string(str[i])
		}
	}

	// Не забываем добавить последнее слово, так как строка обычно не заканчивается пробелом
	if word != "" {
		check[word]++
	}

	return check
}
