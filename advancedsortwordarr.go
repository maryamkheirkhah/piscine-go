package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if f(a[j], a[i]) == -1 {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
}
