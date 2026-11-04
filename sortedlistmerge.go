package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n2 == nil {
		return n1
	}

	if n1 == nil {
		return n2
	}

	dummy := &NodeI{}
	current := dummy

	curr1 := n1
	curr2 := n2

	for curr1 != nil && curr2 != nil {
		if curr1.Data < curr2.Data {
			current.Next = curr1
			curr1 = curr1.Next
		} else {
			current.Next = curr2
			curr2 = curr2.Next
		}
		current = current.Next
	}
	if curr1 != nil {
		current.Next = curr1
	} else {
		current.Next = curr2
	}
	return dummy.Next
}
