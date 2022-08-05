package piscine

func StrLen(s string) int {
	Astr := []rune(s)
	count := 0
	for i := range Astr {
		i++
		count++
	}
	return count
}
