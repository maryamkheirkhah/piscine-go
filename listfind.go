package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	temp := l.Head
	for temp != nil {
		if CompStr(temp.Data, ref) {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}
