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
	for _, fileName := range arg {
		content, error := os.ReadFile(fileName)
		if error == nil {
			for _, i := range content {
				z01.PrintRune(rune(i))
			}
		} else if error != nil {
			for _, char := range error.Error() {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
