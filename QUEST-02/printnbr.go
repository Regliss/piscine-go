package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	last := false
	if n == -9223372036854775808 {
		n = -922337203685477580
		last = true
	}
	if n < 0 {
		n = -1 * n
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var arr [36]int
	div := n
	rem := 0
	i := 0
	for div > 0 {
		rem = div % 10
		div /= 10
		arr[i] = rem
		i++
	}
	for j := i - 1; j >= 0; j-- {
		c := 0
		char := '0'
		for c < arr[j] {
			char++
			c++
		}
		z01.PrintRune(char)
	}
	if last {
		z01.PrintRune('8')
	}
}
