package piscine

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 1 {
		return 2
	}
	p := false
	for !p {
		if nb < 1 {
			return 2
		}
		p = isPrime(nb)
		if p {
			return nb
		}
		nb += 1
	}
	return 1
}
