package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	flag := false
	if len(a) == 0 {
		return false
	}
	if len(a) == 1 {
		return true
	}
	for i := 0; i < len(a)-1 && !flag; i++ {
		if a[i] < a[i+1] {
			flag = true
			for i := range a {
				for j := i + 1; j < len(a); j++ {
					if f(a[i], a[j]) > 0 {
						return false
					}
				}
			}
			return true
		} else if a[i] > a[i+1] {
			flag = true
			for i := range a {
				for j := i + 1; j < len(a); j++ {
					if f(a[i], a[j]) < 0 {
						return false
					}
				}
			}
			return true
		}
	}
	return true
}
