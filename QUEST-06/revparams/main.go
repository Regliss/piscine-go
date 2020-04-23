package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args[1:]
	lenArgs := 0
	for range arguments {
		lenArgs++
	}
	for i := lenArgs - 1; i >= 0; i-- {
		ar := []rune(arguments[i])
		for j := range ar {
			z01.PrintRune(ar[j])
		}
		z01.PrintRune('\n')
	}
}
