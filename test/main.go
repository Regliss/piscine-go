package main

import (
	"fmt"

	piscine ".."
)

func main() {
	// a := 0
	// b := &a
	// n := &b
	// piscine.UltimatePointOne(&n)
	// fmt.Println(a)

	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
