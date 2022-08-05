package main

import (
	"fmt"
	piscine "piscine/quest5"

	"github.com/01-edu/z01"
)

func main() {
	PrintNbrBase(piscine.PrintNbrBase(125, "0123456789"))
}

func PrintNbrBase(nbr int, base string) {
	runes := []rune(base)
	var len int = 0
	for i := range runes {
		len++
		i++
	}
	v := true
	if len < 2 {
		v = false
	} else {
		for i := 0; i < len; i++ {
			if runes[i] == '-' || runes[i] == '+' {
				v = false
			}
			for j := i + 1; j < len; j++ {
				if runes[i] == runes[j] {
					v = false
				}
			}
		}
	}
	if !v {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr == 0 {
			z01.PrintRune(runes[0])
		} else {
			if nbr < 0 {
				z01.PrintRune('-')
			}
			BaseRecursion(nbr, runes, len)
		}
	}
}

func BaseRecursion(n int, runes []rune, len int) {
	if n/len != 0 {
		BaseRecursion(n/len, runes, len)
	}
	m := n % len
	if m < 0 {
		m = -m
	}
	fmt.Println(m)
	z01.PrintRune(runes[m])
}
