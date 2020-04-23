package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func strLen(str string) int {
	a := 0
	for range str {
		a++
	}
	return a
}

func index(s string, toFind string) int {
	arr := []rune(s)
	for i := 0; i < strLen(s); i++ {
		for j := i; j <= strLen(s); j++ {
			if string(arr[i:j]) == toFind {
				return i
			}
		}
	}
	return -1
}

func main() {
	size := 0
	for range os.Args[1:] {
		size++
	}
	if size == 0 {
		printHelp()
		return
	}
	for _, s := range os.Args[1:] {
		if s == "-h" || s == "--help" {
			printHelp()
			return
		}
	}

	isOrder := false
	answer := ""
	for _, s := range os.Args[1:] {
		if index(s, "--order") != -1 || index(s, "-o") != -1 {
			isOrder = true
		}
		if index(s, "-i") == -1 && index(s, "--insert") == -1 && index(s, "--order") == -1 && index(s, "-o") == -1 {
			answer += s
		}
	}
	for _, s := range os.Args[1:] {
		if index(s, "--insert") != -1 {
			answer += s[9:]
		} else if index(s, "-i") != -1 {
			answer += s[3:]
		}
	}

	arrAnswer := []rune(answer)

	if isOrder {
		strsize := 0
		for range answer {
			strsize++
		}

		for i := 0; i < strsize; i++ {
			for j := i + 1; j < strsize; j++ {
				if arrAnswer[i] > arrAnswer[j] {
					arrAnswer[i], arrAnswer[j] = arrAnswer[j], arrAnswer[i]
				}
			}
		}
	}

	fmt.Println(string(arrAnswer))
}
