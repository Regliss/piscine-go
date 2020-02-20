package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	aStringChangeable := []rune(str)

	for i := range aStringChangeable {
		z01.PrintRune(aStringChangeable[i])
	}
	z01.PrintRune('\n')
}
