package piscine

func ActiveBits(n int) int {
	count := 0
	temp := n
	for temp > 0 {
		if temp%2 == 1 {
			count++
		}
		temp /= 2
	}
	return count
}
