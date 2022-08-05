package piscine

func IsAlpha(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if !((r[i] >= 'a' && r[i] <= 'z') || (r[i] >= 'A' && r[i] <= 'Z') || (r[i] >= '0' && r[i] <= '9')) {
			return false
		}
	}
	return true
}
