package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n <= len(r) && n > 0 {
		x := n - 1
		return r[x]
	}
	return 0
}
