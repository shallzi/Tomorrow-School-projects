package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		os.Exit(1)
	}

	count := 0
	for _, r := range os.Args[2] {
		if r < '0' || r > '9' {
			os.Exit(1)
		}
		count = count*10 + int(r-'0')
	}

	hasError := false
	multipleFiles := len(os.Args) > 4
	first := true

	for _, filename := range os.Args[3:] {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("%s\n", err)
			hasError = true
			first = false
			continue
		}

		if multipleFiles {
			if !first {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		start := len(content) - count
		if start < 0 {
			start = 0
		}

		fmt.Printf("%s", content[start:])
		first = false
	}

	if hasError {
		os.Exit(1)
	}
}