package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		s := os.Args[0]
		lastSlash := -1
		for i := 0; i < len(s); i++ {
			if s[i] == '/' {
				lastSlash = i
			}
		}
		filename := s[lastSlash+1:]
		for _, r := range filename {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
