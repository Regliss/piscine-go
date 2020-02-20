package piscine

func BasicAtoi(s string) int {
	var result int = 0
	Array := []rune(s)
	for i := range Array {
		result = result * 10 + int(Array[i]) - '0'
	}
	return result
}
