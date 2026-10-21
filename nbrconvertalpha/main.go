package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	// Check for the --upper flag as the first argument
	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n, err := atoi(arg)
		// If not a valid integer or out of the 1-26 alphabet range, print a space
		if err != nil || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		var r rune
		if upper {
			r = rune('A' + n - 1)
		} else {
			r = rune('a' + n - 1)
		}
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func atoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, os.ErrInvalid
	}

	n := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start == len(s) {
		return 0, os.ErrInvalid
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, os.ErrInvalid
		}
		n = n*10 + int(s[i]-'0')
	}

	return n * sign, nil
}
