package piscine

func Max(a []int) int {
	if a == nil {
		return 0
	}
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
