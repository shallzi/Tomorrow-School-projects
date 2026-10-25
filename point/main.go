package main

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

	print("x = ", points.x, ", y = ", points.y, "\n")
}
