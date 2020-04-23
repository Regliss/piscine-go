package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	arrS := []rune(s)
	arrtoFind := []rune(toFind)
	lenS := 0
	lentoFind := 0
	count := 0
	for range arrS {
		lenS++
	}
	for range arrtoFind {
		lentoFind++
	}
	if lenS == 0 {
		return -1
	} else if lentoFind == 0 {
		return 0
	} else if lentoFind > lenS {
		return -1
	}
	for i := range arrS {
		for j := range arrtoFind {
			if arrS[i] == arrtoFind[j] {
				count++
				break
			}
		}
		if count == lentoFind {
			return i + 1 - count
		}
	}
	return -1
}
