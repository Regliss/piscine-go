package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	runes := []rune{'0' - 1, '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	first := true
	printTwo(1, runes, n, &first)
	z01.PrintRune('\n')
}

func printTwo(curr int, runes []rune, n int, first *bool) {
	if curr == n+1 {
		if !*first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*first = false

		for i := 1; i < n+1; i++ {
			z01.PrintRune(runes[i])
		}
	}

	for a := '0'; a <= '9'; a++ {
		if a > runes[curr-1] {
			runes[curr] = a
			printTwo(curr+1, runes, n, first)
		}
	}
}
