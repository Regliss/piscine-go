package piscine

func Max(arr []int) int {
	for i := 0; i <= len(arr)-2; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[len(arr)-1]
}
