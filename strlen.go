package piscine

func StrLen(str string) int {
	array := []byte(str)
	count := 0
	for _ =range array {
		count++
	}
	return count
}
