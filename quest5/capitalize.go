package piscine

func IsAlphanum(r rune) bool {
	if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
		return false
	}
	return true
}

func UpRune(r rune) rune {
	if !(r < 'a' || r > 'z') {
		r += 'A' - 'a'
	}
	return r
}

func LowRune(r rune) rune {
	if !(r < 'A' || r > 'Z') {
		r += 'a' - 'A'
	}
	return r
}

func Capitalize(s string) string {
	tempStr := []rune(s)
	state := 0
	for i := 0; i < len(tempStr); i++ {
		if state == 0 && IsAlphanum(tempStr[i]) {
			tempStr[i] = UpRune(tempStr[i])
			state = 1
		} else if state == 1 && IsAlphanum(tempStr[i]) {
			tempStr[i] = LowRune(tempStr[i])
		} else {
			state = 0
		}
	}
	return string(tempStr)
}
