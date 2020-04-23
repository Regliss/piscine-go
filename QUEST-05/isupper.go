package piscine

func IsUpper(str string) bool {
	strArr := []rune(str)
	for i := range strArr {
		if strArr[i] < 65 || strArr[i] > 90 {
			return false
		}
	}
	return true
}
