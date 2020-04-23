package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	lenA := max - min
	arr := make([]int, lenA, lenA)
	for i := 0; i < lenA; i++ {
		arr[i] = i + min
	}
	return arr
}
