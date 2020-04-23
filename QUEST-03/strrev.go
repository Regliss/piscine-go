package piscine

func StrRev(s string) string {
	Array := []rune(s)
	newString := ""
	counter := 0
	for range Array {
		counter++
	}
	for i := counter - 1; i >= 0; i-- {
		newString += string(Array[i])
	}
	return newString
}
