package piscine

func ConcatParams(args []string) string {
	var str string
	for i, s := range args {
		str += s
		if i < len(args)-1 {
			str += "\n"
		}

	}
	return str
}
