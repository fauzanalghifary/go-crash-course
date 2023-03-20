package main

import "fmt"

func main() {
	// Pointer
	var myInt int
	myInt = 10

	fmt.Println(myInt)

	var myPointer *int
	myPointer = &myInt
	*myPointer = 20
	fmt.Println(myPointer)
	fmt.Println(myInt)

	changeValueOfPointer(&myInt)
	fmt.Println(myInt)
}

func changeValueOfPointer(p *int) {
	*p = 30
}
