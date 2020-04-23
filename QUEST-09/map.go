package piscine

func Map(f func(int) bool, arr []int) []bool {
	lenArr := 0
	for range arr {
		lenArr++
	}
	result := make([]bool, lenArr)
	for i, num := range arr {
		result[i] = f(num)
	}
	return result
}
