package piscine

type Node struct {
	Data interface{}
	Next *NodeL
}

type Lis struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	uzel := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = uzel
		l.Tail = uzel
	} else {
		l.Tail.Next = uzel
		l.Tail = uzel
	}
}
