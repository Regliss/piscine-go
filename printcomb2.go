package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i == '9' && j == '8' && k == '9' && l == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(32)
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune('\n')
						return
					}
					if int(i) * 10 + int(j) < int(k) * 10 + int(l) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(32)
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}
