package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			for _, char := range err.Error() {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
		return
	}

	hasError := false

	for _, fileName := range arg {
		content, error := os.ReadFile(fileName)
		if error == nil {
			for _, i := range content {
				z01.PrintRune(rune(i))
			}

			continue
		} else if error != nil {
			z01.PrintRune('E')
			z01.PrintRune('R')
			z01.PrintRune('R')
			z01.PrintRune('O')
			z01.PrintRune('R')
			z01.PrintRune(':')
			z01.PrintRune(' ')
			for _, char := range error.Error() {
				z01.PrintRune(char)
			}
			hasError = true
			z01.PrintRune('\n')
		}
	}
	if hasError {
		os.Exit(1)
	}
}
