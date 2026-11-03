package piscine

func ListPushFront(l *List, data interface{}) {
	uzel := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = uzel
		l.Tail = uzel
	} else {
		temp := l.Head
		l.Head = uzel
		l.Head.Next = temp
	}
}
