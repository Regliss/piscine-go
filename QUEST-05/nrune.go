package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	for i := range arr {
		if i == n-1 {
			return arr[i]
		}
	}
	return 0
}
