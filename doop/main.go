package main

import "os"

func main()  {
	arguments := os.Args[1:]
	lenArg := 0
	for range arguments {
		lenArg++
	}
	if lenArg != 3 {
		return
	}
}

func plus()  {
	
}
func minus() {

}
func divide() {

}
func multiply() {}