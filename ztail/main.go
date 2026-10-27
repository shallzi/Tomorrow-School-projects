package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	cCount := -1
	files := []string{}

	if args[0] == "-c" && len(args) >= 2 {
		val := 0
		hasValidInt := true
		for _, r := range args[1] {
			if r < '0' || r > '9' {
				hasValidInt = false
				break
			}
			val = val*10 + int(r-'0')
		}
		if hasValidInt && val > 0 {
			cCount = val
		}
		files = args[2:]
	} else {
		files = args
	}

	hasError := false
	multipleFiles := len(files) > 1
	printedFilesCount := 0

	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}

		if multipleFiles && printedFilesCount > 0 {
			fmt.Println()
		}

		if multipleFiles {
			fmt.Printf("==> %s <==\n", filename)
		}

		if cCount >= 0 {
			if len(content) > cCount {
				content = content[len(content)-cCount:]
			}
		}

		os.Stdout.Write(content)
		printedFilesCount++
	}

	if hasError {
		os.Exit(1)
	}
}
