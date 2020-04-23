package piscine

func IsNumeric(str string) bool {
	strArr := []rune(str)
	for i := range strArr {
		if !(strArr[i] <= '9' && strArr[i] >= '0') {
			return false
		}
	}
	return true
}
