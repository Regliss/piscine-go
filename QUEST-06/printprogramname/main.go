package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args
	runes := []rune(args[0])
	for i := range runes {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
