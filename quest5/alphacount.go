package piscine

func AlphaCount(s string) int {
	count := 0
	r := []rune(s)

	for c := 0; c < len(r); c++ {
		if (r[c] >= 'a' && r[c] <= 'z') || (r[c] >= 'A' && r[c] <= 'Z') {
			count++
		}
	}
	return count
}
