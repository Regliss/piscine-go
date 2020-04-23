package piscine

func ForEach(f func(int), arr []int) {
	for _, element := range arr {
		f(element)
	}
}
