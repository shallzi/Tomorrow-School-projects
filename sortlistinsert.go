package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	uzel := &NodeI{Data: data_ref}
	if l == nil || data_ref < l.Data {
		uzel.Next = l
		return uzel
	}

	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	uzel.Next = current.Next
	current.Next = uzel
	return l
}
