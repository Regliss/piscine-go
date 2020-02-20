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
	// s := "12345"
	// s2 := "0000000012345"
	// s3 := "000000"

	// n := piscine.BasicAtoi(s)
	// n2 := piscine.BasicAtoi(s2)
	// n3 := piscine.BasicAtoi(s3)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// s := "12345"
	// s2 := "0000000012345"
	// s3 := "012 345"
	// s4 := "Hello World!"

	// n := piscine.BasicAtoi2(s)
	// n2 := piscine.BasicAtoi2(s2)
	// n3 := piscine.BasicAtoi2(s3)
	// n4 := piscine.BasicAtoi2(s4)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++123+4"
	s8 := "--123-4"
	s9 := "-Invalid123"

	n := piscine.Atoi(s)
	n2 := piscine.Atoi(s2)
	n3 := piscine.Atoi(s3)
	n4 := piscine.Atoi(s4)
	n5 := piscine.Atoi(s5)
	n6 := piscine.Atoi(s6)
	n7 := piscine.Atoi(s7)
	n8 := piscine.Atoi(s8)
	n9 := piscine.Atoi(s9)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)

}
