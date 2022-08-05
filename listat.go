package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l != nil {
		node := l
		for i := 0; i < pos; i++ {
			if node.Next == nil {
				return nil
			}
			node = node.Next
		}
		return node
	}
	return nil
}
