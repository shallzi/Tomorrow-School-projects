package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(arg) == 1 {
		fileName := arg[0]
		content, error := os.ReadFile(fileName)
		if error != nil {
			fmt.Println("Error:", error)
		}
		fmt.Println(string(content))
	}
}
