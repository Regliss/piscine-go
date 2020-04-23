package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arr := os.Args[1:]
	len := 0
	upper := false
	for range arr {
		len++
	}
	for i := 0; i < len; i++ {
		if arr[i] == "--upper" {
			upper = true
		} else if BasicAtoi(arr[i]) >= 1 && BasicAtoi(arr[i]) <= 26 {
			if upper == true {
				z01.PrintRune(rune(BasicAtoi(arr[i]) + 64))
			} else {
				z01.PrintRune(rune(BasicAtoi(arr[i]) + 96))
			}
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

func BasicAtoi(s string) int {
	var result int = 0
	Array := []rune(s)
	for i := range Array {
		result = result*10 + int(Array[i]) - '0'
	}
	return result
}
