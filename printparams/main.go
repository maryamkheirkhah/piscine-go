package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, res := range args {
		r := []rune(res)
		for i := 0; i < len(r); i++ {
			z01.PrintRune(r[i])
		}
		z01.PrintRune('\n')
	}
}
