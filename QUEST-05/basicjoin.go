package piscine

func BasicJoin(strs []string) string {
	s := ""
	count := 0
	for range strs {
		count++
	}
	if count == 0 {
		return s
	}
	for i := range strs {
		s += strs[i]
	}
	return s
}
