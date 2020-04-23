package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arr := os.Args[1:]
	lenArr := 0
	for range arr {
		lenArr++
	}
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < lenArr-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}
	for i := range arr {
		param := []rune(arr[i])
		for j := range param {
			z01.PrintRune(param[j])
		}
		z01.PrintRune('\n')
	}
}
