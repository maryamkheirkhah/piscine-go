package piscine

func ListSize(l *List) int {
	count := 0
	next := l.Head
	for next != nil {
		count++
		next = next.Next
	}
	return count
}
