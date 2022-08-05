package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb > 20 {
		return 0
	} else {
		fact := 1
		for i := 1; i <= nb; i++ {
			fact *= i
		}
		return fact
	}
}
