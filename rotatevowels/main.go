package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	var vowels []rune
	for _, arg := range args {
		for _, ch := range arg {
			if isVowel(ch) {
				vowels = append(vowels, ch)
			}
		}
	}

	n := len(vowels)
	for i := 0; i < n/2; i++ {
		vowels[i], vowels[n-1-i] = vowels[n-1-i], vowels[i]
	}

	vowelIdx := 0
	for i, arg := range args {
		if i > 0 {
			z01.PrintRune(' ')
		}
		for _, ch := range arg {
			if isVowel(ch) {
				z01.PrintRune(vowels[vowelIdx])
				vowelIdx++
			} else {
				z01.PrintRune(ch)
			}
		}
	}
	z01.PrintRune('\n')
}
