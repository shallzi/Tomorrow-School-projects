package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	output := []rune{
		'x', ' ', '=', ' ',
		'0', '0',
		',', ' ', 'y', ' ', '=', ' ',
		'0', '0',
		'\n',
	}

	for i := range output {
		switch i {
		case 4:
			for n := 0; n < points.x/10; n++ {
				output[i]++
			}
		case 5:
			for n := 0; n < points.x/10; n++ {
				output[i]++
			}
		case 12:
			for n := 0; n < points.y/10; n++ {
				output[i]++
			}
		case 13:
			for n := 0; n < points.y/10; n++ {
				output[i]++
			}
		}
		z01.PrintRune(output[i])
	}
}
