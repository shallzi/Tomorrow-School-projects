package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, char := range i {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
