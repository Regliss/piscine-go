package piscine

func StrLen(str string) int {
	array := []byte(str)
	count := 0
	for range array {
		count++
	}
	return count
}
