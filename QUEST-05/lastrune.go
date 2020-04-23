package piscine

func LastRune(s string) rune {
	arr := []rune(s)
	sum := 0
	for range arr {
		sum += 1
	}
	return arr[sum-1]
}
