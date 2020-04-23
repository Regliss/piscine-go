package main

import (
	"os"
	"io/ioutil"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune('\n')
	arguments := os.Args[1:]
	lenArg := 0
	for range arguments {
		lenArg++
	}
	if lenArg == 0 {
		input , err := ioutil.ReadAll(os.Stdin)
		for _, j := range string(input) {
			z01.PrintRune(j)
		}
		if err != nil {
			for _,k := range err.Error() {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
		return
	}
	for _, arg := range arguments {
		
	}

}
