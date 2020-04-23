package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(res bool) {
	if res == true {
		for _, letter := range "I have an even number of arguments" {
			z01.PrintRune(letter)
		}
	} else {
		for _, letter := range "I have an odd number of arguments" {
			z01.PrintRune(letter)
		}
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	arguments := os.Args[1:]
	lengthOfArg := 0
	for range arguments {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == true {
		printStr(isEven(lengthOfArg))
	} else {
		printStr(isEven(lengthOfArg))
	}
}
