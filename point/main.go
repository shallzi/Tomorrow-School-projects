package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = int('*')
	ptr.y = len("                     ")
}

func printNum(n int) {
	if n >= len("          ") {
		printNum(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNum(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNum(points.y)
}
