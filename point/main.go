func setPoint(ptr *point) {
	ptr.x = '*'
	ptr.y = len("                     ")
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
