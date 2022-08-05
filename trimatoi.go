package piscine

func TrimAtoi(s string) int {
	r := []rune(s)
	temp := 1
	for i := 0; i < len(r); i++ {
		if r[i] == '-' {
			temp = -i - 1
			break
		} else if r[i] == '+' || (r[i] >= '0' && r[i] <= '9') {
			temp = i
			break
		} else if i+1 == len(r) && (r[i] >= '0' && r[i] <= '9') {
			return 0
		}
	}
	if -temp >= len(r) {
		return 0
	}
	b := temp
	if temp < 0 {
		temp = -1
		b = -b
	} else {
		temp = 1
	}
	c := 0
	for i := b; i < len(r); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			c *= 10
			c += int(r[i] - '0')
		}
	}
	return c * temp
}
