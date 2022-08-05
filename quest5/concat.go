package piscine

func Concat(str1 string, str2 string) string {
	r1 := []rune(str1)
	r2 := []rune(str2)
	r3 := append(r1, r2...)
	return string(r3)
}
