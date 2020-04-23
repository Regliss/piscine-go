package piscine

func AlphaCount(str string) int {
	count := 0
	strArr := []rune(str)
	for i := range strArr {
		if strArr[i] >= 97 && strArr[i] <= 122 {
			count++
		} else if strArr[i] >= 65 && strArr[i] <= 90 {
			count++
		}
	}
	return count
}
