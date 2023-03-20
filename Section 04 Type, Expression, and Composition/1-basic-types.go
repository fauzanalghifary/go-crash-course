package main

import "fmt"

// basic types (int, float, string, bool)
var myInt int

// var myInt16 int16
// var myInt32 int32
// var myInt64 int64
var myUint uint
var myFloat32 float32
var myFloat64 float64

func main2() {
	// String in Go are immutable
	myString := "Hello World"
	fmt.Println(myString)
	myString = "Hello Go"
	fmt.Println(myString)

	// Boolean
	var myBool = true
	fmt.Println(myBool)
}
