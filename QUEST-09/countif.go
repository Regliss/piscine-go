package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, word := range tab {
		if f(word) == true {
			counter++
		}
	}
	return counter
}
