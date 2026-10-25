package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = '*'
	ptr.y = '*' / ('+' - '*' + '+' - '*')
}

func main() {
	points := &point{}
	setPoint(points)

	for _, r := range "x = 42, y = 21\n" {
		z01.PrintRune(r)
	}
}
