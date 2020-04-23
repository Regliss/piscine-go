package piscine

func IsPrintable(str string) bool {
	strArr := []rune(str)
	for i := range strArr {
		if !(strArr[i] <= 127 && strArr[i] >= 27) {
			return false
		}
	}
	return true
}
