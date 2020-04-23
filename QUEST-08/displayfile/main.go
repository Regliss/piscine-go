package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	lenArg := 0
	for range arguments {
		lenArg++
	}
	if lenArg == 0 {
		fmt.Println("File name missing")
		return
	} else if lenArg > 1 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := ioutil.ReadFile(arguments[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(file))
}
