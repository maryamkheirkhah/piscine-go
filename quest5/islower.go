package piscine

func IsLower(s string) bool {
	r := []rune(s)
	for c := 0; c < len(r); c++ {
		if r[c] < 'a' || r[c] > 'z' {
			return false
		}
	}
	return true
}
