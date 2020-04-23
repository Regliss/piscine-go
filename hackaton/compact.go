package piscine

func Compact(ptr *[]string) int {
	arrstr := *ptr
	counter := 0
	lenstr := 0
	lenstr2 := 0

	for i := range arrstr {
		if arrstr[i] != "" {
			lenstr2++
		}
	}
	arrstr2 := make([]string, lenstr2)
	for i := range arrstr {
		if arrstr[i] == "" {
			counter++
		} else {
			arrstr2[lenstr] = arrstr[i]
			lenstr++

		}
	}
	*ptr = arrstr2
	return lenstr2
}
