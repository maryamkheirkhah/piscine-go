package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {

		temp := l.Head
		newNode.Next = temp
		l.Head = newNode
	}
}
