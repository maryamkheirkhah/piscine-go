package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if !(len(os.Args) == 1) {
		args := os.Args[1:]
		caps := false
		if args[0] == "--upper" {
			caps = true
			args = args[1:]
		}
		for _, arg := range args {
			num := 0
			for _, v := range arg {
				num = num*10 + int(v-'0')
			}
			if num >= 1 && num <= 26 {
				if caps == false {
					z01.PrintRune(rune(num + 96))
				} else {
					z01.PrintRune(rune(num + 64))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
