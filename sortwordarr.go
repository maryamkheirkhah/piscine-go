package piscine

func SortWordArr(a []string) {
	AsciiBubbleSort(a)
}

func AsciiBubbleSort(a []string) []string {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
	return a
}
