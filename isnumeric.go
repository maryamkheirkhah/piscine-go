package piscine

func IsNumeric(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if !(r[i] >= '0' && r[i] <= '9') {
			return false
		}
	}
	return true
}
