package piscine

func Index(s string, toFind string) int {
	sr := []rune(s)
	fr := []rune(toFind)
	flag := true
	for i, s := range sr {
		if s == fr[0] {
			flag = true
			for j := 0; j < len(fr) && i+j < len(sr); j++ {
				if fr[j] != sr[i+j] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}

		}
	}
	return -1
}
