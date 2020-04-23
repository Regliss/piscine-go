package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var arr []int

	len := 0

	for l := 0; n > 0; l++ {
		arr = append(arr, n%10)
		n /= 10
	}
	for range arr {
		len++
	}

	for i := 1; i < len; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	if len == 0 {
		z01.PrintRune('0')
		return
	}
	for z := range arr {
		z01.PrintRune(rune(arr[z]) + 48)
	}
}
