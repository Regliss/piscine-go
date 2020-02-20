package piscine

func StrLen(str string) int {
	array := []rune(str)
	count := 1
	for range array {
		count += 1
	}
	return count
}
