package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		printNV()
		return
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			printNV()
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				printNV()
				return
			}
		}
	}

	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		printRecursive(922337203685477580, int64(len(base)), base)
		printRecursive(8, int64(len(base)), base)
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	printRecursive(int64(nbr), int64(len(base)), base)
}

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}

func printRecursive(nbr int64, baseLen int64, base string) {
	if nbr >= baseLen {
		printRecursive(nbr/baseLen, baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}
