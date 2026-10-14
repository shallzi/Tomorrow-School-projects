package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := make([]rune, n)
	for i := 0; i < n; i++ {
		digits[i] = rune('0' + i)
	}

	first := true

	for {
		if !first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		first = false

		for i := 0; i < n; i++ {
			z01.PrintRune(digits[i])
		}

		// Ищем цифру с конца, которую можно увеличить
		idx := n - 1
		for idx >= 0 && digits[idx] == rune('0'+10-n+idx) {
			idx--
		}

		if idx < 0 {
			break
		}

		// Увеличиваем найденную цифру и выстраиваем следующие за ней по порядку
		digits[idx]++
		for j := idx + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}

	z01.PrintRune('\n')
}
