package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

func even(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	args := os.Args[1:]
	lengthOfArg := 0
	for i := range args {
		lengthOfArg = i + 1
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
