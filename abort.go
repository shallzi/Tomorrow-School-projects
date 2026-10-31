package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	// Сортируем массив по возрастанию
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				// Меняем элементы местами
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	// Теперь по центру гарантированно будет медиана
	return arr[2]
}
