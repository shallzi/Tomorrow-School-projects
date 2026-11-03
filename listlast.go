package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil || l == nil {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
