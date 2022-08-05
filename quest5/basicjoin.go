package piscine

func BasicJoin(elems []string) string {
	var all string
	all = ""
	for i := 0; i < len(elems); i++ {
		all += elems[i]
	}
	return all
}
