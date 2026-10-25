package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = '*'
	ptr.y = ('*' - ' ') + ('*' - ' ') + ('+' - '*')
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNum(n int) {
	ten := int('*' - ' ')
	if n >= ten {
		printNum(n / ten)
	}
	z01.PrintRune(rune(int('0') + n%ten))
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNum(points.x)
	printStr(", y = ")
	printNum(points.y)
	printStr("\n")
}
