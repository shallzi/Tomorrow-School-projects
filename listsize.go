package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	result := 0

	current := l.Head

	for current != nil {
		result++
		current = current.Next
	}

	return result
}
