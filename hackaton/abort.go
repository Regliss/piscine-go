package piscine

func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}
	for i := 0; i <= 3; i++ {
		for j := i + 1; j <= 4; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[2]
}
