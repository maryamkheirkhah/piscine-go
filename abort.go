package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	arr = BubbleSort(arr)
	return arr[2]
}

func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
