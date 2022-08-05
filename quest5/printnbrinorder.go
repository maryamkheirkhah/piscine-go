package piscine

import (
	"github.com/01-edu/z01"
)

func Sorting(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
		}
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		n *= -1
	}
	temp := n
	l := 0
	if temp == 0 {
		l = 1
	}
	for temp != 0 {
		temp /= 10
		l++
	}
	arr := make([]int, l)
	temp = n
	for i := 0; i < l; i++ {
		arr[i] = temp % 10
		temp /= 10
	}
	Sorting(arr)
	for i := 0; i < l; i++ {
		z01.PrintRune(rune(arr[i] + 48))
	}
}
