package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	current := l.Head
	prev := &NodeL{}

	for current != nil {
		if current.Data == data_ref {
			prev.Next = current.Next
		}

		if current.Next != nil && current.Next.Data == data_ref {
			current.Next = current.Next.Next
		}

		prev = current
		current = current.Next
	}
}
