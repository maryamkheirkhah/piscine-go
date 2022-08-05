package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	a := bubbleSorting(args)
	for _, str := range a {
		r := []rune(str)
		for i := range r {
			z01.PrintRune(r[i])
		}
		z01.PrintRune('\n')
	}
}

func bubbleSorting(arr []string) []string {
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
