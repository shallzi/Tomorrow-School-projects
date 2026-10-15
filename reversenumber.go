package main

import "fmt"

func reverse(s string) {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] < b[j] {
			b[i], b[j] = b[j], b[i]
		}
	}
	for _, char := range b {
		fmt.Printf("%c\n", char)
	}
}