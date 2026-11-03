package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil || l2.Head == nil {
		return
	}
	uzel := l2.Head

	if l1.Head == nil {
		l1.Head = uzel
		l1.Tail = uzel
	} else {
		l1.Tail.Next = uzel
		l1.Tail = uzel
	}
}
