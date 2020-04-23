package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for i, letter := range arr {
		if letter >= 97 && letter <= 122 {
			arr[i] = letter - 32
		}
	}
	return string(arr)
}
