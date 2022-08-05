package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	printResult(points)
}

func printResult(points *point) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.y)
	z01.PrintRune('\n')
}

func printInt(nbr int) {
	if nbr < 0 {
		z01.PrintRune('-')
	}
	temp := nbr
	r := 0
	var nbrArr []rune
	for temp != 0 {
		r = temp % 10
		nbrArr = append(nbrArr, rune(r))
		temp /= 10
	}
	for i := len(nbrArr) - 1; i >= 0; i-- {
		z01.PrintRune(nbrArr[i] + 48)
	}
}
