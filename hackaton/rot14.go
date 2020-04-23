package piscine

func Rot14(str string) string {
	runes := []rune(str)
	for i := range runes {
		x := runes[i] + 14
		if runes[i] >= 65 && runes[i] <= 90 {
			if x > 90 {
				runes[i] = x - 26
			} else {
				runes[i] = x
			}
		} else if runes[i] >= 97 && runes[i] <= 122 {
			if x > 122 {
				runes[i] = x - 26
			} else {
				runes[i] = x
			}
		}
	}
	return string(runes)
}
