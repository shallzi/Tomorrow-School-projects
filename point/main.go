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
		if r >= '0' && r <= '9' {
			if r == '4' {
				z01.PrintRune(rune(points.x))
				continue
			}
			if r == '2' {
				z01.PrintRune(rune(points.y / ('+' - '*' + '+' - '*')))
				continue
			}
			if r == '1' {
				z01.PrintRune(rune(points.y % ('+' - '*' + '+' - '*')))
				continue
			}
		}
		z01.PrintRune(r)
	}
}
