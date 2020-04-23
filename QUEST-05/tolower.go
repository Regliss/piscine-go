package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for i, letter := range arr {
		if letter >= 65 && letter <= 90 {
			arr[i] = letter + 32
		}
	}
	return string(arr)
}
