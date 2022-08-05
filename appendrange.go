package piscine

func AppendRange(min, max int) []int {
	var slices []int
	if min > max {
		return nil
	}
	for i := min; i < max; i++ {
		slices = append(slices, i)
	}
	return slices
}
