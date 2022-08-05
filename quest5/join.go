package piscine

func Join(strs []string, sep string) string {
	var all string
	all = ""
	for i := 0; i < len(strs); i++ {
		all += strs[i]
		if i != len(strs)-1 {
			all += sep
		}
	}
	return all
}
