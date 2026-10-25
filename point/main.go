package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = '*'
	ptr.y = '*' / ('+' - ')')
}

func main() {
	points := &point{}
	setPoint(points)

	arr := [...]rune{
		'x',
		' ',
		'=',
		' ',
		rune(int('0') + (points.x / ('*' - ' '))),
		rune(int('0') + (points.x % ('*' - ' '))),
		',',
		' ',
		'y',
		' ',
		'=',
		' ',
		rune(int('0') + (points.y / ('*' - ' '))),
		rune(int('0') + (points.y % ('*' - ' '))),
		'\n',
	}

	for _, c := range arr {
		z01.PrintRune(c)
	}
}
