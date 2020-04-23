package piscine

func IsAlpha(str string) bool {
	strArr := []rune(str)
	for i := range strArr {
		if !(strArr[i] >= 'a' && strArr[i] <= 'z' || strArr[i] >= 'A' && strArr[i] <= 'Z' || strArr[i] <= '9' && strArr[i] >= '0') {
			return false
		}
	}
	return true
}
