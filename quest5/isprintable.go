package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] < ' ' {
			return false
		}
	}
	return true
}
