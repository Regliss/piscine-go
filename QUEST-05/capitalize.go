package piscine

func Capitalize(s string) string {
	status := 0
	arr := []rune(s)
	for i := range arr {
		if status == 0 {
			if arr[i] >= 'a' && arr[i] <= 'z' {
				status = 1
				arr[i] = arr[i] - 32
			} else if arr[i] <= '9' && arr[i] >= '0' || arr[i] >= 'A' && arr[i] <= 'Z'{
				status = 1
			}
		} else {
			if arr[i] >= 'A' && arr[i] <= 'Z' {
				arr[i] = arr[i] + 32
			} else if !(arr[i] >= 'a' && arr[i] <= 'z' || arr[i] >= 'A' && arr[i] <= 'Z' || arr[i] <= '9' && arr[i] >= '0') {
				status = 0
			}
		}
	}
	return string(arr)
}
