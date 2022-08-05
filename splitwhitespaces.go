package piscine

func SplitWhiteSpaces(s string) []string {
	var args []string
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		var a string
		for runes[i] != ' ' {
			a += string(runes[i])
			i++
			if i >= len(runes) {
				break
			}
		}
		if a != "" {
			args = append(args, a)
		}
	}
	return args
}
