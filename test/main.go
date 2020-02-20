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
	// a := 0
	// b := 1
	// piscine.Swap(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// s := "Hello World!"
	// s = piscine.StrRev(s)
	// fmt.Println(s)
	s := "12345"
	s2 := "0000000012345"
	s3 := "000000"

	n := piscine.BasicAtoi(s)
	n2 := piscine.BasicAtoi(s2)
	n3 := piscine.BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}
