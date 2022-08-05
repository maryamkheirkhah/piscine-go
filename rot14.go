package piscine

func Rot14(s string) string {
	runeArr := []rune(s)
	for i, r := range runeArr {
		if r <= 'z' && r >= 'a' {
			if 'z'-runeArr[i] >= 14 {
				runeArr[i] = runeArr[i] + 14
			} else {
				dif := 14 - ('z' - runeArr[i])
				runeArr[i] = 'a' + dif - 1
			}
		} else if r <= 'Z' && r >= 'A' {
			if 'Z'-runeArr[i] >= 14 {
				runeArr[i] = runeArr[i] + 14
			} else {
				dif := 14 - ('Z' - runeArr[i])
				runeArr[i] = 'A' + dif - 1
			}
		}
	}
	var str string
	for _, r := range runeArr {
		str += string(r)
	}
	return string(str)
}
