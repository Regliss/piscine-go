package main

import (
	"github.com/01-edu/z01"
)

func setPoint(n *point) {
	n.x = -42
	n.y = 21
}

type point struct {
	x int
	y int
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(32)
	z01.PrintRune('=')
	z01.PrintRune(32)
	printNum(points.x)
	z01.PrintRune(',')
	z01.PrintRune(32)
	z01.PrintRune('y')
	z01.PrintRune(32)
	z01.PrintRune('=')
	z01.PrintRune(32)
	printNum(points.y)
	z01.PrintRune('\n')
}
func printNum(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}
func check(n int) {
	c := '0'
	if n == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= n%10; i++ {
		c++
	}
	for i := -1; i >= n%10; i-- {
		c++
	}
	if n/10 != 0 {
		check(n / 10)
	}
	z01.PrintRune(c)
	return
}
