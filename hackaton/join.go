package piscine

func Join(strs []string, sep string) string {
	result := ""
	for _,word := range strs {
		result = result + word + sep
	}
	return result
}
