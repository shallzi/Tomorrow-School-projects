package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	counts := [10]int{}

	if n == 0 {
		counts[0] = 1
	}

	temp := n
	for temp > 0 {
		digit := temp % 10
		counts[digit]++
		temp /= 10
	}
	for i := 0; i <= 9; i++ {
		for j := 0; j < counts[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
