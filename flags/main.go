package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	var insertStr string
	hasInsert := false
	hasOrder := false
	var targetStr string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		} else if hasPrefix(arg, "--insert=") {
			insertStr = trimPrefix(arg, "--insert=")
			hasInsert = true
		} else if hasPrefix(arg, "-i=") {
			insertStr = trimPrefix(arg, "-i=")
			hasInsert = true
		} else if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
				hasInsert = true
			}
		} else if arg == "--order" || arg == "-o" || hasPrefix(arg, "--order=") || hasPrefix(arg, "-o=") {
			hasOrder = true
		} else {
			
			targetStr = arg
		}
	}


	result := targetStr
	if hasInsert {
		result = result + insertStr
	}


	if hasOrder {
		runes := []rune(result)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		result = string(runes)
	}

	
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}


func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

/
func trimPrefix(s, prefix string) string {
	if hasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}