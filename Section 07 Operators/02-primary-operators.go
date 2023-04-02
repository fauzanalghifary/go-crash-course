package main

import "math"

func main() {

	// Multiplication
	var radius = 12.0
	area := math.Pi * radius * radius
	println(area)

	//	Integer division
	half := 1 / 2
	println(half)

	// Floating point division
	halfFloat := 1.0 / 2.0
	println(halfFloat)

	//	Raising
	square := math.Pow(2, 2)
	println(square)

	//	Modulus
	remainder := 13 % 3
	println(remainder)

	//	Unary operators
	//	Increment
	var i = 1
	i++
	println(i)

	//	Decrement
	i--
	println(i)

}
