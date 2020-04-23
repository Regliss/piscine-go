package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for i := range arguments {
		param := []rune(arguments[i])
		for j := range param {
			z01.PrintRune(param[j])
		}
		z01.PrintRune('\n')
	}
}
