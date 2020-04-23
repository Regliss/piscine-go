package piscine

func TrimAtoi(s string) int {
	Negative := 1
	arr := []rune(s)
	var Array []rune
	result := 0
	for i, char := range arr {
		if char == '+' || char == '-' || char >= '0' && char <= '9' {
			Array = append(Array, arr[i])
		}
	}
	for i := range Array {
		if i == 0 {
			if Array[i] == '+' {
				Negative = 1
			} else if Array[i] == '-' {
				Negative = -1
			} else {
				result = result*10 + int(Array[i]) - '0'
			}
		} else {
			if Array[i] >= '0' && Array[i] <= '9' {
				result = result*10 + int(Array[i]) - '0'
			}
		}
	}
	return Negative * result
}
