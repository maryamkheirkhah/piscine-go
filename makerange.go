package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slices := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		slices[i] = min + i
	}
	return slices
}
