package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		r := []rune(args[i])
		for j := 0; j < len(r); j++ {
			z01.PrintRune(r[j])
		}
		z01.PrintRune('\n')
	}
}
