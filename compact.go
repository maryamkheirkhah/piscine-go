package piscine

func Compact(ptr *[]string) int {
	count := 0
	ptrnew := []string(nil)
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			ptrnew = append(ptrnew, (*ptr)[i])
			count++
		}
	}
	*ptr = ptrnew
	return count
}
