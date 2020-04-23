package piscine

func Join(strs []string, sep string) string {
	s := ""
	count := 0
	for range strs {
		count++
	}
	if count == 0 {
		return s
	}
	for i := 0; i < count-1; i++ {
		s += strs[i] + sep
	}
	return s + strs[count-1]
}
