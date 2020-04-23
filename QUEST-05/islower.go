package piscine

func IsLower(str string) bool {
	strArr := []rune(str)
	for i := range strArr {
		if strArr[i] < 97 || strArr[i] > 122 {
			return false
		}
	}
	return true
}
