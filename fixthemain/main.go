package main

import "github.com/01-edu/z01"

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	z01.PrintRune('\n')
	ptrDoor.state = true
}
func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	ptrDoor.state = false
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	if ptrDoor.state == true {
		return true
	} else {
		return false
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	if ptrDoor.state == false {
		return true
	} else {
		return false
	}
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}

type Door struct {
	state bool
}
