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

	// a := 13
	// b := 2
	// var div int
	// var mod int
	// piscine.DivMod(a, b, &div, &mod)
	// fmt.Println(div)
	// fmt.Println(mod)
	// a := 13
	// b := 2
	// piscine.UltimateDivMod(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// aString := "Hello"

	// aStringChangeable := []byte(aString)

	// aStringChangeable[0] = 'A'

	// aStringFinalized := string(aStringChangeable)

	// fmt.Println(aString)
	// fmt.Println(aStringChangeable)
	// fmt.Println(aStringFinalized)
	// str := "Hello World!"
	// piscine.PrintStr(str)
	// str := "Hello World!"
	// nb := piscine.StrLen(str)
	// fmt.Println(nb)
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
