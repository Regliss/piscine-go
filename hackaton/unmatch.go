package piscine

func Unmatch(arr []int) int {
	count := 0
	for _, num := range arr {
		count = 0
		for _, num2 := range arr {
			if num2 == num {
				count++
			}
		}
		if count%2 != 0 {
			return num
		}
	}
	return -1
}
