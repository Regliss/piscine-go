package piscine

func BasicAtoi2(s string) int {
	var result int = 0
	Array := []rune(s)
	for i := range Array {
		if CheckForNum(Array[i]) {
			result = result*10 + int(Array[i]) - '0'
		} else {
			return 0
		}
	}
	return result
}
func CheckForNum(c rune) bool {
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := range numbers {
		if numbers[i] == c {
			return true
		}
	}
	return false
}
