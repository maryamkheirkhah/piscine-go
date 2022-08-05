package piscine

func Split(s, sep string) []string {
	var slices []string
	sRunesArr := []rune(s)
	sepRunesArr := []rune(sep)
	sepFlag := false
	var word string
	var temp string
	for i := 0; i < len(sRunesArr); i++ {
		sepFlag = false
		if sRunesArr[i] == sepRunesArr[0] {
			sepFlag = true
			temp = string(sepRunesArr[0])
			for j := 0; j < len(sepRunesArr) && j+i < len(sRunesArr)-1; j++ {
				if sRunesArr[i+j] != sepRunesArr[j] {
					temp += string(sepRunesArr[j])
					sepFlag = false
					temp = ""
					break
				}
			}
		}
		if sepFlag {
			i += len(sepRunesArr) - 1
			slices = append(slices, word)
			word = ""
		} else {
			word += string(sRunesArr[i])
		}
		if i == len(sRunesArr)-1 {
			slices = append(slices, word)
		}
	}
	return slices
}
