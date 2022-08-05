package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for c := 0; c < len(r); c++ {
		if r[c] < 'A' || r[c] > 'Z' {
			return false
		}
	}
	return true
}
