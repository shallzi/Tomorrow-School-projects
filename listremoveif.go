package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	current := l.Head
	var prev *NodeL

	for current != nil {
		if current.Data == data_ref {
			nextTemp := current.Next
			current.Next = prev
			prev = current
			current = nextTemp
		}
	}
}
