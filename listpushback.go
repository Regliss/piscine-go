package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
	Prev *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	list := &NodeL{
		Next: l.Head,
		Data: data,
	}
	if l.Head != nil {
		l.Head.Prev = list
	}
	l.Head = list

	L := l.Head
	for L.Next != nil {
		L = L.Next
	}
	l.Tail = L
}
